package main

import (
	"bytes"
	"reflect"
	"encoding/json"
	"os"
	"log"
	"github.com/alexeyneu/rino2/on_green"
	"github.com/gofiber/fiber/v2"
	)

type sierra map[string]string

/*{"cash_on_address":{"address":"fftyu"}, "bigtime_transfer":{    "from": "TYMwiDu22V6XG3yk6W9cTVBz48okKLRczh",
    "signatureId": "1f7f7c0c-3906-4aa1-9dfe-4b67c43918f6",
    "to": "TYMwiDu22V6XG3yk6W9cTVBz48okKLRczh",
    "tokenAddress": "TVAEYCmc15awaDRAjUZ1kvcHwQQaoPw2CW",
    "feeLimit": 0.01,    "amount": "100000"}}*/

type limbo struct {
	CashOnAddress struct {
		Address string `json:"address,omitempty"`
	} `json:"cash_on_address,omitempty"`
	BigtimeTransfer struct {
		From         string  `json:"from,omitempty"`
		To           string  `json:"to,omitempty"`
		FeeLimit     float64 `json:"feeLimit,omitempty"`
		Amount       string  `json:"amount,omitempty"`
	} `json:"bigtime_transfer,omitempty"`
}

type train struct {
	From         string  `json:"from"`
	SignatureID  string  `json:"signatureId"`
	To           string  `json:"to"`
	TokenAddress string  `json:"tokenAddress"`
	FeeLimit     float64 `json:"feeLimit"`
	Amount       string  `json:"amount"`
}

type CashOn struct{
	cash string `json:"cash,omitempty"`
}

func main() {
	bz := fiber.New()
   	bz.Post("/", postHandler)
	log.Fatalln(bz.Listen("localhost:8450"))
}

func postHandler(c *fiber.Ctx) error {
	fast := make(sierra)
	data, _:= os.ReadFile("western.json")
	json.Unmarshal(data, &fast)
	var x limbo
	json.NewDecoder(bytes.NewReader(c.Body())).Decode(&x)
        
	if(reflect.ValueOf(x).IsZero()) {
		return c.SendString("bad news")
	}
        if(reflect.ValueOf(x.CashOnAddress).IsZero()) {
		var ferro train
		ferro.From = x.BigtimeTransfer.From
		ferro.SignatureID = fast[ferro.From]
		ferro.To = x.BigtimeTransfer.To
		ferro.TokenAddress = "TR7NHqjeKQxGTCi8Q8ZY4PL8OtSzgjLj6T"
		ferro.FeeLimit = x.BigtimeTransfer.FeeLimit
		ferro.Amount = x.BigtimeTransfer.Amount
		f,_ := json.Marshal(ferro)
		h := on_green.Made_from(string(f))
		return c.SendString(h)
	}
	
	if(reflect.ValueOf(x.BigtimeTransfer).IsZero()) {
		h := on_green.Made(x.CashOnAddress.Address)
		var si CashOn
		si.cash = h
		f,_ := json.Marshal(si)
		return c.SendString(string(f))
	}
	 return c.SendString("yeah")
	
}
