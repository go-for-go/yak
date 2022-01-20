package Pinger

import (
	"fmt"
	"testing"
)

func TestListerAndServe(t *testing.T) {
	l1, err := ListerAndServe(":11111")
	if err != nil {
		t.Fatal("unable to serve")
	}
	defer l1.Close()

	l2, err := ListerAndServe(":11111")
	if err == nil {
		defer l2.Close()
		t.Fatal("something wrong, started server on some port")
	}
	// binding on some port will be return error
	fmt.Println(err.Error())
}
