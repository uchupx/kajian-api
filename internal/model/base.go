package model

type ModelInterface interface {
	TableName() string
}

type BaseModel struct{}

func (BaseModel) TableName() string {
	return ""
}
