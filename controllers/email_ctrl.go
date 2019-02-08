package controllers

import (
	"bytes"
	"emailservice/constants"
	"emailservice/email"
	"emailservice/types"
	"emailservice/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrlr *Controller) SendEmail(c *gin.Context) {

	req := types.SendOrderConfirmationRequest{}
	err := c.ShouldBind(&req)

	fmt.Println(req)

	if err != nil {
		fmt.Println(err)
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	sender := email.NewSender(constants.Username, constants.Password)
	emailBody, _ := render(req)
	sender.SendMail("Your Confirmation Email", emailBody, constants.FromAddress, req.Email)

}

func render(in types.SendOrderConfirmationRequest) (string, error) {
	tmpl := template.Must(template.ParseFiles("template/confirmation.html"))
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, in); err != nil {
		return "", err
	}
	str := tpl.String()
	return str, nil

}
