package okta

type LogDevice struct {
	Id                    string                 `json:"id,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	OsPlatform            *string                `json:"os_platform,omitempty"`
	OsVersion             *string                `json:"os_version,omitempty"`
	Managed               *bool                  `json:"managed,omitempty"`
	Registered            *bool                  `json:"registered,omitempty"`
	DeviceIntegrator      map[string]interface{} `json:"device_integrator,omitempty"`
	DiskEncryptionType    *string                `json:"disk_encryption_type,omitempty"`
	ScreenLockType        *string                `json:"screen_lock_type,omitempty"`
	Jailbreak             *bool                  `json:"jailbreak,omitempty"`
	SecureHardwarePresent *bool                  `json:"secure_hardware_present,omitempty"`
}

/*
	"id": "guob5wtu7rBggkg9G1d7",
    "name": "MacBookPro16,1",
    "os_platform": "OSX",
    "os_version": "14.3.0",
    "managed": false,
    "registered": true,
    "device_integrator": null,
    "disk_encryption_type": "ALL_INTERNAL_VOLUMES",
    "screen_lock_type": "BIOMETRIC",
    "jailbreak": null,
    "secure_hardware_present": true
*/
