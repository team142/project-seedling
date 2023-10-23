package module

import (
	"strings"
	"unicode"
)

// toPascalCase will return the Pascal Case from the param
// PascalCase: Similar to CamelCase, but the first letter of the variable name is also capitalized.
// PascalCase is commonly used for naming types and public variables in some programming languages.
// For example: FirstName, NumberOfStudents, TotalPrice.
func toPascalCase(input string) string {
	var builder strings.Builder
	previousIsSeparator := true

	for _, char := range input {
		if unicode.IsSpace(char) || char == '_' {
			// Skip spaces and underscores, and set the previousIsSeparator flag to true
			previousIsSeparator = true
		} else if previousIsSeparator {
			// Capitalize the character after a separator and add it to the result
			builder.WriteRune(unicode.ToUpper(char))
			previousIsSeparator = false
		} else {
			// Copy other characters as they are
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func isFirstOrLast(index, strLen int) bool {
	return index == 0 || index == strLen-1
}

// toSnakeCase will return the Snake Case from the param
// snake_case: In snake_case, words are written in lowercase, and words are separated by underscores.
// For example: first_name, number_of_students, total_price. This style is often used in Python.
func toSnakeCase(input string) string {
	var builder strings.Builder
	previousIsSpace := false

	for index, char := range input {
		//fmt.Println(index, len(input), !isFirstOrLast(index, len(input)), string(char))
		if unicode.IsSpace(char) || unicode.IsSymbol(char) || unicode.IsPunct(char) {
			// Replace spaces with underscore
			if !previousIsSpace && !isFirstOrLast(index, len(input)) {
				builder.WriteRune('_')
				previousIsSpace = true
			}
		} else if unicode.IsUpper(char) {
			// Convert uppercase letters to lowercase and add an underscore before them
			if !previousIsSpace && !isFirstOrLast(index, len(input)) {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(char))
			previousIsSpace = false
		} else {
			// Copy other characters as they are
			builder.WriteRune(char)
			previousIsSpace = false
		}
	}

	return builder.String()
}

func lowercaseFirstLetter(s string) string {
	if len(s) == 0 {
		return s // Return the empty string if input is empty
	}

	// Convert the first character to uppercase
	firstLetter := strings.ToLower(string(s[0]))

	// Concatenate the first uppercase letter with the rest of the string
	return firstLetter + s[1:]
}

// toCamelCase will return the Camel Case from the param
// CamelCase: In CamelCase, the first letter of the variable name is in lowercase,
// and the first letter of each subsequent concatenated word is capitalized.
// For example: firstName, numberOfStudents, totalPrice.
func toCamelCase(input string) string {
	input = toPascalCase(input)
	return lowercaseFirstLetter(input)
}
