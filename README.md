# Raspberry Pi Card Reader

## Introduction

These are files for setting up a Raspberry Pi (RPi) running Raspbian Jessie Lite
to log timestamped data from a card reader that uses keyboard emulation (i.e.,
most magnetic card readers). Data are stored in CSV files.

Operation is headless; the primary interface for I/O with the RPi (besides the
card reader) is via a FAT-formatted USB stick (Windows-compatible USB sticks).

## Setup

Extract the files in this directory to the RPi's root directory (/). The root
directory is usually on the second partition on the microSD card. Boot the pi. Install the gpsd and gpsd-clients packages (`apt-get install gpsd gpsd-clients`). Also install the ntp package and disable systemd-timesyncd:

   systemctl stop systemd-timesyncd.service
   systemctl disable systemd-timesyncd.service

   apt-get install ntp

## Usage

### Setting Time

Time must be set on the RPi every time it is plugged in (so avoid unplugging it).

To set approximate time on the RPi<sup>1</sup>, create a file called
`current_date.txt` in the root directory<sup>2</sup> of the USB stick with the
current time formatted as shown below:  

    2017-02-13 18:00

To get more accurate timestamps, insert a future time and insert the USB stick
at that future time.

An example `current_date.txt` file is in the /home/pi folder.

### Setting Timezone

In addition to setting time, you may also wish to set the timezone if you'd like
automatic daylight saving time time changes.

Look up your appropriate timezone under the TZ column
[on this page](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

Create a file on the USB stick called `timezone.txt` with the timezone formatted
the same way it is shown on the Wikipedia page (i.e., as shown below):

    America/Los_Angeles

### Obtaining Data

CSV files modified within the last 90 days (or earlier, if no data have been
retrieved for longer than 90 days) are automatically copied to the USB stick
upon insertion.

There is no audible or visual notification for when the process is complete, but
waiting a few seconds after the USB stick stops flashing usually allows
sufficient time for copying to finish. It should then be safe to remove the USB
stick.

Files older than 90 days are compressed and stored in a backup folder on the RPi
after they have been retrieved. The backup files cannot be retrieved using a USB
stick, but can be viewed by removing the RPi's microSD card and viewing the
contents on a different computer. These files should be removed in the unlikely
event that the RPi runs out of disk space.

## Footnotes

<sup>1</sup>The RPi doesn't have a real-time clock  
<sup>2</sup>More accurately, the root directory of the first filesystem on the
USB stick. Most USB sticks only have one filesystem, however.

