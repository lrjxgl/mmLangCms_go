package forumIndex

import (
	"app/access"
	"app/config"
	"app/forum/forumModel"
	"app/index/indexModel"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

/*解决import未使用*/
func ForumNull(c echo.Context) (err error) {

	now := time.Now()
	flashList := indexModel.ListByno("uniapp-forum-index", 4)
	userid := access.UserCheckAccess(c)
	if userid == 0 {
		return config.Success(c, 1000, "请先登录")
	}

	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["now"] = now
	reData["flashList"] = flashList

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumIndex@@*/
func ForumIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")

	flashList := indexModel.ListByno("uniapp-forum-index", 4)
	//fmt.Print(flashList)
	adList := indexModel.ListByno("uniapp-forum-ad", 3)
	navList := indexModel.ListByno("uniapp-forum-nav", 1000)
	recList := forumModel.GetForumByKey("index")

	reData := make(map[string]interface{})
	reData["flashList"] = flashList
	reData["navList"] = navList
	reData["adList"] = adList
	reData["recList"] = recList
	reData["error"] = 0
	reData["message"] = "success"

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumList@@*/
func ForumList(c echo.Context) (err error) {
	gid := c.QueryParam("gid")
	catid := c.QueryParam("catid")
	var db = config.Db
	group := new(forumModel.ForumGroupModel)
	db.Where("gid=?", gid).First(&group)
	var list = []forumModel.ForumModel{}
	catList := new([]forumModel.ForumCategoryModel)
	db.Where("gid=? ", gid).Find(&catList)
	where := " status in(0,1) "
	if catid != "0" {
		where += " AND catid=" + catid
	}
	if gid != "0" {
		where += " AND gid=" + gid
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
	reData["group"] = group
	reData["catList"] = catList
	reData["list"] = forumModel.ForumList(list)
	reData["rscount"] = rscount
	reData["per_page"] = per_page

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumNew@@*/
func ForumNew(c echo.Context) (err error) {
	fmt.Print("forumIndex")

	var db = config.Db
	var list = []forumModel.ForumModel{}

	where := " status in(0,1,2) "
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

/*@@ForumShow@@*/
func ForumShow(c echo.Context) (err error) {

	id := c.QueryParam("id")
	var db = config.Db
	data := new(forumModel.ForumModel)
	res := db.Where("id=?  AND status=1  ", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	author := new(indexModel.UserModel)
	db.Where("userid=? ", data.Userid).First(&author)

	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["data"] = data
	reData["author"] = author
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumMy@@*/
func ForumMy(c echo.Context) (err error) {
	fmt.Print("forumIndex")

	var db = config.Db
	var list = []forumModel.ForumModel{}

	where := " status in(0,1,2) "
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
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"] = rscount
	reData["per_page"] = per_page

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumAdd@@*/
func ForumAdd(c echo.Context) (err error) {

	userid := access.UserCheckAccess(c)
	if userid == 0 {
		return config.Success(c, 1000, "请先登录")
	}

	id, err := strconv.Atoi(c.QueryParam("id"))
	var db = config.Db

	var data = forumModel.ForumModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}

		if data.Userid != userid {
			return config.Success(c, 0, "暂无权限")
		}

	}

	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["data"] = data
	reData["id"] = id

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumSave@@*/
func ForumSave(c echo.Context) (err error) {

	userid := access.UserCheckAccess(c)
	if userid == 0 {
		return config.Success(c, 1000, "请先登录")
	}

	id, err := strconv.Atoi(c.FormValue("id"))
	var db = config.Db
	var data = forumModel.ForumModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}

		if data.Userid != userid {
			return config.Success(c, 0, "暂无权限")
		}

	}
	//新增数据

	postData := map[string]interface{}{}
	postData["title"] = c.FormValue("title")
	postData["description"] = c.FormValue("description")
	now := time.Now()
	postData["createtime"] = now.Format("2006-01-02 15:04:05")
	if id != 0 {
		db.Create(postData)
	} else {
		db.Model(forumModel.ForumModel{}).Where("id=?", id).Updates(postData)
	}

	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["data"] = postData

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumStatus@@*/
func ForumStatus(c echo.Context) (err error) {

	userid := access.UserCheckAccess(c)
	if userid == 0 {
		return config.Success(c, 1000, "请先登录")
	}

	id := c.QueryParam("id")
	var db = config.Db
	data := new(forumModel.ForumModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}

	if data.Userid != userid {
		return config.Success(c, 0, "暂无权限")
	}

	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(forumModel.ForumModel{}).Where("id=?", id).Update("status", status)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["status"] = status

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumDelete@@*/
func ForumDelete(c echo.Context) (err error) {

	userid := access.UserCheckAccess(c)
	if userid == 0 {
		return config.Success(c, 1000, "请先登录")
	}

	id := c.QueryParam("id")
	var db = config.Db
	data := new(forumModel.ForumModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}

	if data.Userid != userid {
		return config.Success(c, 0, "暂无权限")
	}

	db.Model(forumModel.ForumModel{}).Where("id=?", id).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}

/*@@ForumUser@@*/
func ForumUser(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	user := indexModel.UserGet(ssuserid, "userid,nickname,user_head,gold,grade")

	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["user"] = user

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}
