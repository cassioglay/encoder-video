package main

import (
	"log"
	"os"
	"strconv"

	"github.com/cassioglay/encoder/aplication/services"
	"github.com/cassioglay/encoder/framework/database"
	"github.com/cassioglay/encoder/framework/queue"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var db database.Database

func init() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	autoMigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))

	if err != nil {
		log.Fatalf("Error parsing boolean anv variable")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		log.Fatalf("Error parsing boolean anv variable")
	}

	db.AutoMigrateDb = autoMigrateDb
	db.Debug = debug
	db.DsnTest = os.Getenv("DSN_TEST")
	db.Dsn = os.Getenv("DSN")
	db.DbTypeTest = os.Getenv("DB_TYPE_TEST")
	db.DbType = os.Getenv("DB_TYPE")
	db.Env = os.Getenv("ENV")

}

func main() {
	messageChannel := make(chan amqp.Delivery)
	jobReturnbChannel := make(chan services.JobWorkerResult)

	dbConnection, err := db.Connect()

	if err != nil {
		log.Fatalf("error connecting to database")
	}

	defer dbConnection.Close()

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	jobManager := services.NewJobManager(dbConnection, rabbitMQ, jobReturnbChannel, messageChannel)
	jobManager.Start(ch)

}
