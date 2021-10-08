// Copyright (C) Subhajit DasGupta 2021

package data1

import "time"

// Resource is a base type which must be embedded in persisted structs.
type Resource struct {
	ID       string
	Name     string
	Created  time.Time
	Modified time.Time
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
	Resource
	Title     string
	Author    string
	Publisher string
	Editions  []Edition
}
