package proxy_pool

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"

	"demo/internal/dao"
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/os/gctx"
)

// Pool IP 代理池，从数据库加载，轮询分配
type Pool struct {
	proxies []entity.ProxyPools
	index   atomic.Uint64
	mu      sync.RWMutex
}

var defaultPool = &Pool{}

// Reload 从数据库重新加载所有启用的代理
func Reload() error {
	ctx := gctx.New()
	var items []entity.ProxyPools
	err := dao.ProxyPool.Ctx(ctx).
		Where(entity.ProxyPools{Status: 1}).
		Scan(&items)
	if err != nil {
		return err
	}
	defaultPool.mu.Lock()
	defaultPool.proxies = items
	defaultPool.mu.Unlock()
	return nil
}

// Get 轮询获取一个代理，无可用代理返回 nil
func Get() *entity.ProxyPools {
	defaultPool.mu.RLock()
	defer defaultPool.mu.RUnlock()

	if len(defaultPool.proxies) == 0 {
		return nil
	}
	idx := defaultPool.index.Add(1) - 1
	proxy := defaultPool.proxies[idx%uint64(len(defaultPool.proxies))]
	return &proxy
}

// ProxyURL 返回代理的 http URL 字符串
func ProxyURL(p *entity.ProxyPools) string {
	return fmt.Sprintf("http://%s:%s@%s:%d", p.Username, p.Password, p.Host, p.Port)
}

// HTTPClient 获取一个带代理的 http.Client，无代理时返回默认 Client
func HTTPClient() *http.Client {
	proxy := Get()
	if proxy == nil {
		return http.DefaultClient
	}
	proxyURL, err := url.Parse(ProxyURL(proxy))
	if err != nil {
		return http.DefaultClient
	}
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
}

// Count 返回当前可用代理数量
func Count() int {
	defaultPool.mu.RLock()
	defer defaultPool.mu.RUnlock()
	return len(defaultPool.proxies)
}

// ReloadWithCtx 带 context 的重新加载
func ReloadWithCtx(ctx context.Context) error {
	var items []entity.ProxyPools
	err := dao.ProxyPool.Ctx(ctx).
		Where(entity.ProxyPools{Status: 1}).
		Scan(&items)
	if err != nil {
		return err
	}
	defaultPool.mu.Lock()
	defaultPool.proxies = items
	defaultPool.mu.Unlock()
	return nil
}
