# Raspberry Pi Card Reader

## Introduction

These are files for setting up a Raspberry Pi (RPi) 3 running Raspbian Stretch Lite
to log timestamped data from a card reader that uses keyboard emulation (i.e.,
most magnetic card readers). An Adafruit Ultimate GPS module is recommended for
automatic timekeeping. Data are stored in CSV files.

Operation is headless; the primary interface for I/O with the RPi (besides the
card reader) is via a FAT-formatted USB stick (Windows-compatible USB sticks) or
via its web server located at http://10.23.45.1 when connected to the RPi's hotspot.

**NB**: You will not have access to the Internet while connected to the RPi's hotspot.

## Setup

Boot the pi. Install the gpsd and gpsd-clients packages (`apt-get install gpsd gpsd-clients`). Also install the ntp package and disable systemd-timesyncd:

    systemctl stop systemd-timesyncd.service
    systemctl disable systemd-timesyncd.service

    apt-get install ntp

Also install dnsmasq and hostapd for wireless access

    apt-get install dnsmasq hostapd

And install apache2 for showing files

    apt-get install apache2
    
Shutdown the pi. On a linux machine, mount the pi's microSD card's partition where the root directory resides. Extract the files in this directory to the RPi's root directory (/). The root directory is usually on the second partition on the microSD card. 

## Usage

### Setting Time

#### Using the GPS module
Plug in the GPS module with the RPi off. Turn on the RPi.

#### Manually, using a USB stick
Without an RTC or GPS module, time must be set on the RPi every time
it is plugged in (so avoid unplugging it).

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

Create a file on a USB stick called `timezone.txt` with the timezone formatted
the same way it is shown on the Wikipedia page (i.e., as shown below):

    America/Los_Angeles
    
Plug the USB stick in the RPi.

### Changing hotspot name and password

For security purposes, you should change the password in the /etc/hostapd/hostapd.conf
file after connecting for the first time. To do so, log in to the RPi via SSH (if using
Windows, you will first need to install an SSH client, such as PuTTY). On macOS and GNU/Linux,
run:

    ssh pi@10.23.45.1

Enter the password (the default password for Raspbian's pi account is `raspberry`). Your
terminal window should display `[pi@raspberrypi:~ $]` if login was successful. Then, edit
the hostapd.conf file by running nano:

    sudo nano /etc/hostapd/hostapd.conf

Navigate up and down the file using your keyboard's arrow keys.

Find the `ssid=Badger` line and change the hotspot's name to a name of your choice (e.g.,
`ssid=CardReader`).

Find the `wpa_passphrase=BadgerBadgerSnake` line and change the hotspot's password to a
password of your choice (e.g., `wpa_passphrase=AmazingPassword`).

Press Ctrl + O on your keyboard and then hit Return to save the file. Then press Ctrl + X
to exit nano. Reboot the raspberry pi:

    sudo shutdown -r now

### Obtaining Data

#### Using a Web browser

With a Wi-Fi enabled computer, tablet, or phone, connect to the access point (default name is
`Badger`). The default password is `BadgerBadgerSnake`. In your Web browser,
navigate to http://10.23.45.1. Click or tap on any of the CSV file(s) named `cards_read`
(e.g., `cards_read_S2018.csv`) and save the file using your Web browser's save functionality.

#### Using a USB stick

CSV files modified within the last 90 days (or earlier, if no data have been
retrieved for longer than 90 days) are automatically copied to a USB stick
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

