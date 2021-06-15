import time
import sys
import os

print("alice:", sys.argv, os.getcwd(), file=sys.stderr, sep="\t")

while True:
    print("alice", flush=True)
    time.sleep(0.5)
