
# 🦉 Tawny  🦉 
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/tawny-server/community) 
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=nicolasgere_tawny&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=nicolasgere_tawny)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=nicolasgere_tawny&metric=alert_status)](https://sonarcloud.io/dashboard?id=nicolasgere_tawny)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=nicolasgere_tawny&metric=security_rating)](https://sonarcloud.io/dashboard?id=nicolasgere_tawny)
  
🚀 Tawny is an extremely low latency, public facing pubsub. It uses protobuf for end-to-end communication (from browser to server and server to browser) using [grpc-web](https://github.com/grpc/grpc-web) protocol. 
It allows you to build an awesome user experience with realtime features.
  
## Core features  
Tawny was designed with performance/scaling/reliability in mind. It provides multiple keys feature:
- **Publish/Subscribe**, publish a message from anywhere and get your users notified instantly. 
- **Presence**,  know who is connected or not in a channel, perfect for building presense feature like google docs. 
- **State**, you can share a state with your presence, awesome for "is typing" feature.  
- **Binary data**, thanks to profobuf/grpc, your payload doesn't need to be json. It can be every binary file
- **Private channel**, you can control who has access to a channel.  
- **Complete admin library/cli**, manage your Tawny server using a powerful admin api.  
- **Monitoring**, monitor your server easily with our promotheus (and more to come)  metrics exporter.  

## HTTP2/Protobuff advantage over HTTP1/json
- **Network bandwidth** Profobuf allows the app to use less bandwidth as demonstrated there.[auth0 blog](https://auth0.com/blog/beating-json-performance-with-protobuf/)
- **Usage of http2** Http2 is the new version of the famous http protocol, faster, more reliable, allows server push and much more. [cloudfare blog](https://www.cloudflare.com/learning/performance/http2-vs-http1.1/)
- **Binary data** You can send raw binary data with tawny using protobuff. 

## Getting started  
Visit the [documentation](docs)
## Feature roadmap  
  
This is our feature roadmap. If you want to ask for a new feature, please open an issue  
  
##### v0.1 Alpha => On going  
- [x] Push service (Public facing publish/subscrib)
- [x] Presence  service (Presence on channel/topic with state)
- [X] Admin service
- [x] Typescript web npm package
- [ ] Typescript server npm package
- [x] Golang module
- [x] Automatic TLS termination for https
- [ ] GRPC secure mode
- [x] Docker image
- [x] Protected topic (message can only be pushed from authenticated admin request)

  
##### Backlog

- [ ] Secure channel  
- [ ] Cluster mode
- [ ] Admin ui   
- [ ] Clients (java/swift/php/c#/python)
  
## Performance
With 1 core / 1gb of ram on Digital ocean,  Tawny is able to process up to 50k messages seconds within < 5 ms latency.
  
## Community

Join us on [gitter](https://gitter.im/tawny-server/community) 
  
  
## Show your support  
  
Give a ⭐️ if you support the project!  
  
***  
_This README was generated with ❤️ by [readme-md-generator]
