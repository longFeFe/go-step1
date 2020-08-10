/***
** Author: chengxufei
** Date: 2020-04-22 10:10:28
** LastEditors: chengxufei
** LastEditTime: 2020-04-22 10:34:55
** Description: 
**/

package main

import (
	"log"
	"net/http"
	"./api"
)

//http://127.0.0.1/download?key=https://dl.google.com/go/go1.14.7.windows-386.msi
func main() {
	http.HandleFunc("/hello", api.Hello)
	http.HandleFunc("/headers", api.Headers)
	http.HandleFunc("/download", api.Download)
	log.Println("start listening...")
	http.ListenAndServe(":80", nil)
}