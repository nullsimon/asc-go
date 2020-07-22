package submission

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
)

// IDFADeclaration defines model for IDFADeclaration.
type IDFADeclaration struct {
	Attributes *struct {
		AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
		AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
		HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
		ServesAds                             *bool `json:"servesAds,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		AppStoreVersion *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appStoreVersion,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// IDFADeclarationCreateRequest defines model for IDFADeclarationCreateRequest.
type IDFADeclarationCreateRequest struct {
	Data struct {
		Attributes struct {
			AttributesActionWithPreviousAd        bool `json:"attributesActionWithPreviousAd"`
			AttributesAppInstallationToPreviousAd bool `json:"attributesAppInstallationToPreviousAd"`
			HonorsLimitedAdTracking               bool `json:"honorsLimitedAdTracking"`
			ServesAds                             bool `json:"servesAds"`
		} `json:"attributes"`
		Relationships struct {
			AppStoreVersion struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"appStoreVersion"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// IDFADeclarationUpdateRequest defines model for IDFADeclarationUpdateRequest.
type IDFADeclarationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AttributesActionWithPreviousAd        *bool `json:"attributesActionWithPreviousAd,omitempty"`
			AttributesAppInstallationToPreviousAd *bool `json:"attributesAppInstallationToPreviousAd,omitempty"`
			HonorsLimitedAdTracking               *bool `json:"honorsLimitedAdTracking,omitempty"`
			ServesAds                             *bool `json:"servesAds,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// IDFADeclarationResponse defines model for IDFADeclarationResponse.
type IDFADeclarationResponse struct {
	Data  IDFADeclaration        `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// GetIDFADeclarationForAppStoreVersionQuery are query options for GetIDFADeclarationForAppStoreVersion
type GetIDFADeclarationForAppStoreVersionQuery struct {
	FieldsIDFADeclarations *[]string `url:"fields[idfaDeclarations],omitempty"`
}

// CreateIDFADeclaration declares the IDFA usage for an App Store version.
func (s *Service) CreateIDFADeclaration(body *IDFADeclarationCreateRequest) (*IDFADeclarationResponse, *internal.Response, error) {
	res := new(IDFADeclarationResponse)
	resp, err := s.Post("idfaDeclarations", body, res)
	return res, resp, err
}

// UpdateIDFADeclaration updates your declared IDFA usage.
func (s *Service) UpdateIDFADeclaration(id string, body *IDFADeclarationUpdateRequest) (*IDFADeclarationResponse, *internal.Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteIDFADeclaration deletes the IDFA declaration that is associated with a version.
func (s *Service) DeleteIDFADeclaration(id string) (*internal.Response, error) {
	url := fmt.Sprintf("idfaDeclarations/%s", id)
	return s.Delete(url, nil)
}

// GetIDFADeclarationForAppStoreVersion reads your declared Advertising Identifier (IDFA) usage responses.
func (s *Service) GetIDFADeclarationForAppStoreVersion(id string, params *GetIDFADeclarationForAppStoreVersionQuery) (*IDFADeclarationResponse, *internal.Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/idfaDeclaration", id)
	res := new(IDFADeclarationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}