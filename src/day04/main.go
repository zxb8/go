package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

type Cat struct{
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct{
	Name string `json:"name"`
	Type string `json:"Type"`
}

type Hamster struct{
	Name string `json:"name"`
	Type string `json:"Type"`
}

func yallo(c echo.Context) error{
	return c.String(http.StatusOK,"yallo from the web server")
}

func getCats(c echo.Context) error{
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")

	if dataType == "string"{
		return c.String(http.StatusOK,fmt.Sprintf("your cat name is ;%s\nand his type is: %s\n",catName,catType))
	}

	if dataType == "json"{
		return c.JSON(http.StatusOK,map[string]string{
			"name":catName,
			"type":catType,
		})
	}

	return nil
}

func addCat(c echo.Context) error{
	cat := Cat{}
	defer c.Request().Body.Close()
	b,err :=ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Println("Failed reading the request body:%s",err)
		return c.String(http.StatusInternalServerError,"")
	}
	err = json.Unmarshal(b,&cat)
	if err !=  nil{
		log.Println("Failed unmarshaling in addCats ")
		return c.String(http.StatusInternalServerError,"")
	}
	log.Printf("this is your cat:%#v",cat)
	return c.String(http.StatusOK,"we got your cat!")

}

func addDogs(c echo.Context) error{
 dog := Dog{}
 defer c.Request().Body.Close()
 err :=json.NewDecoder(c.Request().Body).Decode(&dog)
 if err != nil{
	log.Printf("Failed processing addDog request:%s",err)
	return echo.NewHTTPError(http.StatusInternalServerError)
 }
 log.Printf("this is your dog:%#v",dog)
 return c.String(http.StatusOK,"this is a dog")
}

func addHamster(c echo.Context) error{
	hamster := Hamster{}
	err :=c.Bind(&hamster)
	if err != nil{
		log.Printf("Failed processinf addHamster request:%#v",hamster)
	}
	return c.String(http.StatusOK,"we get hamster")
}

func mainAdmin(c echo.Context) error{
	return c.String(http.StatusOK,"horay you arte admin")
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error{
		cookie,err :=c.Cookie("sessionID")
		if err != nil{
			return err
		}
		if cookie.Value =="some_string"{
			return next(c)
		}
		return c.String(http.StatusOK,"ok")
	}
}

func main(){
	e := echo.New()
	g:=e.Group("/admin")

	g.Use(middleware.Logger())

	g.GET("/main",mainAdmin)

	e.GET("/",yallo)
	e.GET("/cats/:data",getCats)
	e.POST("/cats",addCat)
	e.POST("/dogs",addDogs)
	e.POST("/hamsters",addHamster)
	e.Start(":8000")
}
