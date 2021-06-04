class BinaryHeap:
    arr = []

    def __init__(self):
        self.arr = []

    def Add(self, other):
        self.arr.append(other)
        index = len(self.arr) - 1
        parent = int((index - 1) / 2)

        while index > 0 and self.arr[parent] < self.arr[index]:
            temp = self.arr[index]
            self.arr[index] = self.arr[parent]
            self.arr[parent] = temp

            index = parent
            parent = int((index - 1) / 2)

    def Heapify(self, other):
        leftChild, rightChild, largestChild = 0, 0, 0
        while True:
            leftChild = 2 * other + 1
            rightChild = 2 * other + 2
            largestChild = other

            if leftChild < len(self.arr) and self.arr[leftChild] > self.arr[largestChild]:
                largestChild = leftChild

            if rightChild < len(self.arr) and self.arr[rightChild] > self.arr[largestChild]:
                largestChild = rightChild

            if largestChild == other:
                break

            temp = self.arr[other]
            self.arr[other] = self.arr[largestChild]
            self.arr[largestChild] = temp
            other = largestChild

    def buildHeap(self, source):
        self.arr = source
        for i in range(int(len(self.arr) / 2), -1, -1):
            self.Heapify(i)

    def getMax(self):
        res = self.arr[0]
        self.arr[0] = self.arr[len(self.arr) - 1]
        self.arr.pop()
        self.Heapify(0)
        return res
