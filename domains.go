package entrust

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// GetDomains lists domains based on a filter
func (c *Client) GetDomains(offset, limit int, expiry *time.Time) (*GetDomainsResponse, error) {
	v := url.Values{}
	v.Set("sort", "domainName:asc")
	v.Set("offset", fmt.Sprintf("%d", offset))
	v.Set("limit", fmt.Sprintf("%d", limit))
	if expiry != nil {
		v.Set("ovExpiry", fmt.Sprintf("lt:%s", expiry.Format(time.RFC3339)))
	}

	resp, err := c.exchange(fmt.Sprintf("/domains?%s", v.Encode()), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var domainsResponse *GetDomainsResponse
	err = json.Unmarshal(resp, &domainsResponse)
	if err != nil {
		return nil, err
	}
	return domainsResponse, nil
}

// GetDomain requests information about a domain in the account
// GET /clients/{clientId}/domains/{domain}
func (c *Client) GetDomain(clientID int, domain string) (*Domain, error) {
	resp, err := c.exchange(fmt.Sprintf("/clients/%d/domains/%s ", clientID, domain), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	var domainResponse *Domain
	err = json.Unmarshal(resp, &domainResponse)
	if err != nil {
		return nil, err
	}
	return domainResponse, nil
}

// ReverifyDomain sends a domain for reverification.
// PUT /clients/{clientId}/domains/{domain}
func (c *Client) ReverifyDomain(clientID int, domain string) error {
	_, err := c.exchange(fmt.Sprintf("/clients/%d/domains/%s ", clientID, domain), http.MethodPut, ReverifyDomainRequest{
		VerificationMethod: "DNS",
	})
	if err != nil {
		return err
	}
	return nil
}
