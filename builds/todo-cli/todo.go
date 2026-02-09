package main

type Todo struct {
	ID          int    `json:"id"` //json tag to specify the field name in JSON
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}