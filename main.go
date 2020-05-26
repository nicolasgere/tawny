package main

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"tawny/core"
	"tawny/tawny"
)

type GrpcWebMiddleware struct {
	*grpcweb.WrappedGrpcServer
}

func (m *GrpcWebMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if true {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewGrpcWebMiddleware(grpcWeb *grpcweb.WrappedGrpcServer) *GrpcWebMiddleware {
	return &GrpcWebMiddleware{grpcWeb}
}

func main() {
	core.Init()
	grpcServer := grpc.NewServer()
	tawny.RegisterPushServiceServer(grpcServer, &core.PushServer{})
	tawny.RegisterPresenceServiceServer(grpcServer, &core.PresenceServer{})
	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))
	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Recoverer,
		NewGrpcWebMiddleware(wrappedGrpc).Handler,
	)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	if err := http.ListenAndServe(":8900", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}

}
