//go:build dev
// +build dev

package goenvconfig

import "github.com/golang/glog"

func init() {
	Env = "dev"
	glog.Infof("Initializing %s environment", Env)
}
