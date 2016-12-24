package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wizjin/weixin"
)

func GuessFib(i string) string {
	number, err := strconv.Atoi(i)
	if err != nil {
		return "换个数字呗"
	}

	return fibNext(number)
}

func main() {
	mux := weixin.New(
		os.Getenv("WB_MP_TOKEN"),
		os.Getenv("WB_MP_APPID"),
		os.Getenv("WB_MP_APPSECRET"),
	)

	mux.HandleFunc(weixin.MsgTypeText, func(w weixin.ResponseWriter, r *weixin.Request) {
		w.ReplyText(GuessFib(r.Content))
	})

	http.Handle("/fibNext/mp", mux)
	http.HandleFunc("/fibNext", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(GuessFib(r.FormValue("number"))))
	})
	log.Fatal(http.ListenAndServe(os.Getenv("WB_BIND"), nil))
}
