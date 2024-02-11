package main

import(
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"go_tutorials/api_tutorial/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r*chi.Mux=chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API services..")

	err:=http.ListenAndServe("localhost:8000",r)
	if err!=nil{
		log.Error(err)
	}
}