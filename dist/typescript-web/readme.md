
# 🦉 Tawny  🦉  [![npm version](https://badge.fury.io/js/tawny-web.svg)](https://badge.fury.io/js/tawny-web)   [![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/gitterHQ/gitter)

  
This repository contains Typescrit/JavaScript client for Nodejs for [Tawny](https://github.com/nicolasgere/tawny). Tawny  is an extremly low latency, public facing pubsub. It use protobuf for end to end communication (from browser to server and server to browser) using [grpc-web](https://github.com/grpc/grpc-web) protocol.   
It allows you to build an awesome realtime features without the pain.  
  
- [Installation](#Installation)  
- [Example](#example)  
- [API](#api)  
- [Community](#community)  
  
## Installation  
You can found instructions for running a Tawny server in the [getting started](https://github.com/nicolasgere/Tawny/blob/master/docs/getting-started.md) page.  
  
---  
Install tawny javascript client for web. Tawny is builded with grpc-web, that's why grpc-web is needed as peer dependency, which also need protobuf.

  
```bash  
npm install tawny-web --save 
#Install peer dependency
npm install google-protobuf @improbable-eng/grpc-web --save
npm install @types/google-protobuf --save-dev
```  
  
## Example  
  
This example show how to use the push service. Push service is used to share message in realtime to multiple user. Like a chat for example.


----

```javascript  
import {Push} from "tawny-web"; 
import {grpc} from "@improbable-eng/grpc-web";  

const transport = grpc.FetchReadableStreamTransport()  
const pushClient = new Push.Client("http://localhost:8080", { transport });
```
 [grpc-web](https://github.com/improbable-eng/grpc-web/tree/master/client/grpc-web) can use multiple transport layer. ( Http2, Http1, websocket etc). You can found more information about which one to use and how to configure a client transport [here](https://github.com/improbable-eng/grpc-web/blob/master/client/grpc-web/docs/transport.md#http/2-based-transports).
 Then we create a client for the push service, which target our Tawny server.
 

```javascript
//Subscribe to a channel/topic  
const subscribeInput = new Push.SubscribeInput();  
subscribeInput.setChannel('channel-1');  
subscribeInput.setTopic('topic-1');  
const stream = this.pushClient.subscribe(subscribeInput);  
stream.on('data', (message) => {  
 console.log(message);
});  
```
This part will subscribe to a specific channel/topic. **Channels have to be created before they can get use.**  [Tawny server documentation](https://github.com/nicolasgere/Tawny/blob/master/docs/getting-started.md) 

 ```javascript
//Publish a message  
const pushInput = new Push.PushInput();  
pushInput.setChannel('channel-1');  
pushInput.setTopic('topic-1');  
  
//Tawny works with binary data, that's why have to encore the json string. 
var encoder = new TextEncoder("utf-8");  
pushInput.setData(encoder.encode(JSON.stringify({  
  message: 'hello world',  
  username: 'eric' 
})))  
pushClient.publish(pushInput, (err) => {  
	if (err) {
		 console.log(err);  
	}}
 );  
```  

Data have to be binary, in this example we are using json format. We could also use protobuf for better performance. Then we use the client to publish our message.

Check the [example](./example) folder for a complete realtime chat with react.  
  
For more information have a look to [grpc-web client](https://github.com/improbable-eng/grpc-web/tree/master/client/grpc-web)

## Api

As the library is generated based on protobuf files. You can found the documentation of the api [here](https://github.com/nicolasgere/tawny/blob/master/docs/proto.md) ***Only push and presence services are available in the web client***

## Community

Join us on [gitter](https://gitter.im/tawny-server/community) 