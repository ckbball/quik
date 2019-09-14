package companies

import (
  "github.com/ckbball/quik/common"
  //"github.com/dgrijalva/jwt-go"
  //"github.com/dgrijalva/jwt-go/request"
  "github.com/gin-gonic/gin"
  //"net/http"
  //"strings"
)

func UpdateContextCompanyModel(c *gin.Context, my_company_id int) {
  var company CompanyModel
  if my_company_id != 0 {
    db := common.GetDB()
    db.First(&CompanyModel{}, my_company_id)
  }
  c.Set("my_company_id", my_company_id)
  c.Set("my_company_model", company)
}
