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

type CT_PCDSDTCEntries struct {
	// Tuple Count
	CountAttr *uint32
	// No Value
	M []*CT_Missing
	// Numeric Value
	N []*CT_Number
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
}

func NewCT_PCDSDTCEntries() *CT_PCDSDTCEntries {
	ret := &CT_PCDSDTCEntries{}
	return ret
}

func (m *CT_PCDSDTCEntries) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "ma:m"}}
		for _, c := range m.M {
			e.EncodeElement(c, sem)
		}
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "ma:n"}}
		for _, c := range m.N {
			e.EncodeElement(c, sen)
		}
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "ma:e"}}
		for _, c := range m.E {
			e.EncodeElement(c, see)
		}
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "ma:s"}}
		for _, c := range m.S {
			e.EncodeElement(c, ses)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PCDSDTCEntries) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PCDSDTCEntries:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "m"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "m"}:
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "n"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "n"}:
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "e"}:
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "s"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "s"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_PCDSDTCEntries %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PCDSDTCEntries
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PCDSDTCEntries and its children
func (m *CT_PCDSDTCEntries) Validate() error {
	return m.ValidateWithPath("CT_PCDSDTCEntries")
}

// ValidateWithPath validates the CT_PCDSDTCEntries and its children, prefixing error messages with path
func (m *CT_PCDSDTCEntries) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
