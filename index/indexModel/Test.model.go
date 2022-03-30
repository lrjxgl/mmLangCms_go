package indexModel

import (
	"app/config"
)

type TestModel struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}

func (TestModel) TableName() string {
	return config.TablePre + "test"
}
