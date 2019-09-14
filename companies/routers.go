package companies

import (
  "errors"
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "net/http"
  //"strconv"
)

// No-Auth Routes
func CompaniesRegister(router *gin.RouterGroup) {
  router.POST("/", CompaniesRegistration)
  router.POST("/login", CompaniesLogin)
  /*
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

func CompaniesLogin(c *gin.Context) {
  login := NewLoginValidator()
  if err := login.Bind(c); err != nil {
    c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Validator error")))
    return
  }

  fmt.Println("Checking login validator binding: --> ", login)

  company, err := FindOneCompany(&CompanyModel{Email: login.companyModel.Email}) //models.go function

  // sending error with token and pass
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("DB: Email not registered or invalid password")))
    return
  }

  if company.checkPassword(login.Company.Pass) != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("Check: Email not registered  or invalid password")))
    return
  }

  fmt.Println("Checking company model found in routers- line 70 - : ", company)
  fmt.Println()
  UpdateContextCompanyModel(c, company.ID)
  fmt.Println("Checking UpdateContextCompanyModel ", c.MustGet("my_company_model").(CompanyModel))
  serializer := CompanySerializer{c}
  c.JSON(http.StatusOK, gin.H{"company": serializer.Response()})
}
