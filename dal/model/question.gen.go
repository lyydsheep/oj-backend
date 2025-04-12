// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameQuestion = "question"

// Question 题目
type Question struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(512);comment:标题" json:"title"`                  // 标题
	Content     string    `gorm:"column:content;type:text;comment:内容" json:"content"`                      // 内容
	Tags        string    `gorm:"column:tags;type:varchar(1024);comment:标签列表（json 数组）" json:"tags"`        // 标签列表（json 数组）
	Answer      string    `gorm:"column:answer;type:text;comment:题解" json:"answer"`                        // 题解
	SubmitNum   int32     `gorm:"column:submit_num;type:int;not null;comment:题目提交数" json:"submit_num"`     // 题目提交数
	AcceptedNum int32     `gorm:"column:accepted_num;type:int;not null;comment:题目通过数" json:"accepted_num"` // 题目通过数
	JudgeCase   string    `gorm:"column:judge_case;type:text;comment:判题用例（json 数组）" json:"judge_case"`     // 判题用例（json 数组）
	JudgeConfig string    `gorm:"column:judge_config;type:text;comment:判题配置（json 对象）" json:"judge_config"` // 判题配置（json 对象）
	ThumbNum    int32     `gorm:"column:thumb_num;type:int;not null;comment:点赞数" json:"thumb_num"`         // 点赞数
	FavourNum   int32     `gorm:"column:favour_num;type:int;not null;comment:收藏数" json:"favour_num"`       // 收藏数
	UserID      string    `gorm:"column:user_id;type:varchar(32);not null;index:idx_userId,priority:1" json:"user_id"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	IsDelete    int32     `gorm:"column:is_delete;type:int;not null;comment:是否删除" json:"is_delete"`                                    // 是否删除
}

// TableName Question's table name
func (*Question) TableName() string {
	return TableNameQuestion
}
