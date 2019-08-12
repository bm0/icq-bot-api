package icqbotapi

import (
	"log"
	"net/http"
	"testing"
)

func TestBot_GetFileInfo(t *testing.T) {
	bot := Bot{
		token,
		apiBaseURL,
		http.DefaultClient,
	}

	data, err := bot.GetFileInfo(FileID("05j5Lk9Oka1VBpeMLS7Qv35d5189ac1af"))
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%#v", data)
}
