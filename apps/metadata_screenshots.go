package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// AppScreenshot defines model for AppScreenshot.
type AppScreenshot struct {
	Attributes *struct {
		AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
		AssetToken         *string             `json:"assetToken,omitempty"`
		AssetType          *string             `json:"assetType,omitempty"`
		FileName           *string             `json:"fileName,omitempty"`
		FileSize           *int                `json:"fileSize,omitempty"`
		ImageAsset         *ImageAsset         `json:"imageAsset,omitempty"`
		SourceFileChecksum *string             `json:"sourceFileChecksum,omitempty"`
		UploadOperations   *[]UploadOperation  `json:"uploadOperations,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		AppScreenshotSet *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"appScreenshotSet,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppScreenshotCreateRequest defines model for AppScreenshotCreateRequest.
type AppScreenshotCreateRequest struct {
	Attributes struct {
		FileName string `json:"fileName"`
		FileSize int    `json:"fileSize"`
	} `json:"attributes"`
	Relationships struct {
		AppScreenshotSet struct {
			Data types.RelationshipsData `json:"data"`
		} `json:"appScreenshotSet"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// AppScreenshotUpdateRequest defines model for AppScreenshotUpdateRequest.
type AppScreenshotUpdateRequest struct {
	Attributes *struct {
		SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
		Uploaded           *bool   `json:"uploaded,omitempty"`
	} `json:"attributes,omitempty"`
	ID   string `json:"id"`
	Type string `json:"type"`
}

// AppScreenshotResponse defines model for AppScreenshotResponse.
type AppScreenshotResponse struct {
	Data  AppScreenshot       `json:"data"`
	Links types.DocumentLinks `json:"links"`
}

// AppScreenshotsResponse defines model for AppScreenshotsResponse.
type AppScreenshotsResponse struct {
	Data  []AppScreenshot          `json:"data"`
	Links types.PagedDocumentLinks `json:"links"`
	Meta  *types.PagingInformation `json:"meta,omitempty"`
}

// GetAppScreenshotQuery are query options for GetAppScreenshot
type GetAppScreenshotQuery struct {
	FieldsAppScreenshots *[]string `url:"fields[appScreenshots],omitempty"`
	Include              *[]string `url:"include,omitempty"`
}

// GetAppScreenshot gets information about an app screenshot and its upload and processing status.
func (s *Service) GetAppScreenshot(id string, params *GetAppScreenshotQuery) (*AppScreenshotResponse, *services.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// CreateAppScreenshot adds a new screenshot to a screenshot set.
func (s *Service) CreateAppScreenshot(body *AppScreenshotCreateRequest) (*AppScreenshotResponse, *services.Response, error) {
	res := new(AppScreenshotResponse)
	resp, err := s.Post("appScreenshots", body, res)
	return res, resp, err
}

// UpdateAppScreenshot commits an app screenshot after uploading it.
func (s *Service) UpdateAppScreenshot(id string, body *AppScreenshotUpdateRequest) (*AppScreenshotResponse, *services.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	res := new(AppScreenshotResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}

// DeleteAppScreenshot deletes an app screenshot that is associated with a screenshot set.
func (s *Service) DeleteAppScreenshot(id string) (*services.Response, error) {
	url := fmt.Sprintf("appScreenshots/%s", id)
	return s.Delete(url, nil)
}
