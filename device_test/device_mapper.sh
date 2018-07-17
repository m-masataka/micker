#!/bin/bash
losetup -f /root/data_device.img
losetup -f /root/metadata_device.img
sleep 1
dmsetup create mypool --table "0 $((1024*1024*1024/512)) thin-pool /dev/loop1 /dev/loop0 512 8192"
sleep 2
dmsetup message /dev/mapper/mypool 0 "create_thin 0"
sleep 2
dmsetup create mydevice --table "0 $((1024*1024*100/512)) thin /dev/mapper/mypool 0"

