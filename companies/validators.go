package companies

import (
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type CompanyModelValidator struct {
  Company struct {
    Name    string `json:"name" form:"name" binding:"exists,alphanum"`
    Size    int    `json:"size" form:"size" binding:"exists"`
    Mission string `json:"mission" form:"mission" binding:"exists,alphanum"` //maybe give this a max
    Pass    string `json:"password" form:"password" binding:"exists,min=8,max255"`
    Email   string `json:"email" form:"email" binding:"exists,email"`
  } `json:"company"`
  companyModel CompanyModel `json:"-"`
}

func (self *CompanyModelValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)

  fmt.Println("/companies/validators-23: Check if common.Bind() binded the body to Company properly: ", self.Company)
  fmt.Println()
  if err != nil {
    return err
  }

  self.companyModel.Name = self.Company.Name
  self.companyModel.Size = self.Company.Size
  self.companyModel.Mission = self.Company.Mission
  self.companyModel.Email = self.Company.Email

  if self.Company.Pass != common.JWTSecretString {
    self.companyModel.setPassword(self.Company.Pass)
  }
  return nil
}

func NewCompanyModelValidator() CompanyModelValidator {
  validator := CompanyModelValidator{}
  return validator
}

func NewCompanyModelValidatorFillWith(company CompanyModel) CompanyModelValidator {
  out := NewCompanyModelValidator()
  out.Company.Name = company.Name
  out.Company.Size = company.Size
  out.Company.Mission = company.Mission
  out.Company.Pass = company.Pass
  out.Company.Email = company.Email

  return out
}

type LoginValidator struct {
  Company struct {
    Email string `json:"email" form:"email" binding:"exists,email`
    Pass  string `json:"pass" form:"password" binding:"exists,min=8,max=255`
  } `json:"company"`
  companyModel CompanyModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)
  if err != nil {
    return err
  }

  self.companyModel.Email = self.Company.Email
  return nil
}

func NewLoginValidator() LoginValidator {
  loginValidator := LoginValidator{}
  return loginValidator
}
