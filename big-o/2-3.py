import time
from random import randrange


def find_min_n2(arr):
    min = arr[0]
    for i in arr:
        smallest = True
        for j in arr:
            if (i > j):
                smallest = False

        if smallest:
            min = i
    return min


def find_min_n(arr):
    min = arr[0]
    for i in arr:
        if (i < min):
            min = i
    return min


for listSize in range(1000, 10001, 1000):
    alist = [randrange(100000) for x in range(listSize)]
    start = time.time()
    print find_min_n2(alist)
    end = time.time()
    print "Size: %d time: %f" % (listSize, end - start)

for listSize in range(1000, 10001, 1000):
    alist = [randrange(100000) for x in range(listSize)]
    start = time.time()
    print find_min_n(alist)
    end = time.time()
    print "Size: %d time: %f" % (listSize, end - start)
