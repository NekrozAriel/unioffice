// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package core_properties

import (
	"encoding/xml"
	"fmt"

	"github.com/NekrozAriel/unioffice"
)

type CT_Keywords struct {
	LangAttr *string
	Value    []*CT_Keyword
}

func NewCT_Keywords() *CT_Keywords {
	ret := &CT_Keywords{}
	return ret
}

func (m *CT_Keywords) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xml:lang"},
			Value: fmt.Sprintf("%v", *m.LangAttr)})
	}
	e.EncodeToken(start)
	if m.Value != nil {
		sevalue := xml.StartElement{Name: xml.Name{Local: "cp:value"}}
		for _, c := range m.Value {
			e.EncodeElement(c, sevalue)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Keywords) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://www.w3.org/XML/1998/namespace" && attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
			continue
		}
	}
lCT_Keywords:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties", Local: "value"}:
				tmp := NewCT_Keyword()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Value = append(m.Value, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Keywords %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Keywords
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Keywords and its children
func (m *CT_Keywords) Validate() error {
	return m.ValidateWithPath("CT_Keywords")
}

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (m *CT_Keywords) ValidateWithPath(path string) error {
	for i, v := range m.Value {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Value[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
