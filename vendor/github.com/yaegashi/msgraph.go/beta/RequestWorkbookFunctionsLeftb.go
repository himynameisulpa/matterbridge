// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsLeftbRequestBuilder struct{ BaseRequestBuilder }

// Leftb action undocumented
func (b *WorkbookFunctionsRequestBuilder) Leftb(reqObj *WorkbookFunctionsLeftbRequestParameter) *WorkbookFunctionsLeftbRequestBuilder {
	bb := &WorkbookFunctionsLeftbRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/leftb"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsLeftbRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsLeftbRequestBuilder) Request() *WorkbookFunctionsLeftbRequest {
	return &WorkbookFunctionsLeftbRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsLeftbRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
