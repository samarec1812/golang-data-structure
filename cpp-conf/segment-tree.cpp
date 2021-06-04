


int Arr[100000];     //массив
int Tree[400004];  //дерево отрезков. в вершинах хранятся суммы

//Построение дерева по изначальному массиву.
//v - номер текущей вершины; tl, tr - границы соответствующего отрезка
void Build_tree(int v, int tl, int tr) {
    if (tl == tr) {
        Tree[v] = Arr[tl];    //сумма отрезка из одного элемента равна этому элементу
    } else {
        //tm - средний элемент отрезка.
        //отрезок разбивается на два отрезка [tl; tm], [tm + 1; tr]
        int tm = (tl + tr) / 2;
        Build_tree(v * 2, tl, tm);
        Build_tree(v * 2 + 1, tm + 1, tr);
        Tree[v] = Tree[v * 2] + Tree[v * 2 + 1];
    }
}

//Запрос суммы
//l, r - границы запроса;
//v - номер текущей вершины; tl, tr - границы соответствующего отрезка
int Get_sum(int l, int r, int v, int tl, int tr) {

    if (l <= tl && tr <= r) {
        return Tree[v];
    }
    if (tr < l || r < tl) {
        return 0;
    }
    int tm = (tl + tr) / 2;
    return Get_sum(l, r, v * 2,     tl,     tm)
           + Get_sum(l, r, v * 2 + 1, tm + 1, tr);
}

void Print(int n) {
    for(int i = 0; i < n; i++) {
        std::cout << Arr[i] << " ";
    }
    std::cout << std::endl;
}

void Delete(int n) {
    for(int i = 0; i < n; i++) {
        Arr[i] = 0;
    }
    for(int i = 0; i < 4*n; i++) {
        Tree[i] = 0;
    }
}