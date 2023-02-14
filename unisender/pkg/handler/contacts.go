package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/samarec1812/sync-unisender/unisender/models"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	priorityBeanstalk  = uint32(10)
	delayBeanstalk     = 0 * time.Second
	timeToRunBeanstalk = 20 * time.Second
)

// swagger:operation POST /contacts  createContact
// Returns status operation
// ---
//
// consumes:
// - application/x-www-form-urlencoded
//
// produces:
// - application/json
//
// schemas: http
// responses:
// 		'200':
//			 description: Successful operation
// 		'400':
// 			 description: Operation failed
// 		'500':
// 			 description: Server error
func (h *Handler) addContact(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()

	params, err := url.ParseQuery(args.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := ""
	var sb strings.Builder

	accountId, _ := strconv.ParseUint(params.Get("account[id]"), 10, 64)
	contactId, _ := strconv.ParseUint(params.Get("contacts[add][0][id]"), 10, 64)

	for key, value := range params {
		if value[0] == "Email" {
			s = key[:len(key)-len("[name]")]
			fmt.Println("KEY: ", s)
		}
		if strings.Contains(key, s) && strings.Contains(key, "[value]") {
			sb.WriteString(value[0])
			sb.WriteString(",")
		}
	}

	f := models.Contact{
		ContactId: contactId,
		Name:      params.Get("contacts[add][0][name]"),
		Email:     (sb.String())[:sb.Len()-1],
		AccountId: accountId,
	}

	id, err := h.services.Contacts.Create(f)
	if err != nil {
		logrus.Errorf("error with create contact")
		ctx.SetStatusCode(http.StatusBadRequest)
		fmt.Fprintf(ctx, "bad request\n")
		return
	}
	fmt.Println("id: ", id)
	fmt.Println(models.BeanstalkData{Id: id})
	bytes, err := json.Marshal(models.BeanstalkData{Id: id})
	if err != nil {
		logrus.Errorf("error with marshaling contact")
		//ctx.SetStatusCode(http.StatusBadRequest)
		//fmt.Fprintf(ctx, "bad request\n")
	}

	jobId, err := h.producer.Put(bytes, priorityBeanstalk, delayBeanstalk, timeToRunBeanstalk)
	if err != nil {
		logrus.Errorf("error with put contact to beanstalk")
		//ctx.SetStatusCode(http.StatusBadRequest)
		//fmt.Fprintf(ctx, "bad request\n")
		//return
	}
	logrus.Printf("beanstalk job is done with id: %d", jobId)

	//fmt.Println(arrContacts)
	logrus.Println("create contacts successful")
	ctx.SetStatusCode(http.StatusOK)
	fmt.Fprintf(ctx, "create contacts successful\n")
}

func (h *Handler) updateContact(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()

	params, err := url.ParseQuery(args.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := ""
	var sb strings.Builder

	accountId, _ := strconv.ParseUint(params.Get("account[id]"), 10, 64)
	contactId, _ := strconv.ParseUint(params.Get("contacts[update][0][id]"), 10, 64)

	for key, value := range params {
		if value[0] == "Email" {
			s = key[:len(key)-len("[name]")]
			fmt.Println("KEY: ", s)
		}
		if strings.Contains(key, s) && strings.Contains(key, "[value]") {
			sb.WriteString(value[0])
			sb.WriteString(",")
		}
	}

	f := models.Contact{
		ContactId: contactId,
		Name:      params.Get("contacts[update][0][name]"),
		Email:     (sb.String())[:sb.Len()-1],
		AccountId: accountId,
	}

	id, err := h.services.Contacts.Update(f)
	if err != nil {
		logrus.Errorf("error with update contact")
		ctx.SetStatusCode(http.StatusBadRequest)
		fmt.Fprintf(ctx, "bad request\n")
		return
	}
	fmt.Println("id: ", id)
	fmt.Println(models.BeanstalkData{Id: id})

	bytes, err := json.Marshal(models.BeanstalkData{Id: id})
	if err != nil {
		logrus.Errorf("error with marshaling contact")
		//ctx.SetStatusCode(http.StatusBadRequest)
		//fmt.Fprintf(ctx, "bad request\n")
	}
	fmt.Println(string(bytes), id)
	jobId, err := h.producer.Put(bytes, priorityBeanstalk, delayBeanstalk, timeToRunBeanstalk)
	if err != nil {
		logrus.Errorf("error with put contact to beanstalk")
		//ctx.SetStatusCode(http.StatusBadRequest)
		//fmt.Fprintf(ctx, "bad request\n")
		//return
	}
	logrus.Printf("beanstalk job is done with id: %d", jobId)

	//fmt.Println(arrContacts)
	logrus.Println("update contacts successful")
	ctx.SetStatusCode(http.StatusOK)
	fmt.Fprintf(ctx, "update contacts successful\n")
}

func (h *Handler) deleteContact(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()

	params, err := url.ParseQuery(args.String())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	contactId, err := strconv.ParseUint(params.Get("contacts[delete][0][id]"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := h.services.Contacts.Delete(contactId)
	if err != nil {
		logrus.Errorf("error with delete contact")
		ctx.SetStatusCode(http.StatusBadRequest)
		fmt.Fprintf(ctx, "bad request\n")
		return
	}
	fmt.Println("id: ", id)
	fmt.Println(models.BeanstalkData{Id: id})
	bytes, err := json.Marshal(models.BeanstalkData{Id: id})
	if err != nil {
		logrus.Errorf("error with marshaling contact")
		//ctx.SetStatusCode(http.StatusBadRequest)
		//fmt.Fprintf(ctx, "bad request\n")
	}

	jobId, err := h.producer.Put(bytes, priorityBeanstalk, delayBeanstalk, timeToRunBeanstalk)
	if err != nil {
		logrus.Errorf("error with put contact to beanstalk")
		//ctx.SetStatusCode(http.StatusBadRequest)
		//fmt.Fprintf(ctx, "bad request\n")
		//return
	}
	logrus.Printf("beanstalk job is done with id: %d", jobId)

	logrus.Println("delete contact successful")
	ctx.SetStatusCode(http.StatusOK)
	fmt.Fprintf(ctx, "delete contact successful\n")
}
