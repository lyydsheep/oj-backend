package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)
import "gorm.io/gen"

func Test(t *testing.T) {

	g := gen.NewGenerator(gen.Config{
		OutPath:          "../dal/model",
		Mode:             gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldWithTypeTag: true,
	})

	gormdb, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:33307)/back?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
