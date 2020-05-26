# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Presence.proto](#Presence.proto)
    - [HeartBeatInput](#tawny.HeartBeatInput)
    - [Presence](#tawny.Presence)
    - [PresenceSubscribeInput](#tawny.PresenceSubscribeInput)
    - [PresenceSubscribeResponse](#tawny.PresenceSubscribeResponse)
  
    - [PresenceSubscribeResponse.Type](#tawny.PresenceSubscribeResponse.Type)
  
    - [PresenceService](#tawny.PresenceService)
  
- [Push.proto](#Push.proto)
    - [Message](#tawny.Message)
    - [PushInput](#tawny.PushInput)
    - [SubscribeInput](#tawny.SubscribeInput)
  
    - [PushService](#tawny.PushService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="Presence.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Presence.proto



<a name="tawny.HeartBeatInput"></a>

### HeartBeatInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [bytes](#bytes) |  | The state for a particular person |
| key | [string](#string) |  | A unique key for presence |
| channel | [string](#string) |  | A reference to the channel |
| topic | [string](#string) |  | A reference to the topic |






<a name="tawny.Presence"></a>

### Presence



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [bytes](#bytes) |  | The state for a particular person |
| key | [string](#string) |  | A unique key for presence |






<a name="tawny.PresenceSubscribeInput"></a>

### PresenceSubscribeInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  | A reference to the channel |
| topic | [string](#string) |  | A reference to the topic |






<a name="tawny.PresenceSubscribeResponse"></a>

### PresenceSubscribeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| presences | [Presence](#tawny.Presence) | repeated | A list of presence |
| update_type | [PresenceSubscribeResponse.Type](#tawny.PresenceSubscribeResponse.Type) |  | For now only full is supported |





 


<a name="tawny.PresenceSubscribeResponse.Type"></a>

### PresenceSubscribeResponse.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| FULL | 0 |  |
| PARTIAL | 1 |  |


 

 


<a name="tawny.PresenceService"></a>

### PresenceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PresenceSubscribe | [PresenceSubscribeInput](#tawny.PresenceSubscribeInput) | [PresenceSubscribeResponse](#tawny.PresenceSubscribeResponse) stream | Subscribe to a channel/topic to get presence update |
| HeartBeat | [HeartBeatInput](#tawny.HeartBeatInput) | [.google.protobuf.Empty](#google.protobuf.Empty) | Heartbeat to a channel |

 



<a name="Push.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## Push.proto



<a name="tawny.Message"></a>

### Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [bytes](#bytes) |  | Your payload |






<a name="tawny.PushInput"></a>

### PushInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  |  |
| topic | [string](#string) |  |  |
| data | [bytes](#bytes) |  | Your payload |






<a name="tawny.SubscribeInput"></a>

### SubscribeInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  | A reference to the channel |
| topic | [string](#string) |  | A reference to the topic |





 

 

 


<a name="tawny.PushService"></a>

### PushService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Subscribe | [SubscribeInput](#tawny.SubscribeInput) | [Message](#tawny.Message) stream | Subscribe to a channel/topic to receive Owl Message |
| Publish | [PushInput](#tawny.PushInput) | [.google.protobuf.Empty](#google.protobuf.Empty) | Subscribe to a channel/topic to receive Owl Message |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

