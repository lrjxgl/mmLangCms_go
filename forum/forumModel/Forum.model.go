package forumModel

import (
	"app/config"
	"app/index/indexModel"
	"encoding/json"
)

type ForumModel struct {
	Id          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `json:"title"`
	Userid      uint    `json:"userid"`
	Gid         uint    `json:"gid"`
	Catid       uint    `json:"catid"`
	Love_num    uint    `json:"love_num"`
	Fav_num     uint    `json:"fav_num"`
	Forward_num uint    `json:"forward_num"`
	Keywords    string  `json:"keywords"`
	Description string  `json:"description"`
	Status      uint    `json:"status"`
	Comment_num uint    `json:"comment_num"`
	Imgurl      string  `json:"imgurl"`
	Grade       uint    `json:"grade"`
	Isrecommend uint    `json:"isrecommend"`
	View_num    uint    `json:"view_num"`
	Isnew       uint    `json:"isnew"`
	Tags        string  `json:"tags"`
	Videourl    string  `json:"videourl"`
	Money       float64 `json:"money"`
	Gold        uint    `json:"gold"`
	Createtime  string  `json:"createtime"`
	Updatetime  string  `json:"updatetime"`
	Imgsdata    string  `json:"imgsdata"`
}

type ForumModels struct {
	ForumModel
	User indexModel.UserModel `json:"user";gorm:"foreignKey:Userid"`
}

func (ForumModel) TableName() string {
	return config.TablePre + "mod_forum"
}

func ForumList(list []ForumModel) []map[string]interface{} {

	slen := len(list)
	if slen == 0 {
		return make([]map[string]interface{}, 0)
	}
	var uids = make([]uint, slen)
	for i := 0; i < slen; i++ {
		uids[i] = list[i].Userid
	}
	us := indexModel.GetUserByIds(uids, "userid,nickname,user_head")
	var nlist = make([]map[string]interface{}, slen)

	for i := 0; i < slen; i++ {
		n := list[i]
		aj, _ := json.Marshal(n)
		var m map[string]interface{}
		_ = json.Unmarshal(aj, &m)
		m["imgurl"] = config.Image_site(m["imgurl"].(string))
		m["user"] = indexModel.OneUserList(uids[i], us)
		nlist[i] = m
	}
	return nlist
}

func GetForumByIds(ids []uint) []map[string]interface{} {
	db := config.Db
	list := []ForumModel{}
	db.Where("id in ?", ids).Find(&list)

	dList := ForumList(list)
	return dList
}
func OneForumList(id uint, list []ForumModel) ForumModel {
	slen := len(list)
	for i := 0; i < slen; i++ {
		if id == list[i].Id {
			return list[i]
		}
	}
	return ForumModel{}
}
