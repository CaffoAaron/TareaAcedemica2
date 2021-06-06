package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type knnNode struct {
	distancia float64
	x         int
	y         int
	estado    string
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
	for i := 0; i < 100; i++ {
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

func mostrarDataset(res http.ResponseWriter, req *http.Request) {
	log.Println("Llamada al endpoint /dataset")
	log.Println(Dataset)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonBytes, _ := json.MarshalIndent(Dataset, "", "\t")
	log.Println(string(jsonBytes))
	io.WriteString(res, string(jsonBytes))
}

func realizarKnn(res http.ResponseWriter, req *http.Request) {
	log.Println("Llamada al endpoint /knn")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonBytes2, _ := json.MarshalIndent(Dataset, "", "\t")
	log.Println(string(jsonBytes2))
	io.WriteString(res, string(jsonBytes2))
}

func handleRequest() {

	http.HandleFunc("/dataset", mostrarDataset)
	http.HandleFunc("/knn", realizarKnn)
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {
	LeerDataSet()
	handleRequest()
}
