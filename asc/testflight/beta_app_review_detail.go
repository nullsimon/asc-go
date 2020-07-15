package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/apps"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// BetaAppReviewDetail defines model for BetaAppReviewDetail.
type BetaAppReviewDetail struct {
	Attributes *struct {
		ContactEmail        *string `json:"contactEmail,omitempty"`
		ContactFirstName    *string `json:"contactFirstName,omitempty"`
		ContactLastName     *string `json:"contactLastName,omitempty"`
		ContactPhone        *string `json:"contactPhone,omitempty"`
		DemoAccountName     *string `json:"demoAccountName,omitempty"`
		DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
		DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
		Notes               *string `json:"notes,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppReviewDetailUpdateRequest defines model for BetaAppReviewDetailUpdateRequest.
type BetaAppReviewDetailUpdateRequest struct {
	Data struct {
		Attributes *struct {
			ContactEmail        *string `json:"contactEmail,omitempty"`
			ContactFirstName    *string `json:"contactFirstName,omitempty"`
			ContactLastName     *string `json:"contactLastName,omitempty"`
			ContactPhone        *string `json:"contactPhone,omitempty"`
			DemoAccountName     *string `json:"demoAccountName,omitempty"`
			DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
			DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
			Notes               *string `json:"notes,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BetaAppReviewDetailResponse defines model for BetaAppReviewDetailResponse.
type BetaAppReviewDetailResponse struct {
	Data     BetaAppReviewDetail  `json:"data"`
	Included *[]apps.App          `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BetaAppReviewDetailsResponse defines model for BetaAppReviewDetailsResponse.
type BetaAppReviewDetailsResponse struct {
	Data     []BetaAppReviewDetail     `json:"data"`
	Included *[]apps.App               `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBetaAppReviewDetailsQuery struct {
	Fields *struct {
		Apps                 *[]string `url:"apps,omitempty"`
		BetaAppReviewDetails *[]string `url:"betaAppReviewDetails,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		App *[]string `url:"app,omitempty"`
	} `url:"filter,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Cursor  *string   `url:"cursor,omitempty"`
}

type GetBetaAppReviewDetailQuery struct {
	Fields *struct {
		Apps                 *[]string `url:"apps,omitempty"`
		BetaAppReviewDetails *[]string `url:"betaAppReviewDetails,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetAppForBetaAppReviewDetailQuery struct {
	Fields *struct {
		Apps *[]string `url:"apps,omitempty"`
	} `url:"fields,omitempty"`
}

type GetBetaAppReviewDetailsForAppQuery struct {
	Fields *struct {
		BetaAppReviewDetails *[]string `url:"betaAppReviewDetails,omitempty"`
	} `url:"fields,omitempty"`
}

// ListBetaAppReviewDetails finds and lists beta app review details for all apps.
func (s *Service) ListBetaAppReviewDetails(params *ListBetaAppReviewDetailsQuery) (*BetaAppReviewDetailsResponse, *common.Response, error) {
	res := new(BetaAppReviewDetailsResponse)
	resp, err := s.GetWithQuery("betaAppReviewDetails", params, res)
	return res, resp, err
}

// GetBetaAppReviewDetail gets beta app review details for a specific app.
func (s *Service) GetBetaAppReviewDetail(id string, params *GetBetaAppReviewDetailQuery) (*BetaAppReviewDetailResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBetaAppReviewDetail gets the app information for a specific beta app review details resource.
func (s *Service) GetAppForBetaAppReviewDetail(id string, params *GetAppForBetaAppReviewDetailQuery) (*apps.AppResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewDetailsForApp gets the beta app review details for a specific app.
func (s *Service) GetBetaAppReviewDetailsForApp(id string, params *GetBetaAppReviewDetailsForAppQuery) (*BetaAppReviewDetailResponse, *common.Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppReviewDetail", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBetaAppReviewDetail updates the details for a specific app's beta app review.
func (s *Service) UpdateBetaAppReviewDetail(id string, body *BetaAppReviewDetailUpdateRequest) (*BetaAppReviewDetailResponse, *common.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}