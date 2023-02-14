// Unisender API
//
// This is unisender API for import contacts for Unisender service. You can find out more about the API.
//
// Schemes: http
// Host: host:8080
// BasePath: /unisender
// Version: 1.0.0
// Contact: Alexey Kondratev <samarec1812@gmail.com> https://github.com/samarec1812
//
//
// Consumes:
// - application/x-www-form-urlencoded
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
	"os"

	pb "github.com/samarec1812/sync-unisender/proto/auth"
	"github.com/samarec1812/sync-unisender/unisender"
	"github.com/samarec1812/sync-unisender/unisender/configs"
	"github.com/samarec1812/sync-unisender/unisender/pkg/handler"
	"github.com/samarec1812/sync-unisender/unisender/pkg/repository"
	"github.com/samarec1812/sync-unisender/unisender/pkg/service"
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

	clientGRPCConn := fmt.Sprintf("%s:%s", viper.GetString("grpc.host"), viper.GetString("grpc.port"))
	cc, err := grpc.Dial(clientGRPCConn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("failed to connect server gRPC: %s", err.Error())
	}

	defer cc.Close()

	clientRPC := pb.NewAuthorizationUnisenderClient(cc)

	beanstalkConn := fmt.Sprintf("%s:%s", viper.GetString("broker.host"), viper.GetString("broker.port"))
	producer, err := beanstalk.Dial("tcp", beanstalkConn)
	if err != nil {
		logrus.Fatalf("failed to connect beanstalk broker: %s", err.Error())
	}
	defer producer.Close()

	handlers := handler.NewHandler(services, clientRPC, producer)

	logrus.Println("Serving HTTP on connection")

	srv := new(unisender.Server)
	if err := srv.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
