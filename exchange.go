package entrust

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func (c *Client) exchange(path, method string, payload interface{}) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", APIServer, path), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.apiKey)

	// Debug log before setting authorization header
	log.WithFields(log.Fields{
		"url":     req.URL.String(),
		"method":  method,
		"request": string(jsonPayload),
	}).Debug("Request")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Slow down when we hit the rate-limit
	if resp.StatusCode == 429 {
		var delay int
		delay, err = strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			delay = 5
		}

		log.WithFields(log.Fields{
			"rateLimit":  resp.Header.Get("X-Rate-Limit-Limit"),
			"retryAfter": resp.Header.Get("Retry-After"),
			"delay":      delay,
		}).Info("Request rate-limited, retrying according to intructions")

		time.Sleep(time.Duration(delay) * time.Second)
		return c.exchange(path, method, payload)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"url":      resp.Request.URL.String(),
		"status":   resp.StatusCode,
		"response": string(body),
	}).Debug("Response")

	return body, nil
}
