extern "C" int GCF(int A, int B);           //g++ -fPIC -c re2.cpp -o d2.o
extern "C" int * Sort(int * array);         //g++ -shared d2.o -o libd2.so

using namespace std;

int min(int a, int b)
{
    return a < b ? a : b;
}

int GCF(int A, int B)
{
    int m = 1;
    for (int i = 1; i <= min(A, B); ++i)
    {
        if ((A % i == 0)&&(B % i == 0)&&(i > m))
        {
            m = i;
        }
    }
    return m;
}

void _sort(int* a, int first, int last)
{
    int i = first, j = last;
    int tmp, x = a[(first + last) / 2];
    do 
    {
        while (a[i] < x)
        {
            i++;
        }
        while (a[j] > x)
        {
            j--;
        }
        if (i <= j) 
        {
            if (i < j)
            {
                tmp = a[i];
                a[i] = a[j];
                a[j] = tmp;
            }
            i++;
            j--;
        }
    } while (i <= j);
    if (i < last)
    {
        _sort(a, i, last);
    }
    if (first < j)
    {
        _sort(a, first, j);
    }
}

int* Sort(int* array)
{
    _sort(array, 1, array[0]);
    return array;
}
