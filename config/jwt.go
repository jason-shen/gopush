package config

type JwtConfig struct {
	Secret []byte
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret: []byte(getIni("jwt_secret", "JWT_SECRET", "awesome")),
	}
}
