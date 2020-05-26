
# 🦉 Tawny  🦉
  
🚀 Tawny is an extremly low latency, public facing pubsub. It use protobuf for end to end communication (from browser to server and server to browser) using [grpc-web](https://github.com/grpc/grpc-web)
It allows you to build an awesome user experience with realtime features.   
  
  
  
  
  
## Core features  
Tawny was designed with performance/scalling/reliability in mind. It provide multiple keys feature:
- **Publish/Subscribe**, publish a message from anywhere and get your users notified instantly  
- **Presence**,  know who is connected or not in a channel, perfect for building presense feature like google docs  
- **State**, you can share a state with your presence, awesome for "is writing" feature.  
- **Private channel**, you can control who have access to a channel.  
- **Complete admin library/cli**, manage your Tawny server using a powerfull admin api.  
- **Monitoring**, monitor your server easily with our promotheus(and more to come)  metrics exporter.  
## Getting started  
### Using free cloud hosted buzz server  
Visit the demo project   
### Using your own server  
Visit the documentation   
## Feature roadmap  
  
This is our feature roadmap. If you want to ask a new feature. Please open an issue  
  
##### v0.1 Alpha => On going  
- [x] Publish/Subscribe  
- [x] Presence  
- [x] State  
- [x] GCP only (bigtable/pubsub)  
- [x] Typescript npm package (client)  
- [ ] Admin api  
  
##### Backlog  
- [ ] Tls terminaison  
- [ ] Secure channel  
- [ ] IOS client  
- [ ] Android client  
- [ ] Admin ui   
- [ ] Hbase support  
- [ ] Kafka support  
- [ ] Rabbitmq support  
  
## Performance  
Buzz rely on hbase(bigtable), and pubsub, it will scale horizontally if you use managed service like google bigtable and pubsub.  
  
## Author  
  
👤 **nicolas gere-lamaysouette**  
  
  
## Show your support  
  
Give a ⭐️ if you support the project!  
  
***  
_This README was generated with ❤️ by [readme-md-generator]
