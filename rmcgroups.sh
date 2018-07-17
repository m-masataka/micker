#!/bin/bash


cPath="/sys/fs/cgroup"
array=(
$cPath"/devices/micker"
$cPath"/pids/micker"
$cPath"/blkio/micker"
$cPath"/freezer/micker"
$cPath"/net_cls,net_prio/micker"
$cPath"/hugetlb/micker"
$cPath"/perf_event/micker"
$cPath"/cpu,cpuacct/micker"
$cPath"/memory/micker"
$cPath"/cpuset/micker"
$cPath"/systemd/micker"
)


for dir in ${array[@]}
do
	rmdir $dir/$1
done
