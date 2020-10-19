package main

import (
	"github.com/zilliztech/milvus-distributed/internal/proxy"
	"log"
)

func main() {
	cfg, err := proxy.ReadProxyOptionsFromConfig()
	if err != nil {
		log.Fatalf("read proxy options form config file , error = %v", err)
	}
	err = proxy.StartProxy(cfg)
	if err != nil {
		log.Fatalf("start proxy failed, error = %v", err)
	}
}
