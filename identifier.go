package stix

import (
  "errors"
  "github.com/google/uuid"
  "strings"
  "fmt"
)

type ObjectIdentifier string

func NewObjectId(objectType string) ObjectIdentifier{
  var buf strings.Builder
  fmt.Fprintf(&buf, "%s--%s", strings.ToLower(objectType), uuid.New().String())
  return ObjectIdentifier(buf.String())
}

func (o *ObjectIdentifier) IsValid() (err error){
  err = nil
  strs := strings.Split(string(*o),"--")
  if len(strs)<2{ // we need an object type and a uuid
    err = errors.New("The id isn't right")
  }else{
    _,err = uuid.Parse(strs[1])//the second part should be the uuid
  }
  return
}
