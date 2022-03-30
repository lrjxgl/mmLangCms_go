package indexIndex

import (
	"app/access"
	"app/config"
	"app/config/cache"
	"app/index/indexModel"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

/*@@RegisterIndex@@*/
func RegisterIndex(c echo.Context) (err error) {
	return config.Success(c, 0, "success")
}

/*@@RegisterSave@@*/
func RegisterSave(c echo.Context) (err error) {
	telephone := c.FormValue("telephone")
	nickname := c.FormValue("nickname")
	password := c.FormValue("password")
	yzm := c.FormValue("yzm")
	cacheKey := "smsReg" + telephone + yzm
	res := cache.CacheGet(cacheKey)
	if res == "" {
		return config.Success(c, 1, "验证码出错")
	}
	user := indexModel.UserModel{}
	db := config.Db
	res2 := db.Where("telephone=?", telephone).First(&user)
	if res2.Error == nil && user.Telephone != "" {
		return config.Success(c, 1, "手机号码已经存在")
	}
	now := time.Now()
	createtime := now.Format("2006-01-02 15:04:05")
	updatetime := createtime
	nUser := indexModel.UserModel{
		Telephone:  telephone,
		Nickname:   nickname,
		Createtime: createtime,
		Updatetime: updatetime,
	}
	var userid uint = 0

	db.Transaction(func(tx *gorm.DB) error {
		db.Create(&nUser)
		userid = nUser.Userid
		salt := strconv.Itoa(rand.Intn(9000) + 999)
		password = config.Umd5(password + salt)
		puser := indexModel.UserPasswordModel{
			Userid:   userid,
			Salt:     salt,
			Password: password,
		}
		db.Create(&puser)

		return nil
	})
	token := access.UserSetToken(userid, password)
	token["error"] = 0
	token["message"] = "success"
	return c.JSON(http.StatusOK, token)

}

/*@@RegisterSendsms@@*/
func RegisterSendsms(c echo.Context) (err error) {
	telephone := c.FormValue("telephone")
	if telephone == "" {
		return config.Success(c, 1, "手机号码错误")
	}
	yzm := strconv.Itoa(rand.Intn(9000) + 999)
	var ops = make(map[string]interface{})
	ops["telephone"] = telephone
	ops["content"] = "您的验证码：" + yzm
	ops["yzm"] = yzm
	cacheKey := "smsReg" + telephone + yzm
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
