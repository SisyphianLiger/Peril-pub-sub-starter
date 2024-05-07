package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

const conn string = "amqp://guest:guest@localhost:5672/"

func main() { 

	fmt.Printf("Starting Peril server...\n")

    connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")

    defer connection.Close()
    fmt.Printf("Peril game server connected to RabbitMQ!\n Connection Port: %s \n", conn)
    
    // wait for ctrl+c
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt)
    <-signalChan

    fmt.Printf("\nEnding Connection to RabbitMQ\n")
}


/*
    Registers Errors in case connection cannot be made
*/
func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}
