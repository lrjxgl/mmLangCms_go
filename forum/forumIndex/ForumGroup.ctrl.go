package forumIndex

import (
	"app/config"
	"app/access"
	"app/forum/forumModel"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
/*解决import未使用*/
func ForumGroupNull(c echo.Context) (err error){
	 
	now := time.Now()
	
			userid := access.UserCheckAccess(c)
			if userid == 0 {
				return config.Success(c, 1000, "请先登录")
			}
		
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["now"]=now;
	 
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@ForumGroupIndex@@*/
func ForumGroupIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []forumModel.ForumGroupModel{}
	 
	where:=" status in(0,1,2) ";
	//统计数量
	start, err := strconv.Atoi(c.QueryParam("per_page"))
	if err!=nil {
		start=0;
	}
	limit, err2 := strconv.Atoi(c.QueryParam("limit"))
	if err2!=nil || limit==0 {
		limit=24;
	}
	var rscount int64;
	db.Model(&forumModel.ForumGroupModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []forumModel.ForumGroupModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = forumModel.ForumGroupList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@ForumGroupList@@*/
func ForumGroupList(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []forumModel.ForumGroupModel{}
	 
	where:=" status in(0,1,2) ";
	//统计数量
	start, err := strconv.Atoi(c.QueryParam("per_page"))
	if err!=nil {
		start=0;
	}
	limit, err2 := strconv.Atoi(c.QueryParam("limit"))
	if err2!=nil || limit==0 {
		limit=24;
	}
	var rscount int64;
	db.Model(&forumModel.ForumGroupModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []forumModel.ForumGroupModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = forumModel.ForumGroupList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@ForumGroupShow@@*/
func ForumGroupShow(c echo.Context) (err error) {
	
	gid := c.QueryParam("gid")
	var db = config.Db
	data := new(forumModel.ForumGroupModel)
	res := db.Where("gid=?  AND status=1  ", gid).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["data"] = data
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}