package nvidiasmijson

type NvidiaSmiLog struct {
	Timestamp      string `xml:"timestamp" json:"timestamp"`
	KmdVersion     string `xml:"kmd_version" json:"kmd_version"`
	CudaUmdVersion string `xml:"cuda_umd_version" json:"cuda_umd_version"`
	AttachedGpus   string `xml:"attached_gpus" json:"attached_gpus"`
	GPUS           []GPU  `xml:"gpu" json:"gpus"`
}

type GPU struct {
	ID                         string                     `xml:"id,attr" json:"id"`
	ProductName                string                     `xml:"product_name" json:"product_name"`
	ProductBrand               string                     `xml:"product_brand" json:"product_brand"`
	ProductArchitecture        string                     `xml:"product_architecture" json:"product_architecture"`
	DisplayAttached            string                     `xml:"display_attached" json:"display_attached"`
	DisplayActive              string                     `xml:"display_active" json:"display_active"`
	PersistenceMode            string                     `xml:"persistence_mode" json:"persistence_mode"`
	AddressingMode             string                     `xml:"addressing_mode" json:"addressing_mode"`
	MigMode                    MigMode                    `xml:"mig_mode" json:"mig_mode"`
	MigDevices                 []MigDevice                `xml:"mig_devices>mig_device" json:"mig_devices"`
	AccountingMode             string                     `xml:"accounting_mode" json:"accounting_mode"`
	AccountingModeBufferSize   string                     `xml:"accounting_mode_buffer_size" json:"accounting_mode_buffer_size"`
	DriverModel                DriverModel                `xml:"driver_model" json:"driver_model"`
	Serial                     string                     `xml:"serial" json:"serial"`
	UUID                       string                     `xml:"uuid" json:"uuid"`
	Pdi                        string                     `xml:"pdi" json:"pdi"`
	MinorNumber                string                     `xml:"minor_number" json:"minor_number"`
	VbiosVersion               string                     `xml:"vbios_version" json:"vbios_version"`
	MultigpuBoard              string                     `xml:"multigpu_board" json:"multigpu_board"`
	BoardID                    string                     `xml:"board_id" json:"board_id"`
	BoardPartNumber            string                     `xml:"board_part_number" json:"board_part_number"`
	GpuPartNumber              string                     `xml:"gpu_part_number" json:"gpu_part_number"`
	GpuFruPartNumber           string                     `xml:"gpu_fru_part_number" json:"gpu_fru_part_number"`
	PlatformInfo               *PlatformInfo              `xml:"platformInfo" json:"platform_info,omitempty"`
	InforomVersion             InforomVersion             `xml:"inforom_version" json:"inforom_version"`
	InforomBbxFlush            *InforomBbxFlush           `xml:"inforom_bbx_flush" json:"inforom_bbx_flush,omitempty"`
	InforomTimeData            *InforomTimeData           `xml:"inforom_time_data" json:"inforom_time_data,omitempty"`
	GpuOperationMode           GpuOperationMode           `xml:"gpu_operation_mode" json:"gpu_operation_mode"`
	C2cMode                    string                     `xml:"c2c_mode" json:"c2c_mode"`
	GpuVirtualizationMode      GpuVirtualizationMode      `xml:"gpu_virtualization_mode" json:"gpu_virtualization_mode"`
	GpuRecoveryAction          string                     `xml:"gpu_recovery_action" json:"gpu_recovery_action"`
	GspFirmwareVersion         string                     `xml:"gsp_firmware_version" json:"gsp_firmware_version"`
	Ibmnpu                     *Ibmnpu                    `xml:"ibmnpu" json:"ibmnpu,omitempty"`
	Pci                        Pci                        `xml:"pci" json:"pci"`
	FanSpeed                   string                     `xml:"fan_speed" json:"fan_speed"`
	PerformanceState           string                     `xml:"performance_state" json:"performance_state"`
	ClocksEventReasons         ClocksEventReasons         `xml:"clocks_event_reasons" json:"clocks_event_reasons"`
	ClocksEventReasonsCounters ClocksEventReasonsCounters `xml:"clocks_event_reasons_counters" json:"clocks_event_reasons_counters"`
	SparseOperationMode        string                     `xml:"sparse_operation_mode" json:"sparse_operation_mode"`
	FbMemoryUsage              MemoryUsage                `xml:"fb_memory_usage" json:"fb_memory_usage"`
	Bar1MemoryUsage            MemoryUsage                `xml:"bar1_memory_usage" json:"bar1_memory_usage"`
	CcProtectedMemoryUsage     *MemoryUsage               `xml:"cc_protected_memory_usage" json:"cc_protected_memory_usage,omitempty"`
	ComputeMode                string                     `xml:"compute_mode" json:"compute_mode"`
	Utilization                Utilization                `xml:"utilization" json:"utilization"`
	EncoderStats               SessionStats               `xml:"encoder_stats" json:"encoder_stats"`
	FbcStats                   SessionStats               `xml:"fbc_stats" json:"fbc_stats"`
	DramEncryptionMode         *DramEncryptionMode        `xml:"dram_encryption_mode" json:"dram_encryption_mode,omitempty"`
	EccMode                    EccMode                    `xml:"ecc_mode" json:"ecc_mode"`
	EccErrors                  EccErrors                  `xml:"ecc_errors" json:"ecc_errors"`
	RetiredPages               RetiredPages               `xml:"retired_pages" json:"retired_pages"`
	RemappedRows               string                     `xml:"remapped_rows" json:"remapped_rows"`
	RemappedBanks              string                     `xml:"remapped_banks" json:"remapped_banks"`
	Temperature                Temperature                `xml:"temperature" json:"temperature"`
	SupportedGpuTargetTemp     SupportedGpuTargetTemp     `xml:"supported_gpu_target_temp" json:"supported_gpu_target_temp"`
	GpuPowerReadings           GpuPowerReadings           `xml:"gpu_power_readings" json:"gpu_power_readings"`
	GpuMemoryPowerReadings     PowerDraw                  `xml:"gpu_memory_power_readings" json:"gpu_memory_power_readings"`
	ModulePowerReadings        ModulePowerReadings        `xml:"module_power_readings" json:"module_power_readings"`
	PowerSmoothing             string                     `xml:"power_smoothing" json:"power_smoothing"`
	PowerProfiles              *PowerProfiles             `xml:"power_profiles" json:"power_profiles,omitempty"`
	EdppMultiplier             string                     `xml:"edpp_multiplier" json:"edpp_multiplier"`
	Clocks                     Clocks                     `xml:"clocks" json:"clocks"`
	DeferredClocks             Clocks                     `xml:"deferred_clocks" json:"deferred_clocks"`
	MaxClocks                  Clocks                     `xml:"max_clocks" json:"max_clocks"`
	MaxCustomerBoostClocks     Clocks                     `xml:"max_customer_boost_clocks" json:"max_customer_boost_clocks"`
	ClockPolicy                ClockPolicy                `xml:"clock_policy" json:"clock_policy"`
	DlaClocks                  *DlaClocks                 `xml:"dla_clocks" json:"dla_clocks,omitempty"`
	SupportedClocks            []SupportedMemClock        `xml:"supported_clocks>supported_mem_clock" json:"supported_clocks"`
	Fabric                     Fabric                     `xml:"fabric" json:"fabric"`
	Processes                  []ProcessInfo              `xml:"processes>process_info" json:"processes"`
	AccountedProcesses         []AccountedProcess         `xml:"accounted_processes>accounted_process_info" json:"accounted_processes"`
	Capabilities               *Capabilities              `xml:"capabilities" json:"capabilities,omitempty"`
}

