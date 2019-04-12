// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var framebuffer_pb = require('./framebuffer_pb.js');

function serialize_DrawResponse(arg) {
  if (!(arg instanceof framebuffer_pb.DrawResponse)) {
    throw new Error('Expected argument of type DrawResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DrawResponse(buffer_arg) {
  return framebuffer_pb.DrawResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_FrameBuffer(arg) {
  if (!(arg instanceof framebuffer_pb.FrameBuffer)) {
    throw new Error('Expected argument of type FrameBuffer');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_FrameBuffer(buffer_arg) {
  return framebuffer_pb.FrameBuffer.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_FrameSequence(arg) {
  if (!(arg instanceof framebuffer_pb.FrameSequence)) {
    throw new Error('Expected argument of type FrameSequence');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_FrameSequence(buffer_arg) {
  return framebuffer_pb.FrameSequence.deserializeBinary(new Uint8Array(buffer_arg));
}


var DrawerService = exports.DrawerService = {
  // DrawFrame draws one frame
  drawFrame: {
    path: '/Drawer/DrawFrame',
    requestStream: false,
    responseStream: false,
    requestType: framebuffer_pb.FrameBuffer,
    responseType: framebuffer_pb.DrawResponse,
    requestSerialize: serialize_FrameBuffer,
    requestDeserialize: deserialize_FrameBuffer,
    responseSerialize: serialize_DrawResponse,
    responseDeserialize: deserialize_DrawResponse,
  },
  // DrawFrames draws a series of frames
  drawFrames: {
    path: '/Drawer/DrawFrames',
    requestStream: false,
    responseStream: false,
    requestType: framebuffer_pb.FrameSequence,
    responseType: framebuffer_pb.DrawResponse,
    requestSerialize: serialize_FrameSequence,
    requestDeserialize: deserialize_FrameSequence,
    responseSerialize: serialize_DrawResponse,
    responseDeserialize: deserialize_DrawResponse,
  },
};

exports.DrawerClient = grpc.makeGenericClientConstructor(DrawerService);
