package tokenService

type ApiToken struct {
	token string
}

var Token ApiToken = ApiToken{
	token: "placeholder",
}

func (apiToken *ApiToken) SetToken(newToken string) {
	apiToken.token = newToken

}

func SetToken(newToken string) {
	Token.token = newToken
}

func GetToken() string {
	return Token.token
}
