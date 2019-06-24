package main

import(
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	date_time_now := time.Now()
	HomePageVars := PageVariables{
		Date: date_time_now.Format("01-02-2006"),
		Time: date_time_now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html")
	if err != nil {
  	  log.Print("template parsing error: ", err)
  	}
	err = t.Execute(res, HomePageVars) 
	if err != nil {
  	  log.Print("template executing error: ", err)
  	}
}

type PageVariables struct {
	Date string
	Time string
}

