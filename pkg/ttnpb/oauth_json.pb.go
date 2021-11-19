// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.1.0
// - protoc             v3.9.1
// source: lorawan-stack/api/oauth.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the OAuthClientAuthorization message to JSON.
func (x *OAuthClientAuthorization) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.UserIds)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("client_ids")
		// NOTE: ClientIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.ClientIds)
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
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the OAuthClientAuthorization message from JSON.
func (x *OAuthClientAuthorization) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIds = v
		case "client_ids", "clientIds":
			s.AddField("client_ids")
			// NOTE: ClientIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ClientIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ClientIds = v
		case "rights":
			s.AddField("rights")
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
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
		}
	})
}

// MarshalProtoJSON marshals the OAuthClientAuthorizations message to JSON.
func (x *OAuthClientAuthorizations) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Authorizations) > 0 || s.HasField("authorizations") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("authorizations")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Authorizations {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("authorizations"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the OAuthClientAuthorizations message from JSON.
func (x *OAuthClientAuthorizations) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "authorizations":
			s.AddField("authorizations")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Authorizations = append(x.Authorizations, nil)
					return
				}
				v := &OAuthClientAuthorization{}
				v.UnmarshalProtoJSON(s.WithField("authorizations", false))
				if s.Err() != nil {
					return
				}
				x.Authorizations = append(x.Authorizations, v)
			})
		}
	})
}

// MarshalProtoJSON marshals the OAuthAuthorizationCode message to JSON.
func (x *OAuthAuthorizationCode) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.UserIds)
	}
	if x.UserSessionId != "" || s.HasField("user_session_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_session_id")
		s.WriteString(x.UserSessionId)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("client_ids")
		// NOTE: ClientIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.ClientIds)
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
	if x.Code != "" || s.HasField("code") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("code")
		s.WriteString(x.Code)
	}
	if x.RedirectUri != "" || s.HasField("redirect_uri") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("redirect_uri")
		s.WriteString(x.RedirectUri)
	}
	if x.State != "" || s.HasField("state") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state")
		s.WriteString(x.State)
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

// UnmarshalProtoJSON unmarshals the OAuthAuthorizationCode message from JSON.
func (x *OAuthAuthorizationCode) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIds = v
		case "user_session_id", "userSessionId":
			s.AddField("user_session_id")
			x.UserSessionId = s.ReadString()
		case "client_ids", "clientIds":
			s.AddField("client_ids")
			// NOTE: ClientIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ClientIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ClientIds = v
		case "rights":
			s.AddField("rights")
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "code":
			s.AddField("code")
			x.Code = s.ReadString()
		case "redirect_uri", "redirectUri":
			s.AddField("redirect_uri")
			x.RedirectUri = s.ReadString()
		case "state":
			s.AddField("state")
			x.State = s.ReadString()
		case "created_at", "createdAt":
			s.AddField("created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
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

// MarshalProtoJSON marshals the OAuthAccessToken message to JSON.
func (x *OAuthAccessToken) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.UserIds)
	}
	if x.UserSessionId != "" || s.HasField("user_session_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_session_id")
		s.WriteString(x.UserSessionId)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("client_ids")
		// NOTE: ClientIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.ClientIds)
	}
	if x.Id != "" || s.HasField("id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id")
		s.WriteString(x.Id)
	}
	if x.AccessToken != "" || s.HasField("access_token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("access_token")
		s.WriteString(x.AccessToken)
	}
	if x.RefreshToken != "" || s.HasField("refresh_token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("refresh_token")
		s.WriteString(x.RefreshToken)
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
	if x.CreatedAt != nil || s.HasField("created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		if x.CreatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.CreatedAt)
		}
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

// UnmarshalProtoJSON unmarshals the OAuthAccessToken message from JSON.
func (x *OAuthAccessToken) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIds = v
		case "user_session_id", "userSessionId":
			s.AddField("user_session_id")
			x.UserSessionId = s.ReadString()
		case "client_ids", "clientIds":
			s.AddField("client_ids")
			// NOTE: ClientIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ClientIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.ClientIds = v
		case "id":
			s.AddField("id")
			x.Id = s.ReadString()
		case "access_token", "accessToken":
			s.AddField("access_token")
			x.AccessToken = s.ReadString()
		case "refresh_token", "refreshToken":
			s.AddField("refresh_token")
			x.RefreshToken = s.ReadString()
		case "rights":
			s.AddField("rights")
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "created_at", "createdAt":
			s.AddField("created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
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

// MarshalProtoJSON marshals the OAuthAccessTokens message to JSON.
func (x *OAuthAccessTokens) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Tokens) > 0 || s.HasField("tokens") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tokens")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Tokens {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("tokens"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the OAuthAccessTokens message from JSON.
func (x *OAuthAccessTokens) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "tokens":
			s.AddField("tokens")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Tokens = append(x.Tokens, nil)
					return
				}
				v := &OAuthAccessToken{}
				v.UnmarshalProtoJSON(s.WithField("tokens", false))
				if s.Err() != nil {
					return
				}
				x.Tokens = append(x.Tokens, v)
			})
		}
	})
}
