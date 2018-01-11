package server

type Info struct {
	Hostname  string `json:"Hostname"`
	UseHTTP   bool   `json:"UseHTTP"`
	UseHTTPS  bool   `json:"UseHTTPS"`
	HTTPPort  int    `json:"HTTPPort"`
	HTTPPorts int    `json:"HTTPPorts"`
	CertFile  string `json:"CertFile"`
	KeyFile   string `json:"KeyFile"`
}
