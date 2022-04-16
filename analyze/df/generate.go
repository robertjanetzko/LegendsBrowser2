package df

import (
	"sort"

	"github.com/robertjanetzko/LegendsBrowser2/backend/util"
)

type Object struct {
	Name      string           `json:"name"`
	Id        bool             `json:"id,omitempty"`
	Named     bool             `json:"named,omitempty"`
	Typed     bool             `json:"typed,omitempty"`
	SubTypes  *[]Subtype       `json:"subtypes,omitempty"`
	SubTypeOf *string          `json:"subtypeof,omitempty"`
	Fields    map[string]Field `json:"fields"`
}

type Subtype struct {
	Name     string `json:"name"`
	BaseType string `json:"base"`
	PlusType string `json:"plus"`
}

type Field struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Multiple    bool    `json:"multiple,omitempty"`
	ElementType *string `json:"elements,omitempty"`
	Legend      string  `json:"legend"`
	Base        bool
	Plus        bool
}

func (f Field) Active(plus bool) bool {
	if plus && f.Plus {
		return true
	}
	if !plus && f.Base {
		return true
	}
	return false
}

type ActiveSubType struct {
	Case    string
	Name    string
	Options []string
}

func (f Object) ActiveSubTypes(plus bool) []*ActiveSubType {
	subs := make(map[string]*ActiveSubType)

	for _, s := range *f.SubTypes {
		if !plus && s.BaseType == "" {
			continue
		}
		if plus && s.PlusType == "" {
			continue
		}

		a := ActiveSubType{}
		if plus {
			a.Case = s.PlusType
		} else {
			a.Case = s.BaseType
		}
		a.Name = f.Name + s.Name
		a.Options = append(a.Options, a.Name)

		if sub, ok := subs[a.Case]; ok {
			sub.Options = append(sub.Options, a.Name)
		} else {
			subs[a.Case] = &a
		}
	}

	list := util.Values(subs)
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Case < list[j].Case
	})

	return list
}
