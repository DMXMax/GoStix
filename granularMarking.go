package stix

type GranularMarking struct {
  MarkingRef ObjectIdentifier `json:"marking_ref,omitempty"`
  Selectors []string `json:"selectors,omitempty"`
}
