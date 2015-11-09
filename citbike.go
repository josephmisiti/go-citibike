package citibike

import(
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    // "os"
    // "runtime"
)


const (
    VERSION = 1
    CITI_BIKE_JSON = "https://www.citibikenyc.com/stations/json"
)

type Station struct {
    Id int64 `json:"id"`
    StationName string `json:"stationName"`
    AvailableDocks int64 `json:"availableDocks"`
    TotalDocks int64 `json:"totalDocks"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    StatusValue string `json:"statusValue"`
    StatusKey int64 `json:"statusKey"`
    AvailableBikes int64 `json:"availableBikes"`
    StAddress1 string `json:"stAddress1"`
    StAddress2 string `json:"stAddress2"`
    City string `json:"city"`
    PostalCode string `json:"postalCode"`
    Location string `json:"location"`
    Altitude string `json:"altitude"`
    TestStation bool `json:"testStation"`
    LastCommunicationTime string `json:"lastCommunicationTime"`
    LandMark string `json:"landMark"`
}

type StationsResponse struct {
    ExecutionTime string `json:"executionTime"`
    StationBeanList []Station `json:"stationBeanList"`
}

type API struct {

}

func (r API) GetStations() (*StationsResponse,error) {

    body, err := MakeRequest("https://www.citibikenyc.com/stations/json")
    if (err != nil) {
        return nil, err
    }
    s, err := ParseStations(body)
    return  s, err
}


func MakeRequest(url string) ([]byte, error) {
    
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    
    return []byte(body), err
}


func ParseStations(body []byte) (*StationsResponse, error) {
    var s = new(StationsResponse)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}