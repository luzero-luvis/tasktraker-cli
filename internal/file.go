package internal

import (
	"encoding/json" // convert json to go data and go data to json
	"os"            // do opration like read file and writefile
)

// the reason we using constant cuz file can not be changed
const file = "task.json"

func LoadTask() ([]Task, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0o644)
}
