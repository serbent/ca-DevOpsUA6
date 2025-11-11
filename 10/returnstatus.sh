#!/bin/bash

echo 'This message should return a status code of 0'
echo "Exit code for the last command is: "$?


ls non_existent_file
echo "Exit code for the last command is: "$?

if [ ! -f "somefile.txt" ]; then
    echo "File does not exist, returning status code 2"
    return 2 2>/dev/null
fi


LIST_OF_FILES=("file1.txt" "file2.txt" "file3.txt")

echo "LIST_OF_FILES contains: ${LIST_OF_FILES[@]}"

echo "Exit code of list of file is "$?

