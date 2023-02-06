package passport_client

import (
	"fmt"
	"testing"
)

func TestQscLogin(t *testing.T) {
	url = "https://cystudio.tech/passport"
	cookies, err := QscLogin("superuser", "123456")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cookies)
}

func TestGetProfile(t *testing.T) {
	url = "https://cystudio.tech/passport"
	cookies, err := QscLogin("superuser", "123456")
	if err != nil {
		t.Error(err)
	}
	user, err := GetProfile(cookies)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)
}
