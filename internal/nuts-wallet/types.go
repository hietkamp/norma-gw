package wallet

type VerifiableCredential struct {
	Context           []string               `json:"@context,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty"`
	Id                string                 `json:"id,omitempty"`
	IssuanceDate      string                 `json:"issuanceDate,omitempty"`
	Issuer            string                 `json:"issuer,omitempty"`
	Type              []string               `json:"type,omitempty"`
}
