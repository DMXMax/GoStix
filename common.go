package stix

import (
 // "strings"
)


type Common struct{
  Type string `json:"type"`
  Id ObjectIdentifier `json:"id"`
  Created string `json:"created"`
  Modified string `json:"modified"`
  CreatedByRef ObjectIdentifier `json:"created_by_ref,omitempty"`
  Revoked bool `json:"revoked,omitempty"`
  Labels []string `json:"labels,omitempty"`
  ExternalReferences []ExternalReference `json:"external_references,omitempty"`
  ObjectMarkingRefs []ObjectIdentifier `json:"object_marking_refs,omitempty"`
  GranularMarkings []GranularMarking `json:"granular_markings,omitempty"`
}

type TestCommon struct{
  Common
  Testfield string
}

func (tc *TestCommon) Initialize(){
  (*tc).Type = "testObject"
  (*tc).Id = NewObjectId((*tc).Type)
}
