package grpc_client_test

import (
	"fmt"
	"testing"

	"demo/internal/grpc_client"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestGetAutoCopyTradeConfigList(t *testing.T) {
	ctx := gctx.New()

	client := grpc_client.NewCopyTradeConfigClient("hypercopy", "your-app-id", "your-app-secret")
	defer client.Close()

	res, err := client.GetAutoCopyTradeConfigList(ctx)
	if err != nil {
		t.Fatalf("GetAutoCopyTradeConfigList failed: %v", err)
	}

	fmt.Printf("Total: %d\n", res.Total)
	for _, item := range res.List {
		fmt.Printf("ID: %d, Wallet: %s, Status: %d\n", item.Id, item.TargetWallet, item.Status)
	}
}

