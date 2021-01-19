package model

import (
	"github.com/krilie/lico_alone/common/com-model"
)

type One struct {
	com_model.Model
}

func (One) TableName() string {
	return "tb_one"
}

type Two struct {
	com_model.Model
}

func (Two) TableName() string {
	return "tb_two"
}
