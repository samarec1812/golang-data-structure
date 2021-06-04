import random
import time

import binaryHeap

if __name__ == '__main__':
    for j in range (0, 100):
        heap = binaryHeap.BinaryHeap()

        arr = []
        for i in range(0, 100000):
            arr.append(random.randint(0, 50000))

        start = time.time()
        heap.buildHeap(arr)
        second = time.time()
        print(f"The time build {(second - start)*1000: 0.4f} ", end=' ')
        # print(len(heap.arr))
        # print(heap.arr)
        start = time.time()
        # start = time.perf_counter()
        for i in range(0, 10000):
            # print(heap.getMax())
            heap.getMax()

        # second = time.perf_counter()
        second = time.time()
        print(f"The time getMax {(second - start)*1000: 0.4f} ", end=' ')

        start = time.time()
        # start = time.perf_counter()
        for i in range(0, 10000):
            # print(heap.getMax())
            heap.Add(random.randint(0, 50000))

        # second = time.perf_counter()
        second = time.time()
        print(f"The time getMax {(second - start) * 1000: 0.4f} ")

