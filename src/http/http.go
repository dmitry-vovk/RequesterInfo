package http

import (
	"encoding/json"
	"geo"
	"io"
	"ip"
	"net/http"
)

type AvailableMethods struct {
	Methods []string `json:"methods"`
}
type AllResponse struct {
	IpAddress string  `json:"ip_address"`
	UserAgent string  `json:"user_agent"`
	GeoLoc    geo.Loc `json:"geo_location"`
}
type IpResponse struct {
	IpAddress string `json:"ip_address"`
}
type UaResponse struct {
	Ua string `json:"ua"`
}
type GeoResponse struct {
	Geo geo.Loc `json:"geo_location`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorNotFoundHandler(w, r)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	methods := AvailableMethods{Methods: []string{"/all", "/ip", "/ua", "/geo"}}
	io.WriteString(w, encodeOutput(methods))
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	ip := ip.GetIp(r)
	all := AllResponse{ip, r.UserAgent(), geo.GetLoc(ip)}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, encodeOutput(all))
}

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ipResponse := IpResponse{ip.GetIp(r)}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, encodeOutput(ipResponse))
}

func uaHandler(w http.ResponseWriter, r *http.Request) {
	uaResponse := UaResponse{r.UserAgent()}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, encodeOutput(uaResponse))
}

func geoHandler(w http.ResponseWriter, r *http.Request) {
	ip := ip.GetIp(r)
	geoResponse := GeoResponse{geo.GetLoc(ip)}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, encodeOutput(geoResponse))
}

func errorNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 Page Not Found")
}

func encodeOutput(content interface{}) string {
	output, err := json.Marshal(content)
	if err == nil {
		return string(output)
	}
	return ""
}

func Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/all", allHandler)
	http.HandleFunc("/ip", ipHandler)
	http.HandleFunc("/ua", uaHandler)
	http.HandleFunc("/geo", geoHandler)
	http.ListenAndServe(":8000", nil)
}
