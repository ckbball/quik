package companies

import (
  //"fmt"
  "github.com/ckbball/quik/common"
  "github.com/dgrijalva/jwt-go"
  "github.com/dgrijalva/jwt-go/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "strings"
)

func UpdateContextCompanyModel(c *gin.Context, my_company_id int) {
  var company CompanyModel
  if my_company_id != 0 {
    db := common.GetDB()
    db.First(&company, my_company_id)
  }
  c.Set("my_company_id", my_company_id)
  c.Set("my_company_model", company)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
  return func(c *gin.Context) {
    UpdateContextCompanyModel(c, 0)
    token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
      b := ([]byte(common.JWTSecretString))
      return b, nil
    })
    if err != nil {
      if auto401 {
        c.AbortWithError(http.StatusUnauthorized, err)
      }
      return
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      my_user_id := int(claims["id"].(float64))
      UpdateContextCompanyModel(c, my_user_id)
    }
  }
}

func stripBearerPrefixFromTokenString(tok string) (string, error) {
  if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
    return tok[6:], nil
  }
  return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
  request.HeaderExtractor{"Authorization"},
  stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
  AuthorizationHeaderExtractor,
  request.ArgumentExtractor{"access_token"},
}
