package handlers

import (
	"episode1/storage"
	"fmt"
	"net/http"
)

func GetKey(db storage.DB) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == ""{
			http.Error(w,"missing key name in query string",http.StatusBadRequest)
			return
		}
		val,err :=db.Get(key)
		if err == storage.ErrNotFound{
			http.Error(w,"not found",http.StatusBadRequest)
			return
		}else if err != nil{
			http.Error(w,fmt.Sprintf("error getting value from database: %s",err),http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
