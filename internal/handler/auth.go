package handler

import (
	"net/http"

	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Returning user name
func (h *Handler) getUserName(ctx *gin.Context) {
	var input struct {
		Id int `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	name, err := h.services.GetUserName(input.Id)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"name": name,
	})
}

func (h *Handler) login(ctx *gin.Context) {
	var input loginInput

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, token, err := h.services.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"token": token,
	})
}

func (h *Handler) register(ctx *gin.Context) {
	var input master_api.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) resetPassword(ctx *gin.Context) {}

func (h *Handler) createSession(ctx *gin.Context) {}
