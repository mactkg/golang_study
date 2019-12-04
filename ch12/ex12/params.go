// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Pack(url string, val interface{}) (string, error) {
	paramStr := ""

	ref := reflect.ValueOf(val)
	for i := 0; i < ref.NumField(); i++ {
		if i == 0 {
			paramStr += "?"
		} else {
			paramStr += "&"
		}

		// build name for param
		fieldInfo := ref.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		v := ref.Field(i)
		switch v.Type().Kind() {
		case reflect.Array, reflect.Slice:
			for a := 0; a < v.Len(); a++ {
				if a != 0 {
					paramStr += "&"
				}
				paramStr += fmt.Sprintf("%s=%v", name, v.Index(a))
			}
		case reflect.Int, reflect.Bool, reflect.String:
			paramStr += fmt.Sprintf("%s=%v", name, v)

		default:
			return "", fmt.Errorf("unsupported kind %s", v.Type())
		}
	}

	return url + paramStr, nil
}

func validate(v string, matcher *regexp.Regexp) bool {
	return matcher.MatchString(v)
}

//!+Unpack

type field struct {
	v reflect.Value
	m *regexp.Regexp
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]field)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}

		// validation
		pattern := tag.Get("validation")
		matcher := regexp.MustCompile(pattern)

		fields[name] = field{v.Field(i), matcher}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.v.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.v.Kind() == reflect.Slice {
				elem := reflect.New(f.v.Type().Elem()).Elem()
				if f.m != nil && !f.m.MatchString(value) {
					return fmt.Errorf("validate error: %v", value)
				}
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.v.Set(reflect.Append(f.v, elem))
			} else {
				if f.m != nil && !f.m.MatchString(value) {
					return fmt.Errorf("validate error: %v", value)
				}
				if err := populate(f.v, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