type MigMode struct {
	CurrentMig string `xml:"current_mig" json:"current_mig"`
	PendingMig string `xml:"pending_mig" json:"pending_mig"`
}

type MigDevice struct {
	Index             string           `xml:"index" json:"index"`
	GpuInstanceID     string           `xml:"gpu_instance_id" json:"gpu_instance_id"`
	ComputeInstanceID string           `xml:"compute_instance_id" json:"compute_instance_id"`
	DeviceAttributes  DeviceAttributes `xml:"device_attributes" json:"device_attributes"`
	EccErrorCount     MigEccErrorCount `xml:"ecc_error_count" json:"ecc_error_count"`
	FbMemoryUsage     MemoryUsage      `xml:"fb_memory_usage" json:"fb_memory_usage"`
	Bar1MemoryUsage   MemoryUsage      `xml:"bar1_memory_usage" json:"bar1_memory_usage"`
}

type DeviceAttributes struct {
	Shared SharedDeviceAttributes `xml:"shared" json:"shared"`
}

type SharedDeviceAttributes struct {
	MultiprocessorCount string `xml:"multiprocessor_count" json:"multiprocessor_count"`
	CopyEngineCount     string `xml:"copy_engine_count" json:"copy_engine_count"`
	EncoderCount        string `xml:"encoder_count" json:"encoder_count"`
	DecoderCount        string `xml:"decoder_count" json:"decoder_count"`
	OfaCount            string `xml:"ofa_count" json:"ofa_count"`
	JpgCount            string `xml:"jpg_count" json:"jpg_count"`
}

