package main

import (
	"flag"
	"fmt"

	"github.com/wuntsong-org/goctlwt/example/rpc/hi/internal/config"
	eventServer "github.com/wuntsong-org/goctlwt/example/rpc/hi/internal/server/event"
	greetServer "github.com/wuntsong-org/goctlwt/example/rpc/hi/internal/server/greet"
	"github.com/wuntsong-org/goctlwt/example/rpc/hi/internal/svc"
	"github.com/wuntsong-org/goctlwt/example/rpc/hi/pb/hi"

	"github.com/wuntsong-org/go-zero-plus/core/conf"
	"github.com/wuntsong-org/go-zero-plus/core/service"
	"github.com/wuntsong-org/go-zero-plus/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/hi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		hi.RegisterGreetServer(grpcServer, greetServer.NewGreetServer(ctx))
		hi.RegisterEventServer(grpcServer, eventServer.NewEventServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
