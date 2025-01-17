package main

import (
	"fmt"
	"strings"
)

// Serialize converts a Go value into a JSON string.
func Serialize(value interface{}) string {
	switch v := value.(type) {

	// Handle string values
	case string:
		return `"` + v + `"`

	// Handle numeric values
	case int, float64:
		return fmt.Sprintf("%v", v)

	// Handle boolean values
	case bool:
		if v {
			return "true"
		}
		return "false"

	// Handle null value
	case nil:
		return "null"

	// Handle arrays
	case []interface{}:
		var items []string
		for _, item := range v {
			items = append(items, Serialize(item))
		}
		return "[" + strings.Join(items, ",") + "]"

	// Handle objects
	case map[string]interface{}:
		var items []string
		for key, val := range v {
			items = append(items, Serialize(key)+":"+Serialize(val))
		}
		return "{" + strings.Join(items, ",") + "}"

	// Handle unsupported types
	default:
		panic(fmt.Sprintf("unsupported type: %T", v))
	}
}
