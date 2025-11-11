getFileMode() {
    if [ -f "$1" ]; then
        echo "It's a regular file."
    elif [ -d "$1" ]; then
        echo "It's a directory."
    else
        echo "It's something else."
    fi
}


FILE_PATH="exitreturn.sh"
getFileMode "$FILE_PATH"

