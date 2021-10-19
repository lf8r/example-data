// # Copyright (C) Subhajit DasGupta 2021

package data2

import (
	"github.com/lf8r/dbgen-common/pkg/common"
)

// AccessToken describes a token used to gain access to a system.
type AccessToken struct {
	common.Resource
	Val      string
	IssuedTo string
	Role     string
	Issuer   string
	Expires  string
}
