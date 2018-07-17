#1/bin/bash

dmsetup suspend /dev/mapper/mydevice
echo "OK1"
dmsetup message /dev/mapper/mypool 0 "create_snap 1 0"
echo "OK2"
dmsetup resume /dev/mapper/mydevice
echo "OK3"
dmsetup create snap --table "0 2097152 mydevice /dev/mapper/mypool 1"
