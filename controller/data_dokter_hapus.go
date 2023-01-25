package controller

import (
	"go_smk/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Data_dokter_hapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dokter model.Dokter
	db.Where("id_user = ?", c.Param("id")).Find(&dokter)
	if dokter.Id_user == 0 {
		c.JSON(400, gin.H{
			"code":    400,
			"data":    "-",
			"message": "Parameter id yang anda masukan salah.",
		})
		return
	}
	db.Where("id_user = ?", c.Param("id")).Delete(&dokter)
	c.JSON(200, gin.H{
		"code":    200,
		"data":    dokter,
		"message": "Data dokter berhasil dihapus.",
	})
}
