// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AdvancedThreatProtectionOnboardingDeviceSettingStates returns request builder for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (b *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates() *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder {
	bb := &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/advancedThreatProtectionOnboardingDeviceSettingStates"
	return bb
}

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder is request builder for AdvancedThreatProtectionOnboardingDeviceSettingState collection
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (b *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder) Request() *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest {
	return &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AdvancedThreatProtectionOnboardingDeviceSettingState item
func (b *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequestBuilder) ID(id string) *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder {
	bb := &AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest is request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
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
	var values []AdvancedThreatProtectionOnboardingDeviceSettingState
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
			value  []AdvancedThreatProtectionOnboardingDeviceSettingState
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

// GetN performs GET request for AdvancedThreatProtectionOnboardingDeviceSettingState collection, max N pages
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) GetN(ctx context.Context, n int) ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Get(ctx context.Context) ([]AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AdvancedThreatProtectionOnboardingDeviceSettingState collection
func (r *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesCollectionRequest) Add(ctx context.Context, reqObj *AdvancedThreatProtectionOnboardingDeviceSettingState) (resObj *AdvancedThreatProtectionOnboardingDeviceSettingState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
