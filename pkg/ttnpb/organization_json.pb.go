// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.1.0
// - protoc             v3.9.1
// source: lorawan-stack/api/organization.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the Organization message to JSON.
func (x *Organization) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil || s.HasField("ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Ids)
	}
	if x.CreatedAt != nil || s.HasField("created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		if x.CreatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.CreatedAt)
		}
	}
	if x.UpdatedAt != nil || s.HasField("updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("updated_at")
		if x.UpdatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.UpdatedAt)
		}
	}
	if x.DeletedAt != nil || s.HasField("deleted_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted_at")
		if x.DeletedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.DeletedAt)
		}
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if x.Description != "" || s.HasField("description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description")
		s.WriteString(x.Description)
	}
	if x.Attributes != nil || s.HasField("attributes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.Attributes {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the Organization message from JSON.
func (x *Organization) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Ids = &v
		case "created_at", "createdAt":
			s.AddField("created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
		case "updated_at", "updatedAt":
			s.AddField("updated_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.UpdatedAt = v
		case "deleted_at", "deletedAt":
			s.AddField("deleted_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.DeletedAt = v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "description":
			s.AddField("description")
			x.Description = s.ReadString()
		case "attributes":
			s.AddField("attributes")
			x.Attributes = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.Attributes[key] = s.ReadString()
			})
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		}
	})
}

// MarshalProtoJSON marshals the Organizations message to JSON.
func (x *Organizations) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Organizations) > 0 || s.HasField("organizations") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organizations")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Organizations {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("organizations"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the Organizations message from JSON.
func (x *Organizations) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organizations":
			s.AddField("organizations")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Organizations = append(x.Organizations, nil)
					return
				}
				v := &Organization{}
				v.UnmarshalProtoJSON(s.WithField("organizations", false))
				if s.Err() != nil {
					return
				}
				x.Organizations = append(x.Organizations, v)
			})
		}
	})
}

// MarshalProtoJSON marshals the CreateOrganizationRequest message to JSON.
func (x *CreateOrganizationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Organization != nil || s.HasField("organization") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization")
		x.Organization.MarshalProtoJSON(s.WithField("organization"))
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		// NOTE: OrganizationOrUserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Collaborator)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the CreateOrganizationRequest message from JSON.
func (x *CreateOrganizationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization":
			if !s.ReadNil() {
				x.Organization = &Organization{}
				x.Organization.UnmarshalProtoJSON(s.WithField("organization", true))
			}
		case "collaborator":
			s.AddField("collaborator")
			// NOTE: OrganizationOrUserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationOrUserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Collaborator = &v
		}
	})
}

// MarshalProtoJSON marshals the UpdateOrganizationRequest message to JSON.
func (x *UpdateOrganizationRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Organization != nil || s.HasField("organization") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization")
		x.Organization.MarshalProtoJSON(s.WithField("organization"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the UpdateOrganizationRequest message from JSON.
func (x *UpdateOrganizationRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization":
			if !s.ReadNil() {
				x.Organization = &Organization{}
				x.Organization.UnmarshalProtoJSON(s.WithField("organization", true))
			}
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// MarshalProtoJSON marshals the CreateOrganizationAPIKeyRequest message to JSON.
func (x *CreateOrganizationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.OrganizationIds != nil || s.HasField("organization_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization_ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.OrganizationIds)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the CreateOrganizationAPIKeyRequest message from JSON.
func (x *CreateOrganizationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.OrganizationIds = &v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "rights":
			s.AddField("rights")
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}

// MarshalProtoJSON marshals the UpdateOrganizationAPIKeyRequest message to JSON.
func (x *UpdateOrganizationAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.OrganizationIds != nil || s.HasField("organization_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization_ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.OrganizationIds)
	}
	if x.ApiKey != nil || s.HasField("api_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		x.ApiKey.MarshalProtoJSON(s.WithField("api_key"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the UpdateOrganizationAPIKeyRequest message from JSON.
func (x *UpdateOrganizationAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.OrganizationIds = &v
		case "api_key", "apiKey":
			if !s.ReadNil() {
				x.ApiKey = &APIKey{}
				x.ApiKey.UnmarshalProtoJSON(s.WithField("api_key", true))
			}
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// MarshalProtoJSON marshals the SetOrganizationCollaboratorRequest message to JSON.
func (x *SetOrganizationCollaboratorRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.OrganizationIds != nil || s.HasField("organization_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("organization_ids")
		// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.OrganizationIds)
	}
	if x.Collaborator != nil || s.HasField("collaborator") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("collaborator")
		x.Collaborator.MarshalProtoJSON(s.WithField("collaborator"))
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the SetOrganizationCollaboratorRequest message from JSON.
func (x *SetOrganizationCollaboratorRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.OrganizationIds = &v
		case "collaborator":
			if !s.ReadNil() {
				x.Collaborator = &Collaborator{}
				x.Collaborator.UnmarshalProtoJSON(s.WithField("collaborator", true))
			}
		}
	})
}
