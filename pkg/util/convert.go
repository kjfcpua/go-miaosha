/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/25 6:23 下午
 * @Description: 转换工具
 */
package util

import (
	"fmt"
	"math"
	"strconv"
)

//FloatToStr float to str 支持指定精度
func FloatToStr(num float64, floatPartLen int) string {
	return strconv.FormatFloat(num, 'f', floatPartLen, 64)
}

//StrToFloat64 支持指定精度
func StrToFloat64(str string, len int) (float64, error) {
	lenstr := "%." + strconv.Itoa(len) + "f"
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	nstr := fmt.Sprintf(lenstr, value) //指定精度
	val, err := strconv.ParseFloat(nstr, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

//StrToFloat64round 支持指定精度， 支持四舍五入
func StrToFloat64round(str string, prec int, round bool) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return FloatPrecision(f, prec, round), nil
}

// FloatPrecision float指定精度; round为true时, 表示支持四舍五入
func FloatPrecision(f float64, prec int, round bool) float64 {
	pow10N := math.Pow10(prec)
	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}
	return math.Trunc((f)*pow10N) / pow10N
}
