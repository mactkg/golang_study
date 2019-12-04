package params

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestParams(t *testing.T) {
	type Test struct {
		Num      int `http:"nyamu"` // tag test
		Str      string
		Bl       bool `http:"💡"` // tag test 2
		NumArray []int
	}

	data := Test{42, "hello", true, []int{1, 2, 4, 8, 16, 32}}
	str, err := Pack("http://example.com/api", data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)

	unpacked := Test{}
	req := httptest.NewRequest("GET", str, nil)
	err = Unpack(req, &unpacked)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(data, unpacked) {
		t.Fatalf("%v != %v", data, unpacked)
	}
}
