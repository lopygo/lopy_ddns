package config

type Config struct {
	TokenId string
	Token   string
	// Email 参见 [关于UserAgent] https://docs.dnspod.cn/api/5f55993d8ae73e11c5b01ce6/
	Email        string
	Lang         string
	ErrorOnEmpty bool
}

func NewConfigDefault() Config {
	return Config{
		TokenId:      "",
		Token:        "",
		Lang:         "cn",
		ErrorOnEmpty: false,
	}
}

func NewConfig(tokenID string, token string) Config {
	d := NewConfigDefault()
	d.TokenId = tokenID
	d.Token = token
	return d
}
