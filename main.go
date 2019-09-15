package main

import (
  // "fmt"
  //"github.com/ckbball/quik/auth"
  "github.com/ckbball/quik/common"
  "github.com/ckbball/quik/companies"
  "github.com/ckbball/quik/jobs"
  "github.com/ckbball/quik/users"
  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
  users.AutoMigrate()
  companies.AutoMigrate()
  //db.AutoMigrate(&companies.CompanyModel{})
  //db.AutoMigrate(&jobs.JobModel{})

}

func main() {

  db := common.Init()
  Migrate(db)
  defer db.Close()

  r := gin.Default()

  v1 := r.Group("/api")

  // no authentication routes
  {
    // /login
    // /register
    users.UsersRegister(v1.Group("/auth"))
    // login and register
    companies.CompaniesRegister(v1.Group("/companies"))
  }

  // basic authentication routes

  {
    companyAuth := r.Group("/api")
    companyAuth.Use(companies.AuthMiddleware(true))
    {
      // create, update, delete jobs
      jobs.Register(companyAuth.Group("/jobs"))
    }
  }

  {
    basicAuth := r.Group("/api")
    basicAuth.Use(users.AuthMiddleware(true))
    {
      users.Register(basicAuth.Group("/users"))
      // user can create, update, delete an application to a job
      //applications.Register(basicAuth.Group("/applications"))
    }
  }

  r.Run()

}
