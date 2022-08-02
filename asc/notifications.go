/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
)

// NotificationsService handles communication with user and role-related methods of the App Store Connect API
// https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification
type NotificationsService service

// Notification https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification
type Notification struct {
}

type TestNotificationResponse struct {
	TestNotificationToken string `json:"testNotificationToken"`
}

type RequestANotificationQuery struct {
}

// NotificationType https://developer.apple.com/documentation/appstoreservernotifications/notificationtype
type NotificationType string

const (
	TEST                NotificationType = "TEST"
	SUBSCRIBED          NotificationType = "SUBSCRIBED"
	CONSUMPTION_REQUEST NotificationType = "CONSUMPTION_REQUEST"
)

type NotificationHistoryRequest struct {
	StartDate             string           `json:"startDate"`
	EndDate               string           `json:"endDate"`
	OriginalTransactionId string           `json:"originalTransactionId"`
	NotificationType      NotificationType `json:"notificationType"`
	NotificationSubtype   string           `json:"notificationSubtype"`
}

type NotificationHistoryResponse struct {
	NotificationHistory []NotificationHistoryResponseItem `json:"data"`
	HasMore             bool                              `json:"hasMore"`
	PaginationToken     string                            `json:"paginationToken"`
}

type NotificationHistoryResponseItem struct {
	FirstSendAttemptResult string `json:"firstSendAttemptResult"`
	SignedPayload          string `json:"signedPayload"`
}

// RequestATestNotification https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification?changes=latest_minor
func (s *NotificationsService) RequestATestNotification(ctx context.Context, params *RequestANotificationQuery) (*TestNotificationResponse, *Response, error) {
	HostProduction := "https://api.storekit.itunes.apple.com"
	url := HostProduction + "/inApps/v1/notifications/test"
	res := new(TestNotificationResponse)
	resp, err := s.client.post(ctx, url, nil, res)

	return res, resp, err
}

// GetTestNotificationStatus
// GET https://api.storekit.itunes.apple.com/inApps/v1/notifications/test/{testNotificationToken}
func (s *NotificationsService) GetTestNotificationStatus(ctx context.Context, testNotificationToken string) (*Response, error) {
	HostProduction := "https://api.storekit.itunes.apple.com"
	url := HostProduction + "/inApps/v1/notifications/test/" + testNotificationToken
	res := new(Response)
	resp, err := s.client.get(ctx, url, nil, res)

	return resp, err
}

// NotificationHistory https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification?changes=latest_minor
func (s *NotificationsService) NotificationHistory(ctx context.Context, params *NotificationHistoryRequest) (*NotificationHistoryResponse, *Response, error) {
	HostProduction := "https://api.storekit.itunes.apple.com"
	url := HostProduction + "/inApps/v1/notifications/history"
	res := new(NotificationHistoryResponse)
	resp, err := s.client.post(ctx, url, nil, res)

	return res, resp, err
}
