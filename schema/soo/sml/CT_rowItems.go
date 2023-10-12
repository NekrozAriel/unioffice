// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/NekrozAriel/unioffice"
)

type CT_rowItems struct {
	// Items in a Row Count
	CountAttr *uint32
	// Row Items
	I []*CT_I
}

func NewCT_rowItems() *CT_rowItems {
	ret := &CT_rowItems{}
	return ret
}

func (m *CT_rowItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sei := xml.StartElement{Name: xml.Name{Local: "ma:i"}}
	for _, c := range m.I {
		e.EncodeElement(c, sei)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_rowItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_rowItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "i"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "i"}:
				tmp := NewCT_I()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.I = append(m.I, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_rowItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_rowItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_rowItems and its children
func (m *CT_rowItems) Validate() error {
	return m.ValidateWithPath("CT_rowItems")
}

// ValidateWithPath validates the CT_rowItems and its children, prefixing error messages with path
func (m *CT_rowItems) ValidateWithPath(path string) error {
	for i, v := range m.I {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/I[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
