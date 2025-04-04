package domain

import "time"

const (
	LANGUAGE_GO   = "Go"
	LANGUAGE_JAVA = "Java"
)

const (
	RESULT_ACCEPTED = "AC"
	RESULT_FAILED   = "WA"
	RESULT_ERROR    = "RE"
)

const (
	STATUS_PENDING = 0
	STATUS_RUNNING = 1
	STATUS_END     = 2
)

type QuestionSubmit struct {
	ID           int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:id" json:"id"`          // id
	CodeLanguage string `gorm:"column:code_language;type:varchar(128);not null;comment:编程语言" json:"code_language"` // 编程语言
	Code         string `gorm:"column:code;type:text;not null;comment:用户代码" json:"code"`                           // 用户代码
	JudgeInfo    `json:"judge_info"`
	//JudgeInfo    *string   `gorm:"column:judge_info;type:text;comment:判题信息（json 对象）" json:"judge_info"`                                      // 判题信息（json 对象）
	Status     int32     `gorm:"column:status;type:int;not null;comment:判题状态（0 - 待判题、1 - 判题中、2 - 判题结束）" json:"status"`                     // 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`      // 创建时间
	QuestionID int64     `gorm:"column:question_id;type:bigint;not null;index:idx_questionId,priority:1;comment:题目 id" json:"question_id"` // 题目 id
	UserID     string    `gorm:"column:user_id;type:varchar(32);not null;index:idx_userId,priority:1" json:"user_id"`
}

type JudgeInfo struct {
	Result   string        `json:"result"`
	CostTime time.Duration `json:"cost_time"`
	CostMem  float64       `json:"cost_mem"`
}
