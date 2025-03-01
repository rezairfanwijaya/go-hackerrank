package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type Manager struct {
	Fullname       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

func main() {
	manager := &Manager{
		Fullname:       "Jack Oliver",
		Position:       "CEO",
		Age:            44,
		YearsInCompany: 34,
	}
	res, err := encodeManager(manager)
	if err != nil {
		log.Println(err)
		return
	}

	readRes, err := io.ReadAll(res)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(readRes))
}

func encodeManager(manager *Manager) (io.Reader, error) {
	managerByte, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(managerByte), nil
}
