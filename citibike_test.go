package citibike

import (
	"testing"
)

func TestVersion(t *testing.T) {
	println("Testing version: ", VERSION)
}

func TestMakeRequest(t *testing.T) {

    body, err := MakeRequest("https://www.citibikenyc.com/stations/json")
    if (err != nil) {
         t.Fatalf("MakeRequest failed")  
    }
    
    s, err := ParseStations(body)
    if (len(s.StationBeanList) < 1) {
         t.Fatalf("it looks like your api call failed")  
    }

}

func TestParsing(t *testing.T) {
    http_body := `
        {
        "executionTime": "2015-11-09 01:39:40 PM",
        "stationBeanList": [
            {
            "id": 72,
            "stationName": "W 52 St & 11 Ave",
            "availableDocks": 37,
            "totalDocks": 39,
            "latitude": 40.76727216,
            "longitude": -73.99392888,
            "statusValue": "In Service",
            "statusKey": 1,
            "availableBikes": 1,
            "stAddress1": "W 52 St & 11 Ave",
            "stAddress2": "",
            "city": "",
            "postalCode": "",
            "location": "",
            "altitude": "",
            "testStation": false,
            "lastCommunicationTime": "2015-11-09 01:38:46 PM",
            "landMark": ""
            }]
        }
        `
    s, err := ParseStations([]byte(http_body))
    if(err != nil) {
        t.Fatalf("getStations failed")  
    }
    
    station := s.StationBeanList[0]
    
    if(s.ExecutionTime != "2015-11-09 01:39:40 PM"){
        t.Fatalf("executionTime parsed incorrectly")
    }
    if(station.Id != 72){
        t.Fatalf("station.id parsed incorrectly")
    }
    if(station.StationName != "W 52 St & 11 Ave"){
        t.Fatalf("station.stationName parsed incorrectly")
    }
    if(station.AvailableDocks != 37){
        t.Fatalf("station.availableDocks parsed incorrectly")
    }
    if(station.Latitude != 40.76727216){
        t.Fatalf("station.latitude parsed incorrectly")
    }
    if(station.Longitude != -73.99392888){
        t.Fatalf("station.id parsed incorrectly")
    }
    if(station.Id != 72){
        t.Fatalf("station.id parsed incorrectly")
    }
}


