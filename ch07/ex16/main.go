package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/mactkg/golang_study/ch07/ex15/eval"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Write expr(return to exec): ")
		if !input.Scan() {
			return
		}
		str := input.Text()

		expr, err := eval.Parse(str)
		if err != nil {
			fmt.Printf("Error happened while parsing: %v\n", err)
			continue
		}

		// var check
		envList := map[eval.Var]bool{}
		env := eval.Env{}
		err = expr.Check(envList)
		if err != nil {
			fmt.Printf("Error happend while checking variables: %v\n", err)
			continue
		}

		// var input
		for k, v := range envList {
			if !v {
				continue
			}
		ASK_INPUT:
			fmt.Printf("%s = ?: ", k)
			if !input.Scan() {
				return
			}

			res := input.Text()
			val, err := strconv.ParseFloat(res, 64)
			if err != nil {
				fmt.Printf("Parse input error: %v\n\n", err)
				goto ASK_INPUT
			}
			env[k] = val
		}

		result := expr.Eval(env)
		fmt.Printf("%s = %v\n\n", str, result)
	}
}
