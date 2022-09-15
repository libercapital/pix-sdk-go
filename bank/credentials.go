package bank

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func NewCredentials(clientId, clientSecret string) Credentials {
	return Credentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}
