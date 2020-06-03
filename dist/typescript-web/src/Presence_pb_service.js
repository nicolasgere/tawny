// package: tawny
// file: Presence.proto

var Presence_pb = require("./Presence_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var PresenceService = (function () {
  function PresenceService() {}
  PresenceService.serviceName = "tawny.PresenceService";
  return PresenceService;
}());

PresenceService.PresenceSubscribe = {
  methodName: "PresenceSubscribe",
  service: PresenceService,
  requestStream: false,
  responseStream: true,
  requestType: Presence_pb.PresenceSubscribeInput,
  responseType: Presence_pb.PresenceSubscribeResponse
};

PresenceService.HeartBeat = {
  methodName: "HeartBeat",
  service: PresenceService,
  requestStream: false,
  responseStream: false,
  requestType: Presence_pb.HeartBeatInput,
  responseType: google_protobuf_empty_pb.Empty
};

exports.PresenceService = PresenceService;

function PresenceServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PresenceServiceClient.prototype.presenceSubscribe = function presenceSubscribe(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(PresenceService.PresenceSubscribe, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

PresenceServiceClient.prototype.heartBeat = function heartBeat(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PresenceService.HeartBeat, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.PresenceServiceClient = PresenceServiceClient;

