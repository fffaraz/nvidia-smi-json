package nvidiasmijson

type GPU struct {
	ID                                         string              `xml:"id,attr" json:"id"`
	ProductName                                string              `xml:"product_name" json:"product_name"`
	ProductBrand                               string              `xml:"product_brand" json:"product_brand"`
	ProductArchitecture                        string              `xml:"product_architecture" json:"product_architecture"`
	DisplayMode                                string              `xml:"display_mode" json:"display_mode"`
	DisplayAttached                            string              `xml:"display_attached" json:"display_attached"`
	DisplayActive                              string              `xml:"display_active" json:"display_active"`
	PersistenceMode                            string              `xml:"persistence_mode" json:"persistence_mode"`
	AddressingMode                             string              `xml:"addressing_mode" json:"addressing_mode"`
	CurrentMig                                 string              `xml:"mig_mode>current_mig" json:"current_mig"`
	PendingMig                                 string              `xml:"mig_mode>pending_mig" json:"pending_mig"`
	MigDevices                                 []MigDevice         `xml:"mig_devices>mig_device" json:"mig_devices"`
	AccountingMode                             string              `xml:"accounting_mode" json:"accounting_mode"`
	AccountingModeBufferSize                   string              `xml:"accounting_mode_buffer_size" json:"accounting_mode_buffer_size"`
	CurrentDm                                  string              `xml:"driver_model>current_dm" json:"current_dm"`
	PendingDm                                  string              `xml:"driver_model>pending_dm" json:"pending_dm"`
	Serial                                     string              `xml:"serial" json:"serial"`
	UUID                                       string              `xml:"uuid" json:"uuid"`
	Pdi                                        string              `xml:"pdi" json:"pdi"`
	MinorNumber                                string              `xml:"minor_number" json:"minor_number"`
	VbiosVersion                               string              `xml:"vbios_version" json:"vbios_version"`
	MultigpuBoard                              string              `xml:"multigpu_board" json:"multigpu_board"`
	BoardID                                    string              `xml:"board_id" json:"board_id"`
	BoardPartNumber                            string              `xml:"board_part_number" json:"board_part_number"`
	GpuPartNumber                              string              `xml:"gpu_part_number" json:"gpu_part_number"`
	GpuFruPartNumber                           string              `xml:"gpu_fru_part_number" json:"gpu_fru_part_number"`
	ChassisSerialNumber                        string              `xml:"platformInfo>chassis_serial_number" json:"chassis_serial_number"`
	SlotNumber                                 string              `xml:"platformInfo>slot_number" json:"slot_number"`
	TrayIndex                                  string              `xml:"platformInfo>tray_index" json:"tray_index"`
	HostID                                     string              `xml:"platformInfo>host_id" json:"host_id"`
	PeerType                                   string              `xml:"platformInfo>peer_type" json:"peer_type"`
	ModuleID                                   string              `xml:"platformInfo>module_id" json:"module_id"`
	GpuFabricGUID                              string              `xml:"platformInfo>gpu_fabric_guid" json:"gpu_fabric_guid"`
	ImgVersion                                 string              `xml:"inforom_version>img_version" json:"img_version"`
	OemObject                                  string              `xml:"inforom_version>oem_object" json:"oem_object"`
	EccObject                                  string              `xml:"inforom_version>ecc_object" json:"ecc_object"`
	PwrObject                                  string              `xml:"inforom_version>pwr_object" json:"pwr_object"`
	InforomBbxFlushLatestTimestamp             string              `xml:"inforom_bbx_flush>latest_timestamp" json:"inforom_bbx_flush_latest_timestamp"`
	InforomBbxFlushLatestDuration              string              `xml:"inforom_bbx_flush>latest_duration" json:"inforom_bbx_flush_latest_duration"`
	InforomTimeRun                             string              `xml:"inforom_time_data>time_run" json:"inforom_time_run"`
	CurrentGom                                 string              `xml:"gpu_operation_mode>current_gom" json:"current_gom"`
	PendingGom                                 string              `xml:"gpu_operation_mode>pending_gom" json:"pending_gom"`
	C2cMode                                    string              `xml:"c2c_mode" json:"c2c_mode"`
	VirtualizationMode                         string              `xml:"gpu_virtualization_mode>virtualization_mode" json:"virtualization_mode"`
	HostVgpuMode                               string              `xml:"gpu_virtualization_mode>host_vgpu_mode" json:"host_vgpu_mode"`
	VgpuHeterogeneousMode                      string              `xml:"gpu_virtualization_mode>vgpu_heterogeneous_mode" json:"vgpu_heterogeneous_mode"`
	GpuRecoveryAction                          string              `xml:"gpu_recovery_action" json:"gpu_recovery_action"`
	GspFirmwareVersion                         string              `xml:"gsp_firmware_version" json:"gsp_firmware_version"`
	RelaxedOrderingMode                        string              `xml:"ibmnpu>relaxed_ordering_mode" json:"relaxed_ordering_mode"`
	PciBus                                     string              `xml:"pci>pci_bus" json:"pci_bus"`
	PciDevice                                  string              `xml:"pci>pci_device" json:"pci_device"`
	PciDomain                                  string              `xml:"pci>pci_domain" json:"pci_domain"`
	PciBaseClass                               string              `xml:"pci>pci_base_class" json:"pci_base_class"`
	PciSubClass                                string              `xml:"pci>pci_sub_class" json:"pci_sub_class"`
	PciDeviceID                                string              `xml:"pci>pci_device_id" json:"pci_device_id"`
	PciBusID                                   string              `xml:"pci>pci_bus_id" json:"pci_bus_id"`
	PciSubSystemID                             string              `xml:"pci>pci_sub_system_id" json:"pci_sub_system_id"`
	MaxLinkGen                                 string              `xml:"pci>pci_gpu_link_info>pcie_gen>max_link_gen" json:"max_link_gen"`
	CurrentLinkGen                             string              `xml:"pci>pci_gpu_link_info>pcie_gen>current_link_gen" json:"current_link_gen"`
	DeviceCurrentLinkGen                       string              `xml:"pci>pci_gpu_link_info>pcie_gen>device_current_link_gen" json:"device_current_link_gen"`
	MaxDeviceLinkGen                           string              `xml:"pci>pci_gpu_link_info>pcie_gen>max_device_link_gen" json:"max_device_link_gen"`
	MaxHostLinkGen                             string              `xml:"pci>pci_gpu_link_info>pcie_gen>max_host_link_gen" json:"max_host_link_gen"`
	MaxLinkWidth                               string              `xml:"pci>pci_gpu_link_info>link_widths>max_link_width" json:"max_link_width"`
	CurrentLinkWidth                           string              `xml:"pci>pci_gpu_link_info>link_widths>current_link_width" json:"current_link_width"`
	BridgeChipType                             string              `xml:"pci>pci_bridge_chip>bridge_chip_type" json:"bridge_chip_type"`
	BridgeChipFw                               string              `xml:"pci>pci_bridge_chip>bridge_chip_fw" json:"bridge_chip_fw"`
	ReplayCounter                              string              `xml:"pci>replay_counter" json:"replay_counter"`
	ReplayRolloverCounter                      string              `xml:"pci>replay_rollover_counter" json:"replay_rollover_counter"`
	PciTxUtil                                  string              `xml:"pci>tx_util" json:"pci_tx_util"`
	PciRxUtil                                  string              `xml:"pci>rx_util" json:"pci_rx_util"`
	AtomicCapsOutbound                         string              `xml:"pci>atomic_caps_outbound" json:"atomic_caps_outbound"`
	AtomicCapsInbound                          string              `xml:"pci>atomic_caps_inbound" json:"atomic_caps_inbound"`
	FanSpeed                                   string              `xml:"fan_speed" json:"fan_speed"`
	PerformanceState                           string              `xml:"performance_state" json:"performance_state"`
	ClocksEventReasonGpuIdle                   string              `xml:"clocks_event_reasons>clocks_event_reason_gpu_idle" json:"clocks_event_reason_gpu_idle"`
	ClocksEventReasonApplicationsClocksSetting string              `xml:"clocks_event_reasons>clocks_event_reason_applications_clocks_setting" json:"clocks_event_reason_applications_clocks_setting"`
	ClocksEventReasonSwPowerCap                string              `xml:"clocks_event_reasons>clocks_event_reason_sw_power_cap" json:"clocks_event_reason_sw_power_cap"`
	ClocksEventReasonHwSlowdown                string              `xml:"clocks_event_reasons>clocks_event_reason_hw_slowdown" json:"clocks_event_reason_hw_slowdown"`
	ClocksEventReasonHwThermalSlowdown         string              `xml:"clocks_event_reasons>clocks_event_reason_hw_thermal_slowdown" json:"clocks_event_reason_hw_thermal_slowdown"`
	ClocksEventReasonHwPowerBrakeSlowdown      string              `xml:"clocks_event_reasons>clocks_event_reason_hw_power_brake_slowdown" json:"clocks_event_reason_hw_power_brake_slowdown"`
	ClocksEventReasonSyncBoost                 string              `xml:"clocks_event_reasons>clocks_event_reason_sync_boost" json:"clocks_event_reason_sync_boost"`
	ClocksEventReasonSwThermalSlowdown         string              `xml:"clocks_event_reasons>clocks_event_reason_sw_thermal_slowdown" json:"clocks_event_reason_sw_thermal_slowdown"`
	ClocksEventReasonDisplayClocksSetting      string              `xml:"clocks_event_reasons>clocks_event_reason_display_clocks_setting" json:"clocks_event_reason_display_clocks_setting"`
	ClocksEventReasonBoardLimit                string              `xml:"clocks_event_reasons>clocks_event_reason_board_limit" json:"clocks_event_reason_board_limit"`
	ClocksEventReasonReliability               string              `xml:"clocks_event_reasons>clocks_event_reason_reliability" json:"clocks_event_reason_reliability"`
	ClocksEventReasonsCountersSwPowerCap       string              `xml:"clocks_event_reasons_counters>clocks_event_reasons_counters_sw_power_cap" json:"clocks_event_reasons_counters_sw_power_cap"`
	ClocksEventReasonsCountersSyncBoost        string              `xml:"clocks_event_reasons_counters>clocks_event_reasons_counters_sync_boost" json:"clocks_event_reasons_counters_sync_boost"`
	ClocksEventReasonsCountersSwThermSlowdown  string              `xml:"clocks_event_reasons_counters>clocks_event_reasons_counters_sw_therm_slowdown" json:"clocks_event_reasons_counters_sw_therm_slowdown"`
	ClocksEventReasonsCountersHwThermSlowdown  string              `xml:"clocks_event_reasons_counters>clocks_event_reasons_counters_hw_therm_slowdown" json:"clocks_event_reasons_counters_hw_therm_slowdown"`
	ClocksEventReasonsCountersHwPowerBrake     string              `xml:"clocks_event_reasons_counters>clocks_event_reasons_counters_hw_power_brake" json:"clocks_event_reasons_counters_hw_power_brake"`
	SparseOperationMode                        string              `xml:"sparse_operation_mode" json:"sparse_operation_mode"`
	FbMemoryUsageTotal                         string              `xml:"fb_memory_usage>total" json:"fb_memory_usage_total"`
	FbMemoryUsageReserved                      string              `xml:"fb_memory_usage>reserved" json:"fb_memory_usage_reserved"`
	FbMemoryUsageUsed                          string              `xml:"fb_memory_usage>used" json:"fb_memory_usage_used"`
	FbMemoryUsageFree                          string              `xml:"fb_memory_usage>free" json:"fb_memory_usage_free"`
	Bar1Total                                  string              `xml:"bar1_memory_usage>total" json:"bar1_total"`
	Bar1Used                                   string              `xml:"bar1_memory_usage>used" json:"bar1_used"`
	Bar1Free                                   string              `xml:"bar1_memory_usage>free" json:"bar1_free"`
	CcProtectedMemoryTotal                     string              `xml:"cc_protected_memory_usage>total" json:"cc_protected_memory_total"`
	CcProtectedMemoryUsed                      string              `xml:"cc_protected_memory_usage>used" json:"cc_protected_memory_used"`
	CcProtectedMemoryFree                      string              `xml:"cc_protected_memory_usage>free" json:"cc_protected_memory_free"`
	ComputeMode                                string              `xml:"compute_mode" json:"compute_mode"`
	GpuUtil                                    string              `xml:"utilization>gpu_util" json:"gpu_util"`
	MemoryUtil                                 string              `xml:"utilization>memory_util" json:"memory_util"`
	EncoderUtil                                string              `xml:"utilization>encoder_util" json:"encoder_util"`
	DecoderUtil                                string              `xml:"utilization>decoder_util" json:"decoder_util"`
	JpegUtil                                   string              `xml:"utilization>jpeg_util" json:"jpeg_util"`
	OfaUtil                                    string              `xml:"utilization>ofa_util" json:"ofa_util"`
	EncoderStats                               SessionStats        `xml:"encoder_stats" json:"encoder_stats"`
	FbcStats                                   SessionStats        `xml:"fbc_stats" json:"fbc_stats"`
	CurrentDramEncryption                      string              `xml:"dram_encryption_mode>current_dram_encryption" json:"current_dram_encryption"`
	PendingDramEncryption                      string              `xml:"dram_encryption_mode>pending_dram_encryption" json:"pending_dram_encryption"`
	CurrentEcc                                 string              `xml:"ecc_mode>current_ecc" json:"current_ecc"`
	PendingEcc                                 string              `xml:"ecc_mode>pending_ecc" json:"pending_ecc"`
	EccErrors                                  EccErrors           `xml:"ecc_errors" json:"ecc_errors"`
	RetiredPages                               RetiredPages        `xml:"retired_pages" json:"retired_pages"`
	RemappedRows                               string              `xml:"remapped_rows" json:"remapped_rows"`
	RemappedBanks                              string              `xml:"remapped_banks" json:"remapped_banks"`
	GpuTemp                                    string              `xml:"temperature>gpu_temp" json:"gpu_temp"`
	GpuTempMax                                 string              `xml:"temperature>gpu_temp_max" json:"gpu_temp_max"`
	GpuTempTlimit                              string              `xml:"temperature>gpu_temp_tlimit" json:"gpu_temp_tlimit"`
	GpuTempMaxTlimitThreshold                  string              `xml:"temperature>gpu_temp_max_tlimit_threshold" json:"gpu_temp_max_tlimit_threshold"`
	GpuTempSlowTlimitThreshold                 string              `xml:"temperature>gpu_temp_slow_tlimit_threshold" json:"gpu_temp_slow_tlimit_threshold"`
	GpuTempMaxGpuTlimitThreshold               string              `xml:"temperature>gpu_temp_max_gpu_tlimit_threshold" json:"gpu_temp_max_gpu_tlimit_threshold"`
	GpuTargetTemp                              string              `xml:"temperature>gpu_target_temperature" json:"gpu_target_temperature"`
	MemoryTemp                                 string              `xml:"temperature>memory_temp" json:"memory_temp"`
	GpuTempMaxMemTlimitThreshold               string              `xml:"temperature>gpu_temp_max_mem_tlimit_threshold" json:"gpu_temp_max_mem_tlimit_threshold"`
	GpuTargetTempMin                           string              `xml:"supported_gpu_target_temp>gpu_target_temp_min" json:"gpu_target_temp_min"`
	GpuTargetTempMax                           string              `xml:"supported_gpu_target_temp>gpu_target_temp_max" json:"gpu_target_temp_max"`
	PowerState                                 string              `xml:"gpu_power_readings>power_state" json:"power_state"`
	AveragePowerDraw                           string              `xml:"gpu_power_readings>average_power_draw" json:"average_power_draw"`
	InstantPowerDraw                           string              `xml:"gpu_power_readings>instant_power_draw" json:"instant_power_draw"`
	GpuCeilingPowerLimit                       PowerLimits         `xml:"gpu_power_readings>gpu_ceiling_power_limit" json:"gpu_ceiling_power_limit"`
	CurrentBasePower                           string              `xml:"gpu_power_readings>gpu_base_power>current_base_power" json:"current_base_power"`
	RequestedBasePower                         string              `xml:"gpu_power_readings>gpu_base_power>requested_base_power" json:"requested_base_power"`
	DefaultBasePower                           string              `xml:"gpu_power_readings>gpu_base_power>default_base_power" json:"default_base_power"`
	AdaptiveTgpAdjustedBasePower               string              `xml:"gpu_power_readings>gpu_base_power>adaptive_tgp_adjusted_base_power" json:"adaptive_tgp_adjusted_base_power"`
	CurrentPowerLimit                          string              `xml:"gpu_power_readings>current_power_limit" json:"current_power_limit"`
	RequestedPowerLimit                        string              `xml:"gpu_power_readings>requested_power_limit" json:"requested_power_limit"`
	DefaultPowerLimit                          string              `xml:"gpu_power_readings>default_power_limit" json:"default_power_limit"`
	MinPowerLimit                              string              `xml:"gpu_power_readings>min_power_limit" json:"min_power_limit"`
	MaxPowerLimit                              string              `xml:"gpu_power_readings>max_power_limit" json:"max_power_limit"`
	GpuMemoryPowerReadings                     PowerDraw           `xml:"gpu_memory_power_readings" json:"gpu_memory_power_readings"`
	ModulePowerReadings                        ModulePowerReadings `xml:"module_power_readings" json:"module_power_readings"`
	PowerSmoothing                             string              `xml:"power_smoothing" json:"power_smoothing"`
	PowerProfileRequestedProfiles              string              `xml:"power_profiles>power_profile_requested_profiles" json:"power_profile_requested_profiles"`
	PowerProfileEnforcedProfiles               string              `xml:"power_profiles>power_profile_enforced_profiles" json:"power_profile_enforced_profiles"`
	EdppMultiplier                             string              `xml:"edpp_multiplier" json:"edpp_multiplier"`
	GraphicsClock                              string              `xml:"clocks>graphics_clock" json:"graphics_clock"`
	SmClock                                    string              `xml:"clocks>sm_clock" json:"sm_clock"`
	MemClock                                   string              `xml:"clocks>mem_clock" json:"mem_clock"`
	VideoClock                                 string              `xml:"clocks>video_clock" json:"video_clock"`
	ApplicationsGraphicsClock                  string              `xml:"applications_clocks>graphics_clock" json:"applications_graphics_clock"`
	ApplicationsMemClock                       string              `xml:"applications_clocks>mem_clock" json:"applications_mem_clock"`
	DefaultApplicationsGraphicsClock           string              `xml:"default_applications_clocks>graphics_clock" json:"default_applications_graphics_clock"`
	DefaultApplicationsMemClock                string              `xml:"default_applications_clocks>mem_clock" json:"default_applications_mem_clock"`
	DeferredMemClock                           string              `xml:"deferred_clocks>mem_clock" json:"deferred_mem_clock"`
	GraphicsClockMax                           string              `xml:"max_clocks>graphics_clock" json:"graphics_clock_max"`
	SmClockMax                                 string              `xml:"max_clocks>sm_clock" json:"sm_clock_max"`
	MemClockMax                                string              `xml:"max_clocks>mem_clock" json:"mem_clock_max"`
	VideoClockMax                              string              `xml:"max_clocks>video_clock" json:"video_clock_max"`
	MaxCustomerBoostGraphicsClock              string              `xml:"max_customer_boost_clocks>graphics_clock" json:"max_customer_boost_graphics_clock"`
	AutoBoost                                  string              `xml:"clock_policy>auto_boost" json:"auto_boost"`
	AutoBoostDefault                           string              `xml:"clock_policy>auto_boost_default" json:"auto_boost_default"`
	DlaCurEngineClock                          string              `xml:"dla_clocks>dla_cur_engine_clock" json:"dla_cur_engine_clock"`
	DlaMaxEngineClock                          string              `xml:"dla_clocks>dla_max_engine_clock" json:"dla_max_engine_clock"`
	SupportedClocks                            []SupportedMemClock `xml:"supported_clocks>supported_mem_clock" json:"supported_clocks"`
	Fabric                                     Fabric              `xml:"fabric" json:"fabric"`
	Processes                                  []ProcessInfo       `xml:"processes>process_info" json:"processes"`
	AccountedProcesses                         []AccountedProcess  `xml:"accounted_processes>accounted_process_info" json:"accounted_processes"`
	Egm                                        string              `xml:"capabilities>egm" json:"egm"`
}

