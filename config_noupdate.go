// +build noupdate

package main

import "os"

func init() {
	os.Setenv("GH_AUTOUPDATE", "never")
}
