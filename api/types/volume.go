package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Volume volume
// swagger:model Volume
type Volume struct {

	// Date/Time the volume was created.
	CreatedAt string `json:"CreatedAt,omitempty"`

	// Name of the volume driver used by the volume.
	// Required: true
	Driver string `json:"Driver"`

	// User-defined key/value metadata.
	// Required: true
	Labels map[string]string `json:"Labels"`

	// Mount path of the volume on the host.
	// Required: true
	Mountpoint string `json:"Mountpoint"`

	// Name of the volume.
	// Required: true
	Name string `json:"Name"`

	// The driver specific options used when creating the volume.
	// Required: true
	Options map[string]string `json:"Options"`

	// The level at which the volume exists. Either `global` for cluster-wide, or `local` for machine level.
	// Required: true
	Scope string `json:"Scope"`

	// Low-level details about the volume, provided by the volume driver.
	// Details are returned as a map with key/value pairs:
	// `{"key":"value","key2":"value2"}`.
	//
	// The `Status` field is optional, and is omitted if the volume driver
	// does not support this feature.
	//
	Status map[string]interface{} `json:"Status,omitempty"`

	// usage data
	UsageData *VolumeUsageData `json:"UsageData,omitempty"`
}

// VolumeUsageData volume usage data
// swagger:model VolumeUsageData
type VolumeUsageData struct {

	// The number of containers referencing this volume.
	// Required: true
	RefCount int64 `json:"RefCount"`

	// The disk space used by the volume (local driver only)
	// Required: true
	Size int64 `json:"Size"`
}
