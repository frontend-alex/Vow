package security

type JWTManager struct {
	secret string
}

func NewJWTManager(secret string) JWTManager {
	return JWTManager{secret: secret}
}
