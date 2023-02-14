package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	pb "github.com/samarec1812/sync-unisender/proto/auth"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func (h *Handler) saveToken(ctx *fasthttp.RequestCtx) {

	args := ctx.PostArgs()
	params, err := url.ParseQuery(args.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//token := models.UnisenderTokenInfo{
	//	Token:     string(args.Peek("token")),
	//	Uname:    ,
	//	AccountId: binary.BigEndian.Uint64(args.Peek("account_id")),
	//}
	//
	//if err := token.Validate(); err != nil {
	//	logrus.Errorf("input parameters token error: %s", err.Error())
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	fmt.Fprintf(ctx, "bad request\n")
	//	return
	//}

	accountId, err := strconv.ParseUint(params.Get("account[id]"), 10, 64)
	if err != nil {
		logrus.Errorf("parameters account id have error")
		ctx.SetStatusCode(http.StatusBadRequest)
		fmt.Fprintf(ctx, "bad request\n")
		return
	}

	authReq := &pb.AuthUnisenderRequest{
		Token:     params.Get("token"),
		Uname:     params.Get("uname"),
		AccountId: accountId,
	}

	resp, err := h.clientRPC.SaveToken(context.Background(), authReq)

	if err != nil {
		logrus.Errorf("input parameters token error: %s", err.Error())
		ctx.SetStatusCode(http.StatusBadRequest)
		fmt.Fprintf(ctx, "bad request\n")
		return
	}

	if resp.AuthCode != "200" {
		logrus.Errorf("error with save unisender token")
		ctx.SetStatusCode(http.StatusBadRequest)
		fmt.Fprintf(ctx, "bad request\n")
		return
	}

	logrus.Println("unisender token save successful")
	ctx.SetStatusCode(http.StatusOK)
	fmt.Fprintf(ctx, "unisender token save successful\n")
}

//func (h *Handler) printToken(ctx *fasthttp.RequestCtx) {
//
//	args := ctx.PostArgs()
//	fmt.Println(args)
//	token := models.UnisenderTokenInfo{
//		Token:     string(args.Peek("token")),
//		Uname:     string(args.Peek("uname")),
//		AccountId: binary.BigEndian.Uint64(args.Peek("account_id")),
//	}
//	fmt.Println(token)
//
//	fmt.Fprintf(ctx, "Hello there %q", ctx.RequestURI())
//}
