package utils

import (
	"encoding/base64"
	"encoding/json"
)

func Jsonfy(som interface{}) string {
	j, _ := json.Marshal(som)
	return string(j)
}

func GenerateToken(name string, password string) string {
	concat := name + password
	token := base64.StdEncoding.EncodeToString([]byte(concat))

	return token
}
