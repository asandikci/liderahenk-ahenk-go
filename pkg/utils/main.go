package utils

import (
	"log"
	"os"
)

func Byte2String(arr []int8) string {
	b := make([]byte, len(arr))
	for i, v := range arr {
		b[i] = byte(v)
	}
	return string(b)
}

func Byte2MiB(b uint64) float64 {
	return float64(float64(b) / (1024 * 1024))
}

func Byte2GiB(b uint64) float64 {
	return float64(float64(b) / (1024 * 1024 * 1024))
}

func CheckPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) {
	if flag, err := CheckPath(path); flag {
		if err != nil {
			log.Fatal(err)
		}
	} else if err := os.Mkdir(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func OpenLogFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	return f
}
