package structural

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "Hello world!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello world!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter2 := PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter2.PrintStored()
	if returnedMsg != msg {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
