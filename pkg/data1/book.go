// Copyright (C) Subhajit DasGupta 2021

package data1

import (
	"github.com/lf8r/dbgen-common/pkg/common"
)

// Resource is a base type which must be embedded in persisted structs.
type Base struct {
	ID       string
	Name     string
	Created  common.Time
	Modified common.Time
}

type Seller struct {
	Name         string
	PurchaseLink string
}

type Edition struct {
	ISBN   string
	Format string
}

type Book struct {
	Base
	Title     string
	Author    string
	Publisher string
	Editions  []Edition
}
