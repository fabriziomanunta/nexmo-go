package nexmo

import "github.com/fabriziomanunta/nexmo-go/sling"

type ReportService struct {
	sling   *sling.Sling
	authSet *AuthSet
}

func newReportService(base *sling.Sling, authSet *AuthSet) *ReportService {
	sling := base.Base("https://api.nexmo.com/v2/reports/")
	return &ReportService{
		sling:   sling,
		authSet: authSet,
	}
}

func (c *ReportService) SetBaseURL(baseURL string) {
	c.sling.Base(baseURL)
}
