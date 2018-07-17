#!/bin/bash

dmsetup remove mydevice
dmsetup remove mypool
losetup -d /dev/loop0
losetup -d /dev/loop1

