/*
 * @Author: chengxufei
 * @Date: 2020-08-10 23:05:57
 * @LastEditors: chengxufei
 * @LastEditTime: 2020-08-11 00:04:13
 * @Description: 
 */ 
 package api
 import (
	"fmt"
	"os"
	"time"
	"../xldl"
	"net/http"
	"log"
	"path/filepath"
    //"path"
)



const (
	kMB int64 = 1024 * 1024 
	kSaveDirectory string = "saveto"
)

func downloadworker(url string) {
	str, _ := os.Getwd()
	str += ("\\" + kSaveDirectory)
	log.Println("download file save diectory:", str)
	if !xldl.InitXLEngine() {
		log.Println("InitXLEngine Error!")
		time.Sleep(1 * 1000 * time.Millisecond)
		return
	}
	defer xldl.UnInitXLEngine()
	dloader := xldl.NewXLDownloader(str)
	_, filename := filepath.Split(url)
	task := dloader.AddTask(url, filename)
	task.Start()
	log.Println("start download task, url:", url, ", filename:", task.FileName)
	
	var exitLoop  = false
	for exitLoop == false {
		info, ret := task.Info()
		if ret {
			log.Printf("TotalSize=%.2f MB, Percent=%.2f %%, Speed=%.2f Mb/s\n", float64(info.TotalSize) / float64(kMB) , info.Percent * 100, float64(info.Speed) / float64(kMB))
			switch info.Stat {
			case xldl.TSC_COMPLETE:
				log.Println("download complete!")
				task.Delete()
				return
			case xldl.TSC_ERROR:
				log.Println("download error!, erno:", info.FailCode)
				exitLoop = true
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}
	task.Stop()
	task.Delete()
	fmt.Println("clean temp directory:", task.DeleteTempFile())
	dloader.RemoveAll()
}


func Download(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	url := req.FormValue("key")
	log.Println("url:", url)
	go downloadworker(url)
	fmt.Fprintf(w, "commit download task ok.\n")
}