#!/bin/bash

set -euo pipefail



curl -sS https://example.com/non_existent_page && echo "Curl command failed as expected"

echo "Exit code for the last command is: $?"

cat non_existent_file.txt | echo "Cat command failed as expected"
echo "Exit code for the last command is: $?"