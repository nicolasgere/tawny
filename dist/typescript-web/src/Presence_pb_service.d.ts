// package: tawny
// file: Presence.proto

import * as Presence_pb from "./Presence_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import {grpc} from "@improbable-eng/grpc-web";

type PresenceServicePresenceSubscribe = {
  readonly methodName: string;
  readonly service: typeof PresenceService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof Presence_pb.PresenceSubscribeInput;
  readonly responseType: typeof Presence_pb.PresenceSubscribeResponse;
};

type PresenceServiceHeartBeat = {
  readonly methodName: string;
  readonly service: typeof PresenceService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof Presence_pb.HeartBeatInput;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class PresenceService {
  static readonly serviceName: string;
  static readonly PresenceSubscribe: PresenceServicePresenceSubscribe;
  static readonly HeartBeat: PresenceServiceHeartBeat;
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

export class PresenceServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  presenceSubscribe(requestMessage: Presence_pb.PresenceSubscribeInput, metadata?: grpc.Metadata): ResponseStream<Presence_pb.PresenceSubscribeResponse>;
  heartBeat(
    requestMessage: Presence_pb.HeartBeatInput,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  heartBeat(
    requestMessage: Presence_pb.HeartBeatInput,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
}

