package ext

import (
	"os"

	"github.com/imroc/req/v3"
)

func OosUpload(url string) {
	file, _ := os.Open(url)
	client := req.C().EnableDumpAllWithoutRequestBody()
	client.R().
		SetFormData(map[string]string{
			"filename": url,
		}).
		SetFileReader("upimg", url, file).
		Post("http://oos.mmlang.com/upload.php")

}
