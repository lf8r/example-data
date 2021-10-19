// Copyright (C) Subhajit DasGupta 2021

package data1

import "github.com/lf8r/dbgen-common/pkg/common"

type Seller struct {
	Name         string
	PurchaseLink string
}

type Edition struct {
	ISBN   string
	Format string
}

type Book struct {
	common.Resource
	Title     string
	Author    string
	Publisher string
	Editions  []Edition
}
