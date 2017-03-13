package populate

import(
	"bytes"
	"io/ioutil"
	"log"
	"fmt"
	"os/exec"
	"html/template"
	"contactform/mail"
	"time"
)



var TemplateFile, _ =  template.ParseFiles("../templates/form.html")

type FormData struct {
	Member string
	CurrentAddr string
	UKAddr string
}



func FillTempl(Member string, CurrentAddr string, UKAddr string, Email string){
	fmt.Printf("FillTempl")
	buff := bytes.NewBufferString("")
	sent_at := time.Now()
	pdf_path := "../templates/form_data.pdf"
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
	err = exec.Command("wkhtmltopdf", "form_compiled.html", pdf_path).Run()
	mail.Send(sent_at, pdf_path, Email);
	if err == nil {
		fmt.Printf("[+ TEMPLATE] Save successful")
	} else {
		fmt.Printf("[- TEMPLATE] Error generating PDF %s", err)
	}
}
