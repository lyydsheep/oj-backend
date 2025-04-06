package test

import (
	"back/dal/dao"
	"back/dal/model"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateQuestion(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(47.120.11.159:3306)/oj?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	query := dao.Use(db)
	title, content := "test", "testContent"
	if err = query.Question.WithContext(context.Background()).Create(&model.Question{
		Title:   &title,
		Content: &content,
		UserID:  "faiz",
	}); err != nil {
		t.Error(err)
	}
}
