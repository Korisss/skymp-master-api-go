package handler

import (
	"net/http"
	"net/mail"

	"github.com/Korisss/skymp-master-api-go/internal/domain"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type requestWithAuth struct {
	Id int `uri:"id" binding:"required"`
}

func checkUserAccess(ctx *gin.Context) (int, bool) {
	ctxId, _ := ctx.Get("id")
	id := ctxId.(int)

	var req requestWithAuth

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return id, false
	}

	if id != req.Id {
		newErrorResponse(ctx, http.StatusForbidden, "no access to create session for this user")
		return id, false
	}

	return id, true
}

// Returning user name
func (h *Handler) getUserName(ctx *gin.Context) {
	id, access := checkUserAccess(ctx)
	if !access {
		return
	}

	name, err := h.services.GetUserName(id)
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
