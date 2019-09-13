package companies

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type CompanySerializer struct {
  c *gin.Context
}

type CompanyResponse struct {
  Name    string `json:"name"`
  Mission string `json:"mission"`
  Size    int    `json:"size"`
  Token   string `json:"token"`
  // need to add job here later or maybe a companyprofileresponse
}

func (self *CompanySerializer) Response() CompanyResponse {
  company := self.c.MustGet("my_company_model").(CompanyModel)
  out := CompanyResponse{
    Name:    company.Name,
    Mission: company.Mission,
    Size:    company.Size,
    Token:   common.GenToken(company.ID),
  }
  return out
}
