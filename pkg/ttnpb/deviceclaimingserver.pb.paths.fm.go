// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var ClaimEndDeviceRequestFieldPathsNested = []string{
	"source_device",
	"source_device.authenticated_identifiers",
	"source_device.authenticated_identifiers.authentication_code",
	"source_device.authenticated_identifiers.dev_eui",
	"source_device.authenticated_identifiers.join_eui",
	"source_device.qr_code",
	"target_application_ids",
	"target_application_ids.application_id",
	"target_device_id",
}

var ClaimEndDeviceRequestFieldPathsTopLevel = []string{
	"source_device",
	"target_application_ids",
	"target_device_id",
}
var AuthorizeApplicationRequestFieldPathsNested = []string{
	"api_key",
	"application_ids",
	"application_ids.application_id",
}

var AuthorizeApplicationRequestFieldPathsTopLevel = []string{
	"api_key",
	"application_ids",
}
var GetInfoByJoinEUIRequestFieldPathsNested = []string{
	"join_eui",
}

var GetInfoByJoinEUIRequestFieldPathsTopLevel = []string{
	"join_eui",
}
var GetInfoByJoinEUIResponseFieldPathsNested = []string{
	"join_eui",
	"supports_claiming",
}

var GetInfoByJoinEUIResponseFieldPathsTopLevel = []string{
	"join_eui",
	"supports_claiming",
}
var GetInfoByJoinEUIsRequestFieldPathsNested = []string{
	"requests",
}

var GetInfoByJoinEUIsRequestFieldPathsTopLevel = []string{
	"requests",
}
var GetInfoByJoinEUIsResponseFieldPathsNested = []string{
	"infos",
}

var GetInfoByJoinEUIsResponseFieldPathsTopLevel = []string{
	"infos",
}
var GetClaimStatusResponseFieldPathsNested = []string{
	"end_device_ids",
	"end_device_ids.application_ids",
	"end_device_ids.application_ids.application_id",
	"end_device_ids.dev_addr",
	"end_device_ids.dev_eui",
	"end_device_ids.device_id",
	"end_device_ids.join_eui",
	"home_net_id",
	"home_ns_id",
	"vendor_specific",
	"vendor_specific.data",
	"vendor_specific.organization_unique_identifier",
}

var GetClaimStatusResponseFieldPathsTopLevel = []string{
	"end_device_ids",
	"home_net_id",
	"home_ns_id",
	"vendor_specific",
}
var BatchUnclaimEndDevicesRequestFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"device_ids",
}

var BatchUnclaimEndDevicesRequestFieldPathsTopLevel = []string{
	"application_ids",
	"device_ids",
}
var BatchUnclaimEndDevicesResponseFieldPathsNested = []string{
	"application_ids",
	"application_ids.application_id",
	"failed",
}

var BatchUnclaimEndDevicesResponseFieldPathsTopLevel = []string{
	"application_ids",
	"failed",
}
var CUPSRedirectionFieldPathsNested = []string{
	"current_gateway_key",
	"gateway_credentials",
	"gateway_credentials.auth_token",
	"gateway_credentials.client_tls",
	"gateway_credentials.client_tls.cert",
	"gateway_credentials.client_tls.key",
	"target_cups_trust",
	"target_cups_uri",
}

var CUPSRedirectionFieldPathsTopLevel = []string{
	"current_gateway_key",
	"gateway_credentials",
	"target_cups_trust",
	"target_cups_uri",
}
var ClaimGatewayRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"cups_redirection",
	"cups_redirection.current_gateway_key",
	"cups_redirection.gateway_credentials",
	"cups_redirection.gateway_credentials.auth_token",
	"cups_redirection.gateway_credentials.client_tls",
	"cups_redirection.gateway_credentials.client_tls.cert",
	"cups_redirection.gateway_credentials.client_tls.key",
	"cups_redirection.target_cups_trust",
	"cups_redirection.target_cups_uri",
	"source_gateway",
	"source_gateway.authenticated_identifiers",
	"source_gateway.authenticated_identifiers.authentication_code",
	"source_gateway.authenticated_identifiers.gateway_eui",
	"source_gateway.qr_code",
	"target_frequency_plan_id",
	"target_gateway_id",
	"target_gateway_server_address",
}

var ClaimGatewayRequestFieldPathsTopLevel = []string{
	"collaborator",
	"cups_redirection",
	"source_gateway",
	"target_frequency_plan_id",
	"target_gateway_id",
	"target_gateway_server_address",
}
var AuthorizeGatewayRequestFieldPathsNested = []string{
	"api_key",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
}

var AuthorizeGatewayRequestFieldPathsTopLevel = []string{
	"api_key",
	"gateway_ids",
}
var GetInfoByGatewayEUIRequestFieldPathsNested = []string{
	"eui",
}

var GetInfoByGatewayEUIRequestFieldPathsTopLevel = []string{
	"eui",
}
var GetInfoByGatewayEUIResponseFieldPathsNested = []string{
	"eui",
	"supports_claiming",
}

var GetInfoByGatewayEUIResponseFieldPathsTopLevel = []string{
	"eui",
	"supports_claiming",
}
var ClaimEndDeviceRequest_AuthenticatedIdentifiersFieldPathsNested = []string{
	"authentication_code",
	"dev_eui",
	"join_eui",
}

var ClaimEndDeviceRequest_AuthenticatedIdentifiersFieldPathsTopLevel = []string{
	"authentication_code",
	"dev_eui",
	"join_eui",
}
var GetClaimStatusResponse_VendorSpecificFieldPathsNested = []string{
	"data",
	"organization_unique_identifier",
}

var GetClaimStatusResponse_VendorSpecificFieldPathsTopLevel = []string{
	"data",
	"organization_unique_identifier",
}
var CUPSRedirection_ClientTLSFieldPathsNested = []string{
	"cert",
	"key",
}

var CUPSRedirection_ClientTLSFieldPathsTopLevel = []string{
	"cert",
	"key",
}
var ClaimGatewayRequest_AuthenticatedIdentifiersFieldPathsNested = []string{
	"authentication_code",
	"gateway_eui",
}

var ClaimGatewayRequest_AuthenticatedIdentifiersFieldPathsTopLevel = []string{
	"authentication_code",
	"gateway_eui",
}
