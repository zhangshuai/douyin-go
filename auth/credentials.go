package auth

type Credentials struct {
	ClientKey    string
	ClientSecret string
}

func New(clientKey, clientSecret string) *Credentials {
	return &Credentials{
		clientKey,
		clientSecret,
	}
}
