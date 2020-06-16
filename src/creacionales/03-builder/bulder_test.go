package builder_test

import (
	"testing"

	builder "github.com/lucaspichi06/patrones-diseno-go/src/creacionales/03-builder"
)

func TestSender_BuildMessage(t *testing.T) {
	//json := &builder.JSONMessageBuilder{} //objeto
	xmlf := &builder.XMLMessageBuilder{} //objeto
	sender := &builder.Sender{}          //director

	sender.SetBuilder(xmlf)
	msg, err := sender.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		t.Fatalf("BuildMessage(): no se esperaba error, se obtuvo: %v", err)
	}

	t.Log(string(msg.Body))
}
