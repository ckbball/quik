package users

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

func UpdateContextUserModel(c *gin.Context, my_user_id int) {
  var user UserModel
  if my_user_id != 0 {
    db := common.GetDB()
    db.First(&user, my_user_id)
  }
  c.Set("my_user_id", my_user_id)
  c.Set("my_user_model", user)
}
