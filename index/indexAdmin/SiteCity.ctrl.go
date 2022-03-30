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
func SiteCityNull(c echo.Context) (err error){
	 
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

/*@@SiteCityIndex@@*/
func SiteCityIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	var db = config.Db
	var list = []indexModel.SiteCityModel{}
	 
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
	db.Model(&indexModel.SiteCityModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.SiteCityModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["list"] = indexModel.SiteCityList(list)
	reData["type"] = reflect.TypeOf(list)
	reData["rscount"]=rscount;
	reData["per_page"]=per_page;
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@SiteCityAdd@@*/
func SiteCityAdd(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	sc_id, err := strconv.Atoi(c.QueryParam("sc_id"))
	var db = config.Db

	var data = indexModel.SiteCityModel{}
	if sc_id != 0 {
		res := db.Where("sc_id=?  AND status<4  ", sc_id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
		
	}

	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["data"] = data
	reData["sc_id"] = sc_id
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@SiteCitySave@@*/
func SiteCitySave(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	sc_id, err := strconv.Atoi(c.FormValue("sc_id"))
	var db = config.Db
	var data = indexModel.SiteCityModel{}
	if sc_id != 0 {
		res := db.Where("sc_id=?  AND status<4  ", sc_id).First(&data)
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
	if sc_id != 0 {
		db.Create(postData)
	} else {
		db.Model(indexModel.SiteCityModel{}).Where("sc_id=?", sc_id).Updates(postData)
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
/*@@SiteCityStatus@@*/
func SiteCityStatus(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	sc_id := c.QueryParam("sc_id")
	var db = config.Db
	data := new(indexModel.SiteCityModel)
	res := db.Where("sc_id=?", sc_id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(indexModel.SiteCityModel{}).Where("sc_id=?", sc_id).Update("status", status)
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

/*@@SiteCityDelete@@*/
func SiteCityDelete(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	sc_id := c.QueryParam("sc_id")
	var db = config.Db
	data := new(indexModel.SiteCityModel)
	res := db.Where("sc_id=?", sc_id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	db.Model(indexModel.SiteCityModel{}).Where("sc_id=?", sc_id).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}