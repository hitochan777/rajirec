package rajirec

import (
	"log"
	"encoding/xml"
	"net/http"
	"io/ioutil"
)


func FetchXML(url string, v interface{}) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal("Failed to fetch XML data")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(body, v); err != nil {
		log.Fatal(err)
	}
}

type Area struct {
	Areajp string `xml:"areajp"`
	Area string `xml:"area"`
	R1 string `xml:"r1"`
	R2 string `xml:"r2"`
	Fm string `xml:"fm"`
}

type Areas map[string]Area

func NewAreas(configUrl string) Areas {
	areas := struct {
		Areas []Area `xml:"stream_url>data"`
	}{}
	areaMap := Areas{}
	FetchXML(configUrl, &areas)
	for _, area := range areas.Areas {
		areaMap[area.Area] = area
	}
	return areaMap
}
