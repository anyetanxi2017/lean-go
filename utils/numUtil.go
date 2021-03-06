package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type NumUtil struct {
}

func (n NumUtil) RandomFloat(min, max float64) float64 {
	if max == min {
		return max
	}
	return n.decimal(rand.Float64()*(max-min) + min)
}
func (n NumUtil) RandomFloatStr(minStr, maxStr string) float64 {
	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		panic(err)
	}
	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		panic(err)
	}
	return n.RandomFloat(min, max)
}
func (n NumUtil) decimal(v float64) float64 {
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
	return val
}

func (n NumUtil) RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func (n NumUtil) StrToFloat(s string) float64 {
	if float, err := strconv.ParseFloat(s, 10); err != nil {
		panic(err)
	} else {
		return float
	}
	return 0
}
func (n NumUtil) NumToInt64(num interface{}) (res int64, err error) {
	switch num.(type) {
	case string:
		res, err = strconv.ParseInt(num.(string), 10, 64)
	case int:
		res = int64(num.(int))
	}
	return
}
