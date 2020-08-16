package main

import (
	"encoding/json"
	"net/http"

	"github.com/distributed-infras/messenger/sandbox/reference/config"
	"github.com/distributed-infras/messenger/sandbox/reference/protocol"
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/service"
)

func main() {
	routes := []router.Route{
		{
			Method:      http.MethodGet,
			MatchPrefix: false,
			Path:        "/allocate-connection-hub",
			Handle: func(w http.ResponseWriter, r *http.Request, params router.Params) {
				cfg := config.ConnectionHub{
					Host:     "ch1",
					Port:     8000,
					Protocol: protocol.WebSocket,
				}
				buf, err := json.Marshal(cfg)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(buf)
			},
		},
	}

	routingService := service.
		NewRoutingBuilder("Example").
		Routes(routes).
		Build()
	routingService.StartAndWait(8080)
}
