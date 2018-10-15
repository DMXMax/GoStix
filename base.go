package stix

import (
  "fmt"
)



//The IsValidInterface supplies a bool whether the object is valid,
//an array of errors, and an array of warnings
type StixObject interface {
  //The Valid call returns your basic bool for whether it is valid to exist as
  //a "published" object. If the result is false, the first error block is the
  //actual errors. The second group of "errors" are violations of the "SHOULD"
  //clause and are treated as warnings
  Valid() (bool, error, error)
  Save() 
}

type HashesType = map[string]string

func CreateNewStixObject(objectType DomainObjectType) (obj StixObject,err error){
  switch objectType{
  case DomainObjects.INDICATOR:
    var i Indicator
    i.Initialize()
    obj = &i 
  case DomainObjects.TEST:
    var t TestCommon
    t.Initialize()
    obj = &t
  default:
    err = fmt.Errorf("CreateNewStixObject: unknown objectType %s", objectType)
  }

  return obj, err
}
