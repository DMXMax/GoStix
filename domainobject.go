package stix

type DomainObjectType string

var DomainObjects = struct{
  TEST DomainObjectType
  INDICATOR DomainObjectType
}{
  "testcommon",
  "indicator",
}
