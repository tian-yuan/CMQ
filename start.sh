#!/bin/bash

cd bin
nohup ./hub mqtt >/dev/null 2>&1&
nohup ./controller controller >/dev/null 2>&1&
nohup ./message-dispatcher rpc >/dev/null 2>&1&
nohup ./publish-engine rpc >/dev/null 2>&1&
nohup ./registry rpc >/dev/null 2>&1&
nohup ./topic-acl rpc >/dev/null 2>&1&
nohup ./topic-manager rpc >/dev/null 2>&1&
