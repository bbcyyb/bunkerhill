package main

import (
	"fmt"
)

var ()

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintf(out, "[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
		end.Format("2006/01/02 - 15:04:05"),
		statusColor, statusCode, resetColor,
		latency,
		clientIP,
		methodColor, method, resetColor,
		path,
		comment,
	)
}
