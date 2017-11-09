package types

import "gopkg.in/mgo.v2/bson"

// WebhookInfo - Physical format of stored webhook info
type WebhookInfo struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	URL            string        `bson:"webhookURL"`
	BaseCurrency   string        `bson:"baseCurrency"`
	TargetCurrency string        `bson:"targetCurrency"`
	MinValue       float64       `bson:"minTriggerValue"`
	MaxValue       float64       `bson:"maxTriggerValue"`
}

// WebhookInv - Payload for invoking the webhook
type WebhookInv struct {
	BaseCurrency   string  `json:"baseCurrency"`
	TargetCurrency string  `json:"targetCurrency"`
	CurrentRate    float64 `json:"currentRate"`
	MinValue       float64 `json:"minTriggerValue"`
	MaxValue       float64 `json:"maxTriggerValue"`
}

// WebhookDisp - Display format for webhooks
type WebhookDisp struct {
	URL            string  `bson:"webhookURL"`
	BaseCurrency   string  `bson:"baseCurrency"`
	TargetCurrency string  `bson:"targetCurrency"`
	MinValue       float64 `bson:"minTriggerValue"`
	MaxValue       float64 `bson:"maxTriggerValue"`
}

// CurrencyInfo - Physical format of stored response from api.fixer.io/latest
type CurrencyInfo struct {
	ID           bson.ObjectId      `bson:"_id,omitempty"`
	BaseCurrency string             `bson:"baseCurrency"`
	Date         string             `bson:"date"` // Date format yyyy-mm-dd
	Rates        map[string]float64 `bson:"rates"`
}

// CurrencyReq - Used variously, requesting exchange, running average, etc.
type CurrencyReq struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}
