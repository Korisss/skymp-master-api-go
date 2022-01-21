package handler

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/gin-gonic/gin"
)

var serversList = make([]master_api.SkyMPServer, 0)

func (h *Handler) addOrUpdateServer(ctx *gin.Context) {
	var req struct {
		Address string `uri:"address" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ipAndPort := strings.Split(req.Address, ":")
	ip := ipAndPort[0]
	port, err := strconv.Atoi(ipAndPort[1])

	if net.ParseIP(ip) == nil || err != nil || port > 65535 || port < 0 {
		newErrorResponse(ctx, http.StatusBadRequest, "bad address")
		return
	}

	if _, serverInList := searchServerInList(ip, port); serverInList != nil {
		serverInList.UpdatedAt = time.Now()
		return
	}

	var request struct {
		Name       string `json:"name" binding:"required"`
		MaxPlayers int    `json:"maxPlayers" binding:"required"`
		Online     int    `json:"online" binding:"required"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	server := master_api.NewSkyMPServer(ip, port, request.Name, request.MaxPlayers, request.Online)

	serversList = append(serversList, server)

	go func() {
		time.Sleep(10 * time.Second)

		checkServerTimeout(ip, port)
	}()
}

func searchServerInList(ip string, port int) (int, *master_api.SkyMPServer) {
	for i := 0; i < len(serversList); i++ {
		if serversList[i].Ip == ip && serversList[i].Port == port {
			return i, &(serversList[i])
		}
	}

	return 0, nil
}

func checkServerTimeout(ip string, port int) {
	if i, serverInList := searchServerInList(ip, port); serverInList != nil && time.Since(serverInList.UpdatedAt) > 5*time.Second {
		if len(serversList) <= 1 {
			serversList = make([]master_api.SkyMPServer, 0)
		} else if i == len(serversList)-1 {
			serversList = serversList[:len(serversList)-1]
		} else {
			serversList[i] = serversList[len(serversList)-1]
			serversList = serversList[:len(serversList)-1]
		}
	}
}

func (h *Handler) getServersList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &serversList)
}

func (h *Handler) getSessionData(ctx *gin.Context) {}
