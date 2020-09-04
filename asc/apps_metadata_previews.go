package asc

import (
	"context"
	"fmt"
)

// PreviewType defines model for PreviewType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/previewtype
type PreviewType string

// List of PreviewType
const (
	PreviewTypeAppleTV        PreviewType = "APPLE_TV"
	PreviewTypeDesktop        PreviewType = "DESKTOP"
	PreviewTypeiPad105        PreviewType = "IPAD_105"
	PreviewTypeiPad97         PreviewType = "IPAD_97"
	PreviewTypeiPadPro129     PreviewType = "IPAD_PRO_129"
	PreviewTypeiPadPro3Gen11  PreviewType = "IPAD_PRO_3GEN_11"
	PreviewTypeiPadPro3Gen129 PreviewType = "IPAD_PRO_3GEN_129"
	PreviewTypeiPhone35       PreviewType = "IPHONE_35"
	PreviewTypeiPhone40       PreviewType = "IPHONE_40"
	PreviewTypeiPhone47       PreviewType = "IPHONE_47"
	PreviewTypeiPhone55       PreviewType = "IPHONE_55"
	PreviewTypeiPhone58       PreviewType = "IPHONE_58"
	PreviewTypeiPhone65       PreviewType = "IPHONE_65"
	PreviewTypeWatchSeries3   PreviewType = "WATCH_SERIES_3"
	PreviewTypeWatchSeries4   PreviewType = "WATCH_SERIES_4"
)

// AppPreview defines model for AppPreview.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreview
type AppPreview struct {
	Attributes    *AppPreviewAttributes    `json:"attributes,omitempty"`
	ID            string                   `json:"id"`
	Links         ResourceLinks            `json:"links"`
	Relationships *AppPreviewRelationships `json:"relationships,omitempty"`
	Type          string                   `json:"type"`
}

// AppPreviewAttributes defines model for AppPreview.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreview/attributes
type AppPreviewAttributes struct {
	AssetDeliveryState   *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	FileName             *string             `json:"fileName,omitempty"`
	FileSize             *int64              `json:"fileSize,omitempty"`
	MimeType             *string             `json:"mimeType,omitempty"`
	PreviewFrameTimeCode *string             `json:"previewFrameTimeCode,omitempty"`
	PreviewImage         *ImageAsset         `json:"previewImage,omitempty"`
	SourceFileChecksum   *string             `json:"sourceFileChecksum,omitempty"`
	UploadOperations     []UploadOperation   `json:"uploadOperations,omitempty"`
	VideoURL             *string             `json:"videoUrl,omitempty"`
}

// AppPreviewRelationships defines model for AppPreview.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreview/relationships
type AppPreviewRelationships struct {
	AppPreviewSet *Relationship `json:"appPreviewSet,omitempty"`
}

// AppPreviewCreateRequest defines model for AppPreviewCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewcreaterequest/data
type appPreviewCreateRequest struct {
	Attributes    appPreviewCreateRequestAttributes    `json:"attributes"`
	Relationships appPreviewCreateRequestRelationships `json:"relationships"`
	Type          string                               `json:"type"`
}

// AppPreviewCreateRequestAttributes are attributes for AppPreviewCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewcreaterequest/data/attributes
type appPreviewCreateRequestAttributes struct {
	FileName             string  `json:"fileName"`
	FileSize             int64   `json:"fileSize"`
	MimeType             *string `json:"mimeType,omitempty"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
}

// AppPreviewCreateRequestRelationships are relationships for AppPreviewCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewcreaterequest/data/relationships
type appPreviewCreateRequestRelationships struct {
	AppPreviewSet relationshipDeclaration `json:"appPreviewSet"`
}

// AppPreviewUpdateRequest defines model for AppPreviewUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewupdaterequest/data
type appPreviewUpdateRequest struct {
	Attributes *appPreviewUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                             `json:"id"`
	Type       string                             `json:"type"`
}

// AppPreviewUpdateRequestAttributes are attributes for AppPreviewUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewupdaterequest/data/attributes
type appPreviewUpdateRequestAttributes struct {
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	SourceFileChecksum   *string `json:"sourceFileChecksum,omitempty"`
	Uploaded             *bool   `json:"uploaded,omitempty"`
}

// AppPreviewResponse defines model for AppPreviewResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewresponse
type AppPreviewResponse struct {
	Data  AppPreview    `json:"data"`
	Links DocumentLinks `json:"links"`
}

// AppPreviewsResponse defines model for AppPreviewsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/apppreviewsresponse
type AppPreviewsResponse struct {
	Data  []AppPreview       `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ImageAsset defines model for ImageAsset.
//
// https://developer.apple.com/documentation/appstoreconnectapi/imageasset
type ImageAsset struct {
	Height      *int    `json:"height,omitempty"`
	TemplateURL *string `json:"templateUrl,omitempty"`
	Width       *int    `json:"width,omitempty"`
}

// GetAppPreviewQuery are query options for GetAppPreview
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_preview_information
type GetAppPreviewQuery struct {
	FieldsAppPreviews []string `url:"fields[appPreviews],omitempty"`
	Include           []string `url:"include,omitempty"`
}

// GetAppPreview gets information about an app preview and its upload and processing status.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_preview_information
func (s *AppsService) GetAppPreview(ctx context.Context, id string, params *GetAppPreviewQuery) (*AppPreviewResponse, *Response, error) {
	url := fmt.Sprintf("appPreviews/%s", id)
	res := new(AppPreviewResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreateAppPreview adds a new preview to a preview set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_preview
func (s *AppsService) CreateAppPreview(ctx context.Context, fileName string, fileSize int64, appPreviewSetID string) (*AppPreviewResponse, *Response, error) {
	req := appPreviewCreateRequest{
		Attributes: appPreviewCreateRequestAttributes{
			FileName: fileName,
			FileSize: fileSize,
		},
		Relationships: appPreviewCreateRequestRelationships{
			AppPreviewSet: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appPreviewSetID,
					Type: "appPreviewSets",
				},
			},
		},
		Type: "appPreviews",
	}
	res := new(AppPreviewResponse)
	resp, err := s.client.post(ctx, "appPreviews", newRequestBody(req), res)
	return res, resp, err
}

// CommitAppPreview commits an app preview after uploading it.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_preview
func (s *AppsService) CommitAppPreview(ctx context.Context, id string, uploaded *bool, sourceFileChecksum *string, previewFrameTimeCode *string) (*AppPreviewResponse, *Response, error) {
	req := appPreviewUpdateRequest{
		ID:   id,
		Type: "appPreviews",
	}
	if uploaded != nil || sourceFileChecksum != nil || previewFrameTimeCode != nil {
		req.Attributes = &appPreviewUpdateRequestAttributes{
			Uploaded:             uploaded,
			SourceFileChecksum:   sourceFileChecksum,
			PreviewFrameTimeCode: previewFrameTimeCode,
		}
	}
	url := fmt.Sprintf("appPreviews/%s", id)
	res := new(AppPreviewResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)
	return res, resp, err
}

// DeleteAppPreview deletes an app preview that is associated with a preview set.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_preview
func (s *AppsService) DeleteAppPreview(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appPreviews/%s", id)
	return s.client.delete(ctx, url, nil)
}
