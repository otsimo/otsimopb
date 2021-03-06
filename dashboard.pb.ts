// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_datasetmodels from "./datasetmodels.pb";

export type CardDecorationSize =  "SMALL"  | "MEDIUM"  | "LARGE" ;
/**
Small is 1x1 block on iphone
*/
export const CardDecorationSize_SMALL: CardDecorationSize = "SMALL";
/**
Medium is 2x1 block on iphone
*/
export const CardDecorationSize_MEDIUM: CardDecorationSize = "MEDIUM";
/**
Large is 2x2 block on iphone
*/
export const CardDecorationSize_LARGE: CardDecorationSize = "LARGE";

export const ALL_CardDecorationSize_VALUES: CardDecorationSize[] = [CardDecorationSize_SMALL,CardDecorationSize_MEDIUM,CardDecorationSize_LARGE];

export type CardDecorationBackgroundStyle =  "EMPTY"  | "IMAGE"  | "CHART_SILHOUETTE" ;
export const CardDecorationBackgroundStyle_EMPTY: CardDecorationBackgroundStyle = "EMPTY";
export const CardDecorationBackgroundStyle_IMAGE: CardDecorationBackgroundStyle = "IMAGE";
export const CardDecorationBackgroundStyle_CHART_SILHOUETTE: CardDecorationBackgroundStyle = "CHART_SILHOUETTE";

export const ALL_CardDecorationBackgroundStyle_VALUES: CardDecorationBackgroundStyle[] = [CardDecorationBackgroundStyle_EMPTY,CardDecorationBackgroundStyle_IMAGE,CardDecorationBackgroundStyle_CHART_SILHOUETTE];

export type ChartType =  "LINE"  | "BAR"  | "PIE"  | "SCATTER"  | "BUBLE"  | "RADAR"  | "GEO"  | "TIMELINE" ;
export const ChartType_LINE: ChartType = "LINE";
export const ChartType_BAR: ChartType = "BAR";
export const ChartType_PIE: ChartType = "PIE";
export const ChartType_SCATTER: ChartType = "SCATTER";
export const ChartType_BUBLE: ChartType = "BUBLE";
export const ChartType_RADAR: ChartType = "RADAR";
export const ChartType_GEO: ChartType = "GEO";
export const ChartType_TIMELINE: ChartType = "TIMELINE";

export const ALL_ChartType_VALUES: ChartType[] = [ChartType_LINE,ChartType_BAR,ChartType_PIE,ChartType_SCATTER,ChartType_BUBLE,ChartType_RADAR,ChartType_GEO,ChartType_TIMELINE];

export interface DashboardItems {
/**
ProfileId
*/
  profileId?: string;
/**
ChildId
*/
  childId?: string;
/**
Created At
*/
  createdAt?: string|number;
  items?: Card[];
}

export const DashboardItems_profileId = "profile_id";
export const DashboardItems_childId = "child_id";
export const DashboardItems_createdAt = "created_at";
export const DashboardItems_items = "items";
export interface DashboardGetRequest {
  profileId?: string;
  childId?: string;
  appVersion?: string;
  language?: string;
  countryCode?: string;
  lastTimeDataFetched?: string|number;
}

export const DashboardGetRequest_profileId = "profile_id";
export const DashboardGetRequest_childId = "child_id";
export const DashboardGetRequest_appVersion = "app_version";
export const DashboardGetRequest_language = "language";
export const DashboardGetRequest_countryCode = "country_code";
export const DashboardGetRequest_lastTimeDataFetched = "last_time_data_fetched";
export interface CardDecoration {
  size?: CardDecorationSize;
  backgroundStyle?: CardDecorationBackgroundStyle;
  imageUrl?: string;
  leftIcon?: string;
  rightIcon?: string;
}

export const CardDecoration_size = "size";
export const CardDecoration_backgroundStyle = "background_style";
export const CardDecoration_imageUrl = "image_url";
export const CardDecoration_leftIcon = "left_icon";
export const CardDecoration_rightIcon = "right_icon";
export interface CardEmpty {
}

export interface CardWebpage {
  url?: string;
}

export const CardWebpage_url = "url";
export interface CardApplink {
  applink?: string;
}

export const CardApplink_applink = "applink";
export interface CardAnalysis {
  data?: apipb_datasetmodels.DataSet;
  chartType?: ChartType;
}

export const CardAnalysis_data = "data";
export const CardAnalysis_chartType = "chart_type";
export interface Card {
  id?: string;
  text?: string;
  expiresAt?: string|number;
  createdAt?: string|number;
  decoration?: CardDecoration;
/**
Score is between 0-500
*/
  providerScore?: number;
/**
ProviderWeight is between 0-2
*/
  providerWeight?: number;
  providerName?: string;
  language?: string;
  empty?: CardEmpty;
  webpage?: CardWebpage;
  applink?: CardApplink;
  analysis?: CardAnalysis;
/**
Title for newer systems
*/
  title?: string;
/**
Subtitle for newer systems
*/
  subtitle?: string;
/**
Labels of the card
*/
  labels?: { [key: string]: string };
		}

export const Card_id = "id";
export const Card_text = "text";
export const Card_expiresAt = "expires_at";
export const Card_createdAt = "created_at";
export const Card_decoration = "decoration";
export const Card_providerScore = "provider_score";
export const Card_providerWeight = "provider_weight";
export const Card_providerName = "provider_name";
export const Card_language = "language";
export const Card_empty = "empty";
export const Card_webpage = "webpage";
export const Card_applink = "applink";
export const Card_analysis = "analysis";
export const Card_title = "title";
export const Card_subtitle = "subtitle";
export const Card_labels = "labels";
