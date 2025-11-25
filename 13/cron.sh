#!/bin/bash

function log {
    # timestamp in ISO-like format
    local ts
    ts=$(date '+%Y-%m-%d %H:%M:%S')
    # support multiple arguments and preserve spacing
    echo "[$ts] $*" | tee -a /home/serpent/logs/cron.log
}

