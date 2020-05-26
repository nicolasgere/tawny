// package: tawny
// file: Presence.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class PresenceSubscribeInput extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): void;

  getTopic(): string;
  setTopic(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PresenceSubscribeInput.AsObject;
  static toObject(includeInstance: boolean, msg: PresenceSubscribeInput): PresenceSubscribeInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PresenceSubscribeInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PresenceSubscribeInput;
  static deserializeBinaryFromReader(message: PresenceSubscribeInput, reader: jspb.BinaryReader): PresenceSubscribeInput;
}

export namespace PresenceSubscribeInput {
  export type AsObject = {
    channel: string,
    topic: string,
  }
}

export class Presence extends jspb.Message {
  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Presence.AsObject;
  static toObject(includeInstance: boolean, msg: Presence): Presence.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Presence, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Presence;
  static deserializeBinaryFromReader(message: Presence, reader: jspb.BinaryReader): Presence;
}

export namespace Presence {
  export type AsObject = {
    state: Uint8Array | string,
    key: string,
  }
}

export class PresenceSubscribeResponse extends jspb.Message {
  clearPresencesList(): void;
  getPresencesList(): Array<Presence>;
  setPresencesList(value: Array<Presence>): void;
  addPresences(value?: Presence, index?: number): Presence;

  getUpdateType(): PresenceSubscribeResponse.TypeMap[keyof PresenceSubscribeResponse.TypeMap];
  setUpdateType(value: PresenceSubscribeResponse.TypeMap[keyof PresenceSubscribeResponse.TypeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PresenceSubscribeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PresenceSubscribeResponse): PresenceSubscribeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PresenceSubscribeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PresenceSubscribeResponse;
  static deserializeBinaryFromReader(message: PresenceSubscribeResponse, reader: jspb.BinaryReader): PresenceSubscribeResponse;
}

export namespace PresenceSubscribeResponse {
  export type AsObject = {
    presencesList: Array<Presence.AsObject>,
    updateType: PresenceSubscribeResponse.TypeMap[keyof PresenceSubscribeResponse.TypeMap],
  }

  export interface TypeMap {
    FULL: 0;
    PARTIAL: 1;
  }

  export const Type: TypeMap;
}

export class HeartBeatInput extends jspb.Message {
  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  getKey(): string;
  setKey(value: string): void;

  getChannel(): string;
  setChannel(value: string): void;

  getTopic(): string;
  setTopic(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HeartBeatInput.AsObject;
  static toObject(includeInstance: boolean, msg: HeartBeatInput): HeartBeatInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HeartBeatInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HeartBeatInput;
  static deserializeBinaryFromReader(message: HeartBeatInput, reader: jspb.BinaryReader): HeartBeatInput;
}

export namespace HeartBeatInput {
  export type AsObject = {
    state: Uint8Array | string,
    key: string,
    channel: string,
    topic: string,
  }
}

