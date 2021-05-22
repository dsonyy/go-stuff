import sys
import time

while True:
    time.sleep(1)
    for i in sys.argv:
        print(i, end=", ")
    print("")
