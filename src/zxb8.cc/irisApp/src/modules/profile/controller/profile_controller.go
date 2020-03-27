package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"io"
	"os"
	"time"
	"zxb8/src/modules/profile/model"
	"zxb8/src/modules/profile/usecase"
)
const ProfileIDKey = "ProfileID"

type ProfileController struct{
	Ctx iris.Context
	Session *sessions.Session
	profileUsecase usecase.ProfileUsecase
}

func (c *ProfileController) getCurrentProfileID()string{
	return c.Session.GetString(ProfileIDKey)
}

func (c *ProfileController) isProfileLoggedIn() bool{
	return c.getCurrentProfileID() != ""
}

func (c *ProfileController) logout(){
	 c.Session.Destroy()
}

//localhost:3000/profile/register
func (c *ProfileController) GetRegister() mvc.Result{
	if c.isProfileLoggedIn(){
		c.logout()
	}
	return mvc.View{
		Name:"profile/register.html",
		Data:iris.Map{"Title":"Profile Registration"},
	}
}

func (c *ProfileController) PostRegister() mvc.Result{
	firstName := c.Ctx.FormValue("first_name")
	lastName  := c.Ctx.FormValue("last_name")
	email := c.Ctx.FormValue("email")
	password :=c.Ctx.FormValue("password")

	if firstName == "" || lastName =="" || email == "" || password ==""{
		return mvc.Response{
			Path:"/profile/register",
		}
	}
	id := uuid.New()
	profileImage,err := c.uploadImage(c.Ctx,id.String())

	if err != nil{
		return mvc.Response{
			Path:"/profile/register",
		}
	}

	var profile model.Profile
	profile.ID = id.String()
	profile.FirstName = firstName
	profile.LastName = lastName
	profile.Email = email
	profile.Password = password
	profile.ImageProfile = profileImage
	profile.CreateAt = time.Now()
	profile.UpdateAt = time.Now()

	_,err =c.profileUsecase.SaveProfile(&profile)
	fmt.Println(err)
	//if err != nil{
	//	return mvc.Response{
	//		Path:"/profile/register",
	//	}
	//}
	c.Session.Set(ProfileIDKey,profile.ID)

	return nil

	//return mvc.Response{
	//	Path:"/profile/me",
	//}
}

func (c *ProfileController) GetLogin() mvc.Result{
	if c.isProfileLoggedIn(){
		c.logout()
	}
	return mvc.View{
		Name:"/profile/login.html",
		Data:iris.Map{
			"Title":"登录",
		},
	}
}

func ( c *ProfileController) PostLogin() mvc.Result{
	email := c.Ctx.FormValue("email")
	password := c.Ctx.FormValue("password")
	if email =="" || password == ""{
		return mvc.Response{
			Path:"/profile/login",
		}
	}
	profile,err:= c.profileUsecase.GetByEmail(email)
	if err != nil{
		return mvc.Response{
			Path:"/profile/login",
		}
	}
	if !profile.IsValidPassword(password){
		return mvc.Response{
			Path:"/profile/login",
		}
	}
	c.Session.Set(ProfileIDKey,profile.ID)
	return mvc.Response{
		Path:"/profile/me",
	}
}

func (c *ProfileController) GetMe() mvc.Result{
	if !c.isProfileLoggedIn(){
		return mvc.Response{
			Path:"/profile/login",
		}
	}
	profile ,err := c.profileUsecase.GetById(c.getCurrentProfileID())
	if err != nil{
		c.logout()
		c.GetMe()
	}
	return mvc.View{
		Name:"/profile/me.html",
		Data:iris.Map{
			"Title":"个人主页",
			"Profile":profile,
		},
	}

}

func (c *ProfileController) AnyLogout(){
	if c.isProfileLoggedIn(){
		c.logout()
	}
	c.Ctx.Redirect("/profile/login")
}

func (c *ProfileController) uploadImage(ctx iris.Context,id string)(string,error){
	file,info,err :=ctx.FormFile("image_profile")
	if err != nil{
		return "",err
	}
	defer file.Close()
	fileName := fmt.Sprintf("%s%s",id,info.Filename)
	out,err := os.OpenFile("./web/public/images/profile/" + fileName,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		return  "",err
	}
	defer out.Close()
	io.Copy(out,file)
	return fileName,nil
}

