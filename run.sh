#!/bin/sh

sh stop.sh

$APRIL_PATH/bin/AdminServer &
sleep 1
$APRIL_PATH/bin/WorldServer &
sleep 1
$APRIL_PATH/bin/GameServer -game_index=1 &
sleep 1
$APRIL_PATH/bin/GameServer -game_index=2 &
sleep 1
$APRIL_PATH/bin/GameServer -game_index=3 &
sleep 1
$APRIL_PATH/bin/GateServer &
sleep 1