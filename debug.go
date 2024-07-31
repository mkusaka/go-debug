// ref. https://x.com/Linda_pp/status/1815586554041311566
package debug

import (
	"fmt"
	"os"
	"runtime"
)

func Dbg[T any](v T) T {
	_, f, l, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "[%s:%d] %#v\n", f, l, v)
	return v
}
