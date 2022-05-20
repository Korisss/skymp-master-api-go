package handler

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/gin-gonic/gin"
)

var servers = make(serversList, 0)

type serversList []domain.SkyMPServer

func (sl *serversList) searchServerInList(ip string, port int) (int, *domain.SkyMPServer) {
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

	var request struct {
		Name       string `json:"name" binding:"required"`
		MaxPlayers int    `json:"maxPlayers" binding:"required"`
		Online     int    `json:"online"` // binding:"required"
	}

	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if _, serverInList := servers.searchServerInList(ip, port); serverInList != nil {
		serverInList.UpdatedAt = time.Now()
		serverInList.Online = request.Online
		return
	}

	server := domain.NewSkyMPServer(ip, port, request.Name, request.MaxPlayers, request.Online)

	servers = append(servers, server)

	go func() {
		time.Sleep(10 * time.Second)

		servers.checkServerTimeout(ip, port)
	}()
}

func (h *Handler) getServersList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &servers)
}
