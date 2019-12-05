package params

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	type Test struct {
		Num      int    `validation:"^([01][0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		Email    string `http:"mail" validation:"[^@]+@[^\\.]+\\..+"`
		Bl       bool
		NumArray []int `validation:"^([0-9])$"`
	}

	data := Test{200, "foobar@gmail.com", true, []int{1, 2, 4, 8}}
	ok, err := Pack("http://example.com/api", data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ok)

	unpacked := Test{}
	req := httptest.NewRequest("GET", ok, nil)
	err = Unpack(req, &unpacked)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(data, unpacked) {
		t.Fatalf("%v != %v", data, unpacked)
	}

	// Error test 1
	data = Test{200, "foobargmail.com", true, []int{1, 2, 4, 8}}
	ng, err := Pack("http://example.com/api", data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ng)

	unpacked = Test{}
	req = httptest.NewRequest("GET", ng, nil)
	err = Unpack(req, &unpacked)
	if err == nil {
		t.Fatal("An error should be raised")
	}
	t.Log(err)

	// Error test 2
	data = Test{800, "foobar@gmail.com", true, []int{1, 2, 4, 8, 16, 32}}
	ng, err = Pack("http://example.com/api", data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ng)

	unpacked = Test{}
	req = httptest.NewRequest("GET", ng, nil)
	err = Unpack(req, &unpacked)
	if err == nil {
		t.Fatal("An error should be raised")
	}
	t.Log(err)
}
