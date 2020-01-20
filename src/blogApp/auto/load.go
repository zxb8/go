package auto

import (
	"blogApp/database"
	"blogApp/models"
	"blogApp/utils/console"
	"log"
)

func Load(){
	db,err := database.Connect()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil{
		log.Fatal(err)
	}

	err =db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil{
		log.Fatal(err)
	}

	for _,user := range users{
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil{
			log.Fatal(err)
		}
		console.Pretty(user)
	}
}
