package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

type req struct {
	text      string
	contentID int64
	clientID  int
	timeStamp time.Time
}

func main() {

	h := requestHandler
	if err := fasthttp.ListenAndServe(":8000", h); err != nil {
		log.Fatal("Error in listen and serve: ", err)
	}

}

func requestHandler(ctx *fasthttp.RequestCtx) {
	reqBody := string(ctx.PostBody())
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(reqJSON))

}
