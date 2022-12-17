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
//  @Description:	获取首字母
//  @param s		字符串
//  @return string
func FirstCharacter(s string) string {
	return strings.ToLower(s)[:1]
}

// FirstLowerWord
//  @Description:	首字母小写
//  @param s		字符串
//  @return string
func FirstLowerWord(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

// RemovePrefix
//  @Description: 移除前缀
//  @param tableName
//  @param pre
//  @return string
func RemovePrefix(tableName string, pre string) string {
	return strings.ReplaceAll(tableName, pre, "")
}

// CamelizeStr
//  @Description: 		转换驼峰格式
//  @param s			字符串
//  @param upperCase	true=首字母大写
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

// camelizeWord 转换驼峰格式
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
//  @Description:	转换特殊字符串
//  @param str
//  @return string
func replaceInvalidChars(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	str = strings.ReplaceAll(str, " ", "_")
	return strings.ReplaceAll(str, ".", "_")
}

// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
//	全大写
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
	//"ID":    true,
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
	//"UID":   true,
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

func getDbNameFromDsn(dsn string) (string, error) {
	strArray := strings.Split(dsn, " ")
	for _, item := range strArray {
		if strings.HasPrefix(item, dbNamePrefix) {
			return strings.TrimPrefix(item, dbNamePrefix), nil
		}
	}
	return "", DSNError
}

// SaveFile
//  @Description:	保存文件
//  @param dirPath	路径
//  @param fileName	文件名
//  @param text		内容
//  @return error
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
//  @Description:	创建目录
//  @param dirPath
//  @return error
func MkdirPathIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, os.ModePerm)
	}
	return nil
}

// CleanUpGenFiles
//  @Description:	删除指定目录下所有文件
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

// RemoveFile
//  @Description: 删除文件
//  @param filePath
//  @return error
func RemoveFile(filePath string) error {
	if res,_ := FileExists(filePath); res == true{
		return os.Remove(filePath)
	}
	return nil
}

// FileExists
//  @Description: 	判断文件是否存在
//  @param filePath	文件路径
//  @return bool
//  @return error
func FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// InSlice
//  @Description: 判断是否在切片中
//  @param items	切片
//  @param item		字符串
//  @return bool
func InSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
