# CFUPDATER

[![Go Report Card](https://goreportcard.com/badge/github.com/j4ng5y/cfupdater)](https://goreportcard.com/report/github.com/j4ng5y/cfupdater)

A fairly simple application that performs Dynamic DNS operations for a CloudFlare Resource Record.

## Installation

* On your linux host, create a user for the service:

    `sudo useradd -r -s /sbin/nologin cfupdater`

* Make a few directories:

    `sudo mkdir -p /opt/cfupdater/bin /opt/cfupdater/config /opt/cfupdater/services`

* Download the binary and the systemd unit files:

    ```bash
    sudo curl -o /opt/cfupdater/bin/cfupdater_linux_amd64 https://github.com/j4ng5y/cfupdater/releases/download/v0.2.2/cfupdater_linux_amd64

    sudo curl -o /opt/cfupdater/services/cfupdater.timer https://github.com/j4ng5y/cfupdater/releases/download/v0.2.2/cfupdater.timer

    sudo curl -o /opt/cfupdater/services/cfupdater.service https://github.com/j4ng5y/cfupdater/releases/download/v0.2.2/cfupdater.service
    ```

* Run the config-file maker part of the application:

    ```bash
    sudo /opt/cfupdater/bin/cfupdater_linux_amd64 configure \
    --cloudflare-api-token <YOUR_API_TOKEN> \
    --cloudflare-dns-zone-id <YOUR_CLOUDFLARE_DNS_ZONE_ID> \
    --cloudflare-dns-record-name <THE RECORD NAME TO UPDATE> \
    --ipinfo-api-token <YOUR IPINFO.IO API TOKEN>
    ```

    This just makes a config.yaml file in the `/opt/cfupdater/config` directory.

* Create/Enable/Start symlinks for the unit files:

    ```bash
    sudo ln -s /opt/cfupdater/services/cfupdater.timer /etc/systemd/system/cfupdater.timer
    
    sudo ln -s /opt/cfupdater/services/cfupdater.service /etc/systemd/system/cfupdater.service

    sudo systemctl daemon-reload

    sudo systemctl enable cfupdater.timer
    ```

    This will simply run the service 5 minutes after a system reboot and also every 24hours thereafter.

    __Note__: You can just as easily run the service itself by running a `sudo systemctl start cfupdater.service` command if you feel the need to do so.

* Move all the files from root's ownership, to the cfupdater user's ownership:

    `sudo chown -R cfupdater:cfupdater /opt/cfupdater`

## TODO

I do plan to make the installation a little less manual, but it serves my purposes as is for the time being, so more to come?