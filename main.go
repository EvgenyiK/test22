package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

type response struct{
	Req string `json:"res_type"`
	Result string `json:"result"`
	Data string `json:"data"`
}


func createAdress(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	var ad Adress
	err:=json.NewDecoder(r.Body).Decode(&ad)
	if err != nil{
		ad.Result = "fail"
		log.Fatalf("Unable to decode the request body. %v", err)
	} else {
		ad.Result = "success"
	}
	
	var data []string
	for _, v := range ad.Data {
		data = append(data, fmt.Sprint(v))
	}
	

	t:= searchSort(data)
	d := strings.Join(t, "||")
	e:= remove_quotes(d)
	

	res:= response{
		Req: ad.Req,
		Result: ad.Result,
		Data: e,
	}
	
	json.NewEncoder(w).Encode(res)
}




func search(docs []string, term string) []string {
	var str []string
    for _, doc := range docs {
        if strings.Contains(doc, term) {
				str = append(str, doc+"...")
		}
	}
	return str
}

func searchSort(t []string) []string {
	var answ []string
	c:= search(t,"CA")
	m:= search(t,"MA")
	o:= search(t,"OK")
	p:= search(t,"PA")
	v:= search(t,"VA")

	answ = append(answ, fmt.Sprintf("California: %s",c))
	answ = append(answ, fmt.Sprintf("Massachusec: %s",m))
	answ = append(answ, fmt.Sprintf("Oklahoma: %s",o))
	answ = append(answ, fmt.Sprintf("Pensylvania: %s",p))
	answ = append(answ, fmt.Sprintf("Virginia: %s",v))

	fmt.Println("California: "+"\n"+".....^",remove_quotes(fmt.Sprintf("%s",c)))
	fmt.Println("Massachusec: "+"\n"+".....^",remove_quotes(fmt.Sprintf("%s",m)))
	fmt.Println("Oklahoma: "+"\n"+".....^",remove_quotes(fmt.Sprintf("%s",o)))
	fmt.Println("Pensylvania: "+"\n"+".....^",remove_quotes(fmt.Sprintf("%s",p)))
	fmt.Println("Virginia: "+"\n"+".....^",remove_quotes(fmt.Sprintf("%s",v)))

	return answ
}

func remove_quotes(s string) string {
    var b bytes.Buffer
    for _, r := range (s) {
        if r != '{' && r != '}' && r != '[' && r != ']' {
            b.WriteRune(r)
        }
    }
 
	return b.String()
}

func main()  {
	r:= mux.NewRouter()
	r.HandleFunc("/", createAdress).Methods("POST")
	log.Fatal(http.ListenAndServe(":8008", r))
}