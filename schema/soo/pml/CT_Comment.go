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
	"time"

	"github.com/NekrozAriel/unioffice"
	"github.com/NekrozAriel/unioffice/schema/soo/dml"
)

type CT_Comment struct {
	// Comment Author ID
	AuthorIdAttr uint32
	// Comment Date/Time
	DtAttr *time.Time
	// Comment Index
	IdxAttr uint32
	// Comment Position
	Pos *dml.CT_Point2D
	// Comment's Text Content
	Text   string
	ExtLst *CT_ExtensionListModify
}

func NewCT_Comment() *CT_Comment {
	ret := &CT_Comment{}
	ret.Pos = dml.NewCT_Point2D()
	return ret
}

func (m *CT_Comment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "authorId"},
		Value: fmt.Sprintf("%v", m.AuthorIdAttr)})
	if m.DtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dt"},
			Value: fmt.Sprintf("%v", *m.DtAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	e.EncodeToken(start)
	sepos := xml.StartElement{Name: xml.Name{Local: "p:pos"}}
	e.EncodeElement(m.Pos, sepos)
	setext := xml.StartElement{Name: xml.Name{Local: "p:text"}}
	unioffice.AddPreserveSpaceAttr(&setext, m.Text)
	e.EncodeElement(m.Text, setext)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Comment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = dml.NewCT_Point2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "authorId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.AuthorIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "dt" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DtAttr = &parsed
			continue
		}
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
			continue
		}
	}
lCT_Comment:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pos"}:
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "text"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "text"}:
				if err := d.DecodeElement(&m.Text, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Comment %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Comment
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Comment and its children
func (m *CT_Comment) Validate() error {
	return m.ValidateWithPath("CT_Comment")
}

// ValidateWithPath validates the CT_Comment and its children, prefixing error messages with path
func (m *CT_Comment) ValidateWithPath(path string) error {
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
