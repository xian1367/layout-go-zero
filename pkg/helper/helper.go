package helper

import (
	"fmt"
	"strings"
)

func VarDump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func Query() {

}

func QueryMap(m map[string][]string, key string) map[string][]string {
	dict := make(map[string][]string)
	for k, v := range m {
		if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == key {
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
				dict[k[i+1:][:j]] = v
			}
		}
	}
	return dict
}
