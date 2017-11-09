package types

import "gopkg.in/mgo.v2/bson"

// WebhookInfo - Physical format of stored webhook info
type WebhookInfo struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	URL            string        `json:"webhookURL"      bson:"webhookURL"`
	BaseCurrency   string        `json:"baseCurrency"    bson:"baseCurrency"`
	TargetCurrency string        `json:"targetCurrency"  bson:"targetCurrency"`
	MinValue       float64       `json:"minTriggerValue" bson:"minTriggerValue"`
	MaxValue       float64       `json:"maxTriggerValue" bson:"maxTriggerValue"`
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
	URL            string  `json:"webhookURL"      bson:"webhookURL"`
	BaseCurrency   string  `json:"baseCurrency"    bson:"baseCurrency"`
	TargetCurrency string  `json:"targetCurrency"  bson:"targetCurrency"`
	MinValue       float64 `json:"minTriggerValue" bson:"minTriggerValue"`
	MaxValue       float64 `json:"maxTriggerValue" bson:"maxTriggerValue"`
}

// CurrencyInfo - Physical format of stored response from api.fixer.io/latest
type CurrencyInfo struct {
	ID           bson.ObjectId      `             bson:"_id,omitempty"`
	BaseCurrency string             `json:"base"  bson:"baseCurrency"`
	Date         string             `json:"date"  bson:"date"` // Date format yyyy-mm-dd
	Rates        map[string]float64 `json:"rates" bson:"rates"`
}

// CurrencyReq - Used variously, requesting exchange, running average, etc.
type CurrencyReq struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}
