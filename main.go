package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

type Adress struct{
	Req string `json:"req_type"`
	Result string 
	Data []Data	`json:"data"`
}

type Data struct{
	Item string `json:"item"`
}



func createAdress(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	var ad Adress
	err:=json.NewDecoder(r.Body).Decode(&ad)
	if err != nil {
		ad.Result = "fail"
		log.Fatalf("Unable to decode the request body. %v", err)
	} else {
		ad.Result = "success"
	}
	
	var data []string
	for _, v := range ad.Data {
		data = append(data, fmt.Sprint(v))
	}
	
	sort.Sort(Alphabetic(data))
	fmt.Println()
	
	d:=strings.Join(data, ".....")

	fmt.Println(d)

	json.NewEncoder(w).Encode(ad)
}

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
    var si string = list[i]
    var sj string = list[j]
    var si_lower = strings.ToLower(si)
    var sj_lower = strings.ToLower(sj)
    if si_lower == sj_lower {
        return si < sj
    }
    return si_lower < sj_lower
}






func main()  {
	r:= mux.NewRouter()
	r.HandleFunc("/", createAdress).Methods("POST")
	fmt.Println("start server")
	log.Fatal(http.ListenAndServe(":8008", r))
}