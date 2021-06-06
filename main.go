package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
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
	Casado                         bool `json:"fever"`
	Hijos                          bool `json:"tiredness"`
	CarreraUniversitaria           bool `json:"dryCough"`
	CasaPropia                     bool `json:"difficultyBrithing"`
	OtroPrestamo                   bool `json:"difficultyBrithing"`
	Mas_4_Años                     bool `json:"soreThroat"`
	Mas_1_Local                    bool `json:"noneSymtons"`
	Mas_10_Empreados               bool `json:"age0_9"`
	PagoIgv_6_Meses                bool `json:"age10_19"`
	DeclaronConfidencialPatrimonio bool `json:"age20_24"`

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
		Dataset[i].Mas_10_Empreados = Mas_10_Empreados

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
	if p.Mas_10_Empreados == true {
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

func main() {
	LeerDataSet()
}