type MigEccErrorCount struct {
	VolatileCount MigEccVolatileCount `xml:"volatile_count" json:"volatile_count"`
}

type MigEccVolatileCount struct {
	SramUncorrectable string `xml:"sram_uncorrectable" json:"sram_uncorrectable"`
}

type DriverModel struct {
	CurrentDm string `xml:"current_dm" json:"current_dm"`
	PendingDm string `xml:"pending_dm" json:"pending_dm"`
}

type PlatformInfo struct {
	ChassisSerialNumber string `xml:"chassis_serial_number" json:"chassis_serial_number"`
	SlotNumber          string `xml:"slot_number" json:"slot_number"`
	TrayIndex           string `xml:"tray_index" json:"tray_index"`
	HostID              string `xml:"host_id" json:"host_id"`
	PeerType            string `xml:"peer_type" json:"peer_type"`
	ModuleID            string `xml:"module_id" json:"module_id"`
	GpuFabricGUID       string `xml:"gpu_fabric_guid" json:"gpu_fabric_guid"`
}

type InforomVersion struct {
	ImgVersion string `xml:"img_version" json:"img_version"`
	OemObject  string `xml:"oem_object" json:"oem_object"`
	EccObject  string `xml:"ecc_object" json:"ecc_object"`
	PwrObject  string `xml:"pwr_object" json:"pwr_object"`
}

type InforomBbxFlush struct {
	LatestTimestamp string `xml:"latest_timestamp" json:"latest_timestamp"`
	LatestDuration  string `xml:"latest_duration" json:"latest_duration"`
}

type InforomTimeData struct {
	TimeRun string `xml:"time_run" json:"time_run"`
}

type GpuOperationMode struct {
	CurrentGom string `xml:"current_gom" json:"current_gom"`
	PendingGom string `xml:"pending_gom" json:"pending_gom"`
}

type GpuVirtualizationMode struct {
	VirtualizationMode    string `xml:"virtualization_mode" json:"virtualization_mode"`
	HostVgpuMode          string `xml:"host_vgpu_mode" json:"host_vgpu_mode"`
	VgpuHeterogeneousMode string `xml:"vgpu_heterogeneous_mode" json:"vgpu_heterogeneous_mode"`
}

type Ibmnpu struct {
	RelaxedOrderingMode string `xml:"relaxed_ordering_mode" json:"relaxed_ordering_mode"`
}

