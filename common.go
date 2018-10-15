package stix

import (
	// "strings"
	"time"
)

type Common struct {
	Type               string              `json:"type"`
	Id                 ObjectIdentifier    `json:"id"`
	Created            time.Time           `json:"created"`
	Modified           time.Time           `json:"modified"`
	CreatedByRef       ObjectIdentifier    `json:"created_by_ref,omitempty"`
	Revoked            bool                `json:"revoked,omitempty"`
	Labels             []string            `json:"labels,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs  []ObjectIdentifier  `json:"object_marking_refs,omitempty"`
	GranularMarkings   []GranularMarking   `json:"granular_markings,omitempty"`
}

type TestCommon struct {
	Common
	Testfield string
}

func (tc *TestCommon) Initialize() {
	(*tc).Type = "testObject"
	(*tc).Id = NewObjectId((*tc).Type)
}

func (tc *TestCommon) Valid() (bool, error, error) {
	return true, nil, nil
}

func (t *TestCommon) Save() {
	tm := time.Now()
	if (*t).Created.IsZero() {
		(*t).Created = tm
		(*t).Modified = tm
	} else {
		(*t).Modified = tm
	}
}
