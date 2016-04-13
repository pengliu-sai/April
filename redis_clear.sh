#!/bin/sh
redis-cli -p 6379 -a liupeng flushdb &
sleep 1
redis-cli -p 6379 -a liupeng flushall &
sleep 1