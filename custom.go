package stix

import (
	"errors"
	"strings"
	"time"
)

type MarkingDefinition struct {
	Type               string              `json:"type"`
	Id                 ObjectIdentifier    `json:"id"`
	Created            time.Time           `json:"created"`
	CreatedByRef       ObjectIdentifier    `json:"created_by_ref,omitempty"`
	Revoked            bool                `json:"revoked,omitempty"`
	Labels             []string            `json:"labels,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs  []ObjectIdentifier  `json:"object_marking_refs,omitempty"`
	GranularMarkings   []GranularMarking   `json:"granular_markings,omitempty"`
	DefinitionType     string              `json:"definition-type"`
	Definition         Definition          `json:"definition"`
}

type Definition struct {
	Statement string `json:"statement"`
}

func (d *Definition) Valid() (bResult bool, err []error, warn []error) {
	stmt := strings.TrimSpace((*d).Statement)
	if len(stmt) == 0 {
		bResult = false
		err = append(err, errors.New("Statement cannot be empty"))

	}else{
    bResult = true
  }
	return
}
func (md *MarkingDefinition) Valid() (bResult bool, err []error, warn []error) {

	bResult = true

	bResult, err, warn = (&(*md).Definition).Valid()
	/*switch (*md).DefinitionType {
	case "tlp":
	case "statement":
	default:
		err = append(err, errors.New("definition-type should be 'tlp' or 'standard' got "+
			DefinitionType))*/
	return
}

func (md *MarkingDefinition) Initialize() {
	(*md).Type = "marking-definition"
	(*md).Id = NewObjectId((*md).Type)
}
