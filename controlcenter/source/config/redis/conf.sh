#!/bin/bash
#-e 判断对象是否存在
#-d 判断对象是否存在，并且为目录
#-f 判断对象是否存在，并且为常规文件
#-L 判断对象是否存在，并且为符号链接
#-h 判断对象是否存在，并且为软链接
#-s 判断对象是否存在，并且长度不为0
#-r 判断对象是否存在，并且可读
#-w 判断对象是否存在，并且可写
#-x 判断对象是否存在，并且可执行
#-O 判断对象是否存在，并且属于当前用户
#-G 判断对象是否存在，并且属于当前用户组
#-nt 判断file1是否比file2新 [ “/data/file1” -nt “/data/file2” ]
#-ot 判断file1是否比file2旧 [ “/data/file1” -ot “/data/file2” ]

# bind 172.100.0.2
# protected-mode no
# port 6391
# tcp-backlog 511
# timeout 0
# tcp-keepalive 300
# daemonize no
# supervised no
# pidfile /var/run/redis_6391.pid
# loglevel notice
# logfile "logs/nodes-6391.log"
# databases 16
# always-show-logo yes
# save 900 1
# save 300 10
# save 60 10000
# stop-writes-on-bgsave-error yes
# rdbcompression yes
# rdbchecksum yes
# dbfilename dump-6391.rdb
# dir ./
# replica-serve-stale-data yes
# replica-read-only yes
# repl-diskless-sync no
# repl-diskless-sync-delay 5
# repl-disable-tcp-nodelay no
# replica-priority 100
# lazyfree-lazy-eviction no
# lazyfree-lazy-expire no
# lazyfree-lazy-server-del no
# replica-lazy-flush no
# appendonly no
# appendfilename "appendonly.aof"
# appendfsync everysec
# no-appendfsync-on-rewrite no
# auto-aof-rewrite-percentage 100
# auto-aof-rewrite-min-size 64mb
# aof-load-truncated yes
# aof-use-rdb-preamble yes
# lua-time-limit 5000
# cluster-enabled yes
# cluster-config-file "cluster/nodes-6391.conf"
# cluster-node-timeout 15000
# slowlog-log-slower-than 10000
# slowlog-max-len 128
# latency-monitor-threshold 0
# notify-keyspace-events ""
# hash-max-ziplist-entries 512
# hash-max-ziplist-value 64
# list-max-ziplist-size -2
# list-compress-depth 0
# set-max-intset-entries 512
# zset-max-ziplist-entries 128
# zset-max-ziplist-value 64
# hll-sparse-max-bytes 3000
# stream-node-max-bytes 4096
# stream-node-max-entries 100
# activerehashing yes
# client-output-buffer-limit normal 0 0 0
# client-output-buffer-limit replica 256mb 64mb 60
# client-output-buffer-limit pubsub 32mb 8mb 60
# hz 10
# dynamic-hz yes
# aof-rewrite-incremental-fsync yes
# rdb-save-incremental-fsync yes



ports="6379 6391 6392 6393 6394 6395 6396"
dir=""
for port in $ports;
do
dir="nodes-$port.conf"
if [ -f "$dir" ];then
rm -r $dir
fi
touch $dir

echo bind 172.100.0.2 >> $dir
echo protected-mode no >> $dir

# port 6391 \
# tcp-backlog 511 \
# timeout 0 \
# tcp-keepalive 300 \
# daemonize no \
# supervised no \
# pidfile /var/run/redis_6391.pid \
# loglevel notice \
# logfile "logs/nodes-$port.log" \
# databases 16 \
# always-show-logo yes \
# save 900 1 \
# save 300 10 \
# save 60 10000 \
# stop-writes-on-bgsave-error yes \
# rdbcompression yes \
# rdbchecksum yes \
# dbfilename dump-$port.rdb \
# dir ./ \
# replica-serve-stale-data yes \
# replica-read-only yes \
# repl-diskless-sync no \
# repl-diskless-sync-delay 5 \
# repl-disable-tcp-nodelay no \
# replica-priority 100 \
# lazyfree-lazy-eviction no \
# lazyfree-lazy-expire no \
# lazyfree-lazy-server-del no \
# replica-lazy-flush no \
# appendonly no \
# appendfilename "appendonly.aof" \
# appendfsync everysec \
# no-appendfsync-on-rewrite no \
# auto-aof-rewrite-percentage 100 \
# auto-aof-rewrite-min-size 64mb \
# aof-load-truncated yes \
# aof-use-rdb-preamble yes \
# lua-time-limit 5000 \
# cluster-enabled yes \
# cluster-config-file "cluster/nodes-$port.conf" \
# cluster-node-timeout 15000 \
# slowlog-log-slower-than 10000 \
# slowlog-max-len 128 \
# latency-monitor-threshold 0 \
# notify-keyspace-events "" \
# hash-max-ziplist-entries 512 \
# hash-max-ziplist-value 64 \
# list-max-ziplist-size -2 \
# list-compress-depth 0 \
# set-max-intset-entries 512 \
# zset-max-ziplist-entries 128 \
# zset-max-ziplist-value 64 \
# hll-sparse-max-bytes 3000 \
# stream-node-max-bytes 4096 \
# stream-node-max-entries 100 \
# activerehashing yes \
# client-output-buffer-limit normal 0 0 0 \
# client-output-buffer-limit replica 256mb 64mb 60 \
# client-output-buffer-limit pubsub 32mb 8mb 60 \
# hz 10 \
# dynamic-hz yes \
# aof-rewrite-incremental-fsync yes \
# rdb-save-incremental-fsync yes  \
# > $dir
done