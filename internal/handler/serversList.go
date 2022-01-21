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

type serversList []master_api.SkyMPServer

func (sl *serversList) searchServerInList(ip string, port int) (int, *master_api.SkyMPServer) {
	for i := 0; i < len(*sl); i++ {
		if (*sl)[i].Ip == ip && (*sl)[i].Port == port {
			return i, &((*sl)[i])
		}
	}

	return 0, nil
}

func (sl *serversList) checkServerTimeout(ip string, port int) {
	if i, serverInList := sl.searchServerInList(ip, port); serverInList != nil && time.Since(serverInList.UpdatedAt) > 5*time.Second {
		if len(*sl) <= 1 {
			*sl = make(serversList, 0)
		} else if i == len(*sl)-1 {
			*sl = (*sl)[:len(*sl)-1]
		} else {
			(*sl)[i] = (*sl)[len(*sl)-1]
			*sl = (*sl)[:len(*sl)-1]
		}
	}
}

var servers serversList = make(serversList, 0)

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

	if _, serverInList := servers.searchServerInList(ip, port); serverInList != nil {
		serverInList.UpdatedAt = time.Now()
		return
	}

	var request struct {
		Name       string `json:"name" binding:"required"`
		MaxPlayers int    `json:"maxPlayers" binding:"required"`
		Online     int    `json:"online"` // TODO: use binding:"required"
	}

	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	server := master_api.NewSkyMPServer(ip, port, request.Name, request.MaxPlayers, request.Online)

	servers = append(servers, server)

	go func() {
		time.Sleep(10 * time.Second)

		servers.checkServerTimeout(ip, port)
	}()
}

func (h *Handler) getServersList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &servers)
}
