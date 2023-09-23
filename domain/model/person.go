package model

type Person struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type ReqPerson struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
