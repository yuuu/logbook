package main

import "testing"

func Test_InputText(t *testing.T) {
	t.Skip()
	text, err := InputText("./")
	if text == nil {
		t.Fatal("text is nil.")
	}
	if err != nil {
		t.Fatal("Error in InputText.(", err.Error(), ")")
	}
	if text.Text() == "" {
		t.Fatal("text.String() is empty.")
	}
}

func Test_InputText_LengthOver(t *testing.T) {
	t.Skip()
	text, err := InputText("./")
	if text != nil {
		t.Fatal("text is not nil.")
	}
	if err == nil {
		t.Fatal("Error is nil.")
	}
}
