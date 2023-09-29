#include <stdio.h>    //g++ prog1.cpp -L. -ld1 -o main1 -Wl,-rpath -Wl,.
#include <vector>        //g++ prog1.cpp -L. -ld2 -o main2 -Wl,-rpath -Wl,.

using namespace std;

extern "C" int GCF(int A, int B);
extern "C" int* Sort(int* array);

int main()
{
    int command;
    char c1 = '1', c2 = '1';
    printf("1 for GCF, 2 for Sort:\n");
    scanf("%d", &command);
    while (1)
    {
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
    return 0;
}
