package indexa

import (
	"reflect"
	"testing"
)

func TestCreateNewClient(t *testing.T) {
	baseURL := "http://foo.com"
	authToken := "some-auth-token"

	got := NewIndexaClient(baseURL, authToken)
	expected := IndexaClient{
		baseURL:   baseURL,
		authToken: authToken,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got '%v', expected '%v'", got, expected)
	}
}
