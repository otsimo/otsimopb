// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!


export interface OtsimoServices {
  environment?: string;
  issuer?: string;
  isProduction?: boolean;
  useTls?: boolean;
  registryGrpc?: string;
  listenerGrpc?: string;
  watchGrpc?: string;
  catalogGrpc?: string;
  contentGrpc?: string;
  dashboardGrpc?: string;
  apiGrpc?: string;
/**
GameContent is registry service http url
*/
  gameContent?: string;
  accounts?: string;
  analyticsGrpc?: string;
  services?: { [key: string]: string };
		  gameStorageProviders?: { [key: string]: string };
		  configs?: { [key: string]: string };
		  inMaintenance?: boolean;
}

export const OtsimoServices_environment = "environment";
export const OtsimoServices_issuer = "issuer";
export const OtsimoServices_isProduction = "is_production";
export const OtsimoServices_useTls = "use_tls";
export const OtsimoServices_registryGrpc = "registry_grpc";
export const OtsimoServices_listenerGrpc = "listener_grpc";
export const OtsimoServices_watchGrpc = "watch_grpc";
export const OtsimoServices_catalogGrpc = "catalog_grpc";
export const OtsimoServices_contentGrpc = "content_grpc";
export const OtsimoServices_dashboardGrpc = "dashboard_grpc";
export const OtsimoServices_apiGrpc = "api_grpc";
export const OtsimoServices_gameContent = "game_content";
export const OtsimoServices_accounts = "accounts";
export const OtsimoServices_analyticsGrpc = "analytics_grpc";
export const OtsimoServices_services = "services";
export const OtsimoServices_gameStorageProviders = "game_storage_providers";
export const OtsimoServices_configs = "configs";
export const OtsimoServices_inMaintenance = "in_maintenance";
export interface DiscoveryRequest {
  environment?: string;
  sdkVersion?: string;
  osName?: string;
  countryCode?: string;
  appBundleId?: string;
  appBundleVersion?: string;
}

export const DiscoveryRequest_environment = "environment";
export const DiscoveryRequest_sdkVersion = "sdk_version";
export const DiscoveryRequest_osName = "os_name";
export const DiscoveryRequest_countryCode = "country_code";
export const DiscoveryRequest_appBundleId = "app_bundle_id";
export const DiscoveryRequest_appBundleVersion = "app_bundle_version";
