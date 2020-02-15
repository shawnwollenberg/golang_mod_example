package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

/*
1) create hello.go
2) create hello_test.go
3) create mod - go mod init v0
4) imported rsc.io/quote in hello.go
5) go test
	-> go.mod updated for dependency
6) go list -m all
	-> shows golang.org/x/text has outdated version
7) to upgrade
	go get golang.org/x/text
8) go list -m all
	-> shows golang.org/x/text has been updated
*/
