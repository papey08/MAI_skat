#include <stdio.h>         //g++ prog2.cpp -L. -ldl -o main -Wl,-rpath -Wl,.
#include <stdlib.h>
#include <dlfcn.h>
#include <vector>

using namespace std;

int main()
{
    void* h = NULL;
    int (*GCF)(int A, int B);
    int* (*Sort)(int* array);
    int lib;
    printf("0 for change libs, 1 for re1.cpp, 2 for re2.cpp, 3 for exit:\n");
    scanf("%d", &lib);
    while ((lib != 1)&&(lib != 2))
    {
        printf("Input error, try again:\n");
        scanf("%d", &lib);
    }
    if (lib == 1)
    {
        h = dlopen("./libd1.so", RTLD_LAZY);
    }
    if (lib == 2)
    {
        h = dlopen("./libd2.so", RTLD_LAZY);
    }
    GCF = (int(*)(int, int))dlsym(h, "GCF");
    Sort = (int*(*)(int*))dlsym(h, "Sort");
    unsigned command;
    printf("1 for GCF, 2 for Sort:\n");
    char c1 = '1', c2 = '1';
    scanf("%d", &command);
    while (1)
    {
        if (command == 0)
        {
            if (lib == 1)
            {
                dlclose(h);
                h = dlopen("./libd2.so", RTLD_LAZY);
                GCF = (int(*)(int, int))dlsym(h, "GCF");
                Sort = (int*(*)(int*))dlsym(h, "Sort");
                lib = 2;
                printf("re1 changed on re2\n");
                scanf("%d", &command);
                continue;
            }
            else
            {
                dlclose(h);
                h = dlopen("./libd1.so", RTLD_LAZY);
                GCF = (int(*)(int, int))dlsym(h, "GCF");
                Sort = (int*(*)(int*))dlsym(h, "Sort");
                lib = 1;
                printf("re2 changed on re1\n");
                scanf("%d", &command);
                continue;
            }
        }
        if (command == 1)
        {
            int a, b;
            scanf("%d%c%d%c", &a, &c1, &b, &c2);
            if ((a < 1)||(b < 1))
            {
                printf("Input error\n");
            }
            else
            {
                printf("%d\n", GCF(a, b));
            }
        }
        if (command == 2)
        {
            int a;
            char c = '1';
            vector<int> v;
            while (c != '\n')
            {
                scanf("%d%c", &a, &c);
                v.push_back(a);
            }
            c = '1';
            int* arr = new int[v.size() + 1];
            arr[0] = v.size();
            for (int i = 0; i < v.size(); ++i)
            {
                arr[i + 1] = v[i];
            }
            Sort(arr);
            for (int i = 1; i < arr[0] + 1; ++i)
            {
                printf("%d ", arr[i]);
            }
            printf("\n");
            delete [] arr;
        }
        if (command == 3)
        {
            break;
        }
        scanf("%d", &command);
    }
    dlclose(h);
    return 0;
}
