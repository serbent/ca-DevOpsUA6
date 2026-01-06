# exmaple 1 
import box

box.boxprint("Hello World")


# exmaple 2
from box import boxprint

boxprint("Hello World")

# exmaple 3
from box import *

boxprint("Hello World")

# exmaple 4
import modules.logging

modules.logging.logger("Hello World")

# exmaple 5
from modules.logging import logger

logger("Hello World")
