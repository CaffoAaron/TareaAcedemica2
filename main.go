package main

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

func main() {

}
