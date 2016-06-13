package http

import (
	"encoding/json"
	"geo"
	"io"
	"ip"
	"net/http"
	"log"
)

type Server struct {
	listen string
	geo    *geo.Geo
}

type (
	AvailableMethods struct {
		Methods []string `json:"methods"`
	}
	AllResponse struct {
		IpAddress string  `json:"ip_address"`
		UserAgent string  `json:"user_agent"`
		GeoLoc    geo.Loc `json:"geo_location"`
	}
	IpResponse struct {
		IpAddress string `json:"ip_address"`
	}
	UaResponse struct {
		Ua string `json:"ua"`
	}
	GeoResponse struct {
		Geo geo.Loc `json:"geo_location`
	}
)

func New(listen string) *Server {
	g, err := geo.New()
	if err != nil {
		log.Fatalf("Error creating geoip: %s", err)
	}
	return &Server{
		listen: listen,
		geo:    g,
	}
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.errorNotFoundHandler(w, r)
		return
	}
	r.Header.Set("Content-Type", "application/json")
	methods := AvailableMethods{Methods: []string{"/all", "/ip", "/ua", "/geo"}}
	io.WriteString(w, s.encodeOutput(methods))
}

func (s *Server) allHandler(w http.ResponseWriter, r *http.Request) {
	ip := ip.GetIp(r)
	all := AllResponse{ip, r.UserAgent(), s.geo.GetLoc(ip)}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, s.encodeOutput(all))
}

func (s *Server) ipHandler(w http.ResponseWriter, r *http.Request) {
	ipResponse := IpResponse{ip.GetIp(r)}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, s.encodeOutput(ipResponse))
}

func (s *Server) uaHandler(w http.ResponseWriter, r *http.Request) {
	uaResponse := UaResponse{r.UserAgent()}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, s.encodeOutput(uaResponse))
}

func (s *Server) geoHandler(w http.ResponseWriter, r *http.Request) {
	ip := ip.GetIp(r)
	geoResponse := GeoResponse{s.geo.GetLoc(ip)}
	r.Header.Set("Content-Type", "application/json")
	io.WriteString(w, s.encodeOutput(geoResponse))
}

func (s *Server) errorNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 Page Not Found")
}

func (s *Server) encodeOutput(content interface{}) string {
	output, err := json.Marshal(content)
	if err == nil {
		return string(output)
	}
	return ""
}

func (s *Server) Start() {
	http.HandleFunc("/", s.homeHandler)
	http.HandleFunc("/all", s.allHandler)
	http.HandleFunc("/ip", s.ipHandler)
	http.HandleFunc("/ua", s.uaHandler)
	http.HandleFunc("/geo", s.geoHandler)
	http.ListenAndServe(s.listen, nil)
}
