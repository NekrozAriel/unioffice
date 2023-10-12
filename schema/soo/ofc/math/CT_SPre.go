// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"

	"github.com/NekrozAriel/unioffice"
)

type CT_SPre struct {
	SPrePr *CT_SPrePr
	Sub    *CT_OMathArg
	Sup    *CT_OMathArg
	E      *CT_OMathArg
}

func NewCT_SPre() *CT_SPre {
	ret := &CT_SPre{}
	ret.Sub = NewCT_OMathArg()
	ret.Sup = NewCT_OMathArg()
	ret.E = NewCT_OMathArg()
	return ret
}

func (m *CT_SPre) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SPrePr != nil {
		sesPrePr := xml.StartElement{Name: xml.Name{Local: "m:sPrePr"}}
		e.EncodeElement(m.SPrePr, sesPrePr)
	}
	sesub := xml.StartElement{Name: xml.Name{Local: "m:sub"}}
	e.EncodeElement(m.Sub, sesub)
	sesup := xml.StartElement{Name: xml.Name{Local: "m:sup"}}
	e.EncodeElement(m.Sup, sesup)
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SPre) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Sub = NewCT_OMathArg()
	m.Sup = NewCT_OMathArg()
	m.E = NewCT_OMathArg()
lCT_SPre:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sPrePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sPrePr"}:
				m.SPrePr = NewCT_SPrePr()
				if err := d.DecodeElement(m.SPrePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sub"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sub"}:
				if err := d.DecodeElement(m.Sub, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sup"}:
				if err := d.DecodeElement(m.Sup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SPre %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SPre
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SPre and its children
func (m *CT_SPre) Validate() error {
	return m.ValidateWithPath("CT_SPre")
}

// ValidateWithPath validates the CT_SPre and its children, prefixing error messages with path
func (m *CT_SPre) ValidateWithPath(path string) error {
	if m.SPrePr != nil {
		if err := m.SPrePr.ValidateWithPath(path + "/SPrePr"); err != nil {
			return err
		}
	}
	if err := m.Sub.ValidateWithPath(path + "/Sub"); err != nil {
		return err
	}
	if err := m.Sup.ValidateWithPath(path + "/Sup"); err != nil {
		return err
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	return nil
}
