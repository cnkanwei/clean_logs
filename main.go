package main

import (
	"github.com/pokeyou/clean_logs/yamlConfig"
	"io/ioutil"
	"strings"
	"path"
	"time"
	"os"
	"flag"
	"fmt"
)

type dirs_config struct {
	Dirs []string 	`yaml:dirs`
	Time int64		`yaml:time`
}

var remain_time int64

var c string

func init(){
	flag.StringVar(&c,"c","/etc/logs_need_clean.yaml","yaml config file path")
}

func main() {
	flag.Parse()

	fmt.Println("config file path:"+c)

	dirs := new(dirs_config)

	yamlConfig.ReadConfig(c,dirs)

	if dirs.Time>0 {
		remain_time=dirs.Time
	}else{
		remain_time=604800
	}

	for _,dir := range dirs.Dirs {
		clean_one_dir(dir)
	}
}

func clean_one_dir(dir string){
	files,_ := ioutil.ReadDir(dir)
	fmt.Println("check dir:"+dir)
	for _,file:= range files {
		if file.IsDir() {
			next := next_dir(dir,file.Name())
			clean_one_dir(next)
		}else{
			if path.Ext(file.Name())==".log" {
				if file.ModTime().Unix() < time.Now().Unix() - remain_time {
					del_real_path:=real_path(dir,file.Name())
					fmt.Println("delete file:"+del_real_path)
					os.Remove(del_real_path)
				}
			}
		}
	}
}

func next_dir(dir string , next string) string {
	dir = strings.TrimRight(dir,"/")
	result:=dir+"/"+next
	return result
}

func real_path(dir string , file string) string {
	dir = strings.TrimRight(dir,"/")
	result:=dir+"/"+file
	return result
}
