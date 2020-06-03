## Tawny Typescript-web client library

The library is generated directly from protobuf files definition. It give access to a low level library

#### Getting started

You need first to start a tawny server and create a channel. [Getting started](https://github.com/nicolasgere/Tawny/blob/master/docs/getting-started.md)

```bash
npm i tawny-web
```

```javascript
//Create a client with http2 transport
const transport = grpc.FetchReadableStreamTransport()
const pushClient = new Push.Client("http://localhost:8080", { transport });

//Subscribe to a channel/topic
const subscribeInput = new Push.SubscribeInput();
subscribeInput.setChannel('channel-1');
subscribeInput.setTopic('topic-1');
const stream = this.pushClient.subscribe(subscribeInput);
stream.on('data', (message) => {
           //new data is received
});

//Publish a message
const pushInput = new Push.PushInput();
pushInput.setChannel('channel-1');
pushInput.setTopic('topic-1');

//Tawny works with binary data, that's why we encode json in binary format.
var encoder = new TextEncoder("utf-8");
pushInput.setData(encoder.encode(JSON.stringify({
            message: 'hello world',
            username:this.state.username
})))
pushClient.publish(pushInput, (err) => {
            if (err) {
                console.log(err);
            }
});
```

Check the example folder for a complete realtime chat with react.

For more information have a look to [grpc-web client](https://github.com/improbable-eng/grpc-web/tree/master/client/grpc-web) and protos definition