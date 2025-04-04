package util

import (
	"back/logic/domain"
	"errors"
	"github.com/jinzhu/copier"
)

func Convert(to, from any) error {
	return copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: domain.JudgeCase{},
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(domain.JudgeCase)
					if !ok {
						return nil, errors.New("src type is not domain.JudgeCase")
					}
					return val.String(), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: domain.JudgeCase{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(string)
					if !ok {
						return nil, errors.New("src type is not string")
					}
					return domain.ConvertJudgeCase(val), nil
				},
			},
			{
				SrcType: domain.JudgeConfig{},
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(domain.JudgeConfig)
					if !ok {
						return nil, errors.New("src type is not domain.JudgeConfig")
					}
					return val.String(), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: domain.JudgeConfig{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(string)
					if !ok {
						return nil, errors.New("src type is not string")
					}
					return domain.ConvertJudgeConfig(val), nil
				},
			},
			{
				SrcType: domain.Answer{},
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(domain.Answer)
					if !ok {
						return nil, errors.New("src type is not domain.Answer")
					}
					return val.String(), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: domain.Answer{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					val, ok := src.(string)
					if !ok {
						return nil, errors.New("src type is not string")
					}
					return domain.ConvertAnswer(val), nil
				},
			},
		},
	})
}
