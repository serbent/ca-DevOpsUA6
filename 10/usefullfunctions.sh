#!/bin/bash 


function log(){
    LOGFILE="script.log"
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $1" >> $LOGFILE
}


function check_command_success(){
    if [ $1 -eq 0 ]; then
        log "Command succeeded."
    else
        log "Command failed with exit code $1."
    fi
}

log "Starting script execution."


ls readme.ms
check_command_success $?



log "Script execution completed."
