// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/NekrozAriel/unioffice/schema/soo/sml"
)

func TestAG_RevDataConstructor(t *testing.T) {
	v := sml.NewAG_RevData()
	if v == nil {
		t.Errorf("sml.NewAG_RevData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.AG_RevData should validate: %s", err)
	}
}

func TestAG_RevDataMarshalUnmarshal(t *testing.T) {
	v := sml.NewAG_RevData()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewAG_RevData()
	xml.Unmarshal(buf, v2)
}
