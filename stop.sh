#!/bin/bash

PID_FOLDER="./bin/pid/*"
for pidfile in $PID_FOLDER; do
	PID=`cat $pidfile`
    if [ "x$PID" != "x" ]
    then
    	echo "kill $PID"
    	kill $PID
    fi
done
