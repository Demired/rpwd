package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
	"github.com/ogier/pflag"
)

var buildStamp = "beta"

func main(){
	var sUpper,sLower,sNumber,sVersion,sCharacter = false,false,false,false,false
	var sLength = 16
	pflag.BoolVarP(&sVersion,"version","v",sVersion,"version")
	pflag.BoolVarP(&sUpper,"Upper","u",sUpper,"upper")
	pflag.BoolVarP(&sLower,"Lower","l",sLower,"lower")
	pflag.BoolVarP(&sNumber,"Number","n",sNumber,"number")
	pflag.BoolVarP(&sCharacter,"Character","c",sCharacter,"character")
	pflag.IntVarP(&sLength,"Length","e",sLength,"length")
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [--Length 15] [--Upper][--Lower][--Number][--Character]\n", os.Args[0])
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if sVersion {
		fmt.Fprintf(os.Stdout, "RandPasswd for Linux/Mac/BSD, buildStamp is :%s.\n(C) 2015 - 2017 IPIP.NET. All Rights Reversed.\n", buildStamp)
		return
	}

	var str = ""

	if sUpper {
		str += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	
	if sLower {
		str += "abcdefghijklmnopqrstuvwxyz"
	}

	if sNumber {
		str += "0123456789"
	}

	if sCharacter {
		str += "!@#$%^&*()"
	}	

	if !sUpper && !sLower && !sNumber && !sCharacter{
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < sLength; i++ {
	   result = append(result, bytes[r.Intn(len(bytes))])
	}
	fmt.Fprintf(os.Stdout,"%s\n",result)
}