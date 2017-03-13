package populate

import(
	"bytes"
	"io/ioutil"
	"log"
	"fmt"
	"os/exec"
	"html/template"
)



var TemplateFile, _ =  template.ParseFiles("../templates/form.html")

type FormData struct {
	Member string
	CurrentAddr string
	UKAddr string
}



func FillTempl(Member string, CurrentAddr string, UKAddr string){
	buff := bytes.NewBufferString("")
	// Compile and allocate in buffer
	err := TemplateFile.Execute(buff, FormData{
		Member: Member,
		CurrentAddr: CurrentAddr,
		UKAddr: UKAddr,
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("form_compiled.html", buff.Bytes(), 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//convert compiled file to pdf
	err = exec.Command("wkhtmltopdf", "form_compiled.html", "../templates/form_data.pdf").Run()
	if err == nil {
		fmt.Printf("Save successful")
	} else {
		fmt.Printf("Error generating PDF %s", err)
	}
}
