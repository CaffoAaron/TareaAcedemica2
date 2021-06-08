package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type knnNode struct {
	Distancia float64
	x         int
	y         int
	estado    string
}

type Respuesta struct {
	Mensaje string
}

type ConsultaBono struct {
	Casado                         bool `json:"casado"`
	Hijos                          bool `json:"hijos"`
	CarreraUniversitaria           bool `json:"carrera_universitaria"`
	CasaPropia                     bool `json:"casa_propia"`
	OtroPrestamo                   bool `json:"otro_prestamo"`
	Mas_4_Años                     bool `json:"mas_de_4_Años_como_empresa"`
	Mas_1_Local                    bool `json:"mas_de_1_Local"`
	Mas_10_Empleados               bool `json:"mas_de_10_Empleados"`
	PagoIgv_6_Meses                bool `json:"Pago_de_Igv_Ultimos_6_Meses"`
	DeclaronConfidencialPatrimonio bool `json:"declaron_confidencial_patrimonio"`

	PuntajePersonal int
	PuntajeEmpresa  int
	Estado          string
}

var Dataset = [1000]ConsultaBono{}

func LeerDataSetFromGit() {
	response, err := http.Get("https://raw.githubusercontent.com/CaffoAaron/DataSet-Programaci-n-Concurrente-y-Distribuida/master/bono_Independiente_trabajaperu.csv") //use package "net/http"
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	reader := csv.NewReader(response.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(data)

	for i, row := range data {

		Casado, _ := strconv.ParseBool(row[0])
		Dataset[i].Casado = Casado

		Hijos, _ := strconv.ParseBool(row[1])
		Dataset[i].Hijos = Hijos

		CarreraUniversitaria, _ := strconv.ParseBool(row[2])
		Dataset[i].CarreraUniversitaria = CarreraUniversitaria

		CasaPropia, _ := strconv.ParseBool(row[3])
		Dataset[i].CasaPropia = CasaPropia

		OtroPrestamo, _ := strconv.ParseBool(row[4])
		Dataset[i].OtroPrestamo = OtroPrestamo

		Mas_4_Años, _ := strconv.ParseBool(row[5])
		Dataset[i].Mas_4_Años = Mas_4_Años

		Mas_1_Local, _ := strconv.ParseBool(row[6])
		Dataset[i].Mas_1_Local = Mas_1_Local

		Mas_10_Empreados, _ := strconv.ParseBool(row[7])
		Dataset[i].Mas_10_Empleados = Mas_10_Empreados

		PagoIgv_6_Meses, _ := strconv.ParseBool(row[8])
		Dataset[i].PagoIgv_6_Meses = PagoIgv_6_Meses

		DeclaronConfidencialPatrimonio, _ := strconv.ParseBool(row[9])
		Dataset[i].DeclaronConfidencialPatrimonio = DeclaronConfidencialPatrimonio
	}
	for i := 0; i < 1000; i++ {
		getEstado(&Dataset[i])
	}
	log.Println(Dataset)
}

func LeerDataSet() {
	data := "bono_Independiente_trabajaperu.csv"
	var i = 0
	file, err := os.Open(data)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	i = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
		}
		Casado, _ := strconv.ParseBool(record[0])
		Dataset[i].Casado = Casado

		Hijos, _ := strconv.ParseBool(record[1])
		Dataset[i].Hijos = Hijos

		CarreraUniversitaria, _ := strconv.ParseBool(record[2])
		Dataset[i].CarreraUniversitaria = CarreraUniversitaria

		CasaPropia, _ := strconv.ParseBool(record[3])
		Dataset[i].CasaPropia = CasaPropia

		OtroPrestamo, _ := strconv.ParseBool(record[4])
		Dataset[i].OtroPrestamo = OtroPrestamo

		Mas_4_Años, _ := strconv.ParseBool(record[5])
		Dataset[i].Mas_4_Años = Mas_4_Años

		Mas_1_Local, _ := strconv.ParseBool(record[6])
		Dataset[i].Mas_1_Local = Mas_1_Local

		Mas_10_Empreados, _ := strconv.ParseBool(record[7])
		Dataset[i].Mas_10_Empleados = Mas_10_Empreados

		PagoIgv_6_Meses, _ := strconv.ParseBool(record[8])
		Dataset[i].PagoIgv_6_Meses = PagoIgv_6_Meses

		DeclaronConfidencialPatrimonio, _ := strconv.ParseBool(record[9])
		Dataset[i].DeclaronConfidencialPatrimonio = DeclaronConfidencialPatrimonio

		i++
	}
	for i := 0; i < 1000; i++ {
		getEstado(&Dataset[i])
	}
	log.Println(Dataset)
}
func getEstado(p *ConsultaBono) {
	contPersonas := 0
	contEmpresa := 0

	if p.Casado == true {
		contPersonas += 3
	}
	if p.Hijos == false {
		contPersonas += 1
	}
	if p.CarreraUniversitaria == true {
		contPersonas += 3
	}
	if p.CasaPropia == true {
		contPersonas += 4
	}
	if p.OtroPrestamo == false {
		contPersonas += 2
	}
	if p.Mas_4_Años == true {
		contEmpresa += 2
	}
	if p.Mas_1_Local == true {
		contEmpresa += 4
	}
	if p.Mas_10_Empleados == true {
		contEmpresa += 4

	}
	if p.PagoIgv_6_Meses == true {
		contEmpresa += 1
	}
	if p.DeclaronConfidencialPatrimonio == true {
		contEmpresa += 1
	}

	p.PuntajeEmpresa = contPersonas
	p.PuntajePersonal = contEmpresa

	if p.PuntajeEmpresa+p.PuntajePersonal > 15 {
		p.Estado = "Pre-Aprobado"
	}
	if p.PuntajeEmpresa+p.PuntajePersonal <= 15 {
		p.Estado = "Denegado"
	}
}

