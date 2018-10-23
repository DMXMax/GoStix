package vocabulary

var IndicatorLabel = struct {
	AnomalousActivity string
	Anonymization     string
	Benign            string
	Compromised       string
	MaliciousActivity string
	Attribution       string
}{
	"anomalous-activity",
	"anonymization",
	"benign",
	"compromised",
	"malicious-activity",
	"attribution",
}

var ToolLabel = struct{
/*denial-of-service, exploitation, information-gathering, network-capture,
* credential-exploitation, remote-access, vulnerability-scanning*/
  DenialOfService string
  Exploitation string
  InformationGathering string
  NetworkCapture string
  CredentialExploitation string
  RemoteAccess string
  VulnerabilityScanning string
}{
  "denial-of-service",
  "exploitation",
  "information-gathering",
  "network-capture",
  "credential-exploitation",
  "remote-access",
  "vulnerability-scanning",
}

var ThreatActorSophistication = struct{
  /*none, minimal, intermediate, advanced, expert, innovator, strategic*/
  None string
  Minimal string
  Intermediate string
  Advanced string
  Expert string
  Innovator string
  Strategic string
}{
  "none",
  "minimal",
  "intermediate",
  "advanced",
  "expert",
  "innovator",
  "strategic",
}

var ThreatActorRole = struct{
  /*agent, director, independent, infrastructure-architect,
  * infrastructure-operator, malware-author, sponsor*/
  Agent string
  Director string
  Independent string
  InfrastructureArchitect string
  InfrastructureOperator string
  MalwareAuthor string
  Sponsor string
}{
  "agent",
  "director",
  "independent",
  "infrastructure-architect",
  "infrastructure-operator",
  "malware-author",
  "sponsor",
}

var ThreatActorLabel = struct{
  /*activist, competitor, crime-syndicate, criminal, hacker,
  * insider-accidental, insider-disgruntled, nation-state, sensationalist, spy,
  * terrorist*/
  Activist string
  Competitor string
  CrimeSyndicate string
  Criminal string
  Hacker string
  InsiderAccidental string
  InsiderDisgruntled string
  NationState string
  Sensationalist string
}{
  "activist",
  "competitor",
  "crime-syndicate",
  "criminal",
  "hacker",
  "insider-accidental",
  "insider-disgruntled",
  "nation-state",
  "sensationalist",
  "spy",
}

var ReportLabel = struct{
/*threat-report, attack-pattern, campaign, identity, indicator, malware,
* observed-data, threat-actor, tool, vulnerability*/
  ThreatReport string
  AttackPattern string
  Campaign string
  Identity string
  Indicator string
  Malware string
  ObservedData string
  ThreatActor string
  Tool string
  Vulnerability string
}{
  "threat-report",
  "attack-pattern",
  "campaign",
  "identity",
  "indicator",
  "malware",
  "observed-data",
  "threat-actor",
  "tool",
  "vulnerability",

}

var MalwareLabel = struct{
/*adware, backdoor, bot, ddos, dropper, exploit-kit, keylogger, ransomware,
* remote-access-trojan, resource-exploitation, rogue-security-software,
* rootkit, screen-capture, spyware, trojan, virus, worm*/
  Adware string
  Backdoor string
  Bot string
  Ddos string
  Dropper string
  ExploitKit string
  Keylogger string
  Ransomware string
  RemoteAccessTrojan string
  ResourceExploitation string
  RogueSecuritySoftware string
  Rootkit string
  ScreenCapture string
  Spyware string
  Trojan string
  Virus string
  Worm string
}{
  "adware",
  "backdoor",
  "bot",
  "ddos",
  "dropper",
  "exploit-kit",
  "keylogger",
  "ransomware",
  "remote-access-trojan",
  "resource-exploitation",
  "rogue-securtiy-software",
  "rootkit",
  "screen-capture",
  "spyware",
  "trojan",
  "virus",
  "worm"
}
