package main

import "os"
import "fmt"
import "io/ioutil"
import "strconv"



			 

			 
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getUseDelay(crystalValue int) string{
	return "#use delay(clock="+strconv.Itoa(crystalValue)+")\n\n\n";
}
func getIncludeControllerName(controllerName string) string{
	return ("#include<"+controllerName+".h>\n")
}
func getDeviceSignature(num int) string{
	return "#device *="+strconv.Itoa(num)+"\n"
}
func getDefaultFuseBits()string{
	return "#fuses HS,NOWDT,NOPROTECT,NOLVP\n"
}
func getInfiniteLoop() string{
	return "\twhile(True){\n\t}\n"
}
func openMainBody() string{
	return "\nvoid main(){\n"
}
func closeBody() string{
	return "\n}\n"
}


func composeAllProgram()string{
  return getIncludeControllerName("PIC16F887")+
  getDeviceSignature(16)+
  getDefaultFuseBits()+
  getUseDelay(8000000)+
  openMainBody()+
  getInfiniteLoop()+
  closeBody()
  
}


func main() {
	cmdList := os.Args[1:]
	if(len(cmdList)< 1){
	
		fmt.Fprintf(os.Stderr, "error: %v\n", "No file name and/or controller name found")
        os.Exit(1)
	
	}
	err := ioutil.WriteFile(cmdList[0], []byte(composeAllProgram()), 0644)
    check(err)
	fmt.Println("File Created Successfully")
}
