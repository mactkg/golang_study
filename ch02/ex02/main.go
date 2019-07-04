package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mactkg/golang_study/ch02/ex01/tempconv"
	"github.com/mactkg/golang_study/ch02/ex02/lengthconv"
	"github.com/mactkg/golang_study/ch02/ex02/weightconv"
)

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			v, err := conv(arg)
			if err != nil {
				printErrorMessage(arg, err)
				continue
			}
			fmt.Printf("%s = %s\n", arg, v)
		}
	} else {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			text := stdin.Text()
			v, err := conv(text)
			if err != nil {
				printErrorMessage(text, err)
				continue
			}
			fmt.Printf("%s = %s\n", text, v)
		}
	}
}

func printErrorMessage(s string, err error) {
	fmt.Printf("Cannot parse: %s. (%v)", s, err)
}

func conv(s string) (string, error) {
	switch {
	case strings.HasSuffix(s, "°C"):
		v, err := strconv.ParseFloat(strings.TrimSuffix(s, "°C"), 64)
		if err != nil {
			return "", err
		}
		return tempconv.CToF(tempconv.Celsius(v)).String(), nil
	case strings.HasSuffix(s, "°F"):
		v, err := strconv.ParseFloat(strings.TrimSuffix(s, "°F"), 64)
		if err != nil {
			return "", err
		}
		return tempconv.FToC(tempconv.Fahrenheit(v)).String(), nil
	case strings.HasSuffix(s, "m"):
		v, err := strconv.ParseFloat(strings.TrimSuffix(s, "m"), 64)
		if err != nil {
			return "", err
		}
		return lengthconv.MToFt(lengthconv.Meter(v)).String(), nil
	case strings.HasSuffix(s, "ft"):
		v, err := strconv.ParseFloat(strings.TrimSuffix(s, "ft"), 64)
		if err != nil {
			return "", err
		}
		return lengthconv.FtToM(lengthconv.Feet(v)).String(), nil
	case strings.HasSuffix(s, "kg"):
		v, err := strconv.ParseFloat(strings.TrimSuffix(s, "kg"), 64)
		if err != nil {
			return "", err
		}
		return weightconv.KgToLbs(weightconv.Kilogram(v)).String(), nil
	case strings.HasSuffix(s, "lbs"):
		v, err := strconv.ParseFloat(strings.TrimSuffix(s, "lbs"), 64)
		if err != nil {
			return "", err
		}
		return weightconv.LbsToKg(weightconv.Pound(v)).String(), nil
	default:
		return "", errors.New("Convert function not found")
	}
}
