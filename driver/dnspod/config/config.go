package config

type Config struct {
	TokenId      string
	Token        string
	Lang         string
	ErrorOnEmpty bool
	// Email 参见 [关于UserAgent] https://docs.dnspod.cn/api/5f55993d8ae73e11c5b01ce6/
	Email string
}

func NewConfigDefault() Config {
	return Config{
		TokenId:      "",
		Token:        "",
		Lang:         "cn",
		ErrorOnEmpty: false,
	}
}

func NewConfig(id string, token string) Config {
	d := NewConfigDefault()
	d.TokenId = id
	d.Token = token
	return d
}
