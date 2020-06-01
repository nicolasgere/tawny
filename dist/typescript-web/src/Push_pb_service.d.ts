// package: tawny
// file: Push.proto

import * as Push_pb from "./Push_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PushServiceSubscribe = {
  readonly methodName: string;
  readonly service: typeof PushService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof Push_pb.SubscribeInput;
  readonly responseType: typeof Push_pb.Message;
};

type PushServicePublish = {
  readonly methodName: string;
  readonly service: typeof PushService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof Push_pb.PushInput;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class PushService {
  static readonly serviceName: string;
  static readonly Subscribe: PushServiceSubscribe;
  static readonly Publish: PushServicePublish;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class PushServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  subscribe(requestMessage: Push_pb.SubscribeInput, metadata?: grpc.Metadata): ResponseStream<Push_pb.Message>;
  publish(
    requestMessage: Push_pb.PushInput,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  publish(
    requestMessage: Push_pb.PushInput,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
}

