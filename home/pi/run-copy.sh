#!/bin/bash
mount /dev/$1 /mnt
cp /home/pi/bis2a* /mnt

# Set date if the system time is before, or at most two hours after, the
# time indicated in the date file (for turning clocks forward or back)
if [ -e /mnt/current_date.txt ]; then
    if [ `date +%s` -le $[`date -d "$(cat /mnt/current_date.txt)" +%s`-7200] ]; then
        date -s "`cat /mnt/current_date.txt`"
    fi
fi

umount /mnt

# Gzip and move files older than 3 months to backup folder
find /home/pi/ -mtime +90 -exec gzip {} \; -exec mv {}.gz /home/pi/backup/ \;
