package routes

import (
	"net/http"
	"time"
	"weather/parser"
	"weather/render"
)

type Static struct {
	Template render.Template
}

type Sensor struct {
	Temperature0  float32
	Temperature1  float32
	Temperature2  float32
	Temperature3  float32
	Temperature4  float32
	Temperature5  float32
	Humidity0     float32
	Humidity1     float32
	Humidity2     float32
	Humidity3     float32
	Humidity4     float32
	Humidity5     float32
	CurrentDate   string
	PreviousDate1 string
	PreviousDate2 string
	PreviousDate3 string
	PreviousDate4 string
	PreviousDate5 string
}

func data() Sensor {

	now := time.Now()
	format := "2006-01-02"
	Current := now.Format(format)
	var previousDates []string
	temperatures := make([]float32, 6)
	humidities := make([]float32, 6)
	for i := 0; i < 6; i++ {
		previousDates = append(previousDates, now.Add(time.Duration(-24*i)*time.Hour).Format(format))
	}

	for i, date := range previousDates {
		result := parser.Parsing(date)
		temperatures[i] = result.Temperature
		humidities[i] = result.Humidity
	}

	current := parser.Parsing(Current)

	data := Sensor{
		Temperature0:  current.Temperature,
		Temperature1:  temperatures[0],
		Temperature2:  temperatures[1],
		Temperature3:  temperatures[2],
		Temperature4:  temperatures[3],
		Temperature5:  temperatures[4],
		Humidity0:     current.Humidity,
		Humidity1:     humidities[0],
		Humidity2:     humidities[1],
		Humidity3:     humidities[2],
		Humidity4:     humidities[3],
		Humidity5:     humidities[4],
		CurrentDate:   Current,
		PreviousDate1: previousDates[0],
		PreviousDate2: previousDates[1],
		PreviousDate3: previousDates[2],
		PreviousDate4: previousDates[3],
		PreviousDate5: previousDates[4],
	}

	return data
}

func StaticHandler(tpl render.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, data())
	}
}
