package main

import (
	"fmt"
	"math/rand"
	"rank/internal/app/model"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	url = "http://127.0.0.1:8080/match/score"
)

type Result struct {
	Code int
	Message string
	Data interface{}
}

type MatchUser struct {
	ID int64
	UserID int64
	MatchID int64
	Score int
}
func main() {
	rand.Seed(time.Now().Unix())
	client := resty.New()
	for i:=0;i<100;i++ {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(NewBody()).
		Post(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", resp)

	}


}

func NewBody() model.MatchUser{
	return model.MatchUser{
		UserID:  rand.Int63n(99) + 1,
		MatchID: rand.Int63n(4) + 1,
		Score:   rand.Intn(10),
	}
}