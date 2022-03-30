package indexAdmin

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
func OpenToutiaoNull(c echo.Context) (err error){
	 
	now := time.Now()
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
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

/*@@OpenToutiaoIndex@@*/
func OpenToutiaoIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	var db = config.Db
	var list = []indexModel.OpenToutiaoModel{}
	 
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
	db.Model(&indexModel.OpenToutiaoModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.OpenToutiaoModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = indexModel.OpenToutiaoList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@OpenToutiaoAdd@@*/
func OpenToutiaoAdd(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	id, err := strconv.Atoi(c.QueryParam("id"))
	var db = config.Db

	var data = indexModel.OpenToutiaoModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
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

/*@@OpenToutiaoSave@@*/
func OpenToutiaoSave(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	id, err := strconv.Atoi(c.FormValue("id"))
	var db = config.Db
	var data = indexModel.OpenToutiaoModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
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
		db.Model(indexModel.OpenToutiaoModel{}).Where("id=?", id).Updates(postData)
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
/*@@OpenToutiaoStatus@@*/
func OpenToutiaoStatus(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.OpenToutiaoModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(indexModel.OpenToutiaoModel{}).Where("id=?", id).Update("status", status)
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

/*@@OpenToutiaoDelete@@*/
func OpenToutiaoDelete(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.OpenToutiaoModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	db.Model(indexModel.OpenToutiaoModel{}).Where("id=?", id).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}