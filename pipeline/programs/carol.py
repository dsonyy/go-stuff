import sys
import os

print("carol:", sys.argv, os.getcwd(), file=sys.stderr, sep="\t")

while True:
    txt = input("carol <- ")
    print(txt, flush=True)
