package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/samarec1812/sync-unisender/auth"
	"github.com/samarec1812/sync-unisender/auth/configs"
	"github.com/samarec1812/sync-unisender/auth/pkg/handler"
	"github.com/samarec1812/sync-unisender/auth/pkg/repository"
	"github.com/samarec1812/sync-unisender/auth/pkg/service"
	pb "github.com/samarec1812/sync-unisender/proto/auth"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetOutput(os.Stdout)

	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewMySqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	list, err := net.Listen("tcp", viper.GetString("app.port"))

	if err != nil {
		logrus.Fatalf("failed to connect tcp: %s", err.Error())
	}

	server := grpc.NewServer()
	auth.NewAuthServerGrpc(server, handlers)

	fmt.Println("ZDEC")
	go func() {
		if err := server.Serve(list); err != nil {
			logrus.Fatalf("failed serve: %s", err.Error())
		}
	}()

	conn, err := grpc.Dial(viper.GetString("app.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()
	// httputil.NewSingleHostReverseProxy()
	mux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterAuthorizationAmoHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = pb.RegisterAuthorizationUnisenderHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())
	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	//
	//mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err = pb.RegisterAuthorizationHandlerFromEndpoint(ctx, mux, viper.GetString("app.port"), opts)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("server listening at 8081")
	//if err := http.ListenAndServe(":8081", mux); err != nil {
	//	panic(err)
	//}
}
