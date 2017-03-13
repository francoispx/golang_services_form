package main


import (
	"contactform/populate"
	"fmt"
)

func main(){
	populate.FillTempl("Foo", "Bar", "dasda")
	fmt.Printf("OK")
}
