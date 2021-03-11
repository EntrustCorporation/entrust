package entrust

import "time"

// Response holds the default API response
type Response struct {
	Status int
	Errors []struct {
		Message string
	}
}

// AddDomainRequest contains the paramters to create a new domain
type AddDomainRequest struct {
	DomainName         string `json:"domainName"`
	VerificationMethod string `json:"verificationMethod"` // enum (DNS, EMAIL, MANUAL, WEB_SERVER)
}

// ReverifyDomainRequest schedule domain for revalidation
type ReverifyDomainRequest struct {
	VerificationMethod string `json:"verificationMethod"` // enum (DNS, EMAIL, MANUAL, WEB_SERVER)
}

// GetDomainsResponse holds a list of domains
type GetDomainsResponse struct {
	Response
	Summary *Summary `json:"summary"`
	Domains []Domain `json:"domains"`
}

// Domain model
type Domain struct {
	ClientID           int        `json:"clientId,omitempty"`           // Client id of the client to which the domain belongs to
	DomainName         string     `json:"domainName,omitempty"`         // Domain name
	EVEligible         bool       `json:"evEligible,omitempty"`         // Whether this domain can be used for EV certificates
	EVExpiry           time.Time  `json:"evExpiry,omitempty"`           // Expiry time of verified EV information
	OVEligible         bool       `json:"ovEligible,omitempty"`         // Whether this domain can be used for OV certificates
	OVExpiry           time.Time  `json:"ovExpiry,omitempty"`           // Expiry time of verified OV information
	VerificationMethod string     `json:"verificationMethod,omitempty"` // (DNS, EMAIL, MANUAL, WEB_SERVER)
	VerificationStatus string     `json:"verificationStatus,omitempty"`
	DNSMethod          *DNSMethod `json:"dnsMethod"`
}

// DNSMethod information
type DNSMethod struct {
	RecordDomain string `json:"recordDomain"`
	RecvordType  string `json:"recvordType"`
	RecordValue  string `json:"recordValue"`
}

// Summary of the response
type Summary struct {
	Timestamp time.Time `json:"timestamp"`
	Elapsed   int       `json:"elapsed"`
	Offset    int       `json:"offset"`
	Limit     int       `json:"limit"`
	Total     int       `json:"total"`
	Sort      string    `json:"sort"`
}
