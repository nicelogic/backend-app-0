package graph_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func TestAddContactsApply(t *testing.T) {

	data := "aaa"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

}

func TestPubSub(t *testing.T){
	client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL:               "pulsar+ssl://broker.pulsar.env0.luojm.com:9443",
        OperationTimeout:  30 * time.Second,
        ConnectionTimeout: 30 * time.Second,
    })
    if err != nil {
        log.Fatalf("Could not instantiate Pulsar client: %v", err)
    }
    defer client.Close()
    log.Printf("pulsar connect success\n")

    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: "tenant-0/contacts/add_contacts_apply",
    })
    if err != nil {
        log.Fatal(err)
    }
    _, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
        Payload: []byte("hello"),
    })
    defer producer.Close()
    if err != nil {
        fmt.Println("Failed to publish message", err)
    }
    fmt.Println("Published message")


    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "tenant-0/contacts/add_contacts_apply",
        SubscriptionName: "my-sub",
        Type:             pulsar.Shared,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer consumer.Close()
    
    for {
        msg, err := consumer.Receive(context.Background())
        if err != nil {
            log.Fatal(err)
        }
    
        fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
            msg.ID(), string(msg.Payload()))
    
        consumer.Ack(msg)
    }
    
    // if err := consumer.Unsubscribe(); err != nil {
    //     log.Fatal(err)
    // }
}
