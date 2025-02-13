// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.5.1
// - protoc             v4.22.2
// source: ttn/lorawan/v3/configuration_services.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the GetPhyVersionsResponse_VersionInfo message to JSON.
func (x *GetPhyVersionsResponse_VersionInfo) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.BandId != "" || s.HasField("band_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("band_id")
		s.WriteString(x.BandId)
	}
	if len(x.PhyVersions) > 0 || s.HasField("phy_versions") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("phy_versions")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.PhyVersions {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetPhyVersionsResponse_VersionInfo to JSON.
func (x *GetPhyVersionsResponse_VersionInfo) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetPhyVersionsResponse_VersionInfo message from JSON.
func (x *GetPhyVersionsResponse_VersionInfo) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "band_id", "bandId":
			s.AddField("band_id")
			x.BandId = s.ReadString()
		case "phy_versions", "phyVersions":
			s.AddField("phy_versions")
			if s.ReadNil() {
				x.PhyVersions = nil
				return
			}
			s.ReadArray(func() {
				var v PHYVersion
				v.UnmarshalProtoJSON(s)
				x.PhyVersions = append(x.PhyVersions, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the GetPhyVersionsResponse_VersionInfo from JSON.
func (x *GetPhyVersionsResponse_VersionInfo) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GetPhyVersionsResponse message to JSON.
func (x *GetPhyVersionsResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.VersionInfo) > 0 || s.HasField("version_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("version_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.VersionInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("version_info"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetPhyVersionsResponse to JSON.
func (x *GetPhyVersionsResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GetPhyVersionsResponse message from JSON.
func (x *GetPhyVersionsResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "version_info", "versionInfo":
			s.AddField("version_info")
			if s.ReadNil() {
				x.VersionInfo = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.VersionInfo = append(x.VersionInfo, nil)
					return
				}
				v := &GetPhyVersionsResponse_VersionInfo{}
				v.UnmarshalProtoJSON(s.WithField("version_info", false))
				if s.Err() != nil {
					return
				}
				x.VersionInfo = append(x.VersionInfo, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the GetPhyVersionsResponse from JSON.
func (x *GetPhyVersionsResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ListBandsRequest message to JSON.
func (x *ListBandsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.BandId != "" || s.HasField("band_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("band_id")
		s.WriteString(x.BandId)
	}
	if x.PhyVersion != 0 || s.HasField("phy_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("phy_version")
		x.PhyVersion.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ListBandsRequest to JSON.
func (x *ListBandsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ListBandsRequest message from JSON.
func (x *ListBandsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "band_id", "bandId":
			s.AddField("band_id")
			x.BandId = s.ReadString()
		case "phy_version", "phyVersion":
			s.AddField("phy_version")
			x.PhyVersion.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the ListBandsRequest from JSON.
func (x *ListBandsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the BandDescription_Beacon message to JSON.
func (x *BandDescription_Beacon) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.DataRateIndex != 0 || s.HasField("data_rate_index") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data_rate_index")
		x.DataRateIndex.MarshalProtoJSON(s)
	}
	if x.CodingRate != "" || s.HasField("coding_rate") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("coding_rate")
		s.WriteString(x.CodingRate)
	}
	if len(x.Frequencies) > 0 || s.HasField("frequencies") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequencies")
		s.WriteUint64Array(x.Frequencies)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the BandDescription_Beacon to JSON.
func (x *BandDescription_Beacon) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the BandDescription_Beacon message from JSON.
func (x *BandDescription_Beacon) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "data_rate_index", "dataRateIndex":
			s.AddField("data_rate_index")
			x.DataRateIndex.UnmarshalProtoJSON(s)
		case "coding_rate", "codingRate":
			s.AddField("coding_rate")
			x.CodingRate = s.ReadString()
		case "frequencies":
			s.AddField("frequencies")
			if s.ReadNil() {
				x.Frequencies = nil
				return
			}
			x.Frequencies = s.ReadUint64Array()
		}
	})
}

// UnmarshalJSON unmarshals the BandDescription_Beacon from JSON.
func (x *BandDescription_Beacon) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the BandDescription_Channel message to JSON.
func (x *BandDescription_Channel) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Frequency != 0 || s.HasField("frequency") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency")
		s.WriteUint64(x.Frequency)
	}
	if x.MinDataRate != 0 || s.HasField("min_data_rate") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("min_data_rate")
		x.MinDataRate.MarshalProtoJSON(s)
	}
	if x.MaxDataRate != 0 || s.HasField("max_data_rate") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("max_data_rate")
		x.MaxDataRate.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the BandDescription_Channel to JSON.
func (x *BandDescription_Channel) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the BandDescription_Channel message from JSON.
func (x *BandDescription_Channel) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "frequency":
			s.AddField("frequency")
			x.Frequency = s.ReadUint64()
		case "min_data_rate", "minDataRate":
			s.AddField("min_data_rate")
			x.MinDataRate.UnmarshalProtoJSON(s)
		case "max_data_rate", "maxDataRate":
			s.AddField("max_data_rate")
			x.MaxDataRate.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the BandDescription_Channel from JSON.
func (x *BandDescription_Channel) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the BandDescription_Rx2Parameters message to JSON.
func (x *BandDescription_Rx2Parameters) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.DataRateIndex != 0 || s.HasField("data_rate_index") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data_rate_index")
		x.DataRateIndex.MarshalProtoJSON(s)
	}
	if x.Frequency != 0 || s.HasField("frequency") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency")
		s.WriteUint64(x.Frequency)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the BandDescription_Rx2Parameters to JSON.
func (x *BandDescription_Rx2Parameters) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the BandDescription_Rx2Parameters message from JSON.
func (x *BandDescription_Rx2Parameters) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "data_rate_index", "dataRateIndex":
			s.AddField("data_rate_index")
			x.DataRateIndex.UnmarshalProtoJSON(s)
		case "frequency":
			s.AddField("frequency")
			x.Frequency = s.ReadUint64()
		}
	})
}

// UnmarshalJSON unmarshals the BandDescription_Rx2Parameters from JSON.
func (x *BandDescription_Rx2Parameters) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the BandDescription_RelayParameters_RelayWORChannel message to JSON.
func (x *BandDescription_RelayParameters_RelayWORChannel) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Frequency != 0 || s.HasField("frequency") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency")
		s.WriteUint64(x.Frequency)
	}
	if x.AckFrequency != 0 || s.HasField("ack_frequency") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ack_frequency")
		s.WriteUint64(x.AckFrequency)
	}
	if x.DataRateIndex != 0 || s.HasField("data_rate_index") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data_rate_index")
		x.DataRateIndex.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the BandDescription_RelayParameters_RelayWORChannel to JSON.
