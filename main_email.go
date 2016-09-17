package main

import (
	"github.com/iris-contrib/mail"
	"github.com/kataras/iris"
)

func main() {
	// change these to your settings

	cfg := mail.Config{
		Host:     "smtp.mailgun.org",
		Username: "postmaster@sandbox661c307650f04e909150b37c0f3b2f09.mailgun.org",
		Password: "38304272b8ee5c176d5961dc155b2417",
		Port:     587,
	}
	// change these to your e-mail to check if that works

	// create the service
	mailService := mail.New(cfg)

	var to = []string{"kataras2006@hotmail.com", "social@ideopod.com"}

	// standalone

	//iris.Must(mailService.Send("iris e-mail test subject", "</h1>outside of context before server's listen!</h1>", to...))

	//inside handler
	iris.Get("/send", func(ctx *iris.Context) {
		content := `<h1>Hello From Iris web framework</h1> <br/><br/> <span style="color:blue"> This is the rich message body </span>`

		err := mailService.Send("iris e-mail just t3st subject", content, to...)

		if err != nil {
			ctx.HTML(200, "<b> Problem while sending the e-mail: "+err.Error())
		} else {
			ctx.HTML(200, "<h1> SUCCESS </h1>")
		}
	})

	// send a body by template
	iris.Get("/send/template", func(ctx *iris.Context) {
		content := iris.TemplateString("body.html", iris.Map{
			"Message": " his is the rich message body sent by a template!!",
			"Footer":  "The footer of this e-mail!",
		}, iris.RenderOptions{"charset": "UTF-8"})
		// iris.RenderOptions are optional parameter,
		// "charset" defaults to UTF-8 but you can change it for a
		// particular mail receiver

		err := mailService.Send("iris e-mail just t3st subject", content, to...)

		if err != nil {
			ctx.HTML(200, "<b> Problem while sending the e-mail: "+err.Error())
		} else {
			ctx.HTML(200, "<h1> SUCCESS </h1>")
		}
	})
	iris.Listen(":8080")
}
