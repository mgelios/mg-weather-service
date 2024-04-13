package orchestrator

import (
	"encoding/json"
	"fmt"
	"mg-weather-service/model"
	"mg-weather-service/storage"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetWeather(city string) model.OneCallWeather {
	geoCoding := GetGeoCoding(city)
	oneCall := GetOneCall(geoCoding)
	return oneCall
}

func GetOneCall(geoCoding model.GeoCoding) model.OneCallWeather {
	oneCallResp, err := storage.GetOneCallRecordByCity(geoCoding.Name)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			oneCallResp = GetOneCallResponse(geoCoding)
			storage.PutOneCallRecord(oneCallResp)
		}
	} else {
		if oneCallResp.Created.Before(time.Now().AddDate(0, 0, -1)) {
			newOneCallResp := GetOneCallResponse(geoCoding)
			storage.UpdateOneCallRecord(oneCallResp.Id, newOneCallResp)
		}
	}
	return oneCallResp
}

func GetOneCallResponse(geoCoding model.GeoCoding) model.OneCallWeather {
	appid := ""
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=hourly,minutely&appid=%s",
		geoCoding.Lat, geoCoding.Lon, appid))
	if err != nil {
		panic(err)
	}
	var oneCallResp model.OneCallWeather
	err = json.NewDecoder(res.Body).Decode(&oneCallResp)
	if err != nil {
		panic(err)
	}
	println("One Call resp")
	defer res.Body.Close()
	oneCallResp.Created = time.Now()
	oneCallResp.City = geoCoding.Name
	return oneCallResp
}

func GetGeoCoding(city string) model.GeoCoding {
	geoRespItem, err := storage.GetGeoCodingRecordByCity(city)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			geoRespItem = GetGeocodingResponse(city)
			storage.PutGeoCodingRecord(geoRespItem)
		}
	} else {
		if geoRespItem.Created.Before(time.Now().AddDate(0, 0, -1)) {
			newGeoRespItem := GetGeocodingResponse(city)
			storage.UpdateGeoCodingRecord(geoRespItem.Id, newGeoRespItem)
		}
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
	defer res.Body.Close()
	geoResp[0].Created = time.Now()
	geoResp[0].QueredCity = city
	return geoResp[0]
}
