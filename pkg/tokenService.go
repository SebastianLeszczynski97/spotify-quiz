package obsoleteTokenService

import (
	"fmt"
	"log"
	"net/http"
)

type ApiToken struct {
	token string
}

var Token ApiToken = ApiToken{
	token: "placeholder",
}

func (apiToken *ApiToken) SetToken(newToken string) {
	fmt.Printf("New token set")
	apiToken.token = newToken

}

func SetToken(newToken string) {
	Token.token = newToken
}

func GetToken() string {
	return Token.token
}

func SetTokenHandler(w http.ResponseWriter, r *http.Request) {
	var newTokenValue = r.PostFormValue("input-token")
	SetToken(newTokenValue)
	log.Printf(fmt.Sprintf("Token provided %s", GetToken()), GetToken())
}
