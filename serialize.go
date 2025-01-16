package main

import (
	"fmt"
	"strings"
)

// Serialize converts a Go value into a JSON string.
func Serialize(value interface{}) string {
	switch v := value.(type) {

	case string:
		return `"` + v + `"`

	case int, float64:
		return fmt.Sprintf("%v", v)

	case bool:
		if v {
			return "true"
		}
		return "false"

	case nil:
		return "null"

	case map[string]interface{}:
		var items []string
		for key, val := range v {
			items = append(items, Serialize(key)+":"+Serialize(val))
		}
		return "{" + strings.Join(items, ",") + "}"

	case []interface{}:
		var items []string
		for _, item := range v {
			items = append(items, Serialize(item))
		}
		return "[" + strings.Join(items, ",") + "]"

	default:
		panic(fmt.Sprintf("unsupported type: %T", v))
	}
}
