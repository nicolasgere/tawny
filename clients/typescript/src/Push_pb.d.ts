// package: tawny
// file: Push.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class SubscribeInput extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): void;

  getTopic(): string;
  setTopic(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubscribeInput.AsObject;
  static toObject(includeInstance: boolean, msg: SubscribeInput): SubscribeInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SubscribeInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubscribeInput;
  static deserializeBinaryFromReader(message: SubscribeInput, reader: jspb.BinaryReader): SubscribeInput;
}

export namespace SubscribeInput {
  export type AsObject = {
    channel: string,
    topic: string,
  }
}

export class PushInput extends jspb.Message {
  getChannel(): string;
  setChannel(value: string): void;

  getTopic(): string;
  setTopic(value: string): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PushInput.AsObject;
  static toObject(includeInstance: boolean, msg: PushInput): PushInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PushInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PushInput;
  static deserializeBinaryFromReader(message: PushInput, reader: jspb.BinaryReader): PushInput;
}

export namespace PushInput {
  export type AsObject = {
    channel: string,
    topic: string,
    data: Uint8Array | string,
  }
}

export class Message extends jspb.Message {
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    data: Uint8Array | string,
  }
}

