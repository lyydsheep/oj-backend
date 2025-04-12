package util

import (
	"back/logic/domain"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
)

func Convert(from, to any) error {
	return copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			JudgeCaseToString(),
			StringToJudgeCase(),
			ConfigToString(),
			StringToConfig(),
			AnswerToString(),
			StringToAnswer(),
			StringToTags(),
			TagsToString(),
		},
	})
}

func StringToTags() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: []domain.Tag{},
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.(string)
			if !ok {
				return nil, errors.New("src type is not string")
			}
			return domain.ConvertTag(val), nil
		},
	}
}

func TagsToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: []domain.Tag{},
		DstType: copier.String,
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.([]domain.Tag)
			if !ok {
				return nil, errors.New("src type is not []domain.Tag")
			}
			marshalData, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			return string(marshalData), nil
		},
	}
}

func JudgeCaseToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: []domain.JudgeCase{},
		DstType: copier.String,
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.([]domain.JudgeCase)
			if !ok {
				return nil, errors.New("src type is not domain.JudgeCase")
			}
			marshalData, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			return string(marshalData), nil
		},
	}
}

func StringToJudgeCase() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: []domain.JudgeCase{},
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.(string)
			if !ok {
				return nil, errors.New("src type is not string")
			}
			return domain.ConvertJudgeCases(val), nil
		},
	}
}

func ConfigToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: domain.JudgeConfig{},
		DstType: copier.String,
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.(domain.JudgeConfig)
			if !ok {
				return nil, errors.New("src type is not domain.JudgeConfig")
			}
			return val.String(), nil
		},
	}
}

func StringToConfig() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: domain.JudgeConfig{},
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.(string)
			if !ok {
				return nil, errors.New("src type is not string")
			}
			return domain.ConvertJudgeConfig(val), nil
		},
	}
}

func AnswerToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: domain.Answer{},
		DstType: copier.String,
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.(domain.Answer)
			if !ok {
				return nil, errors.New("src type is not domain.Answer")
			}
			return val.String(), nil
		},
	}
}

func StringToAnswer() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: copier.String,
		DstType: domain.Answer{},
		Fn: func(src interface{}) (dst interface{}, err error) {
			val, ok := src.(string)
			if !ok {
				return nil, errors.New("src type is not string")
			}
			return domain.ConvertAnswer(val), nil
		},
	}
}
