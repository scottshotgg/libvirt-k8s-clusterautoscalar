#!/bin/sh

# TODO: use this as a general injection mechanism 
# to modify the filesystem 
# Ultimately we will most likely need to figure out
# how to chroot into the drive after mounting to make
# the most generic usage out of itfa
mkdir -p $2

guestmount -a $1 -m /dev/sda2 --rw $2
echo "$3" > $2/etc/hostname
umount $2