package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"strconv"
	"io/ioutil"
	"strings"
	"math"
	
)

var Sucursales []Sucursal
const radio = 3389.5

type Sucursal struct {
    Id string `json:"Id"`
    Direccion string `json:"Direccion"`
    Latitud string `json:"Latitud"`
	Longitud string `json:"Longitud"`
}

type Punto struct {
	Lat string `json:"Lat"`
	Long string `json:"Long"`
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllSucursales)
	myRouter.HandleFunc("/altasucursal", createNewSucursal).Methods("POST")
	myRouter.HandleFunc("/sucursalmascercana", sucursalmMasCercana).Methods("POST")
	myRouter.HandleFunc("/sucursal/{id}", returnSingleSucursal)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func distancia(vpunto1, vpunto2 Punto) float64 {
	sLat1, err := strconv.ParseFloat(vpunto1.Lat, 64)
		if err != nil {
			fmt.Println("Coordenada no válida")
			return 0
		}
	sLong1, err := strconv.ParseFloat(vpunto1.Long, 64)
		if err != nil {
			fmt.Println("Coordenada no válida")
			return 0
		}
	sLat2, err := strconv.ParseFloat(vpunto2.Lat, 64)
		if err != nil {
			fmt.Println("Coordenada no válida")
			return 0
		}
	sLong2, err := strconv.ParseFloat(vpunto2.Long, 64)
		if err != nil {
			fmt.Println("Coordenada no válida")
			return 0
		}

	s1, c1 := math.Sincos(rad(sLat1))
	s2, c2 := math.Sincos(rad(sLat2))
	clong := math.Cos(rad(sLong1 - sLong2))
	return radio * math.Acos(s1*s2+c1*c2*clong)
}
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
func masCerca(deg float64) float64 {
	return deg * math.Pi / 180
}
func sucursalmMasCercana(w http.ResponseWriter, r *http.Request) {
	var vpunto1 Punto
	var vpunto2 Punto
	reqBody, _ := ioutil.ReadAll(r.Body)
	dec := json.NewDecoder(strings.NewReader(string(reqBody)))
	dec.Decode(&vpunto1)
    database, _ := sql.Open("sqlite3", "./dbsucursales.db")
	rows, _ := database.Query("SELECT Id, Direccion, Latitud, Longitud FROM tablesucursales")
	
    var Id string
    var Direccion string
    var Latitud float64
	var Longitud float64
	var DistanciaMenor float64
	var IdSucursalMasCercana string
	var breakFlag int
	DistanciaMenor = 9223372036854775807
	breakFlag = 0
    for rows.Next() {
        rows.Scan(&Id, &Direccion, &Latitud, &Longitud)
		sLatitud := fmt.Sprintf("%f", Latitud)
		sLongitud := fmt.Sprintf("%f", Longitud)
		vpunto2 = Punto{sLatitud, sLongitud}
		distanciaCalculada := distancia(vpunto1, vpunto2)
		sdistancia := fmt.Sprintf("%f", distanciaCalculada)
		if distanciaCalculada > 0 {
			if DistanciaMenor > distanciaCalculada{
				DistanciaMenor = distanciaCalculada
				IdSucursalMasCercana = Direccion
			}
			fmt.Fprintln(w,"La sucursal " + Direccion + " está a " + sdistancia + " Km.")
		} else {
			breakFlag = 1
			break
		}
    }
	if breakFlag > 0 {	
		fmt.Fprintln(w,"Coordenadas no válidas")
	} else {
		sdistanciaMenor := fmt.Sprintf("%f", DistanciaMenor)
		fmt.Fprintln(w,"................ ")
		fmt.Fprintln(w,"Sucursal " + IdSucursalMasCercana + " a " + sdistanciaMenor+ " Km, es la más cercana. ")
		fmt.Println("Verificación de sucursal mas cercana realizada")
	}
}

func createNewSucursal(w http.ResponseWriter, r *http.Request) {
	var suc Sucursal
	reqBody, _ := ioutil.ReadAll(r.Body)
	dec := json.NewDecoder(strings.NewReader(string(reqBody)))
	dec.Decode(&suc)
    database, _ := sql.Open("sqlite3", "./dbsucursales.db")	
	statementInsert, _ := database.Prepare("INSERT INTO tablesucursales (Id, Direccion, Latitud, Longitud) VALUES (?, ?, ?, ?)")
    statementInsert.Exec(suc.Id, suc.Direccion, suc.Latitud, suc.Longitud)
	fmt.Println("Nueva sucursal dada de alta")
}

func returnSingleSucursal(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
	database, _ := sql.Open("sqlite3", "./dbsucursales.db")
    rows, _ := database.Query("SELECT Id, Direccion, Latitud, Longitud FROM tablesucursales WHERE Id = " + key)

    var Id string
    var Direccion string
    var Latitud float64
	var Longitud float64
    for rows.Next() {
        rows.Scan(&Id, &Direccion, &Latitud, &Longitud)
        
		sLatitud := fmt.Sprintf("%f", Latitud)
		sLongitud := fmt.Sprintf("%f", Longitud)
		fmt.Fprintln(w, "Id: " + Id)
		fmt.Fprintln(w, "Direccion: " + Direccion)
		fmt.Fprintln(w, "Latitud: " + sLatitud)
		fmt.Fprintln(w, "Longitud: " + sLongitud)
    }
	fmt.Println("Función returnSingleSucursal ejecutada")
}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Software engineer - Challenge 1 - Sucursal CRUD")
	fmt.Println("Función homePage ejecutada")
}

func returnAllSucursales(w http.ResponseWriter, r *http.Request){
	databaseall, _ := sql.Open("sqlite3", "./dbsucursales.db")
    rows, _ := databaseall.Query("SELECT Id, Direccion, Latitud, Longitud FROM tablesucursales")

	var Id string
    var Direccion string
    var Latitud float64
	var Longitud float64
    for rows.Next() {
        rows.Scan(&Id, &Direccion, &Latitud, &Longitud)
        
		sLatitud := fmt.Sprintf("%f", Latitud)
		sLongitud := fmt.Sprintf("%f", Longitud)
		fmt.Fprintln(w, "Id: " + Id)
		fmt.Fprintln(w, "Direccion: " + Direccion)
		fmt.Fprintln(w, "Latitud: " + sLatitud)
		fmt.Fprintln(w, "Longitud: " + sLongitud)
		fmt.Fprintln(w, "-------------------")
		fmt.Fprintln(w, " ")
    }
	fmt.Println("Función returnAllSucursales ejecutada")
}

func main() {
	fmt.Println("Rest API Inicio")	
    database, _ := sql.Open("sqlite3", "./dbsucursales.db")
	fmt.Println("Base de datos dbsucursales abierta")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tablesucursales (Id INTEGER PRIMARY KEY, Direccion TEXT, Latitud REAL, Longitud REAL)")
    statement.Exec()	
	fmt.Println("Tabla tablesucursales creada")
    handleRequests()
}
