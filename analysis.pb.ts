// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_datasetmodels from "./datasetmodels.pb";

export type QuerySortSortOrder =  "ASC"  | "DSC" ;
export const QuerySortSortOrder_ASC: QuerySortSortOrder = "ASC";
export const QuerySortSortOrder_DSC: QuerySortSortOrder = "DSC";

export const ALL_QuerySortSortOrder_VALUES: QuerySortSortOrder[] = [QuerySortSortOrder_ASC,QuerySortSortOrder_DSC];

export type AggregationAccumulator =  "NONE"  | "COUNT"  | "SUM"  | "MAX"  | "MIN"  | "AVG"  | "STD_SAMP"  | "STD_POP" ;
export const AggregationAccumulator_NONE: AggregationAccumulator = "NONE";
export const AggregationAccumulator_COUNT: AggregationAccumulator = "COUNT";
export const AggregationAccumulator_SUM: AggregationAccumulator = "SUM";
export const AggregationAccumulator_MAX: AggregationAccumulator = "MAX";
export const AggregationAccumulator_MIN: AggregationAccumulator = "MIN";
export const AggregationAccumulator_AVG: AggregationAccumulator = "AVG";
export const AggregationAccumulator_STD_SAMP: AggregationAccumulator = "STD_SAMP";
export const AggregationAccumulator_STD_POP: AggregationAccumulator = "STD_POP";

export const ALL_AggregationAccumulator_VALUES: AggregationAccumulator[] = [AggregationAccumulator_NONE,AggregationAccumulator_COUNT,AggregationAccumulator_SUM,AggregationAccumulator_MAX,AggregationAccumulator_MIN,AggregationAccumulator_AVG,AggregationAccumulator_STD_SAMP,AggregationAccumulator_STD_POP];

export type QueryGroupGroupType =  "Date"  | "Datetime"  | "TimeOfDay"  | "Discrete"  | "ContinuesInterval" ;
/**
Date gives rows each day of given timeRange
*/
export const QueryGroupGroupType_Date: QueryGroupGroupType = "Date";
/**
Datetime gives rows on given interval
*/
export const QueryGroupGroupType_Datetime: QueryGroupGroupType = "Datetime";
/**
TimeofDay gives rows on time of day. minutes and seconds on interval
value determines interval. ex: if minutes and seconds are false then rows
will be 0,1,2,3,4 ex: if minutes is true then rows will be
00:00,00:01,00:02,... when seconds is true than minutes is always true
*/
export const QueryGroupGroupType_TimeOfDay: QueryGroupGroupType = "TimeOfDay";
/**
Discrete should be used for String values
*/
export const QueryGroupGroupType_Discrete: QueryGroupGroupType = "Discrete";
/**
ContinuesInterval is for number values
*/
export const QueryGroupGroupType_ContinuesInterval: QueryGroupGroupType = "ContinuesInterval";

export const ALL_QueryGroupGroupType_VALUES: QueryGroupGroupType[] = [QueryGroupGroupType_Date,QueryGroupGroupType_Datetime,QueryGroupGroupType_TimeOfDay,QueryGroupGroupType_Discrete,QueryGroupGroupType_ContinuesInterval];

export type ActiveUsersRequestType =  "MONTLY"  | "DAILY"  | "NEW"  | "TOTAL" ;
export const ActiveUsersRequestType_MONTLY: ActiveUsersRequestType = "MONTLY";
export const ActiveUsersRequestType_DAILY: ActiveUsersRequestType = "DAILY";
export const ActiveUsersRequestType_NEW: ActiveUsersRequestType = "NEW";
export const ActiveUsersRequestType_TOTAL: ActiveUsersRequestType = "TOTAL";

export const ALL_ActiveUsersRequestType_VALUES: ActiveUsersRequestType[] = [ActiveUsersRequestType_MONTLY,ActiveUsersRequestType_DAILY,ActiveUsersRequestType_NEW,ActiveUsersRequestType_TOTAL];

export type RetentionRequestType =  "ONE"  | "SEVEN"  | "THIRTY" ;
export const RetentionRequestType_ONE: RetentionRequestType = "ONE";
export const RetentionRequestType_SEVEN: RetentionRequestType = "SEVEN";
export const RetentionRequestType_THIRTY: RetentionRequestType = "THIRTY";

