#!/bin/bash
mount /dev/$1 /mnt
cp /home/pi/bis2a* /mnt

# Set date if there's a date file and the system date is before 2017
if [ -e /mnt/current_date.txt ]; then
    if [ `date +%s` -le `date -d "$(cat /mnt/current-Date.txt)" +%s` ]; then
        date -s "`cat /mnt/current_date.txt`"
    fi
fi

umount /mnt

# Gzip and move files older than 3 months to backup folder
find /home/pi/ -mtime +90 -exec gzip {} \; -exec mv {}.gz /home/pi/backup/ \;
