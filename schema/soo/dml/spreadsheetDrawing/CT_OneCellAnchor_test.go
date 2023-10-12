// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/NekrozAriel/unioffice/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_OneCellAnchorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_OneCellAnchor()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_OneCellAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_OneCellAnchor should validate: %s", err)
	}
}

func TestCT_OneCellAnchorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_OneCellAnchor()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_OneCellAnchor()
	xml.Unmarshal(buf, v2)
}