export const ALL_RetentionRequestType_VALUES: RetentionRequestType[] = [RetentionRequestType_ONE,RetentionRequestType_SEVEN,RetentionRequestType_THIRTY];

export type GameInfoResponseFieldType =  "UNKNOWN"  | "STRING"  | "INTEGER"  | "FLOAT"  | "BOOL" ;
export const GameInfoResponseFieldType_UNKNOWN: GameInfoResponseFieldType = "UNKNOWN";
export const GameInfoResponseFieldType_STRING: GameInfoResponseFieldType = "STRING";
export const GameInfoResponseFieldType_INTEGER: GameInfoResponseFieldType = "INTEGER";
export const GameInfoResponseFieldType_FLOAT: GameInfoResponseFieldType = "FLOAT";
export const GameInfoResponseFieldType_BOOL: GameInfoResponseFieldType = "BOOL";

export const ALL_GameInfoResponseFieldType_VALUES: GameInfoResponseFieldType[] = [GameInfoResponseFieldType_UNKNOWN,GameInfoResponseFieldType_STRING,GameInfoResponseFieldType_INTEGER,GameInfoResponseFieldType_FLOAT,GameInfoResponseFieldType_BOOL];

export interface TimeRange {
/**
From is the unix seconds time
*/
  from?: string|number;
/**
To is the unix seconds time
*/
  to?: string|number;
}

export const TimeRange_from = "from";
export const TimeRange_to = "to";
export interface ChildAndProfileIds {
/**
ChildId
*/
  childId?: string;
/**
ProfileId
*/
  profileId?: string;
}

export const ChildAndProfileIds_childId = "child_id";
export const ChildAndProfileIds_profileId = "profile_id";
export interface ChildAndTimeRange {
/**
ChildId
*/
  childId?: string;
/**
ProfileId
*/
  profileId?: string;
/**
Range is the time range
*/
  range?: TimeRange;
}

export const ChildAndTimeRange_childId = "child_id";
export const ChildAndTimeRange_profileId = "profile_id";
export const ChildAndTimeRange_range = "range";
export interface GameWithVersions {
/**
GameId
*/
  gameId?: string;
/**
Versions
*/
  versions?: string[];
}

export const GameWithVersions_gameId = "game_id";
export const GameWithVersions_versions = "versions";
export interface PlayedGamesList {
  games?: GameWithVersions[];
/**
ChildId
*/
  childId?: string;
/**
Range is the time range
*/
  range?: TimeRange;
}

export const PlayedGamesList_games = "games";
export const PlayedGamesList_childId = "child_id";
export const PlayedGamesList_range = "range";
export interface QuerySort {
  fieldName?: string;
  order?: QuerySortSortOrder;
}

export const QuerySort_fieldName = "field_name";
export const QuerySort_order = "order";
export interface Aggregation {
  fieldName?: string;
  outputField?: string;
  accumulator?: AggregationAccumulator;
}

export const Aggregation_fieldName = "field_name";
export const Aggregation_outputField = "output_field";
export const Aggregation_accumulator = "accumulator";
export interface GroupInterval {
/**
For ContinuesInterval
*/
  int?: number;
/**
For ContinuesInterval
*/
  real?: number;
/**
For Datetime
*/
  days?: number;
/**
For Datetime
*/
  hours?: number;
/**
For TimeOfDay
*/
  minutes?: boolean;
/**
For TimeOfDay
*/
  seconds?: boolean;
}

export const GroupInterval_int = "int";
export const GroupInterval_real = "real";
export const GroupInterval_days = "days";
export const GroupInterval_hours = "hours";
export const GroupInterval_minutes = "minutes";
export const GroupInterval_seconds = "seconds";
export interface QueryGroup {
  fieldName?: string;
  type?: QueryGroupGroupType;
/**
Interval is optional for Date and Discrete type
*/
  interval?: GroupInterval;
  outputField?: string;
/**
Timezone Offset in seconds
*/
  timezoneSeconds?: number;
}

export const QueryGroup_fieldName = "field_name";
export const QueryGroup_type = "type";
export const QueryGroup_interval = "interval";
export const QueryGroup_outputField = "output_field";
export const QueryGroup_timezoneSeconds = "timezone_seconds";
export interface Query {
  events?: string[];
  range?: TimeRange;
  sort?: QuerySort[];
  limit?: number;
  offset?: number;
  groupBy?: QueryGroup;
  aggregations?: Aggregation[];
  rawQueries?: string[];
}

