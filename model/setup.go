package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "123"
	dbname   = "manajemen_klinik"
	port     = 5432
)

func SetupModels() *gorm.DB {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	db.AutoMigrate(&Pasien{}, &Dokter{}, &Rekam_medis{}, &Perawat{}, &User{}, &Rawat_jalan{})
	return db
}
