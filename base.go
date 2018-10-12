package stix

import (
  "errors"
  "github.com/google/uuid"
  "strings"
  "fmt"
)


type StixObject interface {
  IsValid() (bool, error)
}

type HashesType = map[string]string


type Base  struct{
  Type string `json:"type"`
  Id string `json:"id"`
  Created string `json:"created"`
  Modified string `json:"modified"`
}

type ExternalReference struct {
  SourceName string `json:"source_name"`
  Description string `json:"description,omitempty"`
  Url string `json:"url,omitempty"`
  Hashes HashesType `json:"hashe,omitempty`
  ExternalId string `json:"external_id,omitempty"`
}
func (er *ExternalReference) AddHash(hashType string, hashValue string){
  if er.Hashes == nil{
    er.Hashes = make(map[string]string)
  }
  er.Hashes[hashType] = hashValue
}
func (er *ExternalReference) IsValid() (bool, error){
  result := false
  var err error  = nil

  if er.SourceName != ""{
    result = true
  }
  if er.Url != ""{
    result = len(er.Hashes) > 0
    if !result{
      err = errors.New("No hashes for URL")
    }
  }
  if( len(er.Url) + len(er.ExternalId) ==0 ){
    result = false
    err = errors.New("Either the url or the external_id property must be filled in.")
  }

  return result,err
}

type ObjectIdentifier string

func NewObjectId(objectType string) ObjectIdentifier{
  var buf strings.Builder
  fmt.Fprintf(&buf, "%s--%s", strings.ToLower(objectType), uuid.New().String())
  return ObjectIdentifier(buf.String())
}

func (o *ObjectIdentifier) Valid() (err error){
  err = nil
  strs := strings.Split(string(*o),"--")
  if len(strs)<2{ // we need an object type and a uuid
    err = errors.New("The id isn't right")
  }else{
    _,err = uuid.Parse(strs[1])//the second part should be the uuid
  }
  return
}
