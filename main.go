package main

import "fmt"

func main() {
	// Example data
	data := map[string]interface{}{
		"name":       "Areeb",
		"age":        21,
		"is_student": true,
		"courses":    []interface{}{"Math", "Science"},
	}

	// Serialize
	serialized := Serialize(data)
	fmt.Println("Serialized JSON:", serialized)

	// Deserialize
	deserialized, _ := Deserialize(serialized)
	fmt.Println("\nDeserialized Go object:", deserialized)
}
