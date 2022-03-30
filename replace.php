<?php
$arr=[
	"forum/forumIndex",
	"forum/forumAdmin",
	"index/indexIndex",
	"index/indexAdmin"
];
$old='return c.JSON(http.StatusOK, reData)';
$new='
	reJson := make(map[string]interface{}) 
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"]=reData;
	return c.JSON(http.StatusOK, reJson) 
';
foreach($arr as $dir){
	$files=glob($dir."/*");
	foreach($files as $file){
		$c=file_get_contents($file);
		$c=str_replace("reJson","reData",$c);
		$c=str_replace($old,$new,$c);
		 
		file_put_contents($file,$c);
	}
}
echo "success";

?>