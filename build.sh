#!/bin/bash

cd controller && GO111MODULE=on make && cd ..
cd hub && GO111MODULE=on make && cd ..
cd message-dispatcher && GO111MODULE=on make && cd ..
cd publish-engine && GO111MODULE=on make && cd ..
cd registry && GO111MODULE=on make && cd ..
cd topic-acl && GO111MODULE=on make && cd ..
cd topic-manager && GO111MODULE=on make && cd ..
