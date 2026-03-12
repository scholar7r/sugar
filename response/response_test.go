package response_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/scholar7r/sugar/response"
)

type data struct {
	Name   string `json:"name"`
	Gender bool   `json:"gender"`
	Age    int    `json:"age"`
}

func TestNew(t *testing.T) {
	want := &response.Response[*data]{
		Code:    http.StatusOK,
		Message: "success",
		Data: &data{
			Name:   "scholar7r",
			Gender: true,
			Age:    7,
		},
	}

	got := response.New[*data](http.StatusOK, "success", &data{
		Name:   "scholar7r",
		Gender: true,
		Age:    7,
	})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
