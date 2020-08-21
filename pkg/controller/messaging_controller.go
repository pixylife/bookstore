package controller

import (
	"context"
)

type Provider interface {
	ReceivingMessage(ctx context.Context,subject string,fn interface{}) error
	SendingMessage(ctx context.Context, subject string, msg interface{}) error
}

type MessageCtrl struct {
	provider Provider
}

func (c MessageCtrl) ReceivingMessage(ctx context.Context,subject string, fn interface{}) error {
	return c.provider.ReceivingMessage(ctx,subject,fn)
}

func (c MessageCtrl) SendingMessage(ctx context.Context,subject string,msg interface{}) error {
	 return c.provider.SendingMessage(ctx,subject, msg)
}

func NewMessageController(p Provider) *MessageCtrl  {
	return &MessageCtrl{provider:p}
}


