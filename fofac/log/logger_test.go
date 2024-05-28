package log

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	Error("查询错i五")
}

func TestString(t *testing.T) {
	str := "html5shiv/r29/html5.min.js"
	bytes, _ := json.Marshal(str)
	str1 := "asn=\"40065\""
	marshal, err := json.Marshal(str1)
	fmt.Print(string(marshal)+"&&"+string(bytes), err)
}