type Pci struct {
	PciBus                string         `xml:"pci_bus" json:"pci_bus"`
	PciDevice             string         `xml:"pci_device" json:"pci_device"`
	PciDomain             string         `xml:"pci_domain" json:"pci_domain"`
	PciBaseClass          string         `xml:"pci_base_class" json:"pci_base_class"`
	PciSubClass           string         `xml:"pci_sub_class" json:"pci_sub_class"`
	PciDeviceID           string         `xml:"pci_device_id" json:"pci_device_id"`
	PciBusID              string         `xml:"pci_bus_id" json:"pci_bus_id"`
	PciSubSystemID        string         `xml:"pci_sub_system_id" json:"pci_sub_system_id"`
	PciGpuLinkInfo        PciGpuLinkInfo `xml:"pci_gpu_link_info" json:"pci_gpu_link_info"`
	PciBridgeChip         *PciBridgeChip `xml:"pci_bridge_chip" json:"pci_bridge_chip,omitempty"`
	ReplayCounter         string         `xml:"replay_counter" json:"replay_counter"`
	ReplayRolloverCounter string         `xml:"replay_rollover_counter" json:"replay_rollover_counter"`
	TxUtil                string         `xml:"tx_util" json:"tx_util"`
	RxUtil                string         `xml:"rx_util" json:"rx_util"`
	AtomicCapsOutbound    string         `xml:"atomic_caps_outbound" json:"atomic_caps_outbound"`
	AtomicCapsInbound     string         `xml:"atomic_caps_inbound" json:"atomic_caps_inbound"`
}

type PciGpuLinkInfo struct {
	PcieGen    PcieGen    `xml:"pcie_gen" json:"pcie_gen"`
	LinkWidths LinkWidths `xml:"link_widths" json:"link_widths"`
}

type PcieGen struct {
	MaxLinkGen           string `xml:"max_link_gen" json:"max_link_gen"`
	CurrentLinkGen       string `xml:"current_link_gen" json:"current_link_gen"`
	DeviceCurrentLinkGen string `xml:"device_current_link_gen" json:"device_current_link_gen"`
	MaxDeviceLinkGen     string `xml:"max_device_link_gen" json:"max_device_link_gen"`
	MaxHostLinkGen       string `xml:"max_host_link_gen" json:"max_host_link_gen"`
}

type LinkWidths struct {
	MaxLinkWidth     string `xml:"max_link_width" json:"max_link_width"`
	CurrentLinkWidth string `xml:"current_link_width" json:"current_link_width"`
}

type PciBridgeChip struct {
	BridgeChipType string `xml:"bridge_chip_type" json:"bridge_chip_type"`
	BridgeChipFw   string `xml:"bridge_chip_fw" json:"bridge_chip_fw"`
}

type ClocksEventReasons struct {
	GpuIdle                   string `xml:"clocks_event_reason_gpu_idle" json:"clocks_event_reason_gpu_idle"`
	ApplicationsClocksSetting string `xml:"clocks_event_reason_applications_clocks_setting" json:"clocks_event_reason_applications_clocks_setting"`
	SwPowerCap                string `xml:"clocks_event_reason_sw_power_cap" json:"clocks_event_reason_sw_power_cap"`
	HwSlowdown                string `xml:"clocks_event_reason_hw_slowdown" json:"clocks_event_reason_hw_slowdown"`
	HwThermalSlowdown         string `xml:"clocks_event_reason_hw_thermal_slowdown" json:"clocks_event_reason_hw_thermal_slowdown"`
	HwPowerBrakeSlowdown      string `xml:"clocks_event_reason_hw_power_brake_slowdown" json:"clocks_event_reason_hw_power_brake_slowdown"`
	SyncBoost                 string `xml:"clocks_event_reason_sync_boost" json:"clocks_event_reason_sync_boost"`
	SwThermalSlowdown         string `xml:"clocks_event_reason_sw_thermal_slowdown" json:"clocks_event_reason_sw_thermal_slowdown"`
	DisplayClocksSetting      string `xml:"clocks_event_reason_display_clocks_setting" json:"clocks_event_reason_display_clocks_setting"`
	BoardLimit                string `xml:"clocks_event_reason_board_limit" json:"clocks_event_reason_board_limit"`
	Reliability               string `xml:"clocks_event_reason_reliability" json:"clocks_event_reason_reliability"`
}

