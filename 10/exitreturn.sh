#!/bin/bash

exit_with_message() {
    local exit_code=$1
    local message=$2
    echo "$message"
    exit "$exit_code"
}

return_with_message() {
    local return_code=$1
    local message=$2
    echo "$message"
    return "$return_code"
}

return_with_message 0 "Returning with success code 0"
# Example usage:
exit_with_message 1 "Exiting with error code 1"

