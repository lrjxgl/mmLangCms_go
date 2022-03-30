package indexIndex

import (
	"app/access"
	"app/config"
	"app/config/cache"
	"app/index/indexModel"
	"math/rand"
	"strconv"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*@@LoginIndex@@*/
func LoginIndex(c echo.Context) (err error) {
	return config.Success(c, 0, "success")
}

/*@@LoginSave@@*/
func LoginSave(c echo.Context) (err error) {
	telephone := c.FormValue("telephone")
	password := c.FormValue("password")
	user := indexModel.UserModel{}
	db := config.Db
	res := db.Where("telephone=?", telephone).First(&user)
	if res.Error != nil {
		return config.Success(c, 1, "用户不存在")
	}
	puser := indexModel.UserPasswordModel{}
	db.Where("userid=?", user.Userid).First(&puser)
	if puser.Password != config.Umd5(password+puser.Salt) {
		fmt.Print(user.Userid)
		return config.Success(c, 1, "密码错误"+password+puser.Salt)
	}
	reData := access.UserSetToken(user.Userid, puser.Password)
	reData["error"] = 0
	reData["message"] = "登录成功"
	
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 

}

/*@@LoginFindpwdsave@@*/
func LoginFindpwdsave(c echo.Context) (err error) {
	telephone := c.FormValue("telephone")
	if telephone == "" {
		return config.Success(c, 1, "手机号码错误")
	}
	yzm := c.FormValue("yzm")
	cacheKey := "smsLogin" + telephone + yzm
	res := cache.CacheGet(cacheKey)
	if res == "" {
		return config.Success(c, 1, "验证码出错")
	}
	db := config.Db
	user := indexModel.UserModel{}
	e1 := db.Where("telephone=?", telephone).First(&user)
	if e1.Error != nil {
		return config.Success(c, 1, "账户不存在")
	}
	userid := user.Userid
	password := c.FormValue("password")

	salt := strconv.Itoa(rand.Intn(9000) + 999)
	password = config.Umd5(password + salt)
	puser := indexModel.UserPasswordModel{
		Userid:   userid,
		Salt:     salt,
		Password: password,
	}
	db.Where("userid=?", userid).Updates(&puser)
	return config.Success(c, 0, "密码修改成功")
}

/*@@LoginSendsms@@*/
func LoginSendsms(c echo.Context) (err error) {
	telephone := c.FormValue("telephone")
	if telephone == "" {
		return config.Success(c, 1, "手机号码错误")
	}
	yzm := strconv.Itoa(rand.Intn(9000) + 999)
	var ops = make(map[string]interface{})
	ops["telephone"] = telephone
	ops["content"] = "您的验证码：" + yzm
	ops["yzm"] = yzm
	cacheKey := "smsLogin" + telephone + yzm
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
