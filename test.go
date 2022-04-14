//go:build test
// +build test

package goenvconfig

import "github.com/golang/glog"

func init() {
	Env = "test"
	glog.Infof("Initializing %s environment", Env)
}
