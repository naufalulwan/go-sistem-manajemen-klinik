package controller

import (
	"go_smk/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Rekam_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var rekam model.Rekam_medis
	db.Where("id = ?", c.Param("id")).Find(&rekam)
	if rekam.Id == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Where("id = ?", c.Param("id")).Delete(&rekam)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    rekam,
		"message": "Data user berhasil dihapus.",
	})
}
