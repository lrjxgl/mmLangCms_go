package forumIndex

import (
	"app/config"
	"app/forum/forumModel"
	"app/index/indexModel"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/*@@ForumHomeIndex@@*/
func ForumHomeIndex(c echo.Context) (err error) {
	uid, _ := strconv.Atoi(c.QueryParam("userid"))
	userid := uint(uid)
	db := config.Db

	user := indexModel.UserGet(userid, "userid,nickname,user_head,follow_num,followed_num,description")
	where := " status userid =?"
	//统计数量
	start, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		start = 0
	}
	limit, err2 := strconv.Atoi(c.QueryParam("limit"))
	if err2 != nil || limit == 0 {
		limit = 24
	}
	var rscount int64
	db.Model(&forumModel.ForumModel{}).Where(where, userid).Count(&rscount)
	list := []forumModel.ForumModel{}
	res := db.Where(where, userid).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []forumModel.ForumModel{}
	}
	var per_page int64 = int64(start + limit)
	if per_page > rscount {
		per_page = 0
	}
	reData := make(map[string]interface{})
	reData["user"] = user
	reData["list"] = list
	reData["per_page"] = per_page
	reData["rscount"] = rscount
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)
}
