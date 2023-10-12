// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"

	"github.com/NekrozAriel/unioffice"
)

type CT_GvmlPictureNonVisual struct {
	CNvPr    *CT_NonVisualDrawingProps
	CNvPicPr *CT_NonVisualPictureProperties
}

func NewCT_GvmlPictureNonVisual() *CT_GvmlPictureNonVisual {
	ret := &CT_GvmlPictureNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvPicPr = NewCT_NonVisualPictureProperties()
	return ret
}

func (m *CT_GvmlPictureNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvPicPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPicPr"}}
	e.EncodeElement(m.CNvPicPr, secNvPicPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlPictureNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvPicPr = NewCT_NonVisualPictureProperties()
lCT_GvmlPictureNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvPicPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvPicPr"}:
				if err := d.DecodeElement(m.CNvPicPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GvmlPictureNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlPictureNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlPictureNonVisual and its children
func (m *CT_GvmlPictureNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlPictureNonVisual")
}

// ValidateWithPath validates the CT_GvmlPictureNonVisual and its children, prefixing error messages with path
func (m *CT_GvmlPictureNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvPicPr.ValidateWithPath(path + "/CNvPicPr"); err != nil {
		return err
	}
	return nil
}
