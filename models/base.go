package models

type BaseInterface interface {
	TableName() string
	Validate() error
	Save() (int64, error)
	GetModelById() error
}

type BaseModel struct {

}