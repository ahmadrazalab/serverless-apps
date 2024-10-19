#!/bin/bash

mv ./mailpit.service /etc/systemd/system
mkdir -p /home/mail
mv ./mailpit /home/mail


service start mailpit.service
service status mailpit.service