export const Query_events = "events";
export const Query_range = "range";
export const Query_sort = "sort";
export const Query_limit = "limit";
export const Query_offset = "offset";
export const Query_groupBy = "group_by";
export const Query_aggregations = "aggregations";
export const Query_rawQueries = "raw_queries";
export interface AnalyzeRequest {
/**
ChildId
*/
  childId?: string;
/**
ProfileId
*/
  profileId?: string;
/**
UseAppData changes data source
*/
  useAppData?: boolean;
/**
Query is calculation query
*/
  query?: Query;
/**
Games are the compute this request on
*/
  games?: GameWithVersions[];
}

export const AnalyzeRequest_childId = "child_id";
export const AnalyzeRequest_profileId = "profile_id";
export const AnalyzeRequest_useAppData = "use_app_data";
export const AnalyzeRequest_query = "query";
export const AnalyzeRequest_games = "games";
export interface AnalyzeResult {
/**
Request
*/
  request?: AnalyzeRequest;
/**
Data
*/
  data?: apipb_datasetmodels.DataSet;
/**
Created At
*/
  createdAt?: string|number;
}

export const AnalyzeResult_request = "request";
export const AnalyzeResult_data = "data";
export const AnalyzeResult_createdAt = "created_at";
/**
Active Users
*/
export interface ActiveUsersRequest {
  type?: ActiveUsersRequestType;
  dates?: string[]|number[];
  appId?: string;
}

export const ActiveUsersRequest_type = "type";
export const ActiveUsersRequest_dates = "dates";
export const ActiveUsersRequest_appId = "app_id";
export interface ActiveUsersResult {
/**
Request
*/
  request?: ActiveUsersRequest;
/**
Data
*/
  data?: apipb_datasetmodels.DataSet;
/**
Created At
*/
  createdAt?: string|number;
}

export const ActiveUsersResult_request = "request";
export const ActiveUsersResult_data = "data";
export const ActiveUsersResult_createdAt = "created_at";
/**
Retention
*/
export interface RetentionRequest {
  type?: RetentionRequestType;
  dates?: string[]|number[];
  appId?: string;
}

export const RetentionRequest_type = "type";
export const RetentionRequest_dates = "dates";
export const RetentionRequest_appId = "app_id";
export interface RetentionResult {
/**
Request
*/
  request?: RetentionRequest;
/**
Data
*/
  data?: apipb_datasetmodels.DataSet;
/**
Created At
*/
  createdAt?: string|number;
}

export const RetentionResult_request = "request";
export const RetentionResult_data = "data";
export const RetentionResult_createdAt = "created_at";
export interface InactiveUsersRequest {
  appId?: string;
  inactiveDuring?: TimeRange;
  activeDuring?: TimeRange;
}

export const InactiveUsersRequest_appId = "app_id";
export const InactiveUsersRequest_inactiveDuring = "inactive_during";
export const InactiveUsersRequest_activeDuring = "active_during";
export interface GetActiveUsersRequest {
  appId?: string;
  period?: TimeRange;
}

export const GetActiveUsersRequest_appId = "app_id";
export const GetActiveUsersRequest_period = "period";
export interface ActiveOnRangeRequest {
  range?: TimeRange;
  countryCodes?: string[];
  appIds?: string[];
}

export const ActiveOnRangeRequest_range = "range";
export const ActiveOnRangeRequest_countryCodes = "country_codes";
export const ActiveOnRangeRequest_appIds = "app_ids";
export interface GameInfoResponse {
  gameId?: string;
  events?: GameInfoResponseEventInfo[];
}

export const GameInfoResponse_gameId = "game_id";
export const GameInfoResponse_events = "events";
export interface GameInfoResponseFieldInfo {
  name?: string;
  type?: GameInfoResponseFieldType;
}

export const GameInfoResponseFieldInfo_name = "name";
export const GameInfoResponseFieldInfo_type = "type";
export interface GameInfoResponseEventInfo {
  name?: string;
  fields?: GameInfoResponseFieldInfo[];
}

export const GameInfoResponseEventInfo_name = "name";
export const GameInfoResponseEventInfo_fields = "fields";
export interface AppDataInfoReq {
}

