package asc // nolint: dupl

import (
	"context"
	"testing"
)

func TestGetAttachment(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.GetAttachment(ctx, "10", &GetAttachmentQuery{})
	})
}

func TestListAttachmentsForReviewDetail(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.ListAttachmentsForReviewDetail(ctx, "10", &ListAttachmentQuery{})
	})
}

func TestCreateAttachment(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CreateAttachment(ctx, "", 0, "")
	})
}

func TestCommitAttachment(t *testing.T) {
	testEndpointWithResponse(t, "{}", &AppStoreReviewAttachmentResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Submission.CommitAttachment(ctx, "10", Bool(true), String("10"))
	})
}

func TestDeleteAttachment(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Submission.DeleteAttachment(ctx, "10")
	})
}
