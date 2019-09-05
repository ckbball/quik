package common

import (
  //"fmt"
  //"math/rand"
  "time"

  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
  //"gopkg.in/go-playground/validator.v8"
)

const JWTSecretString = "hehehoohoo9wf"

func GenToken(id int) string {
  jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))

  jwt_token.Claims = jwt.MapClaims{
    "id":  id,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }

  token, _ := jwt_token.SignedString([]byte(JWTSecretString))
  return token
}

func Bind(c *gin.Context, obj interface{}) error {
  b := binding.Default(c.Request.Method, c.ContentType())
  return c.ShouldBindWith(obj, b)
}

type CommonError struct {
  Errors map[string]interface{} `json:"errors"`
}

/*
func NewValidatorError(err error) CommonError {
  res := CommonError{}
  res.Errors = make(map[string]interface{})
  errs := err.(validator.ValidationErrors)
  for _, v := range errs {
    // can translate each error one at a time.
    //fmt.Println("gg",v.NameNamespace)
    if v.Param != "" {
      res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
    } else {
      res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
    }

  }
  return res
}*/

func NewError(key string, err error) CommonError {
  res := CommonError{}
  res.Errors = make(map[string]interface{})
  res.Errors[key] = err.Error()
  return res
}
