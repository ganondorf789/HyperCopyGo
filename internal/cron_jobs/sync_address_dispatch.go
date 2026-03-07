package cron_jobs

import (
	"context"
	"encoding/json"

	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	redisKeyAssignment = "addr_dispatch:assignment" // Hash: address → server_ip
	maxAddrsPerServer  = 200
	leaderboardMinVlm  = 1000
)

type dispatchNotification struct {
	Subscribe   []string `json:"subscribe"`
	Unsubscribe []string `json:"unsubscribe"`
}

func init() {
	Register("sync_address_dispatch", SyncAddressDispatch)
}

func SyncAddressDispatch(ctx context.Context, _ string) {
	addresses := collectAddresses(ctx)
	if len(addresses) == 0 {
		g.Log().Infof(ctx, "[dispatch] no addresses to dispatch")
		return
	}

	var servers []entity.ServerManagement
	if err := dao.ServerManagement.Ctx(ctx).Scan(&servers); err != nil {
		g.Log().Errorf(ctx, "[dispatch] load servers error: %v", err)
		return
	}
	if len(servers) == 0 {
		g.Log().Warningf(ctx, "[dispatch] no servers available")
		return
	}

	serverIPs := make(map[string]bool, len(servers))
	for _, s := range servers {
		serverIPs[s.Ip] = true
	}

	addrSet := make(map[string]bool, len(addresses))
	for _, a := range addresses {
		addrSet[a] = true
	}

	oldMap, err := g.Redis().HGetAll(ctx, redisKeyAssignment)
	if err != nil {
		g.Log().Errorf(ctx, "[dispatch] redis HGetAll error: %v", err)
		return
	}
	oldAssignment := oldMap.MapStrStr()

	// count: how many addresses each server currently holds (for new round)
	serverCount := make(map[string]int, len(servers))
	for _, s := range servers {
		serverCount[s.Ip] = 0
	}

	newAssignment := make(map[string]string, len(addresses))

	// Phase 1: keep sticky assignments that are still valid
	for _, addr := range addresses {
		if oldIP, ok := oldAssignment[addr]; ok && serverIPs[oldIP] {
			newAssignment[addr] = oldIP
			serverCount[oldIP]++
		}
	}

	// Phase 2: assign remaining addresses to least-loaded servers
	for _, addr := range addresses {
		if _, ok := newAssignment[addr]; ok {
			continue
		}
		ip := pickServer(serverCount)
		if ip == "" {
			g.Log().Warningf(ctx, "[dispatch] all servers full, %d addresses unassigned", len(addresses)-len(newAssignment))
			break
		}
		newAssignment[addr] = ip
		serverCount[ip]++
	}

	// Phase 3: compute diff per server
	// old: per-server address set
	oldPerServer := make(map[string]map[string]bool)
	for addr, ip := range oldAssignment {
		if oldPerServer[ip] == nil {
			oldPerServer[ip] = make(map[string]bool)
		}
		oldPerServer[ip][addr] = true
	}
	// new: per-server address set
	newPerServer := make(map[string]map[string]bool)
	for addr, ip := range newAssignment {
		if newPerServer[ip] == nil {
			newPerServer[ip] = make(map[string]bool)
		}
		newPerServer[ip][addr] = true
	}

	// Phase 4: write new assignment to Redis (full replace)
	if _, err := g.Redis().Del(ctx, redisKeyAssignment); err != nil {
		g.Log().Errorf(ctx, "[dispatch] redis DEL error: %v", err)
		return
	}
	if len(newAssignment) > 0 {
		fields := make(map[string]interface{}, len(newAssignment))
		for addr, ip := range newAssignment {
			fields[addr] = ip
		}
		if err := g.Redis().HMSet(ctx, redisKeyAssignment, fields); err != nil {
			g.Log().Errorf(ctx, "[dispatch] redis HMSet error: %v", err)
			return
		}
	}

	// Phase 5: publish notifications per server
	allIPs := make(map[string]bool)
	for ip := range oldPerServer {
		allIPs[ip] = true
	}
	for ip := range newPerServer {
		allIPs[ip] = true
	}

	for ip := range allIPs {
		oldAddrs := oldPerServer[ip]
		newAddrs := newPerServer[ip]

		var toSubscribe, toUnsubscribe []string
		for addr := range newAddrs {
			if !oldAddrs[addr] {
				toSubscribe = append(toSubscribe, addr)
			}
		}
		for addr := range oldAddrs {
			if !newAddrs[addr] {
				toUnsubscribe = append(toUnsubscribe, addr)
			}
		}

		if len(toSubscribe) == 0 && len(toUnsubscribe) == 0 {
			continue
		}

		notification := dispatchNotification{
			Subscribe:   toSubscribe,
			Unsubscribe: toUnsubscribe,
		}
		payload, _ := json.Marshal(notification)
		channel := "server:" + ip + ":dispatch"
		if _, err := g.Redis().Publish(ctx, channel, string(payload)); err != nil {
			g.Log().Errorf(ctx, "[dispatch] publish to %s error: %v", channel, err)
		} else {
			g.Log().Infof(ctx, "[dispatch] notified %s: +%d -%d", ip, len(toSubscribe), len(toUnsubscribe))
		}
	}

	g.Log().Infof(ctx, "[dispatch] done: %d addresses → %d servers", len(newAssignment), len(servers))
}