type ClocksEventReasonsCounters struct {
	SwPowerCap      string `xml:"clocks_event_reasons_counters_sw_power_cap" json:"clocks_event_reasons_counters_sw_power_cap"`
	SyncBoost       string `xml:"clocks_event_reasons_counters_sync_boost" json:"clocks_event_reasons_counters_sync_boost"`
	SwThermSlowdown string `xml:"clocks_event_reasons_counters_sw_therm_slowdown" json:"clocks_event_reasons_counters_sw_therm_slowdown"`
	HwThermSlowdown string `xml:"clocks_event_reasons_counters_hw_therm_slowdown" json:"clocks_event_reasons_counters_hw_therm_slowdown"`
	HwPowerBrake    string `xml:"clocks_event_reasons_counters_hw_power_brake" json:"clocks_event_reasons_counters_hw_power_brake"`
}

type MemoryUsage struct {
	Total    string `xml:"total" json:"total"`
	Reserved string `xml:"reserved" json:"reserved,omitempty"`
	Used     string `xml:"used" json:"used"`
	Free     string `xml:"free" json:"free"`
}

type Utilization struct {
	GpuUtil     string `xml:"gpu_util" json:"gpu_util"`
	MemoryUtil  string `xml:"memory_util" json:"memory_util"`
	EncoderUtil string `xml:"encoder_util" json:"encoder_util"`
	DecoderUtil string `xml:"decoder_util" json:"decoder_util"`
	JpegUtil    string `xml:"jpeg_util" json:"jpeg_util"`
	OfaUtil     string `xml:"ofa_util" json:"ofa_util"`
}

type SessionStats struct {
	SessionCount   string `xml:"session_count" json:"session_count"`
	AverageFps     string `xml:"average_fps" json:"average_fps"`
	AverageLatency string `xml:"average_latency" json:"average_latency"`
}

type DramEncryptionMode struct {
	CurrentDramEncryption string `xml:"current_dram_encryption" json:"current_dram_encryption"`
	PendingDramEncryption string `xml:"pending_dram_encryption" json:"pending_dram_encryption"`
}

type EccMode struct {
	CurrentEcc string `xml:"current_ecc" json:"current_ecc"`
	PendingEcc string `xml:"pending_ecc" json:"pending_ecc"`
}

type EccErrors struct {
	Volatile                          EccErrorCounts           `xml:"volatile" json:"volatile"`
	Aggregate                         EccErrorCounts           `xml:"aggregate" json:"aggregate"`
	AggregateUncorrectableSramSources UncorrectableSramSources `xml:"aggregate_uncorrectable_sram_sources" json:"aggregate_uncorrectable_sram_sources"`
	ChannelRepairPending              string                   `xml:"channel_repair_pending" json:"channel_repair_pending"`
	TpcRepairPending                  string                   `xml:"tpc_repair_pending" json:"tpc_repair_pending"`
	UnrepairableMemory                string                   `xml:"unrepairable_memory" json:"unrepairable_memory"`
}

type EccErrorCounts struct {
	SramCorrectable         string `xml:"sram_correctable" json:"sram_correctable"`
	SramUncorrectableParity string `xml:"sram_uncorrectable_parity" json:"sram_uncorrectable_parity"`
	SramUncorrectableSecded string `xml:"sram_uncorrectable_secded" json:"sram_uncorrectable_secded"`
	DramCorrectable         string `xml:"dram_correctable" json:"dram_correctable"`
	DramUncorrectable       string `xml:"dram_uncorrectable" json:"dram_uncorrectable"`
	SramThresholdExceeded   string `xml:"sram_threshold_exceeded" json:"sram_threshold_exceeded,omitempty"`
}

type UncorrectableSramSources struct {
	SramL2              string `xml:"sram_l2" json:"sram_l2"`
	SramSm              string `xml:"sram_sm" json:"sram_sm"`
	SramMicrocontroller string `xml:"sram_microcontroller" json:"sram_microcontroller"`
	SramPcie            string `xml:"sram_pcie" json:"sram_pcie"`
	SramOther           string `xml:"sram_other" json:"sram_other"`
}

type RetiredPages struct {
	MultipleSingleBitRetirement PageRetirement `xml:"multiple_single_bit_retirement" json:"multiple_single_bit_retirement"`
	DoubleBitRetirement         PageRetirement `xml:"double_bit_retirement" json:"double_bit_retirement"`
	PendingBlacklist            string         `xml:"pending_blacklist" json:"pending_blacklist"`
	PendingRetirement           string         `xml:"pending_retirement" json:"pending_retirement"`
}

