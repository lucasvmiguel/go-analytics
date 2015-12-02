package model

import "time"

const (
	HIGH   = 3
	MEDIUM = 2
	LOW    = 1
)

const (
	ERROR   = 4
	WARNING = 3
	INFO    = 2
	LOG     = 1
)

type Notification struct {
	Title             string    `json:"title"`
	Tag1              string    `json:"tag1"`
	Tag2              string    `json:"tag2"`
	Tag3              string    `json:"tag3"`
	Info              string    `json:"info"`
	Relevance         uint8     `json:"relevance"`
	Type              uint8     `json:"type"`
	IsTransaction     bool      `json:"isTransaction"`
	Transaction       string    `json:"transaction"`
	TransactionResult bool      `json:"transactionResult"`
	Company           string    `json:"company"`
	Time              time.Time `json:"time"`
}

func (n *Notification) ToMapString() map[string]interface{} {
	return map[string]interface{}{
		"title":             n.Title,
		"tag1":              n.Tag1,
		"tag2":              n.Tag2,
		"tag3":              n.Tag3,
		"info":              n.Info,
		"relevance":         n.Relevance,
		"type":              n.Type,
		"isTransaction":     n.IsTransaction,
		"transaction":       n.Transaction,
		"transactionResult": n.TransactionResult,
		"company":           n.Company,
		"time":              n.Time,
	}
}
