package errorhandling

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// type error interface {
//     Error() string
// }

func ErrorBasicsPractice() {
	errorExample()
	extractErrorAssertingStructFields()
	extractErrorAssertingStructMethods()
	extractErrorThroughDirectComparison()

}

func errorExample() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// type PathError struct {
//     Op   string
//     Path string
//     Err  error
// }

// func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func extractErrorAssertingStructFields() {
	f, err := os.Open("test.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

// type DNSError struct {
//     ...
// }

// func (e *DNSError) Error() string {
//     ...
// }
// func (e *DNSError) Timeout() bool {
//     ...
// }
// func (e *DNSError) Temporary() bool {
//     ...
// }

func extractErrorAssertingStructMethods() {
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}

// var ErrBadPattern = errors.New("syntax error in pattern")
func extractErrorThroughDirectComparison() {
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)
}
