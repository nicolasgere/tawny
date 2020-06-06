// package: tawny
// file: Admin.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class ChannelConfiguration extends jspb.Message {
  getAdminRequiredToPush(): boolean;
  setAdminRequiredToPush(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChannelConfiguration.AsObject;
  static toObject(includeInstance: boolean, msg: ChannelConfiguration): ChannelConfiguration.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChannelConfiguration, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChannelConfiguration;
  static deserializeBinaryFromReader(message: ChannelConfiguration, reader: jspb.BinaryReader): ChannelConfiguration;
}

export namespace ChannelConfiguration {
  export type AsObject = {
    adminRequiredToPush: boolean,
  }
}

export class CreateOrUpdateChannelInput extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasConfiguration(): boolean;
  clearConfiguration(): void;
  getConfiguration(): ChannelConfiguration | undefined;
  setConfiguration(value?: ChannelConfiguration): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOrUpdateChannelInput.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOrUpdateChannelInput): CreateOrUpdateChannelInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateOrUpdateChannelInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOrUpdateChannelInput;
  static deserializeBinaryFromReader(message: CreateOrUpdateChannelInput, reader: jspb.BinaryReader): CreateOrUpdateChannelInput;
}

export namespace CreateOrUpdateChannelInput {
  export type AsObject = {
    name: string,
    configuration?: ChannelConfiguration.AsObject,
  }
}

export class Channel extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasConfiguration(): boolean;
  clearConfiguration(): void;
  getConfiguration(): ChannelConfiguration | undefined;
  setConfiguration(value?: ChannelConfiguration): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Channel.AsObject;
  static toObject(includeInstance: boolean, msg: Channel): Channel.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Channel, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Channel;
  static deserializeBinaryFromReader(message: Channel, reader: jspb.BinaryReader): Channel;
}

export namespace Channel {
  export type AsObject = {
    name: string,
    configuration?: ChannelConfiguration.AsObject,
  }
}

export class ListChannelOutput extends jspb.Message {
  clearChannelsList(): void;
  getChannelsList(): Array<Channel>;
  setChannelsList(value: Array<Channel>): void;
  addChannels(value?: Channel, index?: number): Channel;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListChannelOutput.AsObject;
  static toObject(includeInstance: boolean, msg: ListChannelOutput): ListChannelOutput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListChannelOutput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListChannelOutput;
  static deserializeBinaryFromReader(message: ListChannelOutput, reader: jspb.BinaryReader): ListChannelOutput;
}

export namespace ListChannelOutput {
  export type AsObject = {
    channelsList: Array<Channel.AsObject>,
  }
}

export class DeleteChannelInput extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteChannelInput.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteChannelInput): DeleteChannelInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteChannelInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteChannelInput;
  static deserializeBinaryFromReader(message: DeleteChannelInput, reader: jspb.BinaryReader): DeleteChannelInput;
}

export namespace DeleteChannelInput {
  export type AsObject = {
    name: string,
  }
}

