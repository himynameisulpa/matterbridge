// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OfficeClientCheckinStatus undocumented
type OfficeClientCheckinStatus struct {
	// Object is the base model of OfficeClientCheckinStatus
	Object
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DeviceName undocumented
	DeviceName *string `json:"deviceName,omitempty"`
	// DevicePlatform undocumented
	DevicePlatform *string `json:"devicePlatform,omitempty"`
	// DevicePlatformVersion undocumented
	DevicePlatformVersion *string `json:"devicePlatformVersion,omitempty"`
	// WasSuccessful undocumented
	WasSuccessful *bool `json:"wasSuccessful,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// CheckinDateTime undocumented
	CheckinDateTime *time.Time `json:"checkinDateTime,omitempty"`
	// ErrorMessage undocumented
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// AppliedPolicies undocumented
	AppliedPolicies []string `json:"appliedPolicies,omitempty"`
}

// OfficeClientConfiguration undocumented
type OfficeClientConfiguration struct {
	// Entity is the base model of OfficeClientConfiguration
	Entity
	// UserPreferencePayload undocumented
	UserPreferencePayload *Stream `json:"userPreferencePayload,omitempty"`
	// PolicyPayload undocumented
	PolicyPayload *Stream `json:"policyPayload,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
	// UserCheckinSummary undocumented
	UserCheckinSummary *OfficeUserCheckinSummary `json:"userCheckinSummary,omitempty"`
	// CheckinStatuses undocumented
	CheckinStatuses []OfficeClientCheckinStatus `json:"checkinStatuses,omitempty"`
	// Assignments undocumented
	Assignments []OfficeClientConfigurationAssignment `json:"assignments,omitempty"`
}

// OfficeClientConfigurationAssignment undocumented
type OfficeClientConfigurationAssignment struct {
	// Entity is the base model of OfficeClientConfigurationAssignment
	Entity
	// Target undocumented
	Target *OfficeConfigurationAssignmentTarget `json:"target,omitempty"`
}

// OfficeConfiguration undocumented
type OfficeConfiguration struct {
	// Entity is the base model of OfficeConfiguration
	Entity
	// TenantCheckinStatuses undocumented
	TenantCheckinStatuses []OfficeClientCheckinStatus `json:"tenantCheckinStatuses,omitempty"`
	// TenantUserCheckinSummary undocumented
	TenantUserCheckinSummary *OfficeUserCheckinSummary `json:"tenantUserCheckinSummary,omitempty"`
	// ClientConfigurations undocumented
	ClientConfigurations []OfficeClientConfiguration `json:"clientConfigurations,omitempty"`
}

// OfficeConfigurationAssignmentTarget undocumented
type OfficeConfigurationAssignmentTarget struct {
	// Object is the base model of OfficeConfigurationAssignmentTarget
	Object
}

// OfficeConfigurationGroupAssignmentTarget undocumented
type OfficeConfigurationGroupAssignmentTarget struct {
	// OfficeConfigurationAssignmentTarget is the base model of OfficeConfigurationGroupAssignmentTarget
	OfficeConfigurationAssignmentTarget
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
}

// OfficeGraphInsights undocumented
type OfficeGraphInsights struct {
	// Entity is the base model of OfficeGraphInsights
	Entity
	// Trending undocumented
	Trending []Trending `json:"trending,omitempty"`
	// Shared undocumented
	Shared []SharedInsight `json:"shared,omitempty"`
	// Used undocumented
	Used []UsedInsight `json:"used,omitempty"`
}

// OfficeSuiteApp Contains properties and inherited properties for the Office365 Suite App.
type OfficeSuiteApp struct {
	// MobileApp is the base model of OfficeSuiteApp
	MobileApp
	// AutoAcceptEula The value to accept the EULA automatically on the enduser's device.
	AutoAcceptEula *bool `json:"autoAcceptEula,omitempty"`
	// ProductIDs The Product Ids that represent the Office365 Suite SKU.
	ProductIDs []OfficeProductID `json:"productIds,omitempty"`
	// ExcludedApps The property to represent the apps which are excluded from the selected Office365 Product Id.
	ExcludedApps *ExcludedApps `json:"excludedApps,omitempty"`
	// UseSharedComputerActivation The property to represent that whether the shared computer activation is used not for Office365 app suite.
	UseSharedComputerActivation *bool `json:"useSharedComputerActivation,omitempty"`
	// UpdateChannel The property to represent the Office365 Update Channel.
	UpdateChannel *OfficeUpdateChannel `json:"updateChannel,omitempty"`
	// OfficePlatformArchitecture The property to represent the Office365 app suite version.
	OfficePlatformArchitecture *WindowsArchitecture `json:"officePlatformArchitecture,omitempty"`
	// LocalesToInstall The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard RFC 6033. Ref: https://technet.microsoft.com/en-us/library/cc179219(v=office.16).aspx
	LocalesToInstall []string `json:"localesToInstall,omitempty"`
	// InstallProgressDisplayLevel To specify the level of display for the Installation Progress Setup UI on the Device.
	InstallProgressDisplayLevel *OfficeSuiteInstallProgressDisplayLevel `json:"installProgressDisplayLevel,omitempty"`
	// ShouldUninstallOlderVersionsOfOffice The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the device or not.
	ShouldUninstallOlderVersionsOfOffice *bool `json:"shouldUninstallOlderVersionsOfOffice,omitempty"`
	// TargetVersion The property to represent the specific target version for the Office365 app suite that should be remained deployed on the devices.
	TargetVersion *string `json:"targetVersion,omitempty"`
	// UpdateVersion The property to represent the update version in which the specific target version is available for the Office365 app suite.
	UpdateVersion *string `json:"updateVersion,omitempty"`
	// OfficeConfigurationXML The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence over all other properties. When present, the XML configuration file will be used to create the app.
	OfficeConfigurationXML *Binary `json:"officeConfigurationXml,omitempty"`
}

// OfficeUserCheckinSummary undocumented
type OfficeUserCheckinSummary struct {
	// Object is the base model of OfficeUserCheckinSummary
	Object
	// SucceededUserCount undocumented
	SucceededUserCount *int `json:"succeededUserCount,omitempty"`
	// FailedUserCount undocumented
	FailedUserCount *int `json:"failedUserCount,omitempty"`
}
