#!/bin/sh

function func() {
	killall -9 $1

	killall -0 $1
	while [ $? -ne 1 ]; do
		sleep 1
		killall -0 $1
	done
}

if [ $# -eq 0 ]
	then
		func AdminServer
		func GateServer
		func GameServer
		func WorldServer
	else
		func $1
fi