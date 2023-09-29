extern "C" int GCF(int A, int B);         //g++ -fPIC -c re1.cpp -o d1.o
extern "C" int * Sort(int * array);       //g++ -shared d1.o -o libd1.so

int GCF(int A, int B)
{
    if (B == 0)
    {
        return A;
    }
	else
    {
        return GCF(B, A % B);
    }
}

int* Sort(int* array)
{
    for (int i = 1; i < array[0] + 1; ++i)
    {
        for (int j = 1; j < array[0]; ++j)
        {
            if (array[j] > array[j + 1])
            {
                int a = array[j];
                array[j] = array[j + 1];
                array[j + 1] = a;
            }
        }
    }
    return array;
}
