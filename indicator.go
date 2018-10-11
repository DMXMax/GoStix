package stix 

type KillChainPhase struct{
  Name string `json:"kill_chain_name"`
  Phase string `json:"phase_name"`
}

type Indicator struct{
  Type string `json:"type"`
  Id string `json:"id"`
  CreatedByRef string `json:"created_by_ref,omitempty"`
  Created string `json:"created"`
  Modified string `json:"modified"`
  Labels []string `json:"labels"`
  Name string `json:"name,omitempty"`
  Pattern string `json:"pattern"`
  ValidFrom string `json:"valid_from"`
  ValidUntil string `json:"valid_until,omitempty"`
  Description string `json:"description,omitempty"` 
  KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
}
func Test(){
}
