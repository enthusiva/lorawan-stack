// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.1.0
// - protoc              v4.22.2
// source: ttn/lorawan/v3/packetbrokeragent.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// AddSetFlagsForPacketBrokerRegisterRequest adds flags to select fields in PacketBrokerRegisterRequest.
func AddSetFlagsForPacketBrokerRegisterRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("listed", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the PacketBrokerRegisterRequest message from flags.
func (m *PacketBrokerRegisterRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("listed", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Listed = &wrapperspb.BoolValue{Value: val}
		paths = append(paths, flagsplugin.Prefix("listed", prefix))
	}
	return paths, nil
}

// AddSetFlagsForPacketBrokerGatewayVisibility adds flags to select fields in PacketBrokerGatewayVisibility.
func AddSetFlagsForPacketBrokerGatewayVisibility(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("location", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("antenna-placement", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("antenna-count", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("fine-timestamps", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("contact-info", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("status", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("frequency-plan", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("packet-rates", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the PacketBrokerGatewayVisibility message from flags.
func (m *PacketBrokerGatewayVisibility) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("location", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Location = val
		paths = append(paths, flagsplugin.Prefix("location", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("antenna_placement", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AntennaPlacement = val
		paths = append(paths, flagsplugin.Prefix("antenna_placement", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("antenna_count", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.AntennaCount = val
		paths = append(paths, flagsplugin.Prefix("antenna_count", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("fine_timestamps", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FineTimestamps = val
		paths = append(paths, flagsplugin.Prefix("fine_timestamps", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("contact_info", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.ContactInfo = val
		paths = append(paths, flagsplugin.Prefix("contact_info", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("status", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Status = val
		paths = append(paths, flagsplugin.Prefix("status", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("frequency_plan", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FrequencyPlan = val
		paths = append(paths, flagsplugin.Prefix("frequency_plan", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("packet_rates", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.PacketRates = val
		paths = append(paths, flagsplugin.Prefix("packet_rates", prefix))
	}
	return paths, nil
}
