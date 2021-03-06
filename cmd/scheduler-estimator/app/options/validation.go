package options

import (
	"net"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() field.ErrorList {
	errs := field.ErrorList{}

	newPath := field.NewPath("Options")
	if len(o.ClusterName) == 0 {
		errs = append(errs, field.Invalid(newPath.Child("ClusterName"), o.ClusterName, "clusterName cannot be empty"))
	}

	if net.ParseIP(o.BindAddress) == nil {
		errs = append(errs, field.Invalid(newPath.Child("BindAddress"), o.BindAddress, "not a valid textual representation of an IP address"))
	}

	if o.ServerPort < 0 || o.ServerPort > 65535 {
		errs = append(errs, field.Invalid(newPath.Child("ServerPort"), o.ServerPort, "must be a valid port between 0 and 65535 inclusive"))
	}

	if o.SecurePort < 0 || o.SecurePort > 65535 {
		errs = append(errs, field.Invalid(newPath.Child("SecurePort"), o.SecurePort, "must be a valid port between 0 and 65535 inclusive"))
	}

	return errs
}
