// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.6
// - protoc              v3.9.1
// source: lorawan-stack/api/rights.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForAPIKey adds flags to select fields in APIKey.
func AddSelectFlagsForAPIKey(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("id", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("id", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("key", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("key", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("name", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("rights", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("rights", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("created-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("created-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("updated-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("updated-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("expires-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("expires-at", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forAPIKey message from select flags.
func PathsFromSelectFlagsForAPIKey(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("id", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("id", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("key", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("key", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("rights", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("rights", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("created_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("created_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("updated_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("updated_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("expires_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("expires_at", prefix))
	}
	return paths, nil
}

// AddSetFlagsForAPIKey adds flags to select fields in APIKey.
func AddSetFlagsForAPIKey(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("key", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("rights", prefix), flagsplugin.EnumValueDesc(Right_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("created-at", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("updated-at", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("expires-at", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the APIKey message from flags.
func (m *APIKey) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Id = val
		paths = append(paths, flagsplugin.Prefix("id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("key", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Key = val
		paths = append(paths, flagsplugin.Prefix("key", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Name = val
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("rights", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Rights = make([]Right, len(val))
		for i, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, Right_value)
			if err != nil {
				return nil, err
			}
			m.Rights[i] = Right(enumValue)
		}
		paths = append(paths, flagsplugin.Prefix("rights", prefix))
	}
	if val, changed, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("created_at", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.CreatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("created_at", prefix))
	}
	if val, changed, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("updated_at", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.UpdatedAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("updated_at", prefix))
	}
	if val, changed, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("expires_at", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.ExpiresAt = gogo.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("expires_at", prefix))
	}
	return paths, nil
}
