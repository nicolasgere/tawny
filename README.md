
# 🦉 Tawny  🦉
  
🚀 Tawny is an extremly low latency, public facing pubsub. It use protobuf for end to end communication (from browser to server and server to browser) using [grpc-web](https://github.com/grpc/grpc-web) protocol. 
It allows you to build an awesome user experience with realtime features.
  
## Core features  
Tawny was designed with performance/scalling/reliability in mind. It provide multiple keys feature:
- **Publish/Subscribe**, publish a message from anywhere and get your users notified instantly  
- **Presence**,  know who is connected or not in a channel, perfect for building presense feature like google docs  
- **State**, you can share a state with your presence, awesome for "is writing" feature.  
- **Binary data**, thanks to profobuf/grpc, your payload don't need to be json. It can be every binary fuke
- **Private channel**, you can control who have access to a channel.  
- **Complete admin library/cli**, manage your Tawny server using a powerfull admin api.  
- **Monitoring**, monitor your server easily with our promotheus(and more to come)  metrics exporter.  

## HTTP2/Protobuff advantage over HTTP1/json
- **Network bandwidth** Profobuf allow the app to use less bandwidth as demonstrated there.[auth0 blog](https://auth0.com/blog/beating-json-performance-with-protobuf/)
- **Usage of http2** Http2 is the new version of the famous http protocal, faster, more reliable, allow server push and many more. [cloudfare blog](https://www.cloudflare.com/learning/performance/http2-vs-http1.1/)
- **Binary data** You can send raw binary data using tawny using protobuff. 

## Getting started  
Visit the [documentation](docs)
## Feature roadmap  
  
This is our feature roadmap. If you want to ask a new feature. Please open an issue  
  
##### v0.1 Alpha => On going  
- [x] Push service (Public facing publish/subscrib)
- [x] Presence  service (Presence on channel/topic with state)
- [X] Admin service
- [x] Typescript web npm package
- [x] Typescript server npm package
- [x] Golang module
- [x] Automatique TLS terminaison  
- [x] Docker image
- [x] Protected topic (message can only be push from authenticate admin request)

  
##### Backlog

- [ ] Secure channel  
- [ ] Cluster mode
- [ ] Android client  
- [ ] Admin ui   
- [ ] Clients ( java/php/c#/python)
  
## Performance
With 1 core / 1gb of ram on Digital ocean,  Tawny is able to process up to 50k message seconds within < 5 ms latency.
  
## Author  
  
👤 **nicolas gere-lamaysouette**  
  
  
## Show your support  
  
Give a ⭐️ if you support the project!  
  
***  
_This README was generated with ❤️ by [readme-md-generator]
