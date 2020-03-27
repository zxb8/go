package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main01(){
	user3 :=[]string{
		`3`,`,mary`,`jane`,`83`,`,ary_jane`,`mary@email.com`,`123`,
	}
	user4 :=[]string{
		`4`,`,bob`,`jones`,`93`,`,bob_jones`,`bob_jones@email.com`,`666`,
	}

	f,err:=os.Create("./user.db.csv")



	fmt.Println(user3,user4)

	w :=csv.NewWriter(f)
	w.Write(user3)
	w.Write(user4)
	w.Flush()

	err = w.Error()

	if err != nil{
		log.Fatal(err)
	}
}

func main(){

}
