//
// Created by ПК on 29.05.2021.
//

#include <utility>
#include <vector>

class BinaryHeap {
private:
    std::vector<int> v;
    void heapify(int index);

public:
    BinaryHeap() {
        v = std::vector<int>(0);

    }
    explicit BinaryHeap(std::vector<int> value) {
        v = std::move(value);
    }
    void Add(int value);
    void BuildHeap(std::vector<int> source);
    void Print();
    int GetMax();
};


void BinaryHeap::Add(int value) {
    v.push_back(value);
    int index = v.size() - 1;
    int parent = (index - 1) / 2;

    while(index > 0 && v[parent] < v[index]) {
        int temp = v[index];
        v[index] = v[parent];
        v[parent] = temp;

        index = parent;
        parent = (index - 1)/ 2;
    }
}

void BinaryHeap::heapify(int index) {
    int leftChild = 0, rightChild = 0, largestChild = 0;
    for(;;) {
        leftChild = 2 * index + 1;
        rightChild = 2 * index + 2;
        largestChild = index;

        if(leftChild < v.size() && v[leftChild] > v[largestChild]) {
            largestChild = leftChild;
        }
        if(rightChild < v.size() && v[rightChild] > v[largestChild]) {
            largestChild = rightChild;
        }
        if(largestChild == index) {
            break;
        }
        int temp = v[index];
        v[index] = v[largestChild];
        v[largestChild] = temp;
        index = largestChild;
    }
}

void BinaryHeap::BuildHeap(std::vector<int> source) {
    v = std::move(source);
    for(int i = v.size() / 2; i >= 0; i--) {
        heapify(i);
    }
}

void BinaryHeap::Print() {
    for(auto i: v) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
}

int BinaryHeap::GetMax() {
    int res = v[0];
    v[0] = v[v.size() - 1];
    v.pop_back();
    BinaryHeap::heapify(0);
    return res;
}

