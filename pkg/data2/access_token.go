// # Copyright (C) Subhajit DasGupta 2021

package data2

import (
	"github.com/lf8r/example-data/pkg/common"
)

// Resource is a base type which must be embedded in persisted structs.
type Base struct {
	ID       string
	Name     string
	Created  common.Time
	Modified common.Time
}

// AccessToken describes a token used to gain access to a system.
type AccessToken struct {
	Base
	Val      string
	IssuedTo string
	Role     string
	Issuer   string
	Expires  string
}
