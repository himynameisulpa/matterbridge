// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerBatteryPluggedMode undocumented
type AndroidDeviceOwnerBatteryPluggedMode int

const (
	// AndroidDeviceOwnerBatteryPluggedModeVNotConfigured undocumented
	AndroidDeviceOwnerBatteryPluggedModeVNotConfigured AndroidDeviceOwnerBatteryPluggedMode = 0
	// AndroidDeviceOwnerBatteryPluggedModeVAc undocumented
	AndroidDeviceOwnerBatteryPluggedModeVAc AndroidDeviceOwnerBatteryPluggedMode = 1
	// AndroidDeviceOwnerBatteryPluggedModeVUsb undocumented
	AndroidDeviceOwnerBatteryPluggedModeVUsb AndroidDeviceOwnerBatteryPluggedMode = 2
	// AndroidDeviceOwnerBatteryPluggedModeVWireless undocumented
	AndroidDeviceOwnerBatteryPluggedModeVWireless AndroidDeviceOwnerBatteryPluggedMode = 3
)

// AndroidDeviceOwnerBatteryPluggedModePNotConfigured returns a pointer to AndroidDeviceOwnerBatteryPluggedModeVNotConfigured
func AndroidDeviceOwnerBatteryPluggedModePNotConfigured() *AndroidDeviceOwnerBatteryPluggedMode {
	v := AndroidDeviceOwnerBatteryPluggedModeVNotConfigured
	return &v
}

// AndroidDeviceOwnerBatteryPluggedModePAc returns a pointer to AndroidDeviceOwnerBatteryPluggedModeVAc
func AndroidDeviceOwnerBatteryPluggedModePAc() *AndroidDeviceOwnerBatteryPluggedMode {
	v := AndroidDeviceOwnerBatteryPluggedModeVAc
	return &v
}

// AndroidDeviceOwnerBatteryPluggedModePUsb returns a pointer to AndroidDeviceOwnerBatteryPluggedModeVUsb
func AndroidDeviceOwnerBatteryPluggedModePUsb() *AndroidDeviceOwnerBatteryPluggedMode {
	v := AndroidDeviceOwnerBatteryPluggedModeVUsb
	return &v
}

// AndroidDeviceOwnerBatteryPluggedModePWireless returns a pointer to AndroidDeviceOwnerBatteryPluggedModeVWireless
func AndroidDeviceOwnerBatteryPluggedModePWireless() *AndroidDeviceOwnerBatteryPluggedMode {
	v := AndroidDeviceOwnerBatteryPluggedModeVWireless
	return &v
}