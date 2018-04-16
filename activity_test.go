package sendmail

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	to := "myparther@mail.com"
	tc.SetInput("to", to)
	tc.SetInput("from", "myemail@gmail.com")
	tc.SetInput("password", "MyPassword")
	tc.SetInput("subject", "Test from Flogo")
	tc.SetInput("message", "Hi! This is a message sent from Flogo")
	tc.SetInput("server", "smtp.gmail.com")
	tc.SetInput("port", 587)

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("%v", result)
	assert.Equal(t, result, "The email has been sent to "+to)
}
