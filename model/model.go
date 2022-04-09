package model

import "encoding/xml"

type Region struct {
	XMLName xml.Name `xml:"region" json:"-"`
	NamedObject
	Type string `xml:"type" json:"type"`
}

type UndergroundRegion struct {
	XMLName xml.Name `xml:"underground_region" json:"-"`
	NamedObject
	Type string `xml:"type" json:"type"`
}

type Landmass struct {
	XMLName xml.Name `xml:"landmass" json:"-"`
	NamedObject
}

type Site struct {
	XMLName xml.Name `xml:"site" json:"-"`
	NamedObject
	Type       string      `xml:"type" json:"type"`
	Coords     string      `xml:"coords" json:"coords"`
	Rectangle  string      `xml:"rectangle" json:"rectangle"`
	Structures []Structure `xml:"structures>structure" json:"structures"`

	EventObject
}

// func (obj Site) id() int      { return obj.Id }
// func (obj Site) name() string { return obj.Name }

type Structure struct {
	XMLName xml.Name `xml:"structure" json:"-"`
	LocalId int      `xml:"local_id" json:"localId"`
	Name    string   `xml:"name" json:"name"`
	Type    string   `xml:"type" json:"type"`
}

type WorldConstruction struct {
	XMLName xml.Name `xml:"world_construction" json:"-"`
	NamedObject
}

type Artifact struct {
	XMLName xml.Name `xml:"artifact" json:"-"`
	NamedObject
	SiteId int `xml:"site_id" json:"siteId"`

	EventObject
}

type HistoricalFigure struct {
	XMLName xml.Name `xml:"historical_figure" json:"-"`
	NamedObject
	Race  string `xml:"race" json:"race"`
	Caste string `xml:"caste" json:"caste"`
	OtherElements

	EventObject
}

func (r *HistoricalFigure) Type() string { return "hf" }

type HistoricalEventCollection struct {
	XMLName xml.Name `xml:"historical_event_collection" json:"-"`
	NamedObject
	StartYear    int    `xml:"year"`
	StartSeconds int    `xml:"seconds72"`
	EndYear      int    `xml:"end_year"`
	EndSeconds   int    `xml:"end_seconds72"`
	Type         string `xml:"type" json:"type"`
	EventIds     []int  `xml:"event" json:"eventIds"`
}

type Entity struct {
	XMLName xml.Name `xml:"entity" json:"-"`
	NamedObject
	EventObject
}
