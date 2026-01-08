
# open a file

FILENAME="writefile.log"

logfile = open("app.log", "a") # append
logfile.write("Hello, World!\n")

logfile.close()

writefile = open(FILENAME, "w") # overwrite

writefile.write("Hello, World!\n")

writefile.close()

readfile = open(FILENAME, "r") # read

print(readfile.read())

readfile.close()


# context manager
with open(FILENAME, "a+") as readfile:
    print(readfile.read())
    readfile.write("Hello, World! I'm still here\n")
    readfile.seek(0) # move cursor to the beginning of the file # try commenting it and running the app. see the differencec.
    print(readfile.read())
