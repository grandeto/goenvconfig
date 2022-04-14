//go:build prod
// +build prod

package goenvconfig

import "github.com/golang/glog"

func init() {
	Env = "prod"
	glog.Infof("Initializing %s environment", Env)
}
