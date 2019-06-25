package support

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/**
* @author zhangsheng
* @date 2019/6/12
 */
func LogFile() *os.File {
	filename := fmt.Sprintf("us_%v.log", time.Now().Format("2006-01-02"))
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func MarshalBinary(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func UnmarshalBinary(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}
