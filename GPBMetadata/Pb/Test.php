<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pb/test.proto

namespace GPBMetadata\Pb;

class Test
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
pb/test.protopb"
OtherRequest
name (	"

OtherReply
message (	27
Other.
SayOther.pb.OtherRequest.pb.OtherReply" BZtest/pb�grpc_phpbproto3'
        , true);

        static::$is_initialized = true;
    }
}