type PageRetirement struct {
	RetiredCount    string `xml:"retired_count" json:"retired_count"`
	RetiredPagelist string `xml:"retired_pagelist" json:"retired_pagelist"`
}

type Temperature struct {
	GpuTemp                      string `xml:"gpu_temp" json:"gpu_temp"`
	GpuTempMax                   string `xml:"gpu_temp_max" json:"gpu_temp_max"`
	GpuTempTlimit                string `xml:"gpu_temp_tlimit" json:"gpu_temp_tlimit"`
	GpuTempMaxTlimitThreshold    string `xml:"gpu_temp_max_tlimit_threshold" json:"gpu_temp_max_tlimit_threshold"`
	GpuTempSlowTlimitThreshold   string `xml:"gpu_temp_slow_tlimit_threshold" json:"gpu_temp_slow_tlimit_threshold"`
	GpuTempMaxGpuTlimitThreshold string `xml:"gpu_temp_max_gpu_tlimit_threshold" json:"gpu_temp_max_gpu_tlimit_threshold"`
	GpuTargetTemperature         string `xml:"gpu_target_temperature" json:"gpu_target_temperature"`
	MemoryTemp                   string `xml:"memory_temp" json:"memory_temp"`
	GpuTempMaxMemTlimitThreshold string `xml:"gpu_temp_max_mem_tlimit_threshold" json:"gpu_temp_max_mem_tlimit_threshold"`
}

type SupportedGpuTargetTemp struct {
	GpuTargetTempMin string `xml:"gpu_target_temp_min" json:"gpu_target_temp_min"`
	GpuTargetTempMax string `xml:"gpu_target_temp_max" json:"gpu_target_temp_max"`
}

type GpuPowerReadings struct {
	AveragePowerDraw     string        `xml:"average_power_draw" json:"average_power_draw"`
	InstantPowerDraw     string        `xml:"instant_power_draw" json:"instant_power_draw"`
	GpuCeilingPowerLimit PowerLimits   `xml:"gpu_ceiling_power_limit" json:"gpu_ceiling_power_limit"`
	GpuBasePower         *GpuBasePower `xml:"gpu_base_power" json:"gpu_base_power,omitempty"`
	CurrentPowerLimit    string        `xml:"current_power_limit" json:"current_power_limit"`
	RequestedPowerLimit  string        `xml:"requested_power_limit" json:"requested_power_limit"`
	DefaultPowerLimit    string        `xml:"default_power_limit" json:"default_power_limit"`
	MinPowerLimit        string        `xml:"min_power_limit" json:"min_power_limit"`
	MaxPowerLimit        string        `xml:"max_power_limit" json:"max_power_limit"`
}

type GpuBasePower struct {
	CurrentBasePower             string `xml:"current_base_power" json:"current_base_power"`
	RequestedBasePower           string `xml:"requested_base_power" json:"requested_base_power"`
	DefaultBasePower             string `xml:"default_base_power" json:"default_base_power"`
	AdaptiveTgpAdjustedBasePower string `xml:"adaptive_tgp_adjusted_base_power" json:"adaptive_tgp_adjusted_base_power"`
}

type PowerLimits struct {
	CurrentPowerLimit   string `xml:"current_power_limit" json:"current_power_limit"`
	RequestedPowerLimit string `xml:"requested_power_limit" json:"requested_power_limit"`
	DefaultPowerLimit   string `xml:"default_power_limit" json:"default_power_limit"`
}

type PowerDraw struct {
	AveragePowerDraw string `xml:"average_power_draw" json:"average_power_draw"`
	InstantPowerDraw string `xml:"instant_power_draw" json:"instant_power_draw"`
}