// collectAddresses gathers unique addresses from leaderboard, copy_trade_config, my_track_wallet.
func collectAddresses(ctx context.Context) []string {
	seen := make(map[string]bool)

	// 1. Leaderboard: weekly vlm >= 1000
	var lbAddrs []struct {
		EthAddress string `json:"ethAddress" orm:"eth_address"`
	}
	err := dao.Leaderboard.Ctx(ctx).
		Where("window = ? AND vlm >= ?", "week", leaderboardMinVlm).
		Fields("DISTINCT eth_address").
		Scan(&lbAddrs)
	if err != nil {
		g.Log().Errorf(ctx, "[dispatch] query leaderboard error: %v", err)
	}
	for _, r := range lbAddrs {
		if r.EthAddress != "" {
			seen[r.EthAddress] = true
		}
	}

	// 2. CopyTradeConfig: follow_type = auto(1), status = enabled(1)
	var ctAddrs []struct {
		TargetWallet string `json:"targetWallet" orm:"target_wallet"`
	}
	err = dao.CopyTradeConfig.Ctx(ctx).
		Where("follow_type = ? AND status = ?", consts.FollowTypeAuto, 1).
		Fields("DISTINCT target_wallet").
		Scan(&ctAddrs)
	if err != nil {
		g.Log().Errorf(ctx, "[dispatch] query copy_trade_config error: %v", err)
	}
	for _, r := range ctAddrs {
		if r.TargetWallet != "" {
			seen[r.TargetWallet] = true
		}
	}

	// 3. MyTrackWallet: status = enabled(1)
	var twAddrs []struct {
		Wallet string `json:"wallet" orm:"wallet"`
	}
	err = dao.MyTrackWallet.Ctx(ctx).
		Where("status = ?", 1).
		Fields("DISTINCT wallet").
		Scan(&twAddrs)
	if err != nil {
		g.Log().Errorf(ctx, "[dispatch] query my_track_wallet error: %v", err)
	}
	for _, r := range twAddrs {
		if r.Wallet != "" {
			seen[r.Wallet] = true
		}
	}

	result := make([]string, 0, len(seen))
	for addr := range seen {
		result = append(result, addr)
	}

	g.Log().Infof(ctx, "[dispatch] collected %d unique addresses (leaderboard:%d copy_trade:%d track:%d)",
		len(result), len(lbAddrs), len(ctAddrs), len(twAddrs))
	return result
}

// pickServer returns the server IP with the least addresses (under maxAddrsPerServer), or "" if all full.
func pickServer(counts map[string]int) string {
	minIP := ""
	minCount := maxAddrsPerServer + 1
	for ip, cnt := range counts {
		if cnt < maxAddrsPerServer && cnt < minCount {
			minCount = cnt
			minIP = ip
		}
	}
	return minIP
}
