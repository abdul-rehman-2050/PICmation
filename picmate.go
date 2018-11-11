package main

import "os"
import "fmt"
import "io/ioutil"




func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	cmdList := os.Args[1:]
	if(len(cmdList)< 1){
	
		fmt.Fprintf(os.Stderr, "error: %v\n", "No file name and/or controller name found")
        os.Exit(1)
	
	}
	
	 b, err := ioutil.ReadFile("header.temp") // just pass the file name
	 check(err);
	
	err = ioutil.WriteFile(cmdList[0], b, 0644)
    check(err)
	
	fmt.Println("File Created Successfully")
	
}
