// # Copyright (C) Subhajit DasGupta 2021

package data2

import "time"

// Resource is a base type which must be embedded in persisted structs.
type Base struct {
	ID       string
	Name     string
	Created  time.Time
	Modified time.Time
}

// AccessToken describes a token used to gain access to a system.
type AccessToken struct {
	Val      string
	IssuedTo string
	Role     string
	Issuer   string
	Expires  string
}
