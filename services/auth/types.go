package auth

import "strings"

type GrantType string

const (
	ClientCredentials GrantType = "client_credentials"
)

type Scope string

const (
	ChargeWrite          Scope = "cob.write"
	ChargeRead           Scope = "cob.read"
	PixWrite             Scope = "pix.write"
	PixRead              Scope = "pix.read"
	PixSend              Scope = "pix.send"
	WebhookWrite         Scope = "webhook.write"
	WebhookRead          Scope = "webhook.read"
	PayloadLocationWrite Scope = "payloadlocation.write"
	PayloadLocationRead  Scope = "payloadlocation.read"
)

type AuthorizationRequest struct {
	GrantType GrantType `json:"grant_type" url:"grant_type"`
	Scope     string    `json:"scope" url:"scope"`
}

func NewAuthorizationRequest(grantType GrantType, scopes ...Scope) AuthorizationRequest {
	var finalScopes []string
	for _, scope := range scopes {
		finalScopes = append(finalScopes, string(scope))
	}
	return AuthorizationRequest{
		GrantType: grantType,
		Scope:     strings.Join(finalScopes, ","),
	}
}
