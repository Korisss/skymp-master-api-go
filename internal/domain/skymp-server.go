package domain

import "time"

type SkyMPServer struct {
	Ip         string    `json:"ip"`
	Port       int       `json:"port"`
	Name       string    `json:"name"`
	MaxPlayers int       `json:"maxPlayers"`
	Online     int       `json:"online"`
	UpdatedAt  time.Time `json:"-"`
}

func NewSkyMPServer(ip string, port int, name string, maxPlayers, online int) SkyMPServer {
	return SkyMPServer{
		Ip:         ip,
		Port:       port,
		Name:       name,
		MaxPlayers: maxPlayers,
		Online:     online,
		UpdatedAt:  time.Now(),
	}
}
