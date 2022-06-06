package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"regexp"
	"time"
	"fmt"
	"log"
	"os"
"io"
)

//Struct:      http://json2struct.mervine.net/ 
//Probar Json: https://www.jsonformatter.io/
//Ejemplo:     http://pastebin.com/raw/Fw2P6GLn/
//Crear json   http://objgen.com/json

//series
type Series struct {
	Author string `json:"author"`
		Groups []struct {
			Groups []struct {
				Groups []struct {
					Image    string `json:"image"`
					Name     string `json:"name"`
					Stations []struct {
						Image      string `json:"image"`
						ImageScale string `json:"imageScale"`
						IsHost     string `json:"isHost"`
						Name       string `json:"name"`
						URL        string `json:"url"`
					} `json:"stations"`
				} `json:"groups"`
				Image string `json:"image"`
				Info  string `json:"info"`
				Name  string `json:"name"`
			} `json:"groups"`
			Image string `json:"image"`
			Name  string `json:"name"`
		} `json:"groups"`
		Image string `json:"image"`
		Info  string `json:"info"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	}

//pelis
type Peliculas struct {
	Author string `json:"author"`
	Groups []struct {
		Image    string `json:"image"`
		Name     string `json:"name"`
		Stations []struct {
			Embed           string `json:"embed"`
			Image           string `json:"image"`
			Name            string `json:"name"`
			PlayInNatPlayer string `json:"playInNatPlayer"`
			URL             string `json:"url"`
		} `json:"stations"`
	} `json:"groups"`
	Image string `json:"image"`
	Info  string `json:"info"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

var s,ii, message, foto, video, chtml,lfoto string
var p int
var i int // mi puntero
var t *template.Template
var routeMatch *regexp.Regexp
var pd pageData
var peliculas Peliculas
var series Series
var azaro int

//web
type pageData struct {
	Title       string
	CompanyName string
	Npeli       string //nuevo
	Nfoto       string //nuevo
	Nurl        string //nuevo
	Chtml       string //nuevo
	Lfoto       string //nuevo
}

//web
func root(w http.ResponseWriter, r *http.Request) {
	matches := routeMatch.FindStringSubmatch(r.URL.Path)

	if len(matches) >= 1 {
		page := matches[1] + ".html"
		if t.Lookup(page) != nil {
			w.WriteHeader(200)
			t.ExecuteTemplate(w, page, pd)
			return
		}
	} else if r.URL.Path == "/" {
		w.WriteHeader(200)
		t.ExecuteTemplate(w, "index.html", pd)
		return
	}
	w.WriteHeader(404)
	w.Write([]byte("NOT FOUND "))
}



func pelis(w http.ResponseWriter, r *http.Request) {
// http://127.0.0.1:8000/pelis?id=1
id := r.URL.Query()
log.Println("GET pelis : ", id)
//peli := id.Get("peli")

firstvalue := id["id"]
//pasando de slice []string a int simple
pp, _ := strconv.Atoi(firstvalue[0]) 
p = int(pp)
fmt.Println(p)
//a(p)
c(p)
t.ExecuteTemplate(w, "index.html", pd)
//w.Write([]byte("ok"))
//limpiamos var para no rep la lista
chtml=""
lfoto=""
}

func fseries(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8000/series?id=1&id=1&id=1
	id := r.URL.Query()
	log.Println("GET series : ", id)
	
	//peli := id.Get("peli")
	
	firstvalue := id["id"]
	secondvalue := id["id"]
	tercerovalue := id["id"]
	//pasando de slice []string a int simple
	pp, _ := strconv.Atoi(firstvalue[1]) 
	p = int(pp)
	pa, _ := strconv.Atoi(secondvalue[2]) 
	pe := int(pa)
	ter, _ := strconv.Atoi(tercerovalue[0]) 
	tre := int(ter)
	
	fmt.Println(ter,p,pe)
	aa(p)
	cc(tre,p, pe)
	t.ExecuteTemplate(w, "index.html", pd)
	//w.Write([]byte("ok"))
	//limpiamos var para no rep la lista
	chtml=""
	lfoto=""
	}
// 
func azar() int {
	//Eleccion al azar
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 20
	i = (rand.Intn(max-min) + min)
	return i
}

func a (i int){ //API a lo bestia
	// Datos en Json de servidor remoto
	url := "http://pastebin.com/raw/0qtVr1c4"
	// Recogemos los datos
	res, erro := http.Get(url)
	if erro != nil {
		panic(erro.Error())
	}
	body, erro := ioutil.ReadAll(res.Body)
	if erro != nil {
		panic(erro.Error())
	}

	erro = json.Unmarshal(body, &peliculas)
	//fmt.Println(peliculas.Name)                       //[i].Type)
	
	fmt.Println(peliculas.Groups[0].Stations[i].Name)
	//Capitana Marvel (2019)
	//fmt.Println("User Age: " + strconv.Itoa(peliculas.Peliculass[i].Age))

	if erro != nil {
		panic(erro.Error())
	}
    
}

func aa (i int){ //SERIES
	// Datos en Json de servidor remoto
	//url := "http://pastebin.com/raw/Fw2P6GLn"
	url := "http://127.0.0.1:8000/api/jsonseries.json"
	// Recogemos los datos
	res, erro := http.Get(url)
	if erro != nil {
		panic(erro.Error())
	}
	body, erro := ioutil.ReadAll(res.Body)
	if erro != nil {
		panic(erro.Error())
	}

	erro = json.Unmarshal(body, &series)
	//fmt.Println(peliculas.Name)                       //[i].Type)
	
	fmt.Sprintf(series.Groups[0].Groups[0].Groups[0].Stations[0].Name) //Capitana Marvel (2019)
	//Capitana Marvel (2019)
	//fmt.Println("User Age: " + strconv.Itoa(peliculas.Peliculass[i].Age))

	if erro != nil {
		panic(erro.Error())
	}
    
}


func b(){
	//web

	var err error

	t, err = template.ParseGlob("*.html")

	if err != nil {
		log.Println("Cannot parse templates:", err)
		os.Exit(-1)
	}
}
	

func c(i int){			//lista TODO: FOR
	for i := 0; i < 28; i++ {
		ii = strconv.Itoa(i) //para enlaces
		//soporte 20 //message = fmt.Sprintf(peliculas.Groups[i].Stations[i].Name)
		message = fmt.Sprintf(peliculas.Groups[0].Stations[i].Name) //Capitana Marvel (2019)
		foto = fmt.Sprintf(peliculas.Groups[0].Stations[i].Image)
		video = fmt.Sprintf(peliculas.Groups[0].Stations[i].URL) /// fin añadidos
		chtml = chtml + `<div class="movie">
						<img src="` + foto + ` "alt="` + message + `" title="` + message + `" width="75%" height="75%">
						<br><video  width="75%" height="75%" controls poster="` + foto + `">
						<source src="` + video + `" type="video/mp4">
						</video>`

		lfoto = lfoto + `<a href='pelis?id=` + ii + `'> 
		                 <img src='` + foto + `' alt='`+message+`' width='50' height='50'>`
}

		lfoto = lfoto + `</a></div>`

	routeMatch, _ = regexp.Compile(`^\/(\w+)`)

	pd = pageData{
		"Cine Online",
		peliculas.Name,
		peliculas.Groups[0].Stations[i].Name,
		peliculas.Groups[0].Stations[i].Image,
		peliculas.Groups[0].Stations[i].URL,
		chtml,
		lfoto,
	}
}

func cc(t,i,o int){			
	//for i := 0; i < 11; i++ { //capitulos
		ii = strconv.Itoa(i) //para enlaces int to string
		oo:= strconv.Itoa(o) //para enlaces
		tt:= strconv.Itoa(t) //para enlaces
		//lista TODO: pasar parametros xxx&zzz serie/cap tamaño de imagen 
		
		message = fmt.Sprintf(series.Groups[t].Groups[i].Name) //nombre de la serie
		foto  = fmt.Sprintf(series.Groups[t].Groups[i].Groups[0].Stations[o].Image) //foto
		video  = fmt.Sprintf(series.Groups[t].Groups[i].Groups[0].Stations[o].URL) //video
		chtml = chtml + `<div class="movie">
						<img src="` + foto + ` "alt="` + message + `" title="` + message + `" width="75%" height="75%">
						<br><video  width="75%" height="75%" controls poster="` + foto + `">
						<source src="` + video + `" type="video/mp4">
						</video>`

		lfoto = lfoto + `<a href='series?id=` + tt + `&id=` + ii +  `&id=` + oo + `'> 
		                 <img src='` + foto + `' alt='`+message+`' width='50' height='50'>`
//}
		lfoto = lfoto + `</a></div>`
	routeMatch, _ = regexp.Compile(`^\/(\w+)`)
//Tematica[].NombreSerie[].capitulos[]
	pd = pageData{
		"Series Online",
		series.Groups[t].Groups[i].Name,
		series.Groups[t].Groups[i].Groups[0].Stations[o].Name, //1x1 3 a.m.
		series.Groups[t].Groups[i].Groups[0].Stations[o].Image,  //1x1 3 a.m.
		series.Groups[t].Groups[i].Groups[0].Stations[o].URL,
		chtml,
		lfoto,
	}
}
// 1. Subir ficheros /upload
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Funciona /subir")
	http.ServeFile(w,r,"upload.html")
}

