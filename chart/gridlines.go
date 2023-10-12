// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package chart

import (
	"github.com/NekrozAriel/unioffice/drawing"
	"github.com/NekrozAriel/unioffice/schema/soo/dml"
	crt "github.com/NekrozAriel/unioffice/schema/soo/dml/chart"
)

type GridLines struct {
	x *crt.CT_ChartLines
}

// X returns the inner wrapped XML type.
func (g GridLines) X() *crt.CT_ChartLines {
	return g.x
}

func (g GridLines) Properties() drawing.ShapeProperties {
	if g.x.SpPr == nil {
		g.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(g.x.SpPr)
}
