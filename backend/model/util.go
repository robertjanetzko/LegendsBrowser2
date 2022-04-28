package model

import "encoding/xml"

type NamedObject struct {
	Id_   int    `xml:"id" json:"id"`
	Name_ string `xml:"name" json:"name"`
}

func (r *NamedObject) Id() int      { return r.Id_ }
func (r *NamedObject) Name() string { return r.Name_ }

type Named interface {
	Id() int
	Name() string
}

type Identifiable interface {
	Id() int
	setId(int)
}

type Typed interface {
	Type() string
}

type OtherElements struct {
	Others_ []Element `xml:",any" json:"-"`
}

func (r *OtherElements) Others() []Element { return r.Others_ }

type Others interface {
	Others() []Element
}

type TypedOthers interface {
	Type() string
	Others() []Element
}

type Element struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}
