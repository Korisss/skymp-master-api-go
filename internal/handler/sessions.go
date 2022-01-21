package handler

import (
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/Korisss/skymp-master-api-go/pkg/random"
	"github.com/gin-gonic/gin"
)

type playerSession struct {
	Address string `uri:"address" binding:"required"`
	Session string `uri:"session" binding:"required"`
}

var sessionsStorage = make(map[int]playerSession)

func findUserInSessionsStorage(session playerSession) int {
	for i, s := range sessionsStorage {
		if s == session {
			return i
		}
	}

	return 0
}

// TODO: check id with token match
func (h *Handler) createSession(ctx *gin.Context) {
	_, access := checkUserAccess(ctx)

	if !access {
		return
	}

	var req struct {
		Id      int    `uri:"id" binding:"required"`
		Address string `uri:"serverAddress" binding:"required"`
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

	if _, server := servers.searchServerInList(ip, port); server == nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid server")
		return
	}

	session := playerSession{
		Address: req.Address,
		Session: random.RandString(32),
	}

	sessionsStorage[req.Id] = session

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"session": session.Session,
	})
}

func (h *Handler) getSessionData(ctx *gin.Context) {
	var req playerSession

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := findUserInSessionsStorage(req)

	if userId == 0 {
		newErrorResponse(ctx, http.StatusNotFound, "user not found")
		return
	}

	ctx.JSON(http.StatusOK, map[string]map[string]interface{}{
		"user": {
			"id": userId,
		},
	})
}
