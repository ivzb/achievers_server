package server

type Info struct {
	Version        string `json:"Version"`
	Hostname       string `json:"Hostname"`
	UseHTTP        bool   `json:"UseHTTP"`
	UseHTTPS       bool   `json:"UseHTTPS"`
	HTTPPort       int    `json:"HTTPPort"`
	HTTPPorts      int    `json:"HTTPPorts"`
	FileStorage    string `json:"FileStorage"`
	CertFile       string `json:"CertFile"`
	KeyFile        string `json:"KeyFile"`
	MaxBytesReader int64  `json:"MaxBytesReader"`
}
