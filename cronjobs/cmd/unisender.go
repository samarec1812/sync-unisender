package main

import (
	"encoding/json"
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
	"github.com/samarec1812/sync-unisender/unisender/models"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetOutput(os.Stdout)

	//if err := configs.InitConfig(); err != nil {
	//	logrus.Fatalf("error initializing configs: %s", err.Error())
	//}

	//db, err := repository.NewMySqlDB(repository.Config{
	//	Host:     "localhost",
	//	Port:     "3307",
	//	Username: "root",
	//	Password: "qwerty",
	//	DBName:   "amo_auth",
	//})
	//
	//if err != nil {
	//	logrus.Fatalf("failed to initialize db: %s", err.Error())
	//}
	//
	//repos := repository.NewRepository(db)
	//services := service.NewService(repos)

	consumer, err := beanstalk.Dial("tcp", "localhost:11300")
	if err != nil {
		logrus.Fatalf("failed to connect beanstalk broker: %s", err.Error())
	}
	id, body, err := consumer.Reserve(5 * time.Second)
	if err != nil {
		logrus.Fatalf("failed to reserved contact with beanstalk broker: %s", err.Error())
	}

	var data models.BeanstalkData
	err = json.Unmarshal(body, &data)
	if err != nil {
		logrus.Fatalf("failed to unmarshaling contact structure")
	}
	fmt.Println(data.Id)
	//val, err := services.Contacts.Export(data.Id)
	//if err != nil || val.Email == "" {
	//	logrus.Println("error get contact from DB")
	//	return
	//}

	fmt.Println(id, " : ", data.Id)
}
