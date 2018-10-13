package stix 

type KillChainPhase struct{
  Name string `json:"kill_chain_name"`
  Phase string `json:"phase_name"`
}

type Indicator struct{
  Common
  Name string `json:"name,omitempty"`
  Pattern string `json:"pattern"`
  ValidFrom string `json:"valid_from"`
  ValidUntil string `json:"valid_until,omitempty"`
  Description string `json:"description,omitempty"` 
  KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
}
func Test(){
}

func (item *Indicator) Initialize(){
  (*item).Type = INDICATOR 
  (*item).Id = NewObjectId((item).Type)
}
