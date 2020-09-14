package main

import (
	"fmt"
	//"strings"
	//"os/exec"
	//"encoding/csv"
	//"log"
	"os"
	//"io"
	"time"
	//"github.com/kardianos/service"

)

func doStuff() {
			f, err := os.Create("ServiceTestData.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			
			for {
				f.WriteString("Hello World" + fmt.Sprintln())
				time.Sleep(1 * time.Second)
			}
 }
 
 
func main() {  
	 doStuff()	
}
