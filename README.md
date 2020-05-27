
# 🦉 Tawny  🦉
  
🚀 Tawny is an extremly low latency, public facing pubsub. It use protobuf for end to end communication (from browser to server and server to browser) using [grpc-web](https://github.com/grpc/grpc-web)
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
## Getting started  
Visit the [documentation](docs),
## Feature roadmap  
  
This is our feature roadmap. If you want to ask a new feature. Please open an issue  
  
##### v0.1 Alpha => On going  
- [x] Push service (Public facing publish/subscrib)
- [x] Presence  service (Presence on channel/topic with state)
- [x] Typescript npm package
- [x]  Automatique TLS terminaison  
- [x]  Docker image
- [ ] Secure channel  
- [ ] Admin service
  
##### Backlog  

- [ ] Cluster mode
- [ ] Android client  
- [ ] Admin ui   
- [ ] Clients ( java/php/c#/python)
  
## Performance
With 1 core / 1gb of ram on Digital ocean,  Tawny is able to process up to 100k message seconds with < 5 ms latency.
  
## Author  
  
👤 **nicolas gere-lamaysouette**  
  
  
## Show your support  
  
Give a ⭐️ if you support the project!  
  
***  
_This README was generated with ❤️ by [readme-md-generator]
