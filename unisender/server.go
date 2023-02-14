package unisender

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Server struct {
	httpServer *fasthttp.Server
}

func (s *Server) Run(port string, handler fasthttp.RequestHandler) error {
	s.httpServer = &fasthttp.Server{
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//tun, err := ngrok.Listen(context.Background(), config.HTTPEndpoint(), ngrok.WithAuthtokenFromEnv())
	//if err != nil {
	//	return err
	//}

	return s.httpServer.ListenAndServe(port)
}

func (s *Server) Shutdown() error {
	return s.httpServer.Shutdown()
}
