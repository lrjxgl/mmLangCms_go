package indexIndex

import (
	"app/config"
	"app/access"
	"app/index/indexModel"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
/*解决import未使用*/
func UserAuthNewNull(c echo.Context) (err error){
	 
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

/*@@UserAuthNewIndex@@*/
func UserAuthNewIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []indexModel.UserAuthNewModel{}
	 
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
	db.Model(&indexModel.UserAuthNewModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.UserAuthNewModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = indexModel.UserAuthNewList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserAuthNewList@@*/
func UserAuthNewList(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []indexModel.UserAuthNewModel{}
	 
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
	db.Model(&indexModel.UserAuthNewModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.UserAuthNewModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = indexModel.UserAuthNewList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserAuthNewShow@@*/
func UserAuthNewShow(c echo.Context) (err error) {
	
	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.UserAuthNewModel)
	res := db.Where("id=?  AND status=1  ", id).First(&data)
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

/*@@UserAuthNewMy@@*/
func UserAuthNewMy(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []indexModel.UserAuthNewModel{}
	 
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
	db.Model(&indexModel.UserAuthNewModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.UserAuthNewModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = indexModel.UserAuthNewList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserAuthNewAdd@@*/
func UserAuthNewAdd(c echo.Context) (err error) {
	
			userid := access.UserCheckAccess(c)
			if userid == 0 {
				return config.Success(c, 1000, "请先登录")
			}
		
	id, err := strconv.Atoi(c.QueryParam("id"))
	var db = config.Db

	var data = indexModel.UserAuthNewModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
		
			if(data.Userid!=userid){
				return config.Success(c,0,"暂无权限");
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
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserAuthNewSave@@*/
func UserAuthNewSave(c echo.Context) (err error) {
	
			userid := access.UserCheckAccess(c)
			if userid == 0 {
				return config.Success(c, 1000, "请先登录")
			}
		
	id, err := strconv.Atoi(c.FormValue("id"))
	var db = config.Db
	var data = indexModel.UserAuthNewModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
		
			if(data.Userid!=userid){
				return config.Success(c,0,"暂无权限");
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
		db.Model(indexModel.UserAuthNewModel{}).Where("id=?", id).Updates(postData)
	}

	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["data"] = postData
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}
/*@@UserAuthNewStatus@@*/
func UserAuthNewStatus(c echo.Context) (err error) {
	
			userid := access.UserCheckAccess(c)
			if userid == 0 {
				return config.Success(c, 1000, "请先登录")
			}
		
	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.UserAuthNewModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
			if(data.Userid!=userid){
				return config.Success(c,0,"暂无权限");
			}
		

	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(indexModel.UserAuthNewModel{}).Where("id=?", id).Update("status", status)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["status"] = status
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 


}

/*@@UserAuthNewDelete@@*/
func UserAuthNewDelete(c echo.Context) (err error) {
	
			userid := access.UserCheckAccess(c)
			if userid == 0 {
				return config.Success(c, 1000, "请先登录")
			}
		
	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.UserAuthNewModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
			if(data.Userid!=userid){
				return config.Success(c,0,"暂无权限");
			}
		

	db.Model(indexModel.UserAuthNewModel{}).Where("id=?", id).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}