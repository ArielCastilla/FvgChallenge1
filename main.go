package main

// http://localhost:10000/all

// {
    // "Id": "3", 
    // "Direccion": "Brown 555", 
    // "Latitud": "888888", 
    // "Longitud": "999999" 
// }

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	// "io/ioutil"
)

var Sucursales []Sucursal

type Sucursal struct {
    Id string `json:"Id"`
    Direccion string `json:"Direccion"`
    Latitud string `json:"Latitud"`
	Longitud string `json:"Longitud"`
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllSucursales)
	myRouter.HandleFunc("/sucursal", createNewSucursal).Methods("POST")
	myRouter.HandleFunc("/sucursal/{id}", returnSingleSucursal)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleSucursal(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    // fmt.Fprintf(w, "Key: " + key)
	for _, fsucursal := range Sucursales {
        if fsucursal.Id == key {
            json.NewEncoder(w).Encode(fsucursal)	
        }
    }
}

func createNewSucursal(w http.ResponseWriter, r *http.Request) {
    // reqBody, _ := ioutil.ReadAll(r.Body)
    // fmt.Fprintf(w, "%+v", string(reqBody))
	fmt.Fprintf(w, "%+v", "A ver si anda ")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllSucursales(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllSucursales")
    json.NewEncoder(w).Encode(Sucursales)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	
	Sucursales = []Sucursal{
        Sucursal{Id: "1", Direccion: "Av Fleming 123", Latitud: "444444", Longitud: "111111"},
        Sucursal{Id: "2", Direccion: "Av San Martin 123", Latitud: "555555", Longitud: "222222"},
    }

    handleRequests()
}