func proccesofChossing(k *knnNode, x int, y int, p ConsultaBono) {
	absX := math.Abs(float64(x - p.PuntajeEmpresa))
	absY := math.Abs(float64(y - p.PuntajePersonal))
	distancia := math.Sqrt(math.Pow(absX, 2) + math.Pow(absY, 2))
	k.Distancia = distancia
	k.x = p.PuntajeEmpresa
	k.y = p.PuntajePersonal
	k.estado = p.Estado
}

func knn(usuario *ConsultaBono) bool {
	var getPoints = [100]knnNode{}

	for i := 0; i < 100; i++ {
		go proccesofChossing(&getPoints[i], usuario.PuntajeEmpresa, usuario.PuntajePersonal, Dataset[i])
		time.Sleep(30)
	}
	log.Println(getPoints)
	for i := 1; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			if getPoints[j].Distancia > getPoints[j+1].Distancia {
				getPoints[j], getPoints[j+1] = getPoints[j+1], getPoints[j]
			}
		}
	}
	log.Println(getPoints)
	count := 0
	for i := 0; i < 6; i++ {
		if getPoints[i].estado == "Pre-Aprobado" {
			count++
		}
	}
	if count >= 3 {
		log.Println("Usted esta preaprobado para el bono independiente")
		return true
	} else {
		log.Println("Usted no esta apto para el bono independiente")
		return false
	}
}

func mostrarDataset(res http.ResponseWriter, req *http.Request) {
	log.Println("Llamada al endpoint /dataset")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(Dataset, "", "\t")
	io.WriteString(res, string(jsonBytes))
}

func realizarKnn(res http.ResponseWriter, req *http.Request) {
	log.Println("Llamada al endpoint /knn")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	var usuario = ConsultaBono{}
	var respuesta = Respuesta{}
	usuario.Hijos = true
	usuario.CarreraUniversitaria = true
	usuario.CasaPropia = true
	usuario.OtroPrestamo = false
	usuario.Mas_4_Años = true
	usuario.Mas_1_Local = true
	usuario.Mas_10_Empleados = true
	usuario.PagoIgv_6_Meses = true
	usuario.DeclaronConfidencialPatrimonio = true
	getEstado(&usuario)
	RespuestaKnn := knn(&usuario)
	if RespuestaKnn == true {
		respuesta.Mensaje = "Usted esta preaprobado para el bono independiente"
	} else {
		respuesta.Mensaje = "Usted no esta apto para el bono independiente"
	}
	jsonBytes, _ := json.MarshalIndent(respuesta, "", "\t")
	io.WriteString(res, string(jsonBytes))
}

func handleRequest() {

	http.HandleFunc("/dataset", mostrarDataset)
	http.HandleFunc("/knn", realizarKnn)
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {
	LeerDataSetFromGit()
	//LeerDataSet()
	handleRequest()
}
