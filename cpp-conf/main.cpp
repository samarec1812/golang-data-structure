#include <iostream>
#include <chrono>
#include "binary-heap.cpp"
#include "segment-tree.cpp"



//void Print(std::vector<int> a) {
//    for(auto i : a) {
//        std::cout << i << " ";
//    }
//    std::cout << std::endl;
//}
int main() {
//    for(int j = 0; j < 100; j++) {
//        std::vector<int> a;
//        // a.reserve(500000);
//        for (int i = 0; i < 500000; i++) {
//            a.push_back(rand() % 50000);
//        }
//        // std::vector<int> a = {1,5, 3, 28, 5, 67, 3, 46, 245, 6, 786};
//        BinaryHeap heap = BinaryHeap();
//
//        auto begin1 = std::chrono::steady_clock::now();
//        heap.BuildHeap(a);
//        auto end1 = std::chrono::steady_clock::now();
////
//        auto elapsed_ms1 = std::chrono::duration_cast<std::chrono::milliseconds>(end1 - begin1);
//        std::cout << "The time BUILD: " << elapsed_ms1.count() << " ms; ";
//
//        //  heap.Print();
//        auto begin = std::chrono::steady_clock::now();
//        for (int i = 0; i < 10000; i++) {
//            //  std::cout << heap.GetMax() << std::endl;
//            heap.GetMax();
//        }
//
//        auto end = std::chrono::steady_clock::now();
//
//        auto elapsed_ms = std::chrono::duration_cast<std::chrono::milliseconds>(end - begin);
//        std::cout << "GetMax: " << elapsed_ms.count() << " ms : ";
//
//        auto begin2 = std::chrono::steady_clock::now();
//        for (int i = 0; i < 10000; i++) {
//            //  std::cout << heap.GetMax() << std::endl;
//            heap.Add(rand() % 50000);
//        }
//
//        auto end2 = std::chrono::steady_clock::now();
//
//        auto elapsed_ms2 = std::chrono::duration_cast<std::chrono::milliseconds>(end2 - begin2);
//        std::cout << "Push: " << elapsed_ms.count() << " ms\n";
//    }
    int n = 100000;
    // std::cin >> n;
    for(int j = 0; j < 100; j++) {
        for (int i = 0; i < n; i++) {
            Arr[i] = (rand() % 50000);
        }
        auto begin1 = std::chrono::steady_clock::now();
        Build_tree(1, 0, n - 1);
        auto end1 = std::chrono::steady_clock::now();

        auto elapsed_ms1 = std::chrono::duration_cast<std::chrono::milliseconds>(end1 - begin1);
        std::cout << "The time BUILD: " << elapsed_ms1.count() << " ms; ";

        auto begin = std::chrono::steady_clock::now();
        for (int i = 0; i < 1000; i++) {
            //  std::cout << heap.GetMax() << std::endl;
            Get_sum(10, 70000, 1, 0, n - 1);
        }

        auto end = std::chrono::steady_clock::now();

        auto elapsed_ms = std::chrono::duration_cast<std::chrono::milliseconds>(end - begin);
        std::cout << "SUM: " << elapsed_ms.count() << " ms\n";
       //  Print(n);
        Delete(n);
    }
    return 0;
}