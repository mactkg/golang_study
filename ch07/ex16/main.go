package main

import (
	"fmt"
	"github.com/mactkg/golang_study/ch07/ex15/eval"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func parseEnv(s string) (eval.Env, error) {
	env := eval.Env{}
	split := strings.Split(s, ",")

	for _, v := range split {
		split := strings.Split(v, "=")
		if len(split) != 2 {
			return env, fmt.Errorf("can't parse %s", v)
		}
		parsed, err := strconv.ParseFloat(split[1], 64)
		if err != nil {
			return env, fmt.Errorf("parse error while parsing %v", split[1])
		}
		env[eval.Var(split[0])] = parsed
	}

	return env, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exprStr := r.FormValue("expr")
		if exprStr == "" {
			http.Error(w, "expr required", http.StatusBadRequest)
			return
		}

		envStr := r.FormValue("env")
		env := eval.Env{}
		if envStr != "" {
			e, err := parseEnv(envStr)
			if err != nil {
				http.Error(w, "Can't parse env", http.StatusBadRequest)
				return
			}
			env = e
		}

		expr, err := eval.Parse(exprStr)
		fmt.Println(exprStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, expr.Eval(env))
	})
	fmt.Println(`
command to test:
	curl localhost:8080 --data-urlencode "expr=a+10" --data-urlencode "env=a=10"
	curl localhost:8080 --data-urlencode "expr=a+10"
	curl localhost:8080 --data-urlencode "expr=a*b+10" --data-urlencode "env=a=12,b=20"
	curl localhost:8080 --data-urlencode "expr=10/2" --data-urlencode "env=a=10"`)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
