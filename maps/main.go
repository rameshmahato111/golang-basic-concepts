package main

import (
	"fmt"
	

)


func main(){
	language := make(map[string]string)

	language["py"] = "python"
	language["js"] = "javascript"
	language["rs"] = "rust"
	language["go"] = "golang"

	fmt.Println("this is the map value", language)
	fmt.Println("this is the value getting from key", language["rs"])

	// to delete value we can use delete 

	delete(language, "py")
	fmt.Println("this is the value after deleted ", language)


}