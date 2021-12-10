// Copyright (C) Subhajit DasGupta 2021

package data

import "github.com/lf8r/dbgen-common/pkg/common"

// Address models a United States Street Address with two additional fields.
type Address struct {
	Street1      string
	Street2      string
	City         string
	State        string
	ZIP          string
	SomeIntArray []int
	SomeStrArray []string
}

// Hobby models a hobby.
type Hobby struct {
	Group      bool
	Collecting bool
}

// Person models a person.
type Person struct {
	common.Resource
	Age     int
	Address Address
	Hobbies []Hobby
}

// Person2 models a person with an ID field, used for testing.
type Person2 struct {
	ID       string
	Age      int
	Name     string
	Created  common.Time
	Modified common.Time
	Address  Address
	Hobbies  []Hobby
}

// PersonWithPtr is a type for which DDL's cannot be generated since it contains a
// pointer field.
type PersonWithPtr struct {
	ID      string
	Name    string
	Address *Address
}
