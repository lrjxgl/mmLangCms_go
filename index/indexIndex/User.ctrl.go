package indexIndex

import (
	"app/access"
	"app/config"
	"app/config/cache"
	"app/index/indexModel"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

/*解决import未使用*/
func UserNull(c echo.Context) (err error) {

	now := time.Now()

	userid := access.UserCheckAccess(c)
	if userid == 0 {
		return config.Success(c, 1000, "请先登录")
	}

	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["now"] = now

	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserIndex@@*/
func UserIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")

	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"

	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserSet@@*/
func UserSet(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	user := indexModel.UserGet(ssuserid, "userid,nickname,user_head,gold,grade")

	//输出浏览器
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["user"] = user
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserInfo@@*/
func UserInfo(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	user := indexModel.UserGet(ssuserid, "userid,nickname,telephone,user_head,description,gold,grade")
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	reData["user"] = user
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserSave@@*/
func UserSave(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	db := config.Db
	postData := map[string]interface{}{}
	postData["nickname"] = c.FormValue("nickname")
	postData["description"] = c.FormValue("description")
	now := time.Now()
	postData["updatetime"] = now.Format("2006-01-02 15:04:05")
	db.Model(indexModel.UserModel{}).Where("userid=?", ssuserid).Updates(postData)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserTelephonesave@@*/
func UserTelephonesave(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	user := indexModel.UserGet(ssuserid, "userid,telephone")
	telephone := user.Telephone
	if telephone != "" {
		return config.Success(c, 1, "手机已经绑定了")
	}
	telephone = c.FormValue("telephone")
	yzm := c.FormValue("yzm")
	cacheKey := "smsUser" + telephone + yzm
	res := cache.CacheGet(cacheKey)
	if res == "" {
		return config.Success(c, 1, "验证码出错"+res)
	}
	db := config.Db
	postData := map[string]interface{}{}
	postData["telephone"] = c.FormValue("telephone")

	now := time.Now()
	postData["updatetime"] = now.Format("2006-01-02 15:04:05")
	db.Model(indexModel.UserModel{}).Where("userid=?", ssuserid).Updates(postData)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserPasswordsave@@*/
func UserPasswordsave(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	db := config.Db
	oldpassword := c.FormValue("oldpassword")
	password := c.FormValue("password")
	password2 := c.FormValue("password2")
	if password != password2 {
		return config.Success(c, 1, "两次密码输入不一致")
	}
	puser := indexModel.UserPasswordModel{}
	db.Where("userid=?", ssuserid).First(&puser)
	if puser.Password != config.Umd5(oldpassword+puser.Salt) {
		return config.Success(c, 1, "密码错误")
	}
	salt := strconv.Itoa(rand.Intn(9000) + 999)
	postData := map[string]interface{}{}
	postData["salt"] = salt
	postData["password"] = config.Umd5(password + salt)
	db.Model(indexModel.UserPasswordModel{}).Where("userid=?", ssuserid).Updates(postData)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserPaypwdsave@@*/
func UserPaypwdsave(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	user := indexModel.UserGet(ssuserid, "userid,telephone")
	telephone := user.Telephone

	yzm := c.FormValue("yzm")
	cacheKey := "smsUser" + telephone + yzm
	res := cache.CacheGet(cacheKey)
	if res == "" {
		return config.Success(c, 1, "验证码出错")
	}
	db := config.Db

	paypwd := c.FormValue("password")

	postData := map[string]interface{}{}

	postData["password"] = config.Umd5(paypwd)
	db.Model(indexModel.UserPasswordModel{}).Where("userid=?", ssuserid).Updates(postData)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@UserSendsms@@*/
func UserSendsms(c echo.Context) (err error) {
	ssuserid := access.UserCheckAccess(c)
	if ssuserid == 0 {
		return config.Success(c, 1000, "请先登录")
	}
	user := indexModel.UserGet(ssuserid, "userid,telephone")
	telephone := user.Telephone
	if user.Telephone == "" {
		telephone = c.FormValue("telephone")
	}
	if telephone == "" {
		return config.Success(c, 1, "手机号码错误")
	}
	yzm := strconv.Itoa(rand.Intn(9000) + 999)
	var ops = make(map[string]interface{})
	ops["telephone"] = telephone
	ops["content"] = "您的验证码：" + yzm
	ops["yzm"] = yzm
	cacheKey := "smsUser" + telephone + yzm
	cache.CacheSet(cacheKey, "1", 300)
	if config.SmsTest {
		return config.Success(c, 0, "验证码:"+yzm)
	}

	config.SendSms(ops)
	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}