//2. Subir ficheros
func uploader(w http.ResponseWriter, r *http.Request) {
//limita el tamaño de los archivos a subir
	r.ParseMultipartForm(2000)

	if r.Method == http.MethodPost { 
	
		file, fileinfo, err := r.FormFile("archivo")
// esto es muy importante es la ruta, tiene que terminar en barra "/"
		f,err := os.OpenFile("./files/"+fileinfo.Filename,os.O_WRONLY|os.O_CREATE, 0666)
		
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "error al subir %v",err)
			return
		}

		defer f.Close()

		io.Copy(f,file)	
		fmt.Fprintf(w, "Cargado con exito "+fileinfo.Filename)
	}
}


func main() {
	
	azaro=azar()
	a(azaro)
	b()
	c(azaro)
	//tengo que añadir el api arriba
	
	mux := http.NewServeMux()
	go mux.HandleFunc("/", root) //no lo he tocado
	go mux.HandleFunc("/pelis", pelis) //peliculas
	go mux.HandleFunc("/series", fseries) //series
	
	// Directorio publico contiene jsonseries.json
	FileServer := http.FileServer(http.Dir("api"))
	go mux.Handle("/api/",http.StripPrefix("/api/",FileServer))
    // Subir fuchero
	mux.HandleFunc("/subir",handler)  //1. parte
	mux.HandleFunc("/files",uploader) //2. parte


port := os.Getenv("PORT")

if port ==""{
port = "8000"
}

// escuchando
server := &http.Server{
	Addr: 			":"+port,
	Handler:		mux,
	ReadTimeout:	10 * time.Second,
	WriteTimeout:	10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}

log.Println("Server Active port: "+port)
log.Fatal(server.ListenAndServe())

}
