#!/usr/bin/env bash

echo postinstall.sh
systemctl enable onvif-viewer
systemctl start onvif-viewer