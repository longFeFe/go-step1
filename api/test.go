/*
 * @Author: chengxufei
 * @Date: 2020-08-10 23:13:45
 * @LastEditors: chengxufei
 * @LastEditTime: 2020-08-10 23:14:18
 * @Description: 
 */ 

 package api

 import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
