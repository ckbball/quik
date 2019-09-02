package main

import (
  "fmt"
  "github.com/ckbball/quik/auth"
  "github.com/ckbball/quik/common"
  "github.com/ckbball/quik/companies"
  "github.com/ckbball/quik/jobs"
  "github.com/ckbball/quik/users"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

func main() {

  _, err := common.Init()
  if err != nil {
    fmt.Println(err)
  }
  defer common.Close()

  r := gin.Default()

  v1 := r.Group("/api")

  // no authentication routes
  {
    // /login
    // /register
    auth.Register(v1.Group("/auth"))
  }

  // basic authentication routes
  {
    basicAuth := r.Group('/')
    basicAuth.Use(AuthenticationRequired())
    {
      jobs.Register(basicAuth.Group("/jobs"))
      companies.Register(basicAuth.Group("/companies"))
      users.Register(basicAuth.Group("/users"))
    }
  }

  r.Run()

}
