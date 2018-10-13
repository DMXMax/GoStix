package stix

import (
  "fmt"
)



//The IsValidInterface supplies a bool whether the object is valid,
//an array of errors, and an array of warnings
type StixObject interface {
}

type HashesType = map[string]string

func CreateNewStixObject(objectType DomainObjectType) (obj StixObject,err error){
  switch objectType{
  case DomainObjects.INDICATOR:
    var i Indicator
    i.Initialize()
    obj = StixObject(i)
  case DomainObjects.TEST:
    var t TestCommon
    t.Initialize()
    obj = StixObject(t)
  default:
    err = fmt.Errorf("CreateNewStixObject: unknown objectType %s", objectType)
  }

  return obj, err
}
