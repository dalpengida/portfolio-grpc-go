package config

import (
	"flag"
	"fmt"
)

var (
	Port = flag.Int("port", 9999, "The server port")
	Host = fmt.Sprintf("localhost:%d", *Port)
)
