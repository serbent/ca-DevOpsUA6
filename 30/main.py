import logging
from datetime import datetime

def setup_logging(level, file_name="app.log"):
    """Setup logging configuration."""
    logging.basicConfig(
        level=level,
        format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
        handlers=[
            logging.FileHandler(file_name),
            logging.StreamHandler()
        ]
    )

LOGFILE = "app.log" + datetime.now().strftime("_%Y-%m-%d")
setup_logging(logging.INFO, LOGFILE)

logging.debug("Hello World")
logging.info("Hello World")
logging.warning("Hello World")
logging.error("Hello World")
logging.critical("Hello World")


""" Error handling """
x = input("Enter a number: ")
y = input("Enter a number: ")

try:
    result = int(x) / int(y)
except ZeroDivisionError:
    logging.error("Division by zero")
except ValueError:
    logging.error("Invalid input")
else:
    logging.info("Result: %s", result)
finally:
    logging.info("Finally")



print("We've handled the error")
logging.info("We've handled the error")

