package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Korisss/skymp-master-api-go/pkg/random"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var botToken = os.Getenv("BOT_TOKEN")
var botUri = os.Getenv("BOT_TOKEN")

type verify struct {
	Code int `json:"code" binding:"required"`
}

type sendVerificationCodeReq struct {
	Discord string `json:"discord" binding:"required"`
}

type sendCodeToBotReq struct {
	Discord string `json:"discord"`
	Code    int    `json:"code"`
}

func (h *Handler) sendVerificationCode(ctx *gin.Context) {
	id, access := checkUserAccess(ctx)
	if !access {
		newErrorResponse(ctx, http.StatusUnauthorized, "no access")
		return
	}

	var req sendVerificationCodeReq

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	code := random.RandInt(4)

	h.services.SetVerificationCode(id, code)

	if err := sendCodeToBot(req.Discord, code); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func sendCodeToBot(discord string, code int) error {
	body, err := json.Marshal(sendCodeToBotReq{
		Discord: discord,
		Code:    code,
	})
	if err != nil {
		logrus.Error("Error reading request. ", err.Error())
	}

	req, err := http.NewRequest("POST", "http://"+botUri+"/send-code", bytes.NewBuffer(body))
	if err != nil {
		logrus.Error("Error reading request. ", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", botUri)
	req.Header.Set("Authorization", botToken)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Error("Error reading response. ", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		logrus.Fatal("Bot token is invalid")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("status is not ok")
	}

	return nil
}

func (h *Handler) verify(ctx *gin.Context) {
	id, access := checkUserAccess(ctx)
	if !access {
		newErrorResponse(ctx, http.StatusUnauthorized, "no access")
		return
	}

	var req verify

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	code, err := h.services.GetVerificationCode(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if code != req.Code {
		newErrorResponse(ctx, http.StatusUnauthorized, "verification code is invalid")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"successful": true,
	})
}
