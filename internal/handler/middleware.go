package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "id"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(header)

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}
