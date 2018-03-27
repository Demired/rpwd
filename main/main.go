package main

import (
	"fmt"
	"os"
	"rpwd"

	"github.com/ogier/pflag"
)

var buildStamp = "beta"

func main() {
	var sUpper, sLower, sNumber, sVersion, sCharacter = false, false, false, false, false
	var sLength = 16
	pflag.BoolVarP(&sVersion, "version", "v", sVersion, "version")
	pflag.BoolVarP(&sUpper, "Upper", "u", sUpper, "upper")
	pflag.BoolVarP(&sLower, "Lower", "l", sLower, "lower")
	pflag.BoolVarP(&sNumber, "Number", "n", sNumber, "number")
	pflag.BoolVarP(&sCharacter, "Character", "c", sCharacter, "character")
	pflag.IntVarP(&sLength, "Length", "e", sLength, "length")
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [--Length 15] [--Upper][--Lower][--Number][--Character]\n", os.Args[0])
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if sVersion {
		fmt.Fprintf(os.Stdout, "RandPasswd for Linux/Mac/BSD, buildStamp is :%s.\n(C) 2015 - 2017 IPIP.NET. All Rights Reversed.\n", buildStamp)
		return
	}
	fmt.Println(rpwd.Init(sLength, sUpper, sLower, sNumber, sCharacter))
}
