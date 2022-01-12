package handlers

type HttpHeaderEnvelop struct {
	MessageId  string `header:"Message-ID"`
	Subject    string `header:"Subject"`
	From       string `header:"From"`
	To         string `header:"To"`
	References string `header:"References"`
	ReplyTo    string `header:"Reply-To"`
}

type RequestsReceived struct {
	RequestId      string `json:"request_id,omitempty"`
	Timestamp      string `json:"ts,omitempty"`
	Request        string `json:"request,omitempty"`
	ClientId       string `json:"client_id,omitempty"`
	ClientIP       string `json:"client_ip,omitempty"`
	ControllerId   string `json:"controller_id,omitempty"`
	ControllerHost string `json:"controller_host,omitempty"`
	ConversationId string `json:"conversation_id,omitempty"`
}

type ValidatedQueryCredential struct {
	Context           []string `json:"@context,omitempty"`
	CredentialSubject struct {
		Id             string `json:"id,omitempty"`
		ValidatetQuery struct {
			Ontology string `json:"ontology,omitempty"`
			Profile  string `json:"profile,omitempty"`
			Sparql   string `json:"sparql,omitempty"`
		} `json:"validatedQuery,omitempty"`
	} `json:"credentialSubject,omitempty"`
	Id           string `json:"id,omitempty"`
	IssuanceDate string `json:"issuanceDate,omitempty"`
	Issuer       string `json:"issuer,omitempty"`
	Proof        struct {
		Created            string `json:"created,omitempty"`
		Jws                string `json:"jws,omitempty"`
		ProofPurpose       string `json:"proofPurpose,omitempty"`
		Type               string `json:"type,omitempty"`
		VerificationMethod string `json:"verificationMethod,omitempty"`
	} `json:"proof,omitempty"`
	Type []string `json:"type,omitempty"`
}

type ValidatedQueriesReceived struct {
	Timestamp string                   `json:"ts,omitempty"`
	ClientIP  string                   `json:"client_ip,omitempty"`
	Header    HttpHeaderEnvelop        `json:"header,omitempty"`
	Payload   ValidatedQueryCredential `json:"payload,omitempty"`
}
