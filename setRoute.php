<?php
$javaRoot="./"; 
$mds=["index","forum"];
foreach($mds as $md){
	writeRoute($md,"index");
	writeRoute($md,"admin");
}
 
function writeRoute($mm,$dir){
	global $javaRoot;
	$fDir=$javaRoot.$mm."/".$mm.ucwords($dir);
	$files=glob($fDir."/*.go");
	$routeStr="";
	foreach($files as $file){
		$ex=explode(".",basename($file));
		$fname=$ex[0];
		$routeStr.="		/*".$fname."*/\r\n";
		$fc=file_get_contents($file);
		preg_match_all("/@@(\w+)@@/iUs",$fc,$arr);
		if($dir=='index'){
			 
			foreach($arr[1] as $action){
				$abstr=Abt($action);
				$ex=explode("_",$abstr);
				//动作 index show
				$act=$ex[count($ex)-1];
				//ctrl
				$ctrl=str_replace("_".$act,"",$abstr);
				$rootPath='/'.$ctrl."/".$act;
				if(preg_match("/save/",$act)){
					$routeStr.='		e.POST("'.$rootPath.'", '.$mm.'Index.'.$action.');'."\r\n";
				}else{
					$routeStr.='		e.GET("'.$rootPath.'", '.$mm.'Index.'.$action.');'."\r\n";
				}				
				
			}
			 
		}else{
			
			foreach($arr[1] as $action){
				$abstr=Abt($action);
				$ex=explode("_",$abstr);
				//动作 index show
				$act=$ex[count($ex)-1];
				//ctrl
				$ctrl=str_replace("_".$act,"",$abstr);
				if($mm=="index"){
					$rootPath="/".$dir."/".$ctrl."/".$act;
				}else{
					$rootPath='/'.$dir.'/'.$ctrl."/".$act;
				}	
				$routeStr.='		e.GET("'.$rootPath.'", '.$mm.'Admin.'.$action.');'."\r\n";
			}
		}
	}
 
	$content='
	package router

	import (
		"app/'.$mm.'/'.$mm.ucwords($dir).'"
		"github.com/labstack/echo/v4"
	)

	func '.ucwords($mm).ucwords($dir).'Router(e *echo.Echo) {
'.$routeStr.'
	}

	';
	$fDir=$javaRoot."router";
	$file=$fDir.'/'.$mm.".".$dir.".go";
	umkdir($fDir);
	file_put_contents($file,$content);
}

/*驼峰转下划线*/
function Abt($camelCaps,$separator='_')
{
	return strtolower(preg_replace('/([a-z])([A-Z])/', "$1" . $separator . "$2", $camelCaps));
}

function umkdir($dir)
{
	$arr=explode("/",$dir);
	foreach($arr as $key=>$val)
	{
		$d="";
		for($i=0;$i<=$key;$i++)
		{
			$d.=$arr[$i]."/";
		}
		if(!file_exists($d) && (strpos($val,":")===false))
		{
			mkdir($d,0755);
		}
	}
}