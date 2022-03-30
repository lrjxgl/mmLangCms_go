package forumModel

import (
	"app/config"
	"fmt"
)

type ForumTagsModel struct {
	Tagid     uint   `gorm:"primaryKey" json:"tagid"`
	Title     string `json:"title"`
	Status    uint   `json:"status"`
	Total_num uint   `json:"total_num"`
	View_num  uint   `json:"view_num"`
	Gkey      string `json:"gkey"`
	Gnum      uint   `json:"gnum"`
}

func (ForumTagsModel) TableName() string {
	return config.TablePre + "mod_forum_tags"
}

func ForumTagsList(list []ForumTagsModel) []ForumTagsModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]

		list[i] = m
	}
	return list
}

func GetForumByKey(gkey string) []map[string]interface{} {
	tag := ForumTagsModel{}
	var db = config.Db
	res := db.Where("gkey=?", gkey).First(&tag)
	if res.Error != nil {
		return make([]map[string]interface{}, 0)
	}
	var ids []uint
	res2 := db.Model(ForumTagsIndexModel{}).Limit(int(tag.Gnum)).Where("tagid=?", tag.Tagid).Pluck("objectid", &ids)
	fmt.Print(ids)
	if res2.Error != nil {
		return make([]map[string]interface{}, 0)
	}
	list := GetForumByIds(ids)

	return list
}
