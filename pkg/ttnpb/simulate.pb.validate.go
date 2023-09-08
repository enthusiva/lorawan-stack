// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// ValidateFields checks the field values on SimulateMetadataParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SimulateMetadataParams) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SimulateMetadataParamsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "rssi":
			// no validation rules for Rssi
		case "snr":
			// no validation rules for Snr
		case "timestamp":
			// no validation rules for Timestamp
		case "time":

			if v, ok := interface{}(m.GetTime()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return SimulateMetadataParamsValidationError{
						field:  "time",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "lorawan_version":
			// no validation rules for LorawanVersion
		case "lorawan_phy_version":
			// no validation rules for LorawanPhyVersion
		case "band_id":
			// no validation rules for BandId
		case "frequency":
			// no validation rules for Frequency
		case "channel_index":
			// no validation rules for ChannelIndex
		case "bandwidth":
			// no validation rules for Bandwidth
		case "spreading_factor":
			// no validation rules for SpreadingFactor
		case "data_rate_index":
			// no validation rules for DataRateIndex
		default:
			return SimulateMetadataParamsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SimulateMetadataParamsValidationError is the validation error returned by
// SimulateMetadataParams.ValidateFields if the designated constraints aren't met.
type SimulateMetadataParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SimulateMetadataParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SimulateMetadataParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SimulateMetadataParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SimulateMetadataParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SimulateMetadataParamsValidationError) ErrorName() string {
	return "SimulateMetadataParamsValidationError"
}

// Error satisfies the builtin error interface
func (e SimulateMetadataParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSimulateMetadataParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SimulateMetadataParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SimulateMetadataParamsValidationError{}

// ValidateFields checks the field values on SimulateJoinRequestParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SimulateJoinRequestParams) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SimulateJoinRequestParamsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "join_eui":

			if len(m.GetJoinEui()) > 0 {

				if len(m.GetJoinEui()) != 8 {
					return SimulateJoinRequestParamsValidationError{
						field:  "join_eui",
						reason: "value length must be 8 bytes",
					}
				}

			}

		case "dev_eui":

			if len(m.GetDevEui()) > 0 {

				if len(m.GetDevEui()) != 8 {
					return SimulateJoinRequestParamsValidationError{
						field:  "dev_eui",
						reason: "value length must be 8 bytes",
					}
				}

			}

		case "dev_nonce":

			if len(m.GetDevNonce()) > 0 {

				if len(m.GetDevNonce()) != 2 {
					return SimulateJoinRequestParamsValidationError{
						field:  "dev_nonce",
						reason: "value length must be 2 bytes",
					}
				}

			}

		case "app_key":

			if len(m.GetAppKey()) > 0 {

				if len(m.GetAppKey()) != 16 {
					return SimulateJoinRequestParamsValidationError{
						field:  "app_key",
						reason: "value length must be 16 bytes",
					}
				}

			}

		case "nwk_key":

			if len(m.GetNwkKey()) > 0 {

				if len(m.GetNwkKey()) != 16 {
					return SimulateJoinRequestParamsValidationError{
						field:  "nwk_key",
						reason: "value length must be 16 bytes",
					}
				}

			}

		default:
			return SimulateJoinRequestParamsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SimulateJoinRequestParamsValidationError is the validation error returned by
// SimulateJoinRequestParams.ValidateFields if the designated constraints
// aren't met.
type SimulateJoinRequestParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SimulateJoinRequestParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SimulateJoinRequestParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SimulateJoinRequestParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SimulateJoinRequestParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SimulateJoinRequestParamsValidationError) ErrorName() string {
	return "SimulateJoinRequestParamsValidationError"
}

// Error satisfies the builtin error interface
func (e SimulateJoinRequestParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSimulateJoinRequestParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SimulateJoinRequestParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SimulateJoinRequestParamsValidationError{}

// ValidateFields checks the field values on SimulateDataUplinkParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SimulateDataUplinkParams) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SimulateDataUplinkParamsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "dev_addr":

			if len(m.GetDevAddr()) > 0 {

				if len(m.GetDevAddr()) != 4 {
					return SimulateDataUplinkParamsValidationError{
						field:  "dev_addr",
						reason: "value length must be 4 bytes",
					}
				}

			}

		case "f_nwk_s_int_key":

			if len(m.GetFNwkSIntKey()) > 0 {

				if len(m.GetFNwkSIntKey()) != 16 {
					return SimulateDataUplinkParamsValidationError{
						field:  "f_nwk_s_int_key",
						reason: "value length must be 16 bytes",
					}
				}

			}

		case "s_nwk_s_int_key":

			if len(m.GetSNwkSIntKey()) > 0 {

				if len(m.GetSNwkSIntKey()) != 16 {
					return SimulateDataUplinkParamsValidationError{
						field:  "s_nwk_s_int_key",
						reason: "value length must be 16 bytes",
					}
				}

			}

		case "nwk_s_enc_key":

			if len(m.GetNwkSEncKey()) > 0 {

				if len(m.GetNwkSEncKey()) != 16 {
					return SimulateDataUplinkParamsValidationError{
						field:  "nwk_s_enc_key",
						reason: "value length must be 16 bytes",
					}
				}

			}

		case "app_s_key":

			if len(m.GetAppSKey()) > 0 {

				if len(m.GetAppSKey()) != 16 {
					return SimulateDataUplinkParamsValidationError{
						field:  "app_s_key",
						reason: "value length must be 16 bytes",
					}
				}

			}

		case "adr":
			// no validation rules for Adr
		case "adr_ack_req":
			// no validation rules for AdrAckReq
		case "confirmed":
			// no validation rules for Confirmed
		case "ack":
			// no validation rules for Ack
		case "f_cnt":
			// no validation rules for FCnt
		case "f_port":
			// no validation rules for FPort
		case "frm_payload":
			// no validation rules for FrmPayload
		case "conf_f_cnt":
			// no validation rules for ConfFCnt
		case "tx_dr_idx":
			// no validation rules for TxDrIdx
		case "tx_ch_idx":
			// no validation rules for TxChIdx
		case "f_opts":

			if len(m.GetFOpts()) > 15 {
				return SimulateDataUplinkParamsValidationError{
					field:  "f_opts",
					reason: "value length must be at most 15 bytes",
				}
			}

		default:
			return SimulateDataUplinkParamsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SimulateDataUplinkParamsValidationError is the validation error returned by
// SimulateDataUplinkParams.ValidateFields if the designated constraints
// aren't met.
type SimulateDataUplinkParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SimulateDataUplinkParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SimulateDataUplinkParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SimulateDataUplinkParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SimulateDataUplinkParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SimulateDataUplinkParamsValidationError) ErrorName() string {
	return "SimulateDataUplinkParamsValidationError"
}

// Error satisfies the builtin error interface
func (e SimulateDataUplinkParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSimulateDataUplinkParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SimulateDataUplinkParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SimulateDataUplinkParamsValidationError{}
