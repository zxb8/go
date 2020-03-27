package main

import (
	"database/sql"
	"io"
	"strconv"
	"strings"
	"time"

	//"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct{
	Title string
	Body []byte
}


var(
  tmpView1 = template.Must(template.New("test").ParseFiles("base.html","test.html","index.html"))
  tmpEdit1 = template.Must(template.New("edit").ParseFiles("base.html","edit.html","index.html"))
  db,_ = sql.Open("sqlite3","cache/new.db")
  createDB = "create table if not exists pages(title text,body blob,timestamp text)"
)

func (p *Page) saveCache() error{
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	f := "cache/" + p.Title +".txt"
	db.Exec(createDB)
	tx,_ :=db.Begin()
	stmt,_ := tx.Prepare("insert into page(title,body,timestamp ) values(?,?,?)")
	_,err:=stmt.Exec(p.Title,p.Body,timestamp)
	tx.Commit()
	ioutil.WriteFile(f,p.Body,0600)
	return err
}

func load(title string)(*Page,error){
	f := "cache/" + title + ".txt"
	body,err := ioutil.ReadFile(f)
	if err != nil{
		return nil,err
	}
	return &Page{Title:title,Body:body},nil
}

func loadSource(title string)(*Page,error){
	var name string
	var body []byte
	q,_:=db.Query("select title,body from pages where title='"+title+"'order by timestamp desc limit 1")
	for q.Next(){
		q.Scan(&name,&body)
	}
	return &Page{Title:name,Body:body},nil
}

func view(w http.ResponseWriter,r *http.Request){
	title := r.URL.Path[len("/test/"):]
	//fmt.Println(r.URL.Path)  //  /test/hh
	//fmt.Println(title)
	p,err := loadSource(title)

	if err != nil{
		p,_  = load(title)
	}

	tmpView1.ExecuteTemplate(w,"base",p)

	//t ,_ :=template.ParseFiles("test.html")
	//	//t.Execute(w,p)
}

func edit(w http.ResponseWriter,r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p,_ :=load(title)

	tmpEdit1.ExecuteTemplate(w,"base",p)
	//t,_:=template.ParseFiles("edit.html")
	//t.Execute(w,p)
}

func save(w http.ResponseWriter,r *http.Request){
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title:title,Body:[]byte(body)}
	p.saveCache()
}

func main(){
	p := &Page{Title:"Test",Body:[]byte("Welcome to the Test App")}
	p.save()

	http.HandleFunc("/test/",view)
	http.HandleFunc("/edit",edit)
	http.ListenAndServe(":8080",nil)
}
