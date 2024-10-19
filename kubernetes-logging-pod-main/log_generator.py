import time
import logging
from random import choice

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(message)s')
log_messages = [
    "User login successful",
    "File uploaded successfully",
    "Database connection error",
    "Unauthorized access attempt",
    "Transaction completed",
]

if __name__ == "__main__":
    while True:
        # Randomly select a log message and log it
        log_message = choice(log_messages)
        logging.info(log_message)
        time.sleep(5)  # Log every 5 seconds
