package main

import (
  "log"

  "github.com/streadway/amqp"
)

func handleError (err error, msg string)  {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

//the logic for sending and receiving message
func main() {
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
  handleError(err, "Dialing failed to RabbitMQ broker")
  defer conn.Close()

  channel, err := conn.Channel()
  handleError(err, "Fetching channel failed")
  defer channel.Close()
}

testQueue, err := channel.QueueDeclare
(
  "test", // Name of the queue
  false, // Message is persisted or not
  false, // Delete message when unused
  false, // Exclusive
  false, // No waiting time
  nil,
)

handleError (err, "Queue creation failed")

serverTime := time.Now()
message := amqp.Publishing{
  ContentType: 'text/plain',
  Body:        []byte(serverTime.String()),
}

err = channel.Publish(
  "",             // exchange
  testQueue.Name, // routing key (Queue)
  false,          // mandatory
  false,          // immediate
  message,
)

handleError(err, "Failed to publish a message")
log.Println("Successfully published a message to the queue")
