package services

import (
	"sync"
	"google.golang.org/grpc"
	"github.com/while-loop/remember-me"
	"log"
)

type ConstrucFunc func(app *remme.App) Service
type Service interface {
	Register(rpc *grpc.Server)
}

var services = struct {
	sync.Mutex
	s map[string]ConstrucFunc
}{
	s: map[string]ConstrucFunc{},
}

func Register(name string, construc ConstrucFunc) {
	services.Lock()
	defer services.Unlock()

	services.s[name] = construc
}

func StartServices(app *remme.App, rpc *grpc.Server) {
	services.Lock()
	defer services.Unlock()

	for name, construct := range services.s {
		log.Printf("Starting %s...", name)
		construct(app).Register(rpc)
	}
}
