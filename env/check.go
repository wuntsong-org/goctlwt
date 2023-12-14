package env

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/wuntsong-org/goctlwt/pkg/env"
	"github.com/wuntsong-org/goctlwt/pkg/protoc"
	"github.com/wuntsong-org/goctlwt/pkg/protocgengo"
	"github.com/wuntsong-org/goctlwt/pkg/protocgengogrpc"
	"github.com/wuntsong-org/goctlwt/util/console"
)

type bin struct {
	name   string
	exists bool
	get    func(cacheDir string) (string, error)
}

var bins = []bin{
	{
		name:   "protoc",
		exists: protoc.Exists(),
		get:    protoc.Install,
	},
	{
		name:   "protoc-gen-go",
		exists: protocgengo.Exists(),
		get:    protocgengo.Install,
	},
	{
		name:   "protoc-gen-go-grpc",
		exists: protocgengogrpc.Exists(),
		get:    protocgengogrpc.Install,
	},
}

func check(_ *cobra.Command, _ []string) error {
	return Prepare(boolVarInstall, boolVarForce, boolVarVerbose)
}

func Prepare(install, force, verbose bool) error {
	log := console.NewColorConsole(verbose)
	pending := true
	log.Info("[goctlwt-env]: preparing to check env")
	defer func() {
		if p := recover(); p != nil {
			log.Error("%+v", p)
			return
		}
		if pending {
			log.Success("\n[goctlwt-env]: congratulations! your goctlwt environment is ready!")
		} else {
			log.Error(`
[goctlwt-env]: check env finish, some dependencies is not found in PATH, you can execute
command 'goctlwt env check --install' to install it, for details, please execute command 
'goctlwt env check --help'`)
		}
	}()
	for _, e := range bins {
		time.Sleep(200 * time.Millisecond)
		log.Info("")
		log.Info("[goctlwt-env]: looking up %q", e.name)
		if e.exists {
			log.Success("[goctlwt-env]: %q is installed", e.name)
			continue
		}
		log.Warning("[goctlwt-env]: %q is not found in PATH", e.name)
		if install {
			install := func() {
				log.Info("[goctlwt-env]: preparing to install %q", e.name)
				path, err := e.get(env.Get(env.GoctlCache))
				if err != nil {
					log.Error("[goctlwt-env]: an error interrupted the installation: %+v", err)
					pending = false
				} else {
					log.Success("[goctlwt-env]: %q is already installed in %q", e.name, path)
				}
			}
			if force {
				install()
				continue
			}
			console.Info("[goctlwt-env]: do you want to install %q [y: YES, n: No]", e.name)
			for {
				var in string
				fmt.Scanln(&in)
				var brk bool
				switch {
				case strings.EqualFold(in, "y"):
					install()
					brk = true
				case strings.EqualFold(in, "n"):
					pending = false
					console.Info("[goctlwt-env]: %q installation is ignored", e.name)
					brk = true
				default:
					console.Error("[goctlwt-env]: invalid input, input 'y' for yes, 'n' for no")
				}
				if brk {
					break
				}
			}
		} else {
			pending = false
		}
	}
	return nil
}
