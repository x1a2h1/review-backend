package models

type Project struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      Status  `json:"status"`
	Deadline    int64   `json:"deadline"`
	Member      []int64 `json:"menber"`
}

type Status string
