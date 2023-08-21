package bootstrap

import (
	"github.com/AbdulrahmanMasoud/goblog/pkg/config"
	"github.com/AbdulrahmanMasoud/goblog/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	routing.RegisterRouter()
	routing.Serve()
}
