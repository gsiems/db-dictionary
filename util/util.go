// Package util contains the utility functions used by db-dictionary
package util

import (
	"fmt"
	"os"
)

// Coalesce picks the first non-empty string from a list
func Coalesce(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

func FailOnErr(verbose bool, err error) {
	if err != nil {
		Carp(verbose, err)
		os.Exit(1)
	}
}

func Carp(verbose bool, err error) {
	if err != nil {
		if verbose {
			os.Stderr.WriteString(fmt.Sprintf("%s\n", err))
		}
	}
}
