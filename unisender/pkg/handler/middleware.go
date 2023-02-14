package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net/url"
)

const (
	contactActionAdd    = "contacts[add][0][id]"
	contactActionUpdate = "contacts[update][0][id]"
	contactActionDelete = "contacts[delete][0][id]"
)

func (h *Handler) contactsMiddleware(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()
	params, err := url.ParseQuery(args.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if params.Get(contactActionAdd) != "" {
		h.addContact(ctx)
		return
	} else if params.Get(contactActionUpdate) != "" {
		h.updateContact(ctx)
		return
	} else if params.Get(contactActionDelete) != "" {
		h.deleteContact(ctx)
		return
	} else {
		logrus.Println("handlers with action contacts not found")
		ctx.NotFound()
		return
	}
}
