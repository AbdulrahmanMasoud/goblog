package routing

import (
	"fmt"
	"github.com/AbdulrahmanMasoud/goblog/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Can't serve the application")
	}
}
