package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Manager struct {
	FullName       string `json:"full_name"`
	Position       string `json:"position"`
	Age            int32  `json:"age"`
	YearsInCompany int32  `json:"years_in_company"`
}

func encoderManager(manager *Manager) (io.Reader, error) {
	data, _ := json.Marshal(manager)
	r := strings.NewReader(string(data))
	return r, nil
}

func main() {
	m := &Manager{
		FullName:       "Reza",
		Position:       "BE",
		Age:            20,
		YearsInCompany: 9,
	}

	fmt.Println(encoderManager(m))
}