type ModulePowerReadings struct {
	AveragePowerDraw    string `xml:"average_power_draw" json:"average_power_draw"`
	InstantPowerDraw    string `xml:"instant_power_draw" json:"instant_power_draw"`
	CurrentPowerLimit   string `xml:"current_power_limit" json:"current_power_limit"`
	RequestedPowerLimit string `xml:"requested_power_limit" json:"requested_power_limit"`
	DefaultPowerLimit   string `xml:"default_power_limit" json:"default_power_limit"`
	MinPowerLimit       string `xml:"min_power_limit" json:"min_power_limit"`
	MaxPowerLimit       string `xml:"max_power_limit" json:"max_power_limit"`
}

type PowerProfiles struct {
	RequestedProfiles string `xml:"power_profile_requested_profiles" json:"power_profile_requested_profiles"`
	EnforcedProfiles  string `xml:"power_profile_enforced_profiles" json:"power_profile_enforced_profiles"`
}

// Clocks is shared by clocks, deferred_clocks, max_clocks and
// max_customer_boost_clocks; each section reports only a subset of fields.
type Clocks struct {
	GraphicsClock string `xml:"graphics_clock" json:"graphics_clock,omitempty"`
	SmClock       string `xml:"sm_clock" json:"sm_clock,omitempty"`
	MemClock      string `xml:"mem_clock" json:"mem_clock,omitempty"`
	VideoClock    string `xml:"video_clock" json:"video_clock,omitempty"`
}

type ClockPolicy struct {
	AutoBoost        string `xml:"auto_boost" json:"auto_boost"`
	AutoBoostDefault string `xml:"auto_boost_default" json:"auto_boost_default"`
}

type DlaClocks struct {
	DlaCurEngineClock string `xml:"dla_cur_engine_clock" json:"dla_cur_engine_clock"`
	DlaMaxEngineClock string `xml:"dla_max_engine_clock" json:"dla_max_engine_clock"`
}

type SupportedMemClock struct {
	Value                   string   `xml:"value" json:"value"`
	SupportedGraphicsClocks []string `xml:"supported_graphics_clock" json:"supported_graphics_clocks"`
}

type Fabric struct {
	State       string       `xml:"state" json:"state"`
	Status      string       `xml:"status" json:"status"`
	CliqueID    string       `xml:"cliqueId" json:"clique_id"`
	ClusterUUID string       `xml:"clusterUuid" json:"cluster_uuid"`
	Health      FabricHealth `xml:"health" json:"health"`
}

type FabricHealth struct {
	Summary                 string `xml:"summary" json:"summary"`
	Bandwidth               string `xml:"bandwidth" json:"bandwidth"`
	RouteRecoveryInProgress string `xml:"route_recovery_in_progress" json:"route_recovery_in_progress"`
	RouteUnhealthy          string `xml:"route_unhealthy" json:"route_unhealthy"`
	AccessTimeoutRecovery   string `xml:"access_timeout_recovery" json:"access_timeout_recovery"`
	IncorrectConfiguration  string `xml:"incorrect_configuration" json:"incorrect_configuration"`
	PartitionAssigned       string `xml:"partition_assigned" json:"partition_assigned"`
	GfmState                string `xml:"gfm_state" json:"gfm_state"`
}

type ProcessInfo struct {
	GpuInstanceID     string `xml:"gpu_instance_id" json:"gpu_instance_id"`
	ComputeInstanceID string `xml:"compute_instance_id" json:"compute_instance_id"`
	Pid               string `xml:"pid" json:"pid"`
	Type              string `xml:"type" json:"type"`
	ProcessName       string `xml:"process_name" json:"process_name"`
	UsedMemory        string `xml:"used_memory" json:"used_memory"`
}

type AccountedProcess struct {
	Pid            string `xml:"pid" json:"pid"`
	GpuUtil        string `xml:"gpu_util" json:"gpu_util"`
	MemoryUtil     string `xml:"memory_util" json:"memory_util"`
	MaxMemoryUsage string `xml:"max_memory_usage" json:"max_memory_usage"`
	Time           string `xml:"time" json:"time"`
	IsRunning      string `xml:"is_running" json:"is_running"`
}

type Capabilities struct {
	Egm string `xml:"egm" json:"egm"`
}
