package main

import (
	"errors"
	"strconv"
	"strings"
)

// Deserialize parses a JSON string into a Go value.
func Deserialize(json string) (interface{}, error) {
	json = strings.TrimSpace(json)

	// Handle string values
	if strings.HasPrefix(json, `"`) && strings.HasSuffix(json, `"`) {
		return json[1 : len(json)-1], nil
	}

	// Handle numeric values
	if num, _ := strconv.Atoi(json); num != 0 {
		return num, nil
	}
	if num, _ := strconv.ParseFloat(json, 64); num != 0 {
		return num, nil
	}

	// Handle boolean values
	switch json {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}

	// Handle null value
	if json == "null" {
		return nil, nil
	}

	// Handle arrays
	if strings.HasPrefix(json, "[") && strings.HasSuffix(json, "]") {
		inner := json[1 : len(json)-1]
		elements := splitJSON(inner)
		var array []interface{}

		for _, element := range elements {
			val, _ := Deserialize(element)
			array = append(array, val)
		}

		return array, nil
	}

	// Handle objects
	if strings.HasPrefix(json, "{") && strings.HasSuffix(json, "}") {
		inner := json[1 : len(json)-1]
		pairs := splitJSON(inner)
		obj := make(map[string]interface{})

		for _, pair := range pairs {
			keyValue := strings.SplitN(pair, ":", 2)
			if len(keyValue) != 2 {
				return nil, errors.New("invalid JSON object")
			}

			key, _ := Deserialize(keyValue[0])
			value, _ := Deserialize(keyValue[1])

			strKey, ok := key.(string)
			if !ok {
				return nil, errors.New("invalid object key, must be a string")
			}

			obj[strKey] = value
		}

		return obj, nil
	}

	// Return error for invalid JSON
	return nil, errors.New("invalid JSON string")
}

// splitJSON splits JSON elements while accounting for nested structures.
func splitJSON(json string) []string {
	var elements []string
	start, braces, brackets := 0, 0, 0

	for i, char := range json {
		switch char {
		case '{':
			braces++
		case '}':
			braces--
		case '[':
			brackets++
		case ']':
			brackets--
		case ',':
			if braces == 0 && brackets == 0 {
				elements = append(elements, strings.TrimSpace(json[start:i]))
				start = i + 1
			}
		}
	}

	elements = append(elements, strings.TrimSpace(json[start:]))
	return elements
}
