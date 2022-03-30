package indexModel

import (
	"app/config"
)

type AdModel struct {
	Id         uint    `gorm:"primaryKey" json:"id"`
	Tag_id     uint    `json:"tag_id"`
	Tag_id_2nd uint    `json:"tag_id_2nd"`
	Title      string  `json:"title"`
	Status     uint    `json:"status"`
	Info       string  `json:"info"`
	Link1      string  `json:"link1"`
	Link2      string  `json:"link2"`
	Starttime  string  `json:"starttime"`
	Endtime    string  `json:"endtime"`
	Imgurl     string  `json:"imgurl"`
	Imgurl2    string  `json:"imgurl2"`
	Orderindex uint    `json:"orderindex"`
	Price      float64 `json:"price"`
	Object_id  uint    `json:"object_id"`
	Createtime string  `json:"createtime"`
}

func (AdModel) TableName() string {
	return config.TablePre + "ad"
}

func AdList(list []AdModel) []AdModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		m.Imgurl = config.Image_site(m.Imgurl)
		list[i] = m
	}
	return list
}

func ListByno(tagno string, limit int) []AdModel {
	var db = config.Db
	tags := AdTagsModel{}
	res := db.Where("tagno=? ", tagno).First(&tags)

	if res.Error != nil {
		return []AdModel{}
	}
	list := []AdModel{}
	res2 := db.Where(" tag_id_2nd=? ", tags.Tag_id).Find(&list)
	if res2.Error != nil {
		return []AdModel{}
	}
	list = AdList(list)
	return list
}
