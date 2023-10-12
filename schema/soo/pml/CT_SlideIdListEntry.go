// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/NekrozAriel/unioffice"
)

type CT_SlideIdListEntry struct {
	// Slide Identifier
	IdAttr  uint32
	RIdAttr string
	ExtLst  *CT_ExtensionList
}

func NewCT_SlideIdListEntry() *CT_SlideIdListEntry {
	ret := &CT_SlideIdListEntry{}
	ret.IdAttr = 256
	return ret
}

func (m *CT_SlideIdListEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.RIdAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideIdListEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.IdAttr = 256
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RIdAttr = parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
			continue
		}
	}
lCT_SlideIdListEntry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SlideIdListEntry %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideIdListEntry
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideIdListEntry and its children
func (m *CT_SlideIdListEntry) Validate() error {
	return m.ValidateWithPath("CT_SlideIdListEntry")
}

// ValidateWithPath validates the CT_SlideIdListEntry and its children, prefixing error messages with path
func (m *CT_SlideIdListEntry) ValidateWithPath(path string) error {
	if m.IdAttr < 256 {
		return fmt.Errorf("%s/m.IdAttr must be >= 256 (have %v)", path, m.IdAttr)
	}
	if m.IdAttr >= 2147483648 {
		return fmt.Errorf("%s/m.IdAttr must be < 2147483648 (have %v)", path, m.IdAttr)
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
