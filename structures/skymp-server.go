package master_api

type SkyMPServer struct {
	Ip         string `json:"ip"`
	Port       int    `json:"port"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"maxPlayers"`
	Online     int    `json:"online"`
}
