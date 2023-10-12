// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"

	"github.com/NekrozAriel/unioffice"
	"github.com/NekrozAriel/unioffice/schema/soo/dml"
)

type CT_DateAx struct {
	AxId           *CT_UnsignedInt
	Scaling        *CT_Scaling
	Delete         *CT_Boolean
	AxPos          *CT_AxPos
	MajorGridlines *CT_ChartLines
	MinorGridlines *CT_ChartLines
	Title          *CT_Title
	NumFmt         *CT_NumFmt
	MajorTickMark  *CT_TickMark
	MinorTickMark  *CT_TickMark
	TickLblPos     *CT_TickLblPos
	SpPr           *dml.CT_ShapeProperties
	TxPr           *dml.CT_TextBody
	CrossAx        *CT_UnsignedInt
	Choice         *EG_AxSharedChoice
	Auto           *CT_Boolean
	LblOffset      *CT_LblOffset
	BaseTimeUnit   *CT_TimeUnit
	MajorUnit      *CT_AxisUnit
	MajorTimeUnit  *CT_TimeUnit
	MinorUnit      *CT_AxisUnit
	MinorTimeUnit  *CT_TimeUnit
	ExtLst         *CT_ExtensionList
}

func NewCT_DateAx() *CT_DateAx {
	ret := &CT_DateAx{}
	ret.AxId = NewCT_UnsignedInt()
	ret.Scaling = NewCT_Scaling()
	ret.AxPos = NewCT_AxPos()
	ret.CrossAx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_DateAx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	e.EncodeElement(m.AxId, seaxId)
	sescaling := xml.StartElement{Name: xml.Name{Local: "c:scaling"}}
	e.EncodeElement(m.Scaling, sescaling)
	if m.Delete != nil {
		sedelete := xml.StartElement{Name: xml.Name{Local: "c:delete"}}
		e.EncodeElement(m.Delete, sedelete)
	}
	seaxPos := xml.StartElement{Name: xml.Name{Local: "c:axPos"}}
	e.EncodeElement(m.AxPos, seaxPos)
	if m.MajorGridlines != nil {
		semajorGridlines := xml.StartElement{Name: xml.Name{Local: "c:majorGridlines"}}
		e.EncodeElement(m.MajorGridlines, semajorGridlines)
	}
	if m.MinorGridlines != nil {
		seminorGridlines := xml.StartElement{Name: xml.Name{Local: "c:minorGridlines"}}
		e.EncodeElement(m.MinorGridlines, seminorGridlines)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "c:title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "c:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.MajorTickMark != nil {
		semajorTickMark := xml.StartElement{Name: xml.Name{Local: "c:majorTickMark"}}
		e.EncodeElement(m.MajorTickMark, semajorTickMark)
	}
	if m.MinorTickMark != nil {
		seminorTickMark := xml.StartElement{Name: xml.Name{Local: "c:minorTickMark"}}
		e.EncodeElement(m.MinorTickMark, seminorTickMark)
	}
	if m.TickLblPos != nil {
		setickLblPos := xml.StartElement{Name: xml.Name{Local: "c:tickLblPos"}}
		e.EncodeElement(m.TickLblPos, setickLblPos)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	secrossAx := xml.StartElement{Name: xml.Name{Local: "c:crossAx"}}
	e.EncodeElement(m.CrossAx, secrossAx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.Auto != nil {
		seauto := xml.StartElement{Name: xml.Name{Local: "c:auto"}}
		e.EncodeElement(m.Auto, seauto)
	}
	if m.LblOffset != nil {
		selblOffset := xml.StartElement{Name: xml.Name{Local: "c:lblOffset"}}
		e.EncodeElement(m.LblOffset, selblOffset)
	}
	if m.BaseTimeUnit != nil {
		sebaseTimeUnit := xml.StartElement{Name: xml.Name{Local: "c:baseTimeUnit"}}
		e.EncodeElement(m.BaseTimeUnit, sebaseTimeUnit)
	}
	if m.MajorUnit != nil {
		semajorUnit := xml.StartElement{Name: xml.Name{Local: "c:majorUnit"}}
		e.EncodeElement(m.MajorUnit, semajorUnit)
	}
	if m.MajorTimeUnit != nil {
		semajorTimeUnit := xml.StartElement{Name: xml.Name{Local: "c:majorTimeUnit"}}
		e.EncodeElement(m.MajorTimeUnit, semajorTimeUnit)
	}
	if m.MinorUnit != nil {
		seminorUnit := xml.StartElement{Name: xml.Name{Local: "c:minorUnit"}}
		e.EncodeElement(m.MinorUnit, seminorUnit)
	}
	if m.MinorTimeUnit != nil {
		seminorTimeUnit := xml.StartElement{Name: xml.Name{Local: "c:minorTimeUnit"}}
		e.EncodeElement(m.MinorTimeUnit, seminorTimeUnit)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DateAx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AxId = NewCT_UnsignedInt()
	m.Scaling = NewCT_Scaling()
	m.AxPos = NewCT_AxPos()
	m.CrossAx = NewCT_UnsignedInt()
lCT_DateAx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axId"}:
				if err := d.DecodeElement(m.AxId, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "scaling"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "scaling"}:
				if err := d.DecodeElement(m.Scaling, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "delete"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "delete"}:
				m.Delete = NewCT_Boolean()
				if err := d.DecodeElement(m.Delete, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axPos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axPos"}:
				if err := d.DecodeElement(m.AxPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "majorGridlines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "majorGridlines"}:
				m.MajorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MajorGridlines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minorGridlines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minorGridlines"}:
				m.MinorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MinorGridlines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "title"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "title"}:
				m.Title = NewCT_Title()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numFmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numFmt"}:
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "majorTickMark"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "majorTickMark"}:
				m.MajorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MajorTickMark, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minorTickMark"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minorTickMark"}:
				m.MinorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MinorTickMark, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "tickLblPos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "tickLblPos"}:
				m.TickLblPos = NewCT_TickLblPos()
				if err := d.DecodeElement(m.TickLblPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crossAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crossAx"}:
				if err := d.DecodeElement(m.CrossAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crosses"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crosses"}:
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.Crosses, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crossesAt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crossesAt"}:
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.CrossesAt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "auto"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "auto"}:
				m.Auto = NewCT_Boolean()
				if err := d.DecodeElement(m.Auto, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "lblOffset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "lblOffset"}:
				m.LblOffset = NewCT_LblOffset()
				if err := d.DecodeElement(m.LblOffset, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "baseTimeUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "baseTimeUnit"}:
				m.BaseTimeUnit = NewCT_TimeUnit()
				if err := d.DecodeElement(m.BaseTimeUnit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "majorUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "majorUnit"}:
				m.MajorUnit = NewCT_AxisUnit()
				if err := d.DecodeElement(m.MajorUnit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "majorTimeUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "majorTimeUnit"}:
				m.MajorTimeUnit = NewCT_TimeUnit()
				if err := d.DecodeElement(m.MajorTimeUnit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minorUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minorUnit"}:
				m.MinorUnit = NewCT_AxisUnit()
				if err := d.DecodeElement(m.MinorUnit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minorTimeUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minorTimeUnit"}:
				m.MinorTimeUnit = NewCT_TimeUnit()
				if err := d.DecodeElement(m.MinorTimeUnit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_DateAx %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DateAx
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DateAx and its children
func (m *CT_DateAx) Validate() error {
	return m.ValidateWithPath("CT_DateAx")
}

// ValidateWithPath validates the CT_DateAx and its children, prefixing error messages with path
func (m *CT_DateAx) ValidateWithPath(path string) error {
	if err := m.AxId.ValidateWithPath(path + "/AxId"); err != nil {
		return err
	}
	if err := m.Scaling.ValidateWithPath(path + "/Scaling"); err != nil {
		return err
	}
	if m.Delete != nil {
		if err := m.Delete.ValidateWithPath(path + "/Delete"); err != nil {
			return err
		}
	}
	if err := m.AxPos.ValidateWithPath(path + "/AxPos"); err != nil {
		return err
	}
	if m.MajorGridlines != nil {
		if err := m.MajorGridlines.ValidateWithPath(path + "/MajorGridlines"); err != nil {
			return err
		}
	}
	if m.MinorGridlines != nil {
		if err := m.MinorGridlines.ValidateWithPath(path + "/MinorGridlines"); err != nil {
			return err
		}
	}
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.MajorTickMark != nil {
		if err := m.MajorTickMark.ValidateWithPath(path + "/MajorTickMark"); err != nil {
			return err
		}
	}
	if m.MinorTickMark != nil {
		if err := m.MinorTickMark.ValidateWithPath(path + "/MinorTickMark"); err != nil {
			return err
		}
	}
	if m.TickLblPos != nil {
		if err := m.TickLblPos.ValidateWithPath(path + "/TickLblPos"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if err := m.CrossAx.ValidateWithPath(path + "/CrossAx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.Auto != nil {
		if err := m.Auto.ValidateWithPath(path + "/Auto"); err != nil {
			return err
		}
	}
	if m.LblOffset != nil {
		if err := m.LblOffset.ValidateWithPath(path + "/LblOffset"); err != nil {
			return err
		}
	}
	if m.BaseTimeUnit != nil {
		if err := m.BaseTimeUnit.ValidateWithPath(path + "/BaseTimeUnit"); err != nil {
			return err
		}
	}
	if m.MajorUnit != nil {
		if err := m.MajorUnit.ValidateWithPath(path + "/MajorUnit"); err != nil {
			return err
		}
	}
	if m.MajorTimeUnit != nil {
		if err := m.MajorTimeUnit.ValidateWithPath(path + "/MajorTimeUnit"); err != nil {
			return err
		}
	}
	if m.MinorUnit != nil {
		if err := m.MinorUnit.ValidateWithPath(path + "/MinorUnit"); err != nil {
			return err
		}
	}
	if m.MinorTimeUnit != nil {
		if err := m.MinorTimeUnit.ValidateWithPath(path + "/MinorTimeUnit"); err != nil {
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
