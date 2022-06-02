package main

import (
	"context"
	"fmt"
	"log"

	"github.com/projectcalico/calico/libcalico-go/lib/options"

	"github.com/projectcalico/calico/calicoctl/calicoctl/commands/clientmgr"
)

func main() {
	// Load the client config and connect.
	cf := "./calicoctl.cfg"
	client, err := clientmgr.NewClient(cf)
	if err != nil {
		log.Fatalf("calicoctl 初始化失败: %s", err)
	}

	if ci, err := client.Nodes().List(context.TODO(), options.ListOptions{}); err == nil {
		for _, v := range ci.Items {
			fmt.Printf("节点名称:%s   网段:%s", v.Name, v.Status.PodCIDRs[0])
		}
	}
}
