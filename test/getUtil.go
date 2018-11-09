package test

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)
//uupool 余额
func  GetUupoolBalance(num []byte) string {
	data := string(num[:])
	i, _ := strconv.ParseFloat(data,64)
	s := fmt.Sprintf("%0.8f",i/100000000)
	return s
}
//uupool 已支付
func  GetUupoolPaid(num []byte) string {
	data := string(num[:])
	i, _ := strconv.ParseFloat(data,64)
	s := fmt.Sprintf("%0.8f",i/100000000)
	return s
}
//uupool 总收益
func  GetUupoolAllBalance(paid []byte,balance []byte) string {
	n1,_ := strconv.ParseFloat(GetUupoolBalance(balance),64)
	n2,_ := strconv.ParseFloat(GetUupoolPaid(paid),64)
	s := fmt.Sprintf("%0.8f",n1+n2)
	return s
}
//uupool 实时算力
func  GetUupoolHashrate(num []byte) string {
	data:= string(num[:])
	arr := strings.Fields(data)
	count, _ := strconv.ParseFloat(arr[0], 64)
	switch arr[1] != "" {
	case arr[1] == "M" || arr[1] == "m":
		break
	case arr[1] == "K" || arr[1] == "k":
		count /= 1000
	case arr[1] == "G" || arr[1] == "g":
		count *= 1000
	case arr[1] == "T" || arr[1] == "t":
		count *= 1000000
	case arr[1] == "P" || arr[1] == "p":
		count *= 1000000000
	default:
	}
	s := fmt.Sprintf("%0.2f", count)
	return data
}
func BytesToFloat32(buf []byte) float32 {
	v := binary.LittleEndian.Uint32(buf)
	f := math.Float32frombits(v)
	return f
}
