package forumIndex

import (
	"app/config"
	"app/forum/forumModel"

	"fmt"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

/*@@ForumSearchIndex@@*/
func ForumSearchIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	keyword := c.QueryParam("keyword")
	var db = config.Db
	var list = []forumModel.ForumModel{}

	where := " status in(0,1,2) "
	if keyword != "" {
		where += " AND title like '%" + keyword + "%' "
	}
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
	db.Model(&forumModel.ForumModel{}).Where(where).Count(&rscount)
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []forumModel.ForumModel{}
	}
	//输出浏览器
	var per_page int64 = int64(start + limit)
	if per_page > rscount {
		per_page = 0
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = forumModel.ForumList(list)
	reData["rscount"] = rscount
	reData["per_page"] = per_page

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}
