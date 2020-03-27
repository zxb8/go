package handlers

import (
	"episode1/storage"
	"io/ioutil"
	"net/http"
)

func PutKey(db storage.DB) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == ""{
			http.Error(w,"miss key name in path",http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val,err := ioutil.ReadAll(r.Body)

		if err != nil{
			http.Error(w,"error reading PUT body",http.StatusBadRequest)
			return
		}

		if err := db.Set(key,val);err != nil{
			http.Error(w,"error setting value in DB",http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
