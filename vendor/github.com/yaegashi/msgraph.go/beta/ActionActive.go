// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DomainJoinConfiguration is navigation property
func (b *ActiveDirectoryWindowsAutopilotDeploymentProfileRequestBuilder) DomainJoinConfiguration() *WindowsDomainJoinConfigurationRequestBuilder {
	bb := &WindowsDomainJoinConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/domainJoinConfiguration"
	return bb
}
