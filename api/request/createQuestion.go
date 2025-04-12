package request

import "back/logic/domain"

type CreateQuestion struct {
	Title     *string            `gorm:"column:title;type:varchar(512);comment:标题" json:"title"`              // 标题
	Content   *string            `gorm:"column:content;type:text;comment:内容" json:"content"`                  // 内容
	Tags      []domain.Tag       `gorm:"column:tags;type:varchar(1024);comment:标签列表（json 数组）" json:"tags"`    // 标签列表（json 数组）
	JudgeCase []domain.JudgeCase `gorm:"column:judge_case;type:text;comment:判题用例（json 数组）" json:"judge_case"` // 判题用例（json 数组）
	// TODO 配置如果为空，使用默认配置
	JudgeConfig domain.JudgeConfig `gorm:"column:judge_config;type:text;comment:判题配置（json 对象）" json:"judge_config"` // 判题配置（json 对象）
	UserID      string             `gorm:"column:user_id;type:varchar(32);not null;index:idx_userId,priority:1" json:"user_id"`
}
