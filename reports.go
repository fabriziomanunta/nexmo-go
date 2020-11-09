package nexmo

import (
	"fmt"
	"net/http"

	"github.com/fabriziomanunta/nexmo-go/sling"
)

func (c *ReportErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", c.ErrorTitle, c.Type)
}

// Load Records gets the list of records with the given search criterias
func (c *ReportService) LoadRecords(request LoadRecordsRequest) (*LoadRecordsResponse, *http.Response, error) {
	sling := c.sling.New().Get("records").QueryStruct(request)

	reportResponse := new(LoadRecordsResponse)
	httpResponse, err := c.makeRequest(sling, reportResponse)
	return reportResponse, httpResponse, err
}

func (c *ReportService) makeRequest(s *sling.Sling, successV interface{}) (*http.Response, error) {
	errorV := new(ReportErrorResponse)
	if err := c.authSet.ApplyJWT(s); err != nil {
		return nil, fmt.Errorf("%s - error applying JWT", err)
	}
	httpResponse, err := s.Receive(successV, errorV)
	if err != nil {
		return httpResponse, fmt.Errorf("%s - error receiving data from server", err)
	}
	if errorV.Type != "" {
		return httpResponse, errorV
	}
	return httpResponse, nil
}
