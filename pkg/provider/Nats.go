package provider

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

type NatsMessingProvider struct {
	ES *nats.EncodedConn
}
func NewNatsMessingProvider(es *nats.EncodedConn) *NatsMessingProvider {
	return &NatsMessingProvider{
		ES: es,
	}
}

func (provider *NatsMessingProvider) ReceivingMessage(ctx context.Context,subject string,fn interface{}) error{
	// Subscribe
	 sub, er := provider.ES.Subscribe(subject,fn)
	 if er != nil {
	 	log.Println(er)
	 }else {
	    log.Println("Successfully Subscribed to subject :",sub.Subject)
	 }
	return nil
}

func (provider *NatsMessingProvider) SendingMessage(ctx context.Context,subject string, v interface{}) error {


	//defer ec.Close()
//	cr := event.(*model.CreatedEvent)
	// Publish the message
	if err := provider.ES.Publish(subject, &v); err != nil {
		log.Fatal(err)
	}else {
		log.Println("Successfully published to subject :",subject)
	}

   return nil
}