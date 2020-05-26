var dtoPresence = require('./src/Presence_pb');
var dtoPush = require('./src/Push_pb');

var servicePresence = require('./src/Presence_pb_service');
var servicePush = require('./src/Push_pb_service');



exports.Push = {
    SubscribeInput:dtoPush.SubscribeInput,
    Message: dtoPush.Message,
    PushInput: dtoPush.PushInput,
    Client: servicePush.PushServiceClient
};
exports.Presence = {
    HeartBeatInput: dtoPresence.HeartBeatInput,
    PresenceSubscribeResponse: dtoPresence.PresenceSubscribeResponse,
    PresenceSubscribeInput: dtoPresence.PresenceSubscribeInput,
    Presence: dtoPresence.Presence,
    Client: servicePresence.PresenceServiceClient,
}