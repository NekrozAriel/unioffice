// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"

	"github.com/NekrozAriel/unioffice"
)

type CT_FramesetSplitbar struct {
	// Frameset Splitter Width
	W *CT_TwipsMeasure
	// Frameset Splitter Color
	Color *CT_Color
	// Do Not Display Frameset Splitters
	NoBorder *CT_OnOff
	// Frameset Splitter Border Style
	FlatBorders *CT_OnOff
}

func NewCT_FramesetSplitbar() *CT_FramesetSplitbar {
	ret := &CT_FramesetSplitbar{}
	return ret
}

func (m *CT_FramesetSplitbar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.W != nil {
		sew := xml.StartElement{Name: xml.Name{Local: "w:w"}}
		e.EncodeElement(m.W, sew)
	}
	if m.Color != nil {
		secolor := xml.StartElement{Name: xml.Name{Local: "w:color"}}
		e.EncodeElement(m.Color, secolor)
	}
	if m.NoBorder != nil {
		senoBorder := xml.StartElement{Name: xml.Name{Local: "w:noBorder"}}
		e.EncodeElement(m.NoBorder, senoBorder)
	}
	if m.FlatBorders != nil {
		seflatBorders := xml.StartElement{Name: xml.Name{Local: "w:flatBorders"}}
		e.EncodeElement(m.FlatBorders, seflatBorders)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_FramesetSplitbar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FramesetSplitbar:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "w"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "w"}:
				m.W = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.W, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "color"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "color"}:
				m.Color = NewCT_Color()
				if err := d.DecodeElement(m.Color, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noBorder"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noBorder"}:
				m.NoBorder = NewCT_OnOff()
				if err := d.DecodeElement(m.NoBorder, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "flatBorders"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "flatBorders"}:
				m.FlatBorders = NewCT_OnOff()
				if err := d.DecodeElement(m.FlatBorders, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_FramesetSplitbar %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FramesetSplitbar
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FramesetSplitbar and its children
func (m *CT_FramesetSplitbar) Validate() error {
	return m.ValidateWithPath("CT_FramesetSplitbar")
}

// ValidateWithPath validates the CT_FramesetSplitbar and its children, prefixing error messages with path
func (m *CT_FramesetSplitbar) ValidateWithPath(path string) error {
	if m.W != nil {
		if err := m.W.ValidateWithPath(path + "/W"); err != nil {
			return err
		}
	}
	if m.Color != nil {
		if err := m.Color.ValidateWithPath(path + "/Color"); err != nil {
			return err
		}
	}
	if m.NoBorder != nil {
		if err := m.NoBorder.ValidateWithPath(path + "/NoBorder"); err != nil {
			return err
		}
	}
	if m.FlatBorders != nil {
		if err := m.FlatBorders.ValidateWithPath(path + "/FlatBorders"); err != nil {
			return err
		}
	}
	return nil
}
