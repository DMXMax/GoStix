package stix

import "errors"

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
