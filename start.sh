#!/bin/bash

cd bin
export IOT_HUB_CONFIG_LOG_APPNAME=iothub
./hub mqtt >/dev/null 2>&1&
export IOT_CONTROLLER_CONFIG_LOG_APPNAME=controller
./controller controller >/dev/null 2>&1&
export IOT_DISPATCHER_CONFIG_LOG_APPNAME=dispatcher
./message-dispatcher rpc >/dev/null 2>&1&
export IOT_ENGINE_CONFIG_LOG_APPNAME=engine
./publish-engine rpc >/dev/null 2>&1&
export IOT_REGISTRY_CONFIG_LOG_APPNAME=registry
./registry rpc >/dev/null 2>&1&
export IOT_TOPICACL_CONFIG_LOG_APPNAME=topicacl
./topic-acl rpc >/dev/null 2>&1&
export IOT_TOPICMANAGER_CONFIG_LOG_APPNAME=topicmanager
./topic-manager rpc >/dev/null 2>&1&
