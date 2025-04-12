package domain

import "encoding/json"

type Question struct {
	ID          int64       `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Title       *string     `gorm:"column:title;type:varchar(512);comment:标题" json:"title"`                  // 标题
	Content     *string     `gorm:"column:content;type:text;comment:内容" json:"content"`                      // 内容
	Tags        []Tag       `gorm:"column:tags;type:varchar(1024);comment:标签列表（json 数组）" json:"tags"`        // 标签列表（json 数组）
	Answer      []Answer    `gorm:"column:answer;type:text;comment:题解" json:"answer"`                        // 题解
	SubmitNum   int32       `gorm:"column:submit_num;type:int;not null;comment:题目提交数" json:"submit_num"`     // 题目提交数
	AcceptedNum int32       `gorm:"column:accepted_num;type:int;not null;comment:题目通过数" json:"accepted_num"` // 题目通过数
	JudgeCase   []JudgeCase `gorm:"column:judge_case;type:text;comment:判题用例（json 数组）" json:"judge_case"`     // 判题用例（json 数组）
	JudgeConfig JudgeConfig `gorm:"column:judge_config;type:text;comment:判题配置（json 对象）" json:"judge_config"` // 判题配置（json 对象）
	ThumbNum    int32       `gorm:"column:thumb_num;type:int;not null;comment:点赞数" json:"thumb_num"`         // 点赞数
	FavourNum   int32       `gorm:"column:favour_num;type:int;not null;comment:收藏数" json:"favour_num"`       // 收藏数
	UserID      string      `gorm:"column:user_id;type:varchar(32);not null;index:idx_userId,priority:1" json:"user_id"`
}

type Answer struct {
	Content string `json:"content"`
}

func (a *Answer) String() string {
	jsonStr, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func ConvertAnswer(answer string) Answer {
	var a Answer
	err := json.Unmarshal([]byte(answer), &a)
	if err != nil {
		panic(err)
	}
	return a
}

type Tag struct {
	Tag string `json:"tag"`
}

func (t *Tag) String() string {
	jsonStr, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func ConvertTag(tag string) []Tag {
	var t []Tag
	err := json.Unmarshal([]byte(tag), &t)
	if err != nil {
		panic(err)
	}
	return t
}

type JudgeCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

func (j *JudgeCase) String() string {
	jsonStr, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func ConvertJudgeCases(caseStr string) []JudgeCase {
	var jc []JudgeCase
	err := json.Unmarshal([]byte(caseStr), &jc)
	if err != nil {
		panic(err)
	}
	return jc
}

type JudgeConfig struct {
	TimeLimit   int32 `json:"time_limit"`
	MemoryLimit int32 `json:"memory_limit"`
}

func (j *JudgeConfig) String() string {
	jsonStr, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return string(jsonStr)
}

func ConvertJudgeConfig(config string) JudgeConfig {
	var judgeConfig JudgeConfig
	err := json.Unmarshal([]byte(config), &judgeConfig)
	if err != nil {
		panic(err)
	}
	return judgeConfig
}
