package sendmail

import (
	"log"
	"net/smtp"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// set inputs
	to := context.GetInput("to").(string)
	from := context.GetInput("from").(string)
	password := context.GetInput("password").(string)
	subject := context.GetInput("subject").(string)
	message := context.GetInput("message").(string)
	server := context.GetInput("server").(string)
	port := context.GetInput("port").(int)

	url := server + ":" + strconv.Itoa(port)

	// do eval
	body := "To: " + to + "\r\nSubject: " +
		subject + "\r\n\r\n" + message
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err = smtp.SendMail(url, auth, from,
		[]string{to}, []byte(body))
	context.SetOutput("result", "The email has been sent to "+to)
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
