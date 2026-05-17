package main

type Event struct {
	Type string     `json:"type"`
	Repo Repo      	`json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	Name string 	`json:"name"`
}

type Payload struct {
	Commits []Commit`json:"commits"`
	Action  string  `json:"action"`
}

type Commit struct {
	Message string 	`json:"message"`
}