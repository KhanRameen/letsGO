package main

import (
	"encoding/json"
	"os"
)

const fileName = "todos.json"

func loadTodos() ([]Todo, error) {
	var todos []Todo
	
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		
		if os.IsNotExist(err) {
			return []Todo{}, nil // Return an empty slice if the file doesn't exist
		}
		return nil, err
}


	err = json.Unmarshal(fileBytes, &todos) //Unmarshal is used to convert JSON data into Go data structures. It takes a byte slice of JSON data and a pointer to the variable where the decoded data should be stored.
	if err != nil {
		return nil, err
	}
	return todos, err
	
}

func saveTodos(todos []Todo) error {
	data , err := json.MarshalIndent(todos, "", "  ") //turns go data to json but with formatiing , (data, newline prefix, indentation)
	if err!= nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644) //0644 is the file permission
}