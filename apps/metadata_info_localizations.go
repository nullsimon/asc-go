package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal"
)

// AppInfoLocalization defines model for AppInfoLocalization.
type AppInfoLocalization struct {
	Attributes *struct {
		Locale            *string `json:"locale,omitempty"`
		Name              *string `json:"name,omitempty"`
		PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
		PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
		Subtitle          *string `json:"subtitle,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string                 `json:"id"`
	Links         internal.ResourceLinks `json:"links"`
	Relationships *struct {
		AppInfo *struct {
			Data  *internal.RelationshipsData  `json:"data,omitempty"`
			Links *internal.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appInfo,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppInfoLocalizationCreateRequest defines model for AppInfoLocalizationCreateRequest.
type AppInfoLocalizationCreateRequest struct {
	Data struct {
		Attributes struct {
			Locale            string  `json:"locale"`
			Name              *string `json:"name,omitempty"`
			PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
			PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
			Subtitle          *string `json:"subtitle,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			AppInfo struct {
				Data internal.RelationshipsData `json:"data"`
			} `json:"appInfo"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppInfoLocalizationResponse defines model for AppInfoLocalizationResponse.
type AppInfoLocalizationResponse struct {
	Data  AppInfoLocalization    `json:"data"`
	Links internal.DocumentLinks `json:"links"`
}

// AppInfoLocalizationUpdateRequest defines model for AppInfoLocalizationUpdateRequest.
type AppInfoLocalizationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			Name              *string `json:"name,omitempty"`
			PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
			PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
			Subtitle          *string `json:"subtitle,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AppInfoLocalizationsResponse defines model for AppInfoLocalizationsResponse.
type AppInfoLocalizationsResponse struct {
	Data  []AppInfoLocalization       `json:"data"`
	Links internal.PagedDocumentLinks `json:"links"`
	Meta  *internal.PagingInformation `json:"meta,omitempty"`
}

// ListAppInfoLocalizationsForAppInfoQuery are query options for ListAppInfoLocalizationsForAppInfo
type ListAppInfoLocalizationsForAppInfoQuery struct {
	FieldsAppInfos             *[]string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations *[]string `url:"fields[appInfoLocalizations],omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	FilterLocale               *[]string `url:"filter[locale],omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// GetAppInfoLocalizationQuery are query options for GetAppInfoLocalization
type GetAppInfoLocalizationQuery struct {
	FieldsAppInfoLocalizations *[]string `url:"fields[appInfoLocalizations],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
}

// ListAppInfoLocalizationsForAppInfo gets a list of localized, app-level information for an app.
func (s *Service) ListAppInfoLocalizationsForAppInfo(id string, params *ListAppInfoLocalizationsForAppInfoQuery) (*AppInfoLocalizationsResponse, *internal.Response, error) {
	url := fmt.Sprintf("appInfos/%s/appInfoLocalizations", id)
	res := new(AppInfoLocalizationsResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppInfoLocalization reads localized app-level information.
func (s *Service) GetAppInfoLocalization(id string, params *GetAppInfoLocalizationQuery) (*AppInfoLocalizationResponse, *internal.Response, error) {
	url := fmt.Sprintf("appInfoLocalizations/%s", id)
	res := new(AppInfoLocalizationResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppInfoLocalization adds app-level localized information for a new locale.
func (s *Service) CreateAppInfoLocalization(body *AppInfoLocalizationCreateRequest) (*AppInfoLocalizationResponse, *internal.Response, error) {
	res := new(AppInfoLocalizationResponse)
	resp, err := s.Post("appInfoLocalizations", body, res)
	return res, resp, err
}

// UpdateAppInfoLocalization modifies localized app-level information for a particular language.
func (s *Service) UpdateAppInfoLocalization(id string, body *AppInfoLocalizationUpdateRequest) (*AppInfoLocalizationResponse, *internal.Response, error) {
	url := fmt.Sprintf("appInfoLocalizations/%s", id)
	res := new(AppInfoLocalizationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppInfoLocalization deletes an app information localization that is associated with an app.
func (s *Service) DeleteAppInfoLocalization(id string) (*internal.Response, error) {
	url := fmt.Sprintf("appInfoLocalizations/%s", id)
	return s.Delete(url, nil)
}