package main

import (
	"fmt"
	"sort"
)

func main(){
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "Calicanoa"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states,"OR")
	fmt.Println(states)

	states["NR"] = "New York"

	for k,v := range states{
		fmt.Printf("%v:%v\n",k,v)
	}

	keys := make([]string,len(states))
	i:=0

	for k:= range states{
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i:= range keys{
		fmt.Printf("%v",i)
	}


}
