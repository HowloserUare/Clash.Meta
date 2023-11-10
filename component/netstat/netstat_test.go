package netstat

import (
	"fmt"
	"testing"
)

func TestLookupPort(t *testing.T) {
	data := LookupPort("::1", "38484")
	fmt.Println(data)
	t.Fatal(nil)
}
