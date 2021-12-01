package toss

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"K-BANK/model"
)

func Account_check(aID string) (aid string, name string, err error) {
	type response struct {
		Msg    string `json:"msg"`
		Status int    `json:"status"`
		Data   struct {
			AccountNumber int    `json:"accountNumber"`
			Name          string `json:"name"`
		} `json:"data"`
	}

	res, err := http.Get(model.TossURL + "account/account-number/" + aID)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		data := new(response)
		if err := json.NewDecoder(res.Body).Decode(data); err != nil {
			panic(err)
		}
		return strconv.Itoa(data.Data.AccountNumber), data.Data.Name, nil
	case 400:
		return "", "", errors.New("not found")
	default:
		panic("Toss error")
	}
}
