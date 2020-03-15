// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// UserRequestBuilder is request builder for User
type UserRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserRequest
func (b *UserRequestBuilder) Request() *UserRequest {
	return &UserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserRequest is request for User
type UserRequest struct{ BaseRequest }

// Get performs GET request for User
func (r *UserRequest) Get(ctx context.Context) (resObj *User, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for User
func (r *UserRequest) Update(ctx context.Context, reqObj *User) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for User
func (r *UserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserAccountInformationRequestBuilder is request builder for UserAccountInformation
type UserAccountInformationRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAccountInformationRequest
func (b *UserAccountInformationRequestBuilder) Request() *UserAccountInformationRequest {
	return &UserAccountInformationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserAccountInformationRequest is request for UserAccountInformation
type UserAccountInformationRequest struct{ BaseRequest }

// Get performs GET request for UserAccountInformation
func (r *UserAccountInformationRequest) Get(ctx context.Context) (resObj *UserAccountInformation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserAccountInformation
func (r *UserAccountInformationRequest) Update(ctx context.Context, reqObj *UserAccountInformation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserAccountInformation
func (r *UserAccountInformationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserActivityRequestBuilder is request builder for UserActivity
type UserActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserActivityRequest
func (b *UserActivityRequestBuilder) Request() *UserActivityRequest {
	return &UserActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserActivityRequest is request for UserActivity
type UserActivityRequest struct{ BaseRequest }

// Get performs GET request for UserActivity
func (r *UserActivityRequest) Get(ctx context.Context) (resObj *UserActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserActivity
func (r *UserActivityRequest) Update(ctx context.Context, reqObj *UserActivity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserActivity
func (r *UserActivityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserAnalyticsRequestBuilder is request builder for UserAnalytics
type UserAnalyticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAnalyticsRequest
func (b *UserAnalyticsRequestBuilder) Request() *UserAnalyticsRequest {
	return &UserAnalyticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserAnalyticsRequest is request for UserAnalytics
type UserAnalyticsRequest struct{ BaseRequest }

// Get performs GET request for UserAnalytics
func (r *UserAnalyticsRequest) Get(ctx context.Context) (resObj *UserAnalytics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserAnalytics
func (r *UserAnalyticsRequest) Update(ctx context.Context, reqObj *UserAnalytics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserAnalytics
func (r *UserAnalyticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserAppInstallStatusRequestBuilder is request builder for UserAppInstallStatus
type UserAppInstallStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAppInstallStatusRequest
func (b *UserAppInstallStatusRequestBuilder) Request() *UserAppInstallStatusRequest {
	return &UserAppInstallStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserAppInstallStatusRequest is request for UserAppInstallStatus
type UserAppInstallStatusRequest struct{ BaseRequest }

// Get performs GET request for UserAppInstallStatus
func (r *UserAppInstallStatusRequest) Get(ctx context.Context) (resObj *UserAppInstallStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserAppInstallStatus
func (r *UserAppInstallStatusRequest) Update(ctx context.Context, reqObj *UserAppInstallStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserAppInstallStatus
func (r *UserAppInstallStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserConfigurationRequestBuilder is request builder for UserConfiguration
type UserConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserConfigurationRequest
func (b *UserConfigurationRequestBuilder) Request() *UserConfigurationRequest {
	return &UserConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserConfigurationRequest is request for UserConfiguration
type UserConfigurationRequest struct{ BaseRequest }

// Get performs GET request for UserConfiguration
func (r *UserConfigurationRequest) Get(ctx context.Context) (resObj *UserConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserConfiguration
func (r *UserConfigurationRequest) Update(ctx context.Context, reqObj *UserConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserConfiguration
func (r *UserConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserCredentialUsageDetailsRequestBuilder is request builder for UserCredentialUsageDetails
type UserCredentialUsageDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserCredentialUsageDetailsRequest
func (b *UserCredentialUsageDetailsRequestBuilder) Request() *UserCredentialUsageDetailsRequest {
	return &UserCredentialUsageDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserCredentialUsageDetailsRequest is request for UserCredentialUsageDetails
type UserCredentialUsageDetailsRequest struct{ BaseRequest }

// Get performs GET request for UserCredentialUsageDetails
func (r *UserCredentialUsageDetailsRequest) Get(ctx context.Context) (resObj *UserCredentialUsageDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserCredentialUsageDetails
func (r *UserCredentialUsageDetailsRequest) Update(ctx context.Context, reqObj *UserCredentialUsageDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserCredentialUsageDetails
func (r *UserCredentialUsageDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsBaselineRequestBuilder is request builder for UserExperienceAnalyticsBaseline
type UserExperienceAnalyticsBaselineRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsBaselineRequest
func (b *UserExperienceAnalyticsBaselineRequestBuilder) Request() *UserExperienceAnalyticsBaselineRequest {
	return &UserExperienceAnalyticsBaselineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsBaselineRequest is request for UserExperienceAnalyticsBaseline
type UserExperienceAnalyticsBaselineRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsBaseline, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsBaseline) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsCategoryRequestBuilder is request builder for UserExperienceAnalyticsCategory
type UserExperienceAnalyticsCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsCategoryRequest
func (b *UserExperienceAnalyticsCategoryRequestBuilder) Request() *UserExperienceAnalyticsCategoryRequest {
	return &UserExperienceAnalyticsCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsCategoryRequest is request for UserExperienceAnalyticsCategory
type UserExperienceAnalyticsCategoryRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsCategory
func (r *UserExperienceAnalyticsCategoryRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsCategory
func (r *UserExperienceAnalyticsCategoryRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsCategory
func (r *UserExperienceAnalyticsCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsDevicePerformanceRequestBuilder is request builder for UserExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsDevicePerformanceRequest
func (b *UserExperienceAnalyticsDevicePerformanceRequestBuilder) Request() *UserExperienceAnalyticsDevicePerformanceRequest {
	return &UserExperienceAnalyticsDevicePerformanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsDevicePerformanceRequest is request for UserExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsDevicePerformance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsDevicePerformance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder is request builder for UserExperienceAnalyticsDeviceStartupHistory
type UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsDeviceStartupHistoryRequest
func (b *UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) Request() *UserExperienceAnalyticsDeviceStartupHistoryRequest {
	return &UserExperienceAnalyticsDeviceStartupHistoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsDeviceStartupHistoryRequest is request for UserExperienceAnalyticsDeviceStartupHistory
type UserExperienceAnalyticsDeviceStartupHistoryRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsDeviceStartupHistory
func (r *UserExperienceAnalyticsDeviceStartupHistoryRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsDeviceStartupHistory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsDeviceStartupHistory
func (r *UserExperienceAnalyticsDeviceStartupHistoryRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsDeviceStartupHistory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsDeviceStartupHistory
func (r *UserExperienceAnalyticsDeviceStartupHistoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsMetricRequestBuilder is request builder for UserExperienceAnalyticsMetric
type UserExperienceAnalyticsMetricRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsMetricRequest
func (b *UserExperienceAnalyticsMetricRequestBuilder) Request() *UserExperienceAnalyticsMetricRequest {
	return &UserExperienceAnalyticsMetricRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsMetricRequest is request for UserExperienceAnalyticsMetric
type UserExperienceAnalyticsMetricRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsMetric
func (r *UserExperienceAnalyticsMetricRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsMetric, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsMetric
func (r *UserExperienceAnalyticsMetricRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsMetric) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsMetric
func (r *UserExperienceAnalyticsMetricRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsOverviewRequestBuilder is request builder for UserExperienceAnalyticsOverview
type UserExperienceAnalyticsOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsOverviewRequest
func (b *UserExperienceAnalyticsOverviewRequestBuilder) Request() *UserExperienceAnalyticsOverviewRequest {
	return &UserExperienceAnalyticsOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsOverviewRequest is request for UserExperienceAnalyticsOverview
type UserExperienceAnalyticsOverviewRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsOverview
func (r *UserExperienceAnalyticsOverviewRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsOverview
func (r *UserExperienceAnalyticsOverviewRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsOverview
func (r *UserExperienceAnalyticsOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsRegressionSummaryRequestBuilder is request builder for UserExperienceAnalyticsRegressionSummary
type UserExperienceAnalyticsRegressionSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsRegressionSummaryRequest
func (b *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Request() *UserExperienceAnalyticsRegressionSummaryRequest {
	return &UserExperienceAnalyticsRegressionSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsRegressionSummaryRequest is request for UserExperienceAnalyticsRegressionSummary
type UserExperienceAnalyticsRegressionSummaryRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsRegressionSummary
func (r *UserExperienceAnalyticsRegressionSummaryRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsRegressionSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsRegressionSummary
func (r *UserExperienceAnalyticsRegressionSummaryRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsRegressionSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsRegressionSummary
func (r *UserExperienceAnalyticsRegressionSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserInstallStateSummaryRequestBuilder is request builder for UserInstallStateSummary
type UserInstallStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserInstallStateSummaryRequest
func (b *UserInstallStateSummaryRequestBuilder) Request() *UserInstallStateSummaryRequest {
	return &UserInstallStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserInstallStateSummaryRequest is request for UserInstallStateSummary
type UserInstallStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Get(ctx context.Context) (resObj *UserInstallStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Update(ctx context.Context, reqObj *UserInstallStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserPFXCertificateRequestBuilder is request builder for UserPFXCertificate
type UserPFXCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserPFXCertificateRequest
func (b *UserPFXCertificateRequestBuilder) Request() *UserPFXCertificateRequest {
	return &UserPFXCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserPFXCertificateRequest is request for UserPFXCertificate
type UserPFXCertificateRequest struct{ BaseRequest }

// Get performs GET request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Get(ctx context.Context) (resObj *UserPFXCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Update(ctx context.Context, reqObj *UserPFXCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSecurityProfileRequestBuilder is request builder for UserSecurityProfile
type UserSecurityProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSecurityProfileRequest
func (b *UserSecurityProfileRequestBuilder) Request() *UserSecurityProfileRequest {
	return &UserSecurityProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSecurityProfileRequest is request for UserSecurityProfile
type UserSecurityProfileRequest struct{ BaseRequest }

// Get performs GET request for UserSecurityProfile
func (r *UserSecurityProfileRequest) Get(ctx context.Context) (resObj *UserSecurityProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSecurityProfile
func (r *UserSecurityProfileRequest) Update(ctx context.Context, reqObj *UserSecurityProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSecurityProfile
func (r *UserSecurityProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSettingsRequestBuilder is request builder for UserSettings
type UserSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSettingsRequest
func (b *UserSettingsRequestBuilder) Request() *UserSettingsRequest {
	return &UserSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSettingsRequest is request for UserSettings
type UserSettingsRequest struct{ BaseRequest }

// Get performs GET request for UserSettings
func (r *UserSettingsRequest) Get(ctx context.Context) (resObj *UserSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSettings
func (r *UserSettingsRequest) Update(ctx context.Context, reqObj *UserSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSettings
func (r *UserSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserTeamworkRequestBuilder is request builder for UserTeamwork
type UserTeamworkRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTeamworkRequest
func (b *UserTeamworkRequestBuilder) Request() *UserTeamworkRequest {
	return &UserTeamworkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTeamworkRequest is request for UserTeamwork
type UserTeamworkRequest struct{ BaseRequest }

// Get performs GET request for UserTeamwork
func (r *UserTeamworkRequest) Get(ctx context.Context) (resObj *UserTeamwork, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTeamwork
func (r *UserTeamworkRequest) Update(ctx context.Context, reqObj *UserTeamwork) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTeamwork
func (r *UserTeamworkRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type UserAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumented
func (b *UserRequestBuilder) AssignLicense(reqObj *UserAssignLicenseRequestParameter) *UserAssignLicenseRequestBuilder {
	bb := &UserAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserAssignLicenseRequest struct{ BaseRequest }

//
func (b *UserAssignLicenseRequestBuilder) Request() *UserAssignLicenseRequest {
	return &UserAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserAssignLicenseRequest) Post(ctx context.Context) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserChangePasswordRequestBuilder struct{ BaseRequestBuilder }

// ChangePassword action undocumented
func (b *UserRequestBuilder) ChangePassword(reqObj *UserChangePasswordRequestParameter) *UserChangePasswordRequestBuilder {
	bb := &UserChangePasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/changePassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserChangePasswordRequest struct{ BaseRequest }

//
func (b *UserChangePasswordRequestBuilder) Request() *UserChangePasswordRequest {
	return &UserChangePasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserChangePasswordRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserInvalidateAllRefreshTokensRequestBuilder struct{ BaseRequestBuilder }

// InvalidateAllRefreshTokens action undocumented
func (b *UserRequestBuilder) InvalidateAllRefreshTokens(reqObj *UserInvalidateAllRefreshTokensRequestParameter) *UserInvalidateAllRefreshTokensRequestBuilder {
	bb := &UserInvalidateAllRefreshTokensRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/invalidateAllRefreshTokens"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserInvalidateAllRefreshTokensRequest struct{ BaseRequest }

//
func (b *UserInvalidateAllRefreshTokensRequestBuilder) Request() *UserInvalidateAllRefreshTokensRequest {
	return &UserInvalidateAllRefreshTokensRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserInvalidateAllRefreshTokensRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserRevokeSignInSessionsRequestBuilder struct{ BaseRequestBuilder }

// RevokeSignInSessions action undocumented
func (b *UserRequestBuilder) RevokeSignInSessions(reqObj *UserRevokeSignInSessionsRequestParameter) *UserRevokeSignInSessionsRequestBuilder {
	bb := &UserRevokeSignInSessionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeSignInSessions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserRevokeSignInSessionsRequest struct{ BaseRequest }

//
func (b *UserRevokeSignInSessionsRequestBuilder) Request() *UserRevokeSignInSessionsRequest {
	return &UserRevokeSignInSessionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserRevokeSignInSessionsRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserReprocessLicenseAssignmentRequestBuilder struct{ BaseRequestBuilder }

// ReprocessLicenseAssignment action undocumented
func (b *UserRequestBuilder) ReprocessLicenseAssignment(reqObj *UserReprocessLicenseAssignmentRequestParameter) *UserReprocessLicenseAssignmentRequestBuilder {
	bb := &UserReprocessLicenseAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reprocessLicenseAssignment"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserReprocessLicenseAssignmentRequest struct{ BaseRequest }

//
func (b *UserReprocessLicenseAssignmentRequestBuilder) Request() *UserReprocessLicenseAssignmentRequest {
	return &UserReprocessLicenseAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserReprocessLicenseAssignmentRequest) Post(ctx context.Context) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserFindMeetingTimesRequestBuilder struct{ BaseRequestBuilder }

// FindMeetingTimes action undocumented
func (b *UserRequestBuilder) FindMeetingTimes(reqObj *UserFindMeetingTimesRequestParameter) *UserFindMeetingTimesRequestBuilder {
	bb := &UserFindMeetingTimesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/findMeetingTimes"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserFindMeetingTimesRequest struct{ BaseRequest }

//
func (b *UserFindMeetingTimesRequestBuilder) Request() *UserFindMeetingTimesRequest {
	return &UserFindMeetingTimesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserFindMeetingTimesRequest) Post(ctx context.Context) (resObj *MeetingTimeSuggestionsResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type UserSendMailRequestBuilder struct{ BaseRequestBuilder }

// SendMail action undocumented
func (b *UserRequestBuilder) SendMail(reqObj *UserSendMailRequestParameter) *UserSendMailRequestBuilder {
	bb := &UserSendMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserSendMailRequest struct{ BaseRequest }

//
func (b *UserSendMailRequestBuilder) Request() *UserSendMailRequest {
	return &UserSendMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserSendMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserGetMailTipsRequestBuilder struct{ BaseRequestBuilder }

// GetMailTips action undocumented
func (b *UserRequestBuilder) GetMailTips(reqObj *UserGetMailTipsRequestParameter) *UserGetMailTipsRequestBuilder {
	bb := &UserGetMailTipsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getMailTips"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserGetMailTipsRequest struct{ BaseRequest }

//
func (b *UserGetMailTipsRequestBuilder) Request() *UserGetMailTipsRequest {
	return &UserGetMailTipsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserGetMailTipsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MailTips, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MailTips
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []MailTips
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *UserGetMailTipsRequest) PostN(ctx context.Context, n int) ([]MailTips, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *UserGetMailTipsRequest) Post(ctx context.Context) ([]MailTips, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type UserTranslateExchangeIDsRequestBuilder struct{ BaseRequestBuilder }

// TranslateExchangeIDs action undocumented
func (b *UserRequestBuilder) TranslateExchangeIDs(reqObj *UserTranslateExchangeIDsRequestParameter) *UserTranslateExchangeIDsRequestBuilder {
	bb := &UserTranslateExchangeIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/translateExchangeIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserTranslateExchangeIDsRequest struct{ BaseRequest }

//
func (b *UserTranslateExchangeIDsRequestBuilder) Request() *UserTranslateExchangeIDsRequest {
	return &UserTranslateExchangeIDsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserTranslateExchangeIDsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConvertIDResult, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ConvertIDResult
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ConvertIDResult
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *UserTranslateExchangeIDsRequest) PostN(ctx context.Context, n int) ([]ConvertIDResult, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *UserTranslateExchangeIDsRequest) Post(ctx context.Context) ([]ConvertIDResult, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type UserRemoveAllDevicesFromManagementRequestBuilder struct{ BaseRequestBuilder }

// RemoveAllDevicesFromManagement action undocumented
func (b *UserRequestBuilder) RemoveAllDevicesFromManagement(reqObj *UserRemoveAllDevicesFromManagementRequestParameter) *UserRemoveAllDevicesFromManagementRequestBuilder {
	bb := &UserRemoveAllDevicesFromManagementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeAllDevicesFromManagement"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserRemoveAllDevicesFromManagementRequest struct{ BaseRequest }

//
func (b *UserRemoveAllDevicesFromManagementRequestBuilder) Request() *UserRemoveAllDevicesFromManagementRequest {
	return &UserRemoveAllDevicesFromManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserRemoveAllDevicesFromManagementRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserWipeManagedAppRegistrationByDeviceTagRequestBuilder struct{ BaseRequestBuilder }

// WipeManagedAppRegistrationByDeviceTag action undocumented
func (b *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag(reqObj *UserWipeManagedAppRegistrationByDeviceTagRequestParameter) *UserWipeManagedAppRegistrationByDeviceTagRequestBuilder {
	bb := &UserWipeManagedAppRegistrationByDeviceTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/wipeManagedAppRegistrationByDeviceTag"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserWipeManagedAppRegistrationByDeviceTagRequest struct{ BaseRequest }

//
func (b *UserWipeManagedAppRegistrationByDeviceTagRequestBuilder) Request() *UserWipeManagedAppRegistrationByDeviceTagRequest {
	return &UserWipeManagedAppRegistrationByDeviceTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserWipeManagedAppRegistrationByDeviceTagRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct{ BaseRequestBuilder }

// WipeManagedAppRegistrationsByDeviceTag action undocumented
func (b *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag(reqObj *UserWipeManagedAppRegistrationsByDeviceTagRequestParameter) *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder {
	bb := &UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/wipeManagedAppRegistrationsByDeviceTag"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserWipeManagedAppRegistrationsByDeviceTagRequest struct{ BaseRequest }

//
func (b *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Request() *UserWipeManagedAppRegistrationsByDeviceTagRequest {
	return &UserWipeManagedAppRegistrationsByDeviceTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserWipeManagedAppRegistrationsByDeviceTagRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type UserExportPersonalDataRequestBuilder struct{ BaseRequestBuilder }

// ExportPersonalData action undocumented
func (b *UserRequestBuilder) ExportPersonalData(reqObj *UserExportPersonalDataRequestParameter) *UserExportPersonalDataRequestBuilder {
	bb := &UserExportPersonalDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/exportPersonalData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type UserExportPersonalDataRequest struct{ BaseRequest }

//
func (b *UserExportPersonalDataRequestBuilder) Request() *UserExportPersonalDataRequest {
	return &UserExportPersonalDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *UserExportPersonalDataRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
