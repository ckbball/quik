package companies

import (
  "errors"
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
)

// No-Auth Routes
func CompaniesRegister(router *gin.RouterGroup) {
  router.POST("/", CompaniesRegistration)
  /*
    router.POST("/login", CompaniesLogin)
    router.GET("/:id", CompaniesGet) // idk if this should be un-auth'd or third partied or first only
  */
}

// Auth Routes
/*
func Register(router *gin.RouterGroup) {
  router.POST("", CompanyUpdate) // update company info
  // maybe the routes to look at candidates for jobs and such should be here or under jobs or applications

}
*/

func CompaniesRegistration(c *gin.Context) {
  company := NewCompanyModelValidator()
  if err := company.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }

  fmt.Println("check if validator validated company: ", company.companyModel)

  if err := SaveOne(&company.companyModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  c.Set("my_company_model", company.companyModel)
  serializer := CompanySerializer{c}
  c.JSON(http.StatusCreated, gin.H{"company": serializer.Response()})
}
