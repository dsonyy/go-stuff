import sys
import time

f = open("i-was-here.txt", "w")
f.write("I was here.")
f.close()


for i in sys.argv:
    print(i, end=", ")

while True:
    print("aaaa")
    time.sleep(0.5)
