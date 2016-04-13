#!/bin/sh

#需要配置Redis环境变量
#export PATH=$PATH:/opt/redis-3.0.7/src
redis-server ./data/redis/redis.conf &
sleep 1
ps -ef |grep redis
