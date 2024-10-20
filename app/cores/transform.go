package cores

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func ToTitleCase(value string) string {
	if IsNoneOrEmptyWhiteSpace(value) {
		return EmptyString
	}

	temp := RegexReplaceAllString("([A-Z])", strings.TrimSpace(value), " $1")
	temp = RegexReplaceAllString("[-_\\s]+", temp, " ")

	transform := cases.Title(language.English)
	return transform.String(strings.ToLower(temp))
}

func ToStartCharUpper(value string) string {
	size := len(value)
	switch size {
	case 0:
		return EmptyString
	case 1:
		return strings.ToUpper(value)
	default:
		return strings.ToUpper(value[:1]) + value[1:]
	}
}

func ToStartCharLower(value string) string {
	size := len(value)
	switch size {
	case 0:
		return EmptyString
	case 1:
		return strings.ToLower(value)
	default:
		return strings.ToLower(value[:1]) + value[1:]
	}
}

func ToPascalCase(value string) string {
	temp := ToTitleCase(value)
	temp = strings.ReplaceAll(temp, " ", "")
	return ToStartCharUpper(temp)
}

func ToCamelCase(value string) string {
	temp := ToPascalCase(value)
	return ToStartCharLower(temp)
}

func ToSnakeCaseRaw(value string) string {
	if IsNoneOrEmptyWhiteSpace(value) {
		return EmptyString
	}

	temp := RegexReplaceAllString("([A-Z])", strings.TrimSpace(value), "_$1")
	temp = RegexReplaceAllString("[-_\\s]+", temp, "_")

	return temp
}

func ToSnakeCase(value string) string {
	temp := ToSnakeCaseRaw(value)
	return strings.ToLower(temp)
}

func ToSnakeCaseUpper(value string) string {
	temp := ToSnakeCaseRaw(value)
	return strings.ToUpper(temp)
}

func ToKebabCaseRaw(value string) string {
	if IsNoneOrEmptyWhiteSpace(value) {
		return EmptyString
	}

	temp := RegexReplaceAllString("([A-Z])", strings.TrimSpace(value), "-$1")
	temp = RegexReplaceAllString("[-_\\s]+", temp, "-")

	return temp
}

func ToKebabCase(value string) string {
	temp := ToKebabCaseRaw(value)
	return strings.ToLower(temp)
}

func ToKebabCaseUpper(value string) string {
	temp := ToKebabCaseRaw(value)
	return strings.ToUpper(temp)
}
