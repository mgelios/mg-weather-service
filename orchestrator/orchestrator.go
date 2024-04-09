package orchestrator

import (
	"encoding/json"
	"fmt"
	"mg-weather-service/model"
	"mg-weather-service/storage"
	"net/http"
	"reflect"
	"time"
)

func GetWeather(city string) model.OneCallWeather {
	geoResp := GetGeocodingResponse(city)

	oneCallResp := GetOneCallResponse(geoResp.Lat, geoResp.Lon)

	return oneCallResp
}

func GetOneCallResponse(lat float64, lon float64) model.OneCallWeather {
	appid := ""
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=hourly,minutely&appid=%s", lat, lon, appid))
	if err != nil {
		panic(err)
	}
	var oneCallResp model.OneCallWeather
	err = json.NewDecoder(res.Body).Decode(&oneCallResp)
	if err != nil {
		panic(err)
	}
	println("One Call resp")
	fmt.Printf("%v \r\n", oneCallResp)
	defer res.Body.Close()
	return oneCallResp
}

func GetGeoCoding(city string) model.GeoCoding {
	geoRespItem := storage.GetGeoCodingRecordByCity(city)
	if reflect.ValueOf(geoRespItem).IsNil() {
		geoRespItem = GetGeoCoding(city)
		storage.PutGeoCodingRecord(geoRespItem)
	}
	return geoRespItem
}

func GetGeocodingResponse(city string) model.GeoCoding {
	var geoResp []model.GeoCoding
	appid := ""
	res, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, appid))
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(res.Body).Decode(&geoResp)
	if err != nil {
		panic(err)
	}
	println("Geo: resp")
	fmt.Printf("%v \r\n", geoResp)
	println(len(geoResp))
	defer res.Body.Close()
	geoResp[0].Created = time.Now()
	geoResp[0].QueredCity = city
	return geoResp[0]
}
