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

	"github.com/NekrozAriel/unioffice"
	"github.com/NekrozAriel/unioffice/schema/soo/dml"
)

type CT_GroupShape struct {
	// Non-Visual Properties for a Group Shape
	NvGrpSpPr *CT_GroupShapeNonVisual
	// Group Shape Properties
	GrpSpPr *dml.CT_GroupShapeProperties
	Choice  []*CT_GroupShapeChoice
	ExtLst  *CT_ExtensionListModify
}

func NewCT_GroupShape() *CT_GroupShape {
	ret := &CT_GroupShape{}
	ret.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	ret.GrpSpPr = dml.NewCT_GroupShapeProperties()
	return ret
}

func (m *CT_GroupShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	senvGrpSpPr := xml.StartElement{Name: xml.Name{Local: "p:nvGrpSpPr"}}
	e.EncodeElement(m.NvGrpSpPr, senvGrpSpPr)
	segrpSpPr := xml.StartElement{Name: xml.Name{Local: "p:grpSpPr"}}
	e.EncodeElement(m.GrpSpPr, segrpSpPr)
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	m.GrpSpPr = dml.NewCT_GroupShapeProperties()
lCT_GroupShape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "nvGrpSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "nvGrpSpPr"}:
				if err := d.DecodeElement(m.NvGrpSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "grpSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "grpSpPr"}:
				if err := d.DecodeElement(m.GrpSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sp"}:
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.Sp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "grpSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "grpSp"}:
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.GrpSp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "graphicFrame"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "graphicFrame"}:
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.GraphicFrame, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cxnSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cxnSp"}:
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.CxnSp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pic"}:
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.Pic, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "contentPart"}:
				tmp := NewCT_GroupShapeChoice()
				if err := d.DecodeElement(&tmp.ContentPart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GroupShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShape and its children
func (m *CT_GroupShape) Validate() error {
	return m.ValidateWithPath("CT_GroupShape")
}

// ValidateWithPath validates the CT_GroupShape and its children, prefixing error messages with path
func (m *CT_GroupShape) ValidateWithPath(path string) error {
	if err := m.NvGrpSpPr.ValidateWithPath(path + "/NvGrpSpPr"); err != nil {
		return err
	}
	if err := m.GrpSpPr.ValidateWithPath(path + "/GrpSpPr"); err != nil {
		return err
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
