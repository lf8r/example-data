// Copyright (C) Subhajit DasGupta 2021

package data1

import "github.com/lf8r/dbgen-common/pkg/common"

type Periodical struct {
	common.Resource
	Title     string
	Author    string
	Issue     common.Time
	Publisher string
	Editions  []Edition
}
