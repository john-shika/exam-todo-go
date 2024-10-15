package cores

import "regexp"

type ReOrStrImpl interface {
	~string | *regexp.Regexp
}

func GetRegexPattern[T ReOrStrImpl](pattern T) *regexp.Regexp {
	var ok bool
	var re *regexp.Regexp
	var reStr string
	if re, ok = CastPtr[regexp.Regexp](pattern); !ok {
		if reStr, ok = Cast[string](pattern); !ok {
			panic("pattern must be 'regexp.Regexp' or string type")
		}
		re = regexp.MustCompile(reStr)
	}
	return re
}

func RegexMatch[T ReOrStrImpl](pattern T, value []byte) bool {
	re := GetRegexPattern(pattern)
	return re.Match(value)
}

func RegexMatchString[T ReOrStrImpl](pattern T, value string) bool {
	re := GetRegexPattern(pattern)
	return re.MatchString(value)
}

func RegexFind[T ReOrStrImpl](pattern T, value []byte) []byte {
	re := GetRegexPattern(pattern)
	return re.Find(value)
}

func RegexFindString[T ReOrStrImpl](pattern T, value string) string {
	re := GetRegexPattern(pattern)
	return re.FindString(value)
}

func RegexFindAll[T ReOrStrImpl](pattern T, value []byte, n int) [][]byte {
	re := GetRegexPattern(pattern)
	return re.FindAll(value, n)
}

func RegexFindAllString[T ReOrStrImpl](pattern T, value string, n int) []string {
	re := GetRegexPattern(pattern)
	return re.FindAllString(value, n)
}

func RegexReplaceAll[T ReOrStrImpl](pattern T, value []byte, replace []byte) []byte {
	re := GetRegexPattern(pattern)
	return re.ReplaceAll(value, replace)
}

func RegexReplaceAllString[T ReOrStrImpl](pattern T, value string, replace string) string {
	re := GetRegexPattern(pattern)
	return re.ReplaceAllString(value, replace)
}

func RegexSplit[T ReOrStrImpl](pattern T, value string, n int) []string {
	re := GetRegexPattern(pattern)
	return re.Split(value, n)
}
