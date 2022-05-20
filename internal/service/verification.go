package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Korisss/skymp-master-api-go/internal/repository"
	"github.com/Korisss/skymp-master-api-go/pkg/random"
	"github.com/sirupsen/logrus"
)

var botToken = os.Getenv("BOT_TOKEN")
var botUri = os.Getenv("BOT_URI")

type VerificationService struct {
	repo repository.Verification
}

type sendCodeToBotReq struct {
	Discord string `json:"discord"`
	Code    int    `json:"code"`
}

type sendCodeToBotRes struct {
	DiscordId string `json:"id"`
}

func NewVerificationService(repo repository.Verification) *VerificationService {
	return &VerificationService{repo: repo}
}

func (s *VerificationService) GetVerificationCode(id int64) (int, error) {
	return s.repo.GetVerificationCode(id)
}

func (s *VerificationService) SendCodeToBot(id int64, discord string) error {
	code := random.RandInt(4)

	if err := s.repo.SetVerificationCode(id, code); err != nil {
		return err
	}

	data, err := json.Marshal(sendCodeToBotReq{
		Discord: discord,
		Code:    code,
	})
	if err != nil {
		logrus.Error("Error reading request. ", err.Error())
	}

	req, err := http.NewRequest("POST", "http://"+botUri+"/send-code", bytes.NewBuffer(data))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("can't read body")
	}

	var res sendCodeToBotRes

	err = json.Unmarshal(body, &res)
	if err != nil {
		return errors.New("error when reading body")
	}

	s.repo.SetDiscord(id, res.DiscordId)

	return nil
}
