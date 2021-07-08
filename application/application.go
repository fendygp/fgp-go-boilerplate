package application

import (
	"github.com/fgp-go-boilerplate/config"
	"github.com/fgp-go-boilerplate/router"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
)

func StartApp() {
	addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.NewRouter(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
