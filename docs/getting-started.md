## Getting started


#### Using docker
Only docker is supported for now. To start a new tawny server for local development it's pretty simple:

```bash
docker run -e TAWNY_ADMIN_KEY="A_SECRET_KEY" -p 8080:80 -p 4000:4000 elmazout/tawny
```
This will start a tawny server listening on port 8080 for http and 4000 for grpc.

#### Connect with grpcurl
[grpcurl](https://github.com/fullstorydev/grpcurl) is a tool to communicate with a grpc server using the command line. We will use that as an admin interface with the tawny server.

Listing services:
```bash
~/ grpcurl -plaintext localhost:4000 list

grpc.reflection.v1alpha.ServerReflection
tawny.AdminService
tawny.PresenceService
tawny.PushService
```

We will need to create a channel. For that we will use the admin service.
```bash
~/ grpcurl -plaintext localhost:4000 describe tawny.AdminService

tawny.AdminService is a service:
service AdminService {
  rpc CreateChannelOrUpdate ( .tawny.CreateOrUpdateChannelInput ) returns ( .google.protobuf.Empty );
  rpc DeleteChannel ( .tawny.DeleteChannelInput ) returns ( .google.protobuf.Empty );
  rpc ListChannel ( .google.protobuf.Empty ) returns ( .tawny.ListChannelOutput );
}
```
``CreateChannelOrUpdate`` is the grpc method we want to use. All these services are defined in the [protofile](https://github.com/nicolasgere/Tawny/tree/master/protos).
Let's create the channel now based on the proto definition.

```bash
~/ grpcurl -plaintext -d '{"name":"test-channel", "configuration":{"admin_required_to_push":false}}' localhost:4000 tawny.AdminService/CreateChannelOrUpdate
```
 to list channel.
 
 ```bash
~/ grpcurl -plaintext localhost:4000 tawny.AdminService/ListChannel
{
  "channels": [
    {
      "name": "test-channel",
      "configuration": {

      }
    }
  ]
}
 ```
Voila !

As tawny does not provide any admin UI for now I recommand to use [bloomrpc](https://github.com/uw-labs/bloomrpc). It's a postman like UI for grpc. Using proto definition it gives you an easy way to test tawny.

