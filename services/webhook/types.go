package webhook

import "time"

type Webhook struct {
	URL       string    `json:"webhookUrl"`
	Key       string    `json:"chave"`
	CreatedAt time.Time `json:"criacao"`
}

type CreateWebhook struct {
	URL string `json:"webhookUrl"`
}
