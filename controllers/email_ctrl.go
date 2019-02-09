package controllers

import (
	"bytes"
	"emailservice/constants"
	"emailservice/email"
	"emailservice/types"
	"emailservice/utils"
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

	utils.Info.Println("request body:")
	utils.Info.Println(req)

	if err != nil {
		utils.Info.Println(err)
		utils.NewError(c, http.StatusBadRequest, err)
		return
	}

	sender := email.NewSender(constants.Username, constants.Password)
	emailBody, err := render(req)

	if err != nil {
		utils.Info.Println("Error in render:")
		utils.Info.Println(err)
	}

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
