package controller

import (
	"go_smk/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var login login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Request harus dalam bentuk JSON.",
		})
		return
	}
	if (login.Email == "") || (login.Password == "") {
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Email dan password tidak boleh kosong.",
			})
			return
		}
	}
	var user model.User
	db.Where("email = ?", login.Email).Where("password = ?", login.Password).Find(&user)
	if (login.Email == user.Email) && (login.Password == user.Password) {
		c.JSON(200, gin.H{
			"code":  200,
			"id":    user.Id,
			"level": user.Level,
		})
	} else {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Email atau Password salah",
		})
	}
}
