// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_messages from "./messages.pb";
import * as apipb_models from "./models.pb";
import * as google_protobuf_google_protobuf_field_mask from "./google/protobuf/field_mask.pb";

export interface StatisticsReq {
  timeFrom?: string|number;
  timeTo?: string|number;
}

export interface StatisticsRes {
  statistics?: { [key: string]: string|number };
		}

export interface UpdateProfileReq {
  fields?: google_protobuf_google_protobuf_field_mask.FieldMask;
  profile?: apipb_models.Profile;
}

export interface UpdateChildReq {
  fields?: google_protobuf_google_protobuf_field_mask.FieldMask;
  child?: apipb_models.Child;
}

