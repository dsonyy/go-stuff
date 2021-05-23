import sys
import time

f = open("i-was-here.txt", "w")
f.write("I was here.")
f.close()


for i in sys.argv:
    print(i, end=", ")

for i in range(10000000):
    print("aaaa", flush=True)
    time.sleep(0.5)
