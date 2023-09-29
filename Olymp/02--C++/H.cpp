#include <iostream>
#include <cmath>
#include <vector>

using namespace std;

int main()
{
    int N, n;
    cin >> N;
    for(int i = 0; i < N; i++)
    {
        cin >> n;
        vector<int> v(n);
        int a = 0;
        for (int g = 0; g < n; g++)
            cin >> v[g];
        for (int j = 0; j < n - 1; j++)
        {
            for(int l = 0; l < n - j - 1; l++)
            {
                if (v[l] > v[l + 1])
                {
                    int b;
                    b = v[l];
                    v[l] = v[l + 1];
                    v[l + 1] = b;
                    a++;
                }
            }
        }
        cout << a << endl;
    }
    return 0;
}
