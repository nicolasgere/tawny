// package: tawny
// file: Push.proto

var Push_pb = require("./Push_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var PushService = (function () {
  function PushService() {}
  PushService.serviceName = "tawny.PushService";
  return PushService;
}());

PushService.Subscribe = {
  methodName: "Subscribe",
  service: PushService,
  requestStream: false,
  responseStream: true,
  requestType: Push_pb.SubscribeInput,
  responseType: Push_pb.Message
};

PushService.Publish = {
  methodName: "Publish",
  service: PushService,
  requestStream: false,
  responseStream: false,
  requestType: Push_pb.PushInput,
  responseType: google_protobuf_empty_pb.Empty
};

exports.PushService = PushService;

function PushServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

PushServiceClient.prototype.subscribe = function subscribe(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(PushService.Subscribe, {
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

PushServiceClient.prototype.publish = function publish(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(PushService.Publish, {
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

exports.PushServiceClient = PushServiceClient;

