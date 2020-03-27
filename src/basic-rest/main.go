package main

import (
	"fmt"
	"os"
	"net/http"
)

func SetJSONResp(res http.ResponseWriter,message []byte,httpCode int){
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func main(){
	http.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
		message := []byte(`{"message":"server is up"}`)
		SetJSONResp(res,message,http.StatusOK)
	})

	err :=http.ListenAndServe(":9000",nil)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
