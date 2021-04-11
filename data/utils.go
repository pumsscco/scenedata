package data

import (
	"os"
	"fmt"
	"bytes"
	"encoding/binary"
)
func GetInt(fp *os.File) int {
    tmpS := make([]byte,4)
    fp.Read(tmpS)
    var i int32
    binary.Read(bytes.NewReader(tmpS),binary.LittleEndian,&i)
    return int(i)
}
func GetFloat(fp *os.File)(f float32){
    tmpS := make([]byte,4)
    fp.Read(tmpS)
    binary.Read(bytes.NewReader(tmpS),binary.LittleEndian,&f)
    return
}
func GetStr(fp *os.File,l int) string {
    tmpS := make([]byte,l)
    fp.Read(tmpS)
    return string(tmpS)
}
func GetFlag(fp *os.File,l int) string {
    tmpS := make([]byte,l)
    fp.Read(tmpS)
    return fmt.Sprintf("%0X", tmpS)
}