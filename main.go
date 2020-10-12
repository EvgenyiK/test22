package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"


	"github.com/gorilla/mux"
)

type Adress struct{
	Req string `json:"req_type"`
	Result string 
	Data []interface{}	`json:"data"`
}

type Users struct{
	Item string `json:"item"`
}


var adress []Adress


func createAdress(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var ad Adress
	err:=json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		ad.Result = "fail"
		log.Fatalf("Unable to decode the request body. %v", err)
	} else {
		ad.Result = "success"
	}
	
	
	//var us Users

	for _, i := range ad.Data {
		k,_ := json.Marshal(i)
		fmt.Println(k)
	}

	

	//fmt.Println(ad.Data)
	json.NewEncoder(w).Encode(ad)
}

/*var state []string

func ToSlice(m map[string]interface{}) []state {
	for k, v := range m {
		v.
	}
}*/

func main()  {
	r:= mux.NewRouter()
	r.HandleFunc("/", createAdress).Methods("POST")
	fmt.Println("start server")
	log.Fatal(http.ListenAndServe(":8000", r))
	
}