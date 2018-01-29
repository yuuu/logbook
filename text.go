package main

import (
	"os"
	"io/ioutil"

	"errors"

	editor "github.com/yuuu/go-editor"
)

const TEXT_LENGTH_MAX = 10240
const TEXT_TMPORARY_FILE_PATH = "tmp.txt"

type Text struct {
	text string
}

func editSub() (string, error) {
	var err error

	err = os.Remove(TEXT_TMPORARY_FILE_PATH)
	if err != nil {
		// already removed.
		err = nil
	}

	var entryEditor *editor.Editor
	entryEditor = editor.NewEditor("")
	entryEditor.Launch(TEXT_TMPORARY_FILE_PATH)

	var textBytes []byte
	textBytes, err = ioutil.ReadFile(TEXT_TMPORARY_FILE_PATH)
	if err != nil {
		return "", errors.New("ReadFile Error.")
	}

	textString := string(textBytes)
	if len([]rune(textString)) > TEXT_LENGTH_MAX {
		return "", errors.New("text length over.")
	}

	return textString, nil
}

func InputText(workpath string) (*Text, error) {
	var err error
	var textString string

	textString, err = editSub()
	if err != nil {
		return nil, err
	}

	return &Text{textString}, nil
}

func CreateText(text string) *Text {
	return &Text{text}
}

func (self *Text) Edit() error {
	var err error
	var textString string

	err = ioutil.WriteFile(TEXT_TMPORARY_FILE_PATH, []byte(self.text), 0777)
	if err != nil {
		return err
	}
	textString, err = editSub()
	if err != nil {
		return err
	}
	self.text = textString
	return nil
}

func (self *Text) Text() string {
	return self.text
}
