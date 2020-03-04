package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"go-disk/common/rpcinterface/uploadinterface"
	"go-disk/config"
	"go-disk/services/upload/router"
	"go-disk/services/upload/rpc"
	"log"
)

func main() {
	uploadRouter := router.Router()

	go startRpcService()

	err := uploadRouter.Run(":9000")
	if err != nil {
		panic(err)
	}
}

func startRpcService() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			config.ConsulAddress,
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.upload"))

	service.Init()

	err := uploadinterface.RegisterUploadServiceHandler(service.Server(), new(rpc.EndPoint))
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}

	log.Printf("start upload service success")
}