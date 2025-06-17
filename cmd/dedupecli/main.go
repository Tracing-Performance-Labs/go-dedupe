package main

import (
	"flag"
	"fmt"

	"github.com/Tracing-Performance-Labs/go-dedupe"
)

func main() {
	var repr string
	flag.StringVar(&repr, "repr", "default", "Object representation to use (default, murmur)")

	flag.Parse()

	s := flag.Arg(0)
	if s == "" {
		flag.Usage()
		return
	}

	var oRepr dedupe.ObjectRepr[string]

	switch repr {
	case "default":
		oRepr = dedupe.NewDefaultObjectRepr()
	case "murmur":
		oRepr = dedupe.NewMurmurRepr()
	default:
		panic("Invalid representation: " + repr)
	}

	result := oRepr.GetRepr(s)

	fmt.Printf("= %s\n\t%d bytes\n", result, len(result))
}
