package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidationErrorResponse represents the error response structure
type ValidationErrorResponse struct {
	Error  string   `json:"error"`
	Fields []string `json:"fields"`
}

// getJSONFieldNames retrieves all JSON field names from the struct
func getJSONFieldNames(structType reflect.Type) map[string]bool {
	fieldNames := make(map[string]bool)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if tag := field.Tag.Get("json"); tag != "" {
			name := strings.Split(tag, ",")[0]
			if name != "-" {
				fieldNames[name] = true
			}
		}
	}
	return fieldNames
}

// getJSONFieldName retrieves the JSON field name from the struct tag
func getJSONFieldName(structType reflect.Type, fieldName string) string {
	if field, ok := structType.FieldByName(fieldName); ok {
		if tag := field.Tag.Get("json"); tag != "" {
			return strings.Split(tag, ",")[0]
		}
	}
	return strings.ToLower(fieldName)
}

// CheckMissingFields validates the input struct for missing required fields,
// extra fields, and ensures at least one field from optionalFields is present.
func CheckMissingFields(c *gin.Context, input any, optionalFields ...string) (ValidationErrorResponse, bool) {
	// Read the raw request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return ValidationErrorResponse{
			Error:  "Failed to read request body",
			Fields: nil,
		}, true
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// Attempt to bind JSON to the input struct
	if err := c.ShouldBindJSON(input); err != nil {
		// Check for validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			invalidFields := make([]string, 0)
			for _, e := range validationErrors {
				if e.Tag() == "required" {
					jsonFieldName := getJSONFieldName(reflect.TypeOf(input).Elem(), e.Field())
					invalidFields = append(invalidFields, jsonFieldName)
				}
			}
			if len(invalidFields) > 0 {
				return ValidationErrorResponse{
					Error:  "Missing required fields",
					Fields: invalidFields,
				}, true
			}
		}
		return ValidationErrorResponse{
			Error:  "Invalid input format",
			Fields: nil,
		}, true
	}

	// Check that at least one of the optionalFields is present
	structType := reflect.TypeOf(input).Elem()
	structValue := reflect.ValueOf(input).Elem()
	atLeastOnePresent := false
	for _, field := range optionalFields {
		// Find the struct field corresponding to the JSON field name
		for i := 0; i < structType.NumField(); i++ {
			jsonTag := structType.Field(i).Tag.Get("json")
			jsonFieldName := strings.Split(jsonTag, ",")[0]
			if jsonFieldName == field {
				// Check if the field value is non-empty
				fieldValue := structValue.Field(i).String()
				if fieldValue != "" {
					atLeastOnePresent = true
					break
				}
			}
		}
		if atLeastOnePresent {
			break
		}
	}

	// if !atLeastOnePresent {
	// 	return ValidationErrorResponse{
	// 		Error:  "At least one field required",
	// 		Fields: optionalFields,
	// 	}, true
	// }

	// Check for extra fields
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(body, &jsonMap); err != nil {
		return ValidationErrorResponse{
			Error:  "Invalid JSON format",
			Fields: nil,
		}, true
	}

	// Get expected JSON field names from the struct
	expectedFields := getJSONFieldNames(structType)

	// Collect extra fields
	extraFields := make([]string, 0)
	for key := range jsonMap {
		if !expectedFields[key] {
			extraFields = append(extraFields, key)
		}
	}

	if len(extraFields) > 0 {
		return ValidationErrorResponse{
			Error:  "Unexpected fields",
			Fields: extraFields,
		}, true
	}

	return ValidationErrorResponse{}, false
}
