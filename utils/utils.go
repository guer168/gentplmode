package utils

import (
	"errors"
	"go/format"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var DSNError = errors.New("dsn string error")

// FirstCharacter
//  @Description:	gets the lower-case first character
//  @param s		character string
//  @return string
func FirstCharacter(s string) string {
	return strings.ToLower(s)[:1]
}

// FirstLowerWord
//  @Description:	first character in lowercase
//  @param s		character string
//  @return string
func FirstLowerWord(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

// CamelizeStr
//  @Description: 		convert the hump string
//  @param s			character string
//  @param upperCase	true=Uppercase first character
//  @return string
func CamelizeStr(s string, upperCase bool) string {
	if len(s) == 0 {
		return s
	}
	s = replaceInvalidChars(s)
	var result string
	words := strings.Split(s, "_")
	for i, word := range words {
		if upper := strings.ToUpper(word); commonInitialisms[upper] {
			result += upper
			continue
		}
		if i > 0 || upperCase {
			result += camelizeWord(word)
		} else {
			result += word
		}
	}
	return result
}

// camelizeWord
//  @Description:	convert the hump string
//  @param word		character string
//  @return string
func camelizeWord(word string) string {
	runes := []rune(word)
	for i, r := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(r)
		} else {
			runes[i] = unicode.ToLower(r)
		}
	}
	return string(runes)
}

// replaceInvalidChars
//  @Description:
//  @param str
//  @return string
func replaceInvalidChars(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	str = strings.ReplaceAll(str, " ", "_")
	return strings.ReplaceAll(str, ".", "_")
}

// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"ETA":   true,
	"GPU":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"OS":    true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"OAuth": true,
}

// GetDbNameFromDSN
//  @Description:
//  @param dsn
//  @return string
//  @return error
func GetDbNameFromDSN(dsn string) (string, error) {
	if len(strings.Split(dsn, " ")) > 1 {
		return getDbNameFromDsn(dsn)
	}
	index := strings.LastIndex(dsn, "/")
	if index <= 0 {
		return getDbNameFromDsn(dsn)
	}
	str := dsn[index:]
	urlStr, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return strings.Trim(urlStr.Path, "/"), nil
}

// host=127.0.0.1 dbname=test sslmode=disable Timezone=Asia/Shanghai
const dbNamePrefix = "dbname="

// getDbNameFromDsn
//  @Description:
//  @param dsn
//  @return string
//  @return error
func getDbNameFromDsn(dsn string) (string, error) {
	strArray := strings.Split(dsn, " ")
	for _, item := range strArray {
		if strings.HasPrefix(item, dbNamePrefix) {
			return strings.TrimPrefix(item, dbNamePrefix), nil
		}
	}
	return "", DSNError
}

func SaveFile(dirPath, fileName string, text []byte) error {
	file, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	defer file.Close()
	p, err := format.Source(text)
	if err != nil {
		return err
	}
	_, err = file.Write(p)
	return err
}

// MkdirPathIfNotExist
//  @Description:
//  @param dirPath
//  @return error
func MkdirPathIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, os.ModePerm)
	}
	return nil
}

// CleanUpGenFiles
//  @Description:
//  @param dir
//  @return error
func CleanUpGenFiles(dir string) error {
	exist, err := FileExists(dir)
	if err != nil {
		return err
	}
	if exist {
		return os.RemoveAll(dir)
	}
	return nil
}

// FileExists
//  @Description: Check whether the file exists
//  @param name
//  @return bool
//  @return error
func FileExists(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// FileIsExist
//  @Description: Check file exists
//  @param fieldPath
//  @return bool
func FileIsExist(fieldPath string) bool {
	_, err := os.Stat(fieldPath)
	if err == nil{
		//fmt.Println("File exist")
		return true
	}
	if os.IsNotExist(err){
		//fmt.Println("File not exist")
		return false
	}
	return false
}