func (x *BandDescription_RelayParameters_RelayWORChannel) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the BandDescription_RelayParameters_RelayWORChannel message from JSON.
func (x *BandDescription_RelayParameters_RelayWORChannel) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "frequency":
			s.AddField("frequency")
			x.Frequency = s.ReadUint64()
		case "ack_frequency", "ackFrequency":
			s.AddField("ack_frequency")
			x.AckFrequency = s.ReadUint64()
		case "data_rate_index", "dataRateIndex":
			s.AddField("data_rate_index")
			x.DataRateIndex.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the BandDescription_RelayParameters_RelayWORChannel from JSON.
func (x *BandDescription_RelayParameters_RelayWORChannel) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the BandDescription_RelayParameters message to JSON.
func (x *BandDescription_RelayParameters) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.WorChannels) > 0 || s.HasField("wor_channels") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("wor_channels")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.WorChannels {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("wor_channels"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the BandDescription_RelayParameters to JSON.
func (x *BandDescription_RelayParameters) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the BandDescription_RelayParameters message from JSON.
func (x *BandDescription_RelayParameters) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "wor_channels", "worChannels":
			s.AddField("wor_channels")
			if s.ReadNil() {
				x.WorChannels = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.WorChannels = append(x.WorChannels, nil)
					return
				}
				v := &BandDescription_RelayParameters_RelayWORChannel{}
				v.UnmarshalProtoJSON(s.WithField("wor_channels", false))
				if s.Err() != nil {
					return
				}
				x.WorChannels = append(x.WorChannels, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the BandDescription_RelayParameters from JSON.
func (x *BandDescription_RelayParameters) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the BandDescription message to JSON.
func (x *BandDescription) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Id != "" || s.HasField("id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id")
		s.WriteString(x.Id)
	}
	if x.Beacon != nil || s.HasField("beacon") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("beacon")
		x.Beacon.MarshalProtoJSON(s.WithField("beacon"))
	}
	if len(x.PingSlotFrequencies) > 0 || s.HasField("ping_slot_frequencies") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ping_slot_frequencies")
		s.WriteUint64Array(x.PingSlotFrequencies)
	}
	if x.MaxUplinkChannels != 0 || s.HasField("max_uplink_channels") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("max_uplink_channels")
		s.WriteUint32(x.MaxUplinkChannels)
	}
	if len(x.UplinkChannels) > 0 || s.HasField("uplink_channels") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("uplink_channels")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.UplinkChannels {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("uplink_channels"))
		}
		s.WriteArrayEnd()
	}
	if x.MaxDownlinkChannels != 0 || s.HasField("max_downlink_channels") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("max_downlink_channels")
		s.WriteUint32(x.MaxDownlinkChannels)
	}
	if len(x.DownlinkChannels) > 0 || s.HasField("downlink_channels") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("downlink_channels")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.DownlinkChannels {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("downlink_channels"))
		}
		s.WriteArrayEnd()
	}
	if len(x.SubBands) > 0 || s.HasField("sub_bands") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("sub_bands")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.SubBands {
			s.WriteMoreIf(&wroteElement)
			// NOTE: BandDescription_SubBandParameters does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, element)
		}
		s.WriteArrayEnd()
	}
	if x.DataRates != nil || s.HasField("data_rates") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("data_rates")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.DataRates {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectUint32Field(k)
			// NOTE: BandDescription_BandDataRate does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, v)
		}
		s.WriteObjectEnd()
	}
	if x.FreqMultiplier != 0 || s.HasField("freq_multiplier") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("freq_multiplier")
		s.WriteUint64(x.FreqMultiplier)
	}
	if x.ImplementsCfList || s.HasField("implements_cf_list") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("implements_cf_list")
		s.WriteBool(x.ImplementsCfList)
	}
	if x.CfListType != 0 || s.HasField("cf_list_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("cf_list_type")
		x.CfListType.MarshalProtoJSON(s)
	}
	if x.ReceiveDelay_1 != nil || s.HasField("receive_delay_1") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("receive_delay_1")
		if x.ReceiveDelay_1 == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.ReceiveDelay_1)
		}
	}
	if x.ReceiveDelay_2 != nil || s.HasField("receive_delay_2") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("receive_delay_2")
		if x.ReceiveDelay_2 == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.ReceiveDelay_2)
		}
	}
	if x.JoinAcceptDelay_1 != nil || s.HasField("join_accept_delay_1") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_accept_delay_1")
		if x.JoinAcceptDelay_1 == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.JoinAcceptDelay_1)
		}
	}
	if x.JoinAcceptDelay_2 != nil || s.HasField("join_accept_delay_2") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_accept_delay_2")
		if x.JoinAcceptDelay_2 == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.JoinAcceptDelay_2)
		}
	}
	if x.MaxFcntGap != 0 || s.HasField("max_fcnt_gap") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("max_fcnt_gap")
		s.WriteUint64(x.MaxFcntGap)
	}
	if x.SupportsDynamicAdr || s.HasField("supports_dynamic_adr") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("supports_dynamic_adr")
		s.WriteBool(x.SupportsDynamicAdr)
	}
	if x.AdrAckLimit != 0 || s.HasField("adr_ack_limit") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("adr_ack_limit")
		x.AdrAckLimit.MarshalProtoJSON(s)
	}
	if x.MinRetransmitTimeout != nil || s.HasField("min_retransmit_timeout") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("min_retransmit_timeout")
		if x.MinRetransmitTimeout == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.MinRetransmitTimeout)
		}
	}
	if x.MaxRetransmitTimeout != nil || s.HasField("max_retransmit_timeout") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("max_retransmit_timeout")
		if x.MaxRetransmitTimeout == nil {
			s.WriteNil()
		} else {
			golang.MarshalDuration(s, x.MaxRetransmitTimeout)
		}
	}
	if len(x.TxOffset) > 0 || s.HasField("tx_offset") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tx_offset")
		s.WriteFloat32Array(x.TxOffset)
	}
	if x.MaxAdrDataRateIndex != 0 || s.HasField("max_adr_data_rate_index") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("max_adr_data_rate_index")
		x.MaxAdrDataRateIndex.MarshalProtoJSON(s)
	}
	if x.TxParamSetupReqSupport || s.HasField("tx_param_setup_req_support") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tx_param_setup_req_support")
		s.WriteBool(x.TxParamSetupReqSupport)
	}
	if x.DefaultMaxEirp != 0 || s.HasField("default_max_eirp") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("default_max_eirp")
		s.WriteFloat32(x.DefaultMaxEirp)
	}
	if x.DefaultRx2Parameters != nil || s.HasField("default_rx2_parameters") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("default_rx2_parameters")
		x.DefaultRx2Parameters.MarshalProtoJSON(s.WithField("default_rx2_parameters"))
	}
	if x.BootDwellTime != nil || s.HasField("boot_dwell_time") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("boot_dwell_time")
		// NOTE: BandDescription_DwellTime does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.BootDwellTime)
	}
	if x.Relay != nil || s.HasField("relay") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("relay")
		x.Relay.MarshalProtoJSON(s.WithField("relay"))
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the BandDescription to JSON.
func (x *BandDescription) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the BandDescription message from JSON.
func (x *BandDescription) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "id":
			s.AddField("id")
			x.Id = s.ReadString()
		case "beacon":
			if s.ReadNil() {
				x.Beacon = nil
				return
			}
			x.Beacon = &BandDescription_Beacon{}
			x.Beacon.UnmarshalProtoJSON(s.WithField("beacon", true))
		case "ping_slot_frequencies", "pingSlotFrequencies":
			s.AddField("ping_slot_frequencies")
			if s.ReadNil() {
				x.PingSlotFrequencies = nil
				return
			}
			x.PingSlotFrequencies = s.ReadUint64Array()
		case "max_uplink_channels", "maxUplinkChannels":
			s.AddField("max_uplink_channels")
			x.MaxUplinkChannels = s.ReadUint32()
		case "uplink_channels", "uplinkChannels":
			s.AddField("uplink_channels")
			if s.ReadNil() {
				x.UplinkChannels = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.UplinkChannels = append(x.UplinkChannels, nil)
					return
				}
				v := &BandDescription_Channel{}
				v.UnmarshalProtoJSON(s.WithField("uplink_channels", false))
				if s.Err() != nil {
					return
				}
				x.UplinkChannels = append(x.UplinkChannels, v)
			})
		case "max_downlink_channels", "maxDownlinkChannels":
			s.AddField("max_downlink_channels")
			x.MaxDownlinkChannels = s.ReadUint32()
		case "downlink_channels", "downlinkChannels":
			s.AddField("downlink_channels")
			if s.ReadNil() {
				x.DownlinkChannels = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.DownlinkChannels = append(x.DownlinkChannels, nil)
					return
				}
				v := &BandDescription_Channel{}
				v.UnmarshalProtoJSON(s.WithField("downlink_channels", false))
				if s.Err() != nil {
					return
				}
				x.DownlinkChannels = append(x.DownlinkChannels, v)
			})
		case "sub_bands", "subBands":
			s.AddField("sub_bands")
			if s.ReadNil() {
				x.SubBands = nil
				return
			}
			s.ReadArray(func() {
				// NOTE: BandDescription_SubBandParameters does not seem to implement UnmarshalProtoJSON.
				var v BandDescription_SubBandParameters
				golang.UnmarshalMessage(s, &v)
				x.SubBands = append(x.SubBands, &v)
			})
		case "data_rates", "dataRates":
			s.AddField("data_rates")
			if s.ReadNil() {
				x.DataRates = nil
				return
			}
			x.DataRates = make(map[uint32]*BandDescription_BandDataRate)
			s.ReadUint32Map(func(key uint32) {
				// NOTE: BandDescription_BandDataRate does not seem to implement UnmarshalProtoJSON.
				var v BandDescription_BandDataRate
				golang.UnmarshalMessage(s, &v)
				x.DataRates[key] = &v
			})
		case "freq_multiplier", "freqMultiplier":
			s.AddField("freq_multiplier")
			x.FreqMultiplier = s.ReadUint64()
		case "implements_cf_list", "implementsCfList":
			s.AddField("implements_cf_list")
			x.ImplementsCfList = s.ReadBool()
		case "cf_list_type", "cfListType":
			s.AddField("cf_list_type")
			x.CfListType.UnmarshalProtoJSON(s)
		case "receive_delay_1", "receiveDelay1":
			s.AddField("receive_delay_1")
			if s.ReadNil() {
				x.ReceiveDelay_1 = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.ReceiveDelay_1 = v
		case "receive_delay_2", "receiveDelay2":
			s.AddField("receive_delay_2")
			if s.ReadNil() {
				x.ReceiveDelay_2 = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.ReceiveDelay_2 = v
		case "join_accept_delay_1", "joinAcceptDelay1":
			s.AddField("join_accept_delay_1")
			if s.ReadNil() {
				x.JoinAcceptDelay_1 = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.JoinAcceptDelay_1 = v
		case "join_accept_delay_2", "joinAcceptDelay2":
			s.AddField("join_accept_delay_2")
			if s.ReadNil() {
				x.JoinAcceptDelay_2 = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.JoinAcceptDelay_2 = v
		case "max_fcnt_gap", "maxFcntGap":
			s.AddField("max_fcnt_gap")
			x.MaxFcntGap = s.ReadUint64()
		case "supports_dynamic_adr", "supportsDynamicAdr":
			s.AddField("supports_dynamic_adr")
			x.SupportsDynamicAdr = s.ReadBool()
		case "adr_ack_limit", "adrAckLimit":
			s.AddField("adr_ack_limit")
			x.AdrAckLimit.UnmarshalProtoJSON(s)
		case "min_retransmit_timeout", "minRetransmitTimeout":
			s.AddField("min_retransmit_timeout")
			if s.ReadNil() {
				x.MinRetransmitTimeout = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.MinRetransmitTimeout = v
		case "max_retransmit_timeout", "maxRetransmitTimeout":
			s.AddField("max_retransmit_timeout")
			if s.ReadNil() {
				x.MaxRetransmitTimeout = nil
				return
			}
			v := golang.UnmarshalDuration(s)
			if s.Err() != nil {
				return
			}
			x.MaxRetransmitTimeout = v
		case "tx_offset", "txOffset":
			s.AddField("tx_offset")
			if s.ReadNil() {
				x.TxOffset = nil
				return
			}
			x.TxOffset = s.ReadFloat32Array()
		case "max_adr_data_rate_index", "maxAdrDataRateIndex":
			s.AddField("max_adr_data_rate_index")
			x.MaxAdrDataRateIndex.UnmarshalProtoJSON(s)
		case "tx_param_setup_req_support", "txParamSetupReqSupport":
			s.AddField("tx_param_setup_req_support")
			x.TxParamSetupReqSupport = s.ReadBool()
		case "default_max_eirp", "defaultMaxEirp":
			s.AddField("default_max_eirp")
			x.DefaultMaxEirp = s.ReadFloat32()
		case "default_rx2_parameters", "defaultRx2Parameters":
			if s.ReadNil() {
				x.DefaultRx2Parameters = nil
				return
			}
			x.DefaultRx2Parameters = &BandDescription_Rx2Parameters{}
			x.DefaultRx2Parameters.UnmarshalProtoJSON(s.WithField("default_rx2_parameters", true))
		case "boot_dwell_time", "bootDwellTime":
			s.AddField("boot_dwell_time")
			if s.ReadNil() {
				x.BootDwellTime = nil
				return
			}
			// NOTE: BandDescription_DwellTime does not seem to implement UnmarshalProtoJSON.
			var v BandDescription_DwellTime
			golang.UnmarshalMessage(s, &v)
			x.BootDwellTime = &v
		case "relay":
			if s.ReadNil() {
				x.Relay = nil
				return
			}
			x.Relay = &BandDescription_RelayParameters{}
			x.Relay.UnmarshalProtoJSON(s.WithField("relay", true))
		}
	})
}

// UnmarshalJSON unmarshals the BandDescription from JSON.
func (x *BandDescription) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
