package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/samarec1812/sync-unisender/auth/models"
	pb "github.com/samarec1812/sync-unisender/proto/auth"
	"github.com/sirupsen/logrus"
)

const (
	getAccessTokenURI = "https://clientalexey.amocrm.ru/oauth2/access_token"
)

type getTokenQueryParam struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
}

func (h *Handler) GetToken(ctx context.Context, in *pb.AuthAmoRequest) (*pb.AuthAmoResponse, error) {
	fmt.Println(in.ClientId)

	postBody := &getTokenQueryParam{
		ClientId:     in.ClientId,
		ClientSecret: in.ClientSecret,
		GrantType:    in.GrantType,
		Code:         in.Code,
		RedirectUri:  in.RedirectUri,
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(postBody)
	if err != nil {
		return &pb.AuthAmoResponse{AuthCode: "z"}, errors.New("zalupa2")
	}
	resp, err := http.Post(getAccessTokenURI, "application/json", &buf)
	if err != nil {
		return &pb.AuthAmoResponse{AuthCode: "zalupa3"}, errors.New("zalupa3")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil || resp.StatusCode == 400 {
		logrus.Println(string(body))
		return &pb.AuthAmoResponse{AuthCode: "zalupa4"}, errors.New("zalupa4")
	}

	var data models.Token
	err = json.Unmarshal(body, &data)
	if err != nil {
		logrus.Println("error read response body")
		return &pb.AuthAmoResponse{AuthCode: "zalupa5"}, errors.New("zalupa5")
	}
	user := models.User{
		ClientId:     in.ClientId,
		ClientSecret: in.ClientSecret,
		GrantType:    in.GrantType,
		RedirectUri:  in.RedirectUri,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}
	token, err := h.services.AuthorizationAmo.GetAuthCode(user)
	if err != nil {
		logrus.Println("error save token info")
		return &pb.AuthAmoResponse{AuthCode: "zalupa6"}, errors.New("zalupa6")

	} //err = h.services.Authorization.Create("6e872c6d-5478-4bc3-b083-5eb4ba203933", token)
	//if err != nil {
	//	logrus.Println(err.Error())
	//	responseError(w, r, "Bad request", http.StatusBadRequest)
	//
	//}

	return &pb.AuthAmoResponse{AuthCode: token}, nil
}
