package reply

import "back/logic/domain"

type DemoOrder struct {
	OrderId   string `json:"orderId"`
	UserId    string `json:"userId"`
	BillMoney int64  `json:"billMoney"`
	PaidAt    string `json:"paidAt"`
	CreatedAt string `json:"createAt"`
	UpdatedAt string `json:"updateAt"`
}

type Question struct {
	ID          int64              `gorm:"column:question_id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Title       *string            `gorm:"column:title;type:varchar(512);comment:标题" json:"title"`                  // 标题
	Content     *string            `gorm:"column:content;type:text;comment:内容" json:"content"`                      // 内容
	Tags        []domain.Tag       `gorm:"column:tags;type:varchar(1024);comment:标签列表（json 数组）" json:"tags"`        // 标签列表（json 数组）
	SubmitNum   int32              `gorm:"column:submit_num;type:int;not null;comment:题目提交数" json:"submit_num"`     // 题目提交数
	AcceptedNum int32              `gorm:"column:accepted_num;type:int;not null;comment:题目通过数" json:"accepted_num"` // 题目通过数
	JudgeCase   []domain.JudgeCase `gorm:"column:judge_case;type:text;comment:判题用例（json 数组）" json:"judge_case"`     // 判题用例（json 数组）
	JudgeConfig domain.JudgeConfig `gorm:"column:judge_config;type:text;comment:判题配置（json 对象）" json:"judge_config"` // 判题配置（json 对象）
}
