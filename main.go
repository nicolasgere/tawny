package main

import (
	"fmt"
	"github.com/caddyserver/certmagic"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"net/http"
	badgerstorage "oya.to/certmagic-badgerstorage"
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
	core.InitConfig()
	core.Init()
	grpcServer := grpc.NewServer()
	tawny.RegisterAdminServiceServer(grpcServer, &core.AdminServer{})
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

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString(core.GRPC_PORT)))
		if err != nil {
			grpclog.Fatalf("failed starting grpc server: %v", err)
		}
		grpcServer.Serve(lis)
	}()

	if viper.GetBool(core.HTTPS_ENABLE) {
		domain := viper.GetString(core.HTTPS_DOMAIN)
		if domain == "" {
			grpclog.Fatalf("failed to start, if HTTPS_ENABLE then HTTPS_DOMAIN is needed")
		}
		email := viper.GetString(core.HTTPS_EMAIL)
		if email == "" {
			grpclog.Fatalf("failed to start, if HTTPS_ENABLE then HTTPS_EMAIL is needed")
		}
		grpclog.Infof("HTTPS configuration %s %s", domain, email)
		certmagic.DefaultACME.Agreed = true
		certmagic.DefaultACME.Email = email
		certmagic.HTTPPort = viper.GetInt(core.HTTP_PORT)
		certmagic.HTTPSPort = viper.GetInt(core.HTTPS_PORT)
		certmagic.Default.Storage = badgerstorage.New(core.Db)
		if err := certmagic.HTTPS([]string{domain}, router); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		if err := http.ListenAndServe(":"+viper.GetString(core.HTTP_PORT), router); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	}

}
