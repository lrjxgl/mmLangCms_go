package access

import (
	"app/config/cache"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserCheckAccess(c echo.Context) uint {

	str := c.QueryParam("token")

	if str == "" {
		str = c.FormValue("token")
	}
	if str == "" {
		return 0
	}
	token := cache.CacheGet(str)
	if token == "" {
		return 0
	} else {
		i, e := strconv.Atoi(token)
		if e != nil {
			return 0
		}
		return uint(i)
	}

}

func UserSetToken(userid uint, password string) map[string]interface{} {
	var useridStr = strconv.Itoa(int(userid))
	salt := strconv.Itoa(rand.Intn(10000))
	data := []byte(password + salt)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	nstr := md5str[0:16]
	var token string = useridStr + ".user." + salt + "." + nstr
	var token_expire int64 = 3600 * 24 * 3
	//refreh_token生成
	salt = strconv.Itoa(rand.Intn(10000))
	data = []byte(password[0:16] + salt)
	has = md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	var refresh_token string = useridStr + ".user." + salt + "." + md5str[0:16]
	var refresh_expire_time int64 = 3600 * 24 * 14
	reJson := make(map[string]interface{})
	reJson["token"] = token
	reJson["token_expire"] = token_expire
	reJson["refresh_token"] = refresh_token
	reJson["refresh_expire_time"] = refresh_expire_time
	cache.CacheSet(token, useridStr, token_expire)
	cache.CacheSet(refresh_token, useridStr, refresh_expire_time)
	return reJson
}
