package main

import (
	"os"
	"fmt"
)
func main() {
	var s,sep string 
	for index := 0; index < len(os.Args); index++ {
		s+=sep+os.Args[index]
		sep=" "
	}
	fmt.Println(s)
	//print(os.Args)
	//print(s)
	//print(len(os.Args))
	for _,arg:=range os.Args[1:]{
		s+=sep+arg
		sep =" "
	}
	fmt.Println(s)
}