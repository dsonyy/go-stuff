import sys
import os

print("bob:", sys.argv, os.getcwd(), file=sys.stderr, sep="\t")

while True:
    txt = input("bob <- ")
    print(txt, flush=True)
