package config

var AppConfig map[string]interface{} = make(map[string]interface{})

func init() {

	AppConfig["images_site"] = "http://oos.mmlang.com"

}
