package auth

// Credentials 认证结构体
type Credentials struct {
	ClientKey    string
	ClientSecret string
}

// New 新的认证
func New(clientKey, clientSecret string) *Credentials {
	return &Credentials{
		clientKey,
		clientSecret,
	}
}
