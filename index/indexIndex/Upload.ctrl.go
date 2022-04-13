package indexIndex

import (
	"app/access"
	"app/config"
	"app/ext"
	"io"
	"net/http"
	"os"

	"time"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
)

/*解决import未使用*/
func UploadNull(c echo.Context) (err error) {

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
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}

/*@@UploadImg@@*/
func UploadImg(c echo.Context) (err error) {
	file, err := c.FormFile("upimg")
	if err != nil {
		return config.Success(c, 0, "上传出错"+err.Error())
	}
	url, err := file.Open()
	if err != nil {
		return config.Success(c, 0, "上传出错"+err.Error())
	}
	defer url.Close()
	fileurl := "attach/" + file.Filename
	dst, err := os.Create(fileurl)
	if err != nil {
		return config.Success(c, 0, "上传出错"+err.Error())
	}
	defer dst.Close()

	// 下面将源拷贝到目标文件
	if _, err = io.Copy(dst, url); err != nil {
		return config.Success(c, 0, "上传出错"+err.Error())
	}
	//裁剪图片
	src, err := imaging.Open(fileurl)
	if err != nil {
		return config.Success(c, 0, "上传出错"+err.Error())
	}

	// Crop the original image to 300x300px size using the center anchor.
	img100 := imaging.CropAnchor(src, 200, 200, imaging.Center)
	err = imaging.Save(img100, fileurl+".100x100.png")
	imgSmall := imaging.Resize(src, 400, 0, imaging.Lanczos)
	err = imaging.Save(imgSmall, fileurl+".small.png")
	imgMiddle := imaging.Resize(src, 750, 0, imaging.Lanczos)
	err = imaging.Save(imgMiddle, fileurl+".middle.png")
	if err != nil {
		return config.Success(c, 0, "上传出错"+err.Error())
	}
	// 上传成功 同步OOs
	ext.OosUpload(fileurl)
	ext.OosUpload(fileurl + ".100x100.png")
	ext.OosUpload(fileurl + ".small.png")
	ext.OosUpload(fileurl + ".middle.png")
	reData := make(map[string]interface{})
	reData["imgurl"] = fileurl
	reData["trueimgurl"] = config.Image_site(fileurl)
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)
}
