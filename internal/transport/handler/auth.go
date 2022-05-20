package handler

import (
	"net/http"
	"net/mail"
	"strconv"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type requestWithId struct {
	Id string `uri:"id" binding:"required"`
}

func checkUserAccess(ctx *gin.Context) (int64, bool) {
	ctxId, _ := ctx.Get("id")
	id := ctxId.(int64)

	var req requestWithId

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return id, false
	}

	reqId, err := strconv.Atoi(req.Id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return id, false
	}

	if id != int64(reqId) {
		newErrorResponse(ctx, http.StatusForbidden, "no access to create session for this user")
		return id, false
	}

	return id, true
}

// Returning user name
func (h *Handler) getUserName(ctx *gin.Context) {
	var req requestWithId

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	reqId, err := strconv.Atoi(req.Id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	name, err := h.services.GetUserName(int64(reqId))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func (h *Handler) login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, token, err := h.services.GenerateToken(req.Email, req.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"token": token,
	})
}

func (h *Handler) register(ctx *gin.Context) {
	var req domain.User

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(req.Password) < 6 {
		newErrorResponse(ctx, http.StatusBadRequest, "password must contain at least 6 characters")
		return
	}

	id, err := h.services.CreateUser(req)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// TODO: create with verification
func (h *Handler) resetPassword(ctx *gin.Context) {}
