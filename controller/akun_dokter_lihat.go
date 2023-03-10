package controller

import (
	"go_smk/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Lihat_akun_dokter(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mDokter []model.Dokter
	db.Raw("select * from dokters order by id_user desc").Scan(&mDokter)
	c.JSON(200, gin.H{
		"code": 200,
		"data": mDokter,
	})
}
