// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing

import (
	"encoding/xml"

	"github.com/NekrozAriel/unioffice"
	"github.com/NekrozAriel/unioffice/schema/soo/dml"
)

type CT_ConnectorNonVisual struct {
	CNvPr      *dml.CT_NonVisualDrawingProps
	CNvCxnSpPr *dml.CT_NonVisualConnectorProperties
}

func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	ret := &CT_ConnectorNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvCxnSpPr = dml.NewCT_NonVisualConnectorProperties()
	return ret
}

func (m *CT_ConnectorNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "xdr:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "xdr:cNvCxnSpPr"}}
	e.EncodeElement(m.CNvCxnSpPr, secNvCxnSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConnectorNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvCxnSpPr = dml.NewCT_NonVisualConnectorProperties()
lCT_ConnectorNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cNvCxnSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cNvCxnSpPr"}:
				if err := d.DecodeElement(m.CNvCxnSpPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ConnectorNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectorNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConnectorNonVisual and its children
func (m *CT_ConnectorNonVisual) Validate() error {
	return m.ValidateWithPath("CT_ConnectorNonVisual")
}

// ValidateWithPath validates the CT_ConnectorNonVisual and its children, prefixing error messages with path
func (m *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvCxnSpPr.ValidateWithPath(path + "/CNvCxnSpPr"); err != nil {
		return err
	}
	return nil
}
