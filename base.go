package stix

import (
  "fmt"
)


//The IsValidInterface supplies a bool whether the object is valid,
//an array of errors, and an array of warnings
type StixObject interface {
  IsValid() (bool, []error, []error)
}

type HashesType = map[string]string

func CreateNewStixObject(objectType string) (obj interface{},err error){
  switch objectType{
  case INDICATOR:
    var i Indicator
    i.Initialize()
    obj = interface{}(i)

  default:
    err = fmt.Errorf("CreateNewStixObject: unknown objectType %s", objectType)
  }

  return obj, err
}
