// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SecureScoreRequestBuilder is request builder for SecureScore
type SecureScoreRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecureScoreRequest
func (b *SecureScoreRequestBuilder) Request() *SecureScoreRequest {
	return &SecureScoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecureScoreRequest is request for SecureScore
type SecureScoreRequest struct{ BaseRequest }

// Get performs GET request for SecureScore
func (r *SecureScoreRequest) Get(ctx context.Context) (resObj *SecureScore, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecureScore
func (r *SecureScoreRequest) Update(ctx context.Context, reqObj *SecureScore) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecureScore
func (r *SecureScoreRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecureScoreControlProfileRequestBuilder is request builder for SecureScoreControlProfile
type SecureScoreControlProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecureScoreControlProfileRequest
func (b *SecureScoreControlProfileRequestBuilder) Request() *SecureScoreControlProfileRequest {
	return &SecureScoreControlProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecureScoreControlProfileRequest is request for SecureScoreControlProfile
type SecureScoreControlProfileRequest struct{ BaseRequest }

// Get performs GET request for SecureScoreControlProfile
func (r *SecureScoreControlProfileRequest) Get(ctx context.Context) (resObj *SecureScoreControlProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecureScoreControlProfile
func (r *SecureScoreControlProfileRequest) Update(ctx context.Context, reqObj *SecureScoreControlProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecureScoreControlProfile
func (r *SecureScoreControlProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