type MigDevice struct {
	Index                 string `xml:"index" json:"index"`
	GpuInstanceID         string `xml:"gpu_instance_id" json:"gpu_instance_id"`
	ComputeInstanceID     string `xml:"compute_instance_id" json:"compute_instance_id"`
	MultiprocessorCount   string `xml:"device_attributes>shared>multiprocessor_count" json:"multiprocessor_count"`
	CopyEngineCount       string `xml:"device_attributes>shared>copy_engine_count" json:"copy_engine_count"`
	EncoderCount          string `xml:"device_attributes>shared>encoder_count" json:"encoder_count"`
	DecoderCount          string `xml:"device_attributes>shared>decoder_count" json:"decoder_count"`
	OfaCount              string `xml:"device_attributes>shared>ofa_count" json:"ofa_count"`
	JpgCount              string `xml:"device_attributes>shared>jpg_count" json:"jpg_count"`
	SramUncorrectable     string `xml:"ecc_error_count>volatile_count>sram_uncorrectable" json:"sram_uncorrectable"`
	FbMemoryUsageTotal    string `xml:"fb_memory_usage>total" json:"fb_memory_usage_total"`
	FbMemoryUsageReserved string `xml:"fb_memory_usage>reserved" json:"fb_memory_usage_reserved"`
	FbMemoryUsageUsed     string `xml:"fb_memory_usage>used" json:"fb_memory_usage_used"`
	FbMemoryUsageFree     string `xml:"fb_memory_usage>free" json:"fb_memory_usage_free"`
	Bar1Total             string `xml:"bar1_memory_usage>total" json:"bar1_total"`
	Bar1Used              string `xml:"bar1_memory_usage>used" json:"bar1_used"`
	Bar1Free              string `xml:"bar1_memory_usage>free" json:"bar1_free"`
}

