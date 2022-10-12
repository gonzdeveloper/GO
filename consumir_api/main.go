package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type ResponsePokemon struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

type ResponseClima struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Temp_min   float64 `json:"temp_min"`
	Temp_max   float64 `json:"temp_max"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
	Sea_level  int     `json:"sea_level"`
	Grnd_level int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

// api.openweathermap.org/data/2.5/weather?lat=-30.8335&lon=-54.1674&cnt=5&appid=cc1a974e19551a68d8b555682cd871e3
// api.openweathermap.org/data/2.5/weather?q=Montevideo,uy&APPID=cc1a974e19551a68d8b555682cd871e3
func main() {

	codigos_paises := []Coord{{Lat: -34.7869983, Lon: -55.4757374}, {Lat: -34.6158037, Lon: -58.5033386}, {Lat: -33.6158037, Lon: -57.5033386},
		{Lat: -32.6158037, Lon: -56.5033386}, {Lat: -31.6158037, Lon: -55.5033386}}

	temp := ""

	fmt.Println(codigos_paises)

	for _, pais := range codigos_paises {
		fmt.Println("URL:", "http://api.openweathermap.org/data/2.5/weather?lat="+fmt.Sprintf("%v", pais.Lat)+"&lon="+fmt.Sprintf("%v", pais.Lon)+"&cnt=5&appid=cc1a974e19551a68d8b555682cd871e3")
		response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?lat=" + fmt.Sprintf("%v", pais.Lat) + "&lon=" + fmt.Sprintf("%v", pais.Lon) + "&cnt=5&appid=cc1a974e19551a68d8b555682cd871e3")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Con coordenadas:\n", string(responseData))

		var responseObjectClima ResponseClima
		json.Unmarshal(responseData, &responseObjectClima)

		fmt.Println(responseObjectClima)
		fmt.Println(responseObjectClima.Sys.Country)
		temp += "\n" + fmt.Sprintf("%v", responseObjectClima.Name) + ", " + fmt.Sprintf("%v", responseObjectClima.Sys.Country) + ";" + fmt.Sprintf("%v", responseObjectClima.Main.Temp_max) + ";" + fmt.Sprintf("%v", responseObjectClima.Coord.Lat) + ";" + fmt.Sprintf("%v", responseObjectClima.Coord.Lon)

	}

	fmt.Println(temp)

}
