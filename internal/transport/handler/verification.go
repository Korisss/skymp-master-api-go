package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type verify struct {
// 	Code int `json:"code" binding:"required"`
// }

// type sendVerificationCodeReq struct {
// 	Discord string `json:"discord" binding:"required"`
// }

// func (h *Handler) sendVerificationCode(ctx *gin.Context) {
// 	id, access := checkUserAccess(ctx)
// 	if !access {
// 		newErrorResponse(ctx, http.StatusUnauthorized, "no access")
// 		return
// 	}

// 	var req sendVerificationCodeReq

// 	if err := ctx.BindJSON(&req); err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	err := h.services.SendCodeToBot(id, req.Discord)

// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// }

// func (h *Handler) verify(ctx *gin.Context) {
// 	id, access := checkUserAccess(ctx)
// 	if !access {
// 		newErrorResponse(ctx, http.StatusUnauthorized, "no access")
// 		return
// 	}

// 	var req verify

// 	if err := ctx.BindJSON(&req); err != nil {
// 		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	code, err := h.services.GetVerificationCode(id)
// 	if err != nil {
// 		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	if code != req.Code {
// 		newErrorResponse(ctx, http.StatusUnauthorized, "verification code is invalid")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"successful": true,
// 	})
// }