type SessionStats struct {
	SessionCount   string `xml:"session_count" json:"session_count"`
	AverageFps     string `xml:"average_fps" json:"average_fps"`
	AverageLatency string `xml:"average_latency" json:"average_latency"`
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
	PowerState          string `xml:"power_state" json:"power_state"`
	AveragePowerDraw    string `xml:"average_power_draw" json:"average_power_draw"`
	InstantPowerDraw    string `xml:"instant_power_draw" json:"instant_power_draw"`
	CurrentPowerLimit   string `xml:"current_power_limit" json:"current_power_limit"`
	RequestedPowerLimit string `xml:"requested_power_limit" json:"requested_power_limit"`
	DefaultPowerLimit   string `xml:"default_power_limit" json:"default_power_limit"`
	MinPowerLimit       string `xml:"min_power_limit" json:"min_power_limit"`
	MaxPowerLimit       string `xml:"max_power_limit" json:"max_power_limit"`
}

type SupportedMemClock struct {
	Value                   string   `xml:"value" json:"value"`
	SupportedGraphicsClocks []string `xml:"supported_graphics_clock" json:"supported_graphics_clocks"`
}

type Fabric struct {
	State                   string `xml:"state" json:"state"`
	Status                  string `xml:"status" json:"status"`
	CliqueID                string `xml:"cliqueId" json:"clique_id"`
	ClusterUUID             string `xml:"clusterUuid" json:"cluster_uuid"`
	HealthSummary           string `xml:"health>summary" json:"health_summary"`
	HealthBandwidth         string `xml:"health>bandwidth" json:"health_bandwidth"`
	RouteRecoveryInProgress string `xml:"health>route_recovery_in_progress" json:"route_recovery_in_progress"`
	RouteUnhealthy          string `xml:"health>route_unhealthy" json:"route_unhealthy"`
	AccessTimeoutRecovery   string `xml:"health>access_timeout_recovery" json:"access_timeout_recovery"`
	IncorrectConfiguration  string `xml:"health>incorrect_configuration" json:"incorrect_configuration"`
	PartitionAssigned       string `xml:"health>partition_assigned" json:"partition_assigned"`
	GfmState                string `xml:"health>gfm_state" json:"gfm_state"`
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

type NvidiaSmiLog struct {
	Timestamp      string `xml:"timestamp" json:"timestamp"`
	DriverVersion  string `xml:"driver_version" json:"driver_version"`
	CudaVersion    string `xml:"cuda_version" json:"cuda_version"`
	KmdVersion     string `xml:"kmd_version" json:"kmd_version"`
	CudaUmdVersion string `xml:"cuda_umd_version" json:"cuda_umd_version"`
	AttachedGpus   string `xml:"attached_gpus" json:"attached_gpus"`
	GPUS           []GPU  `xml:"gpu" json:"gpus"`
}
