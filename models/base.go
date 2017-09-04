package models

type BaseInterface interface {
	TableName() string
	Validate() error
	Save() (int64, error)
}

type BaseModel struct {

}