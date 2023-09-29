#include <iostream>
#include <vector>
#include <thread>
#include <string>

using namespace std;

vector<double> res;

void bebra(double (*buffer)[3], double(*conv)[3], int current)
{
    for (int i = 0; i < 3; ++i)
    {
        for (int j = 0; j < 3; ++j)
        {
            res[current] += buffer[i][j] * conv[i][j];
        }
    }
}

int main(int argc, char *argv[])
{
    int thread_amount;
    if (argc < 2)
    {
        cout << "Enter thread amount:\n";
        cin >> thread_amount;
    }
    else
    {
        thread_amount = stoi(argv[1]);
    }
    cout << "Thread amount is " << thread_amount << endl;
    res.resize(thread_amount);
    for (int i = 0; i < thread_amount; ++i)
    {
        res[i] = 0.0;
    }
    vector<thread> th(thread_amount);
    int current = 0;
    int k;
    cout << "Enter k:\n";
    cin >> k;
    int lines, columns;
    cout << "Enter amount of lines and columns:\n";
    do
    {
        cin >> lines >> columns;
        if ((lines - 2 * k <= 0)||(columns - 2 * k <= 0))
        {
            cout << "Error, try again:\n";
        }
    } while ((lines - 2 * k <= 0)||(columns - 2 * k <= 0));
    vector<vector<double>> orig(lines, vector<double> (columns, 0.0));
    cout << "Enter original matrix:\n";
    for (int i = 0; i < lines; ++i)
    {
        for (int j = 0; j < columns; ++j)
        {
            cin >> orig[i][j];
        }
    }
    cout << "Enter conv. 3x3 matrix:\n";
    double conv[3][3];
    for (int i = 0; i < 3; ++i)
    {
        for (int j = 0; j < 3; ++j)
        {
            cin >> conv[i][j];
        }
    }
    vector<vector <double>> result(lines, vector<double> (columns, 0.0));
    for (int t = 1; t <= k; ++t)
    {
        for (int I = 0; I < lines - 2 * t; ++I)
        {
            for (int J = 0; J < columns - 2 * t; ++J)
            {
                double buffer[3][3];
                for (int i = 0; i < 3; ++i)
                {
                    for (int j = 0; j < 3; ++j)
                    {
                        buffer[i][j] = orig[i + I][j + J];
                    }
                }
                th[current] = thread(bebra, buffer, conv, current);
                result[I][J] = res[current];
                ++current;
                if (current == thread_amount)
                {
                    current = 0;
                    for (int i = 0; i < thread_amount; ++i)
                    {
                        th[i].join();
                        res[i] = 0.0;
                    }
                }
            }
            for (int i = 0; i < current; ++i)
            {
                th[i].join();
                res[i] = 0.0;
            }
            current = 0;
        }
        for (int i = 0; i < lines - 2 * t; ++i)
        {
            for (int j = 0; j < columns - 2 * t; ++j)
            {
                orig[i][j] = result[i][j];
            }
        }
    }
    cout << "\nResult:\n";
    for (int i = 0; i < lines - 2 * k; ++i)
    {
        for (int j = 0; j < columns - 2 * k; ++j)
        {
            cout << orig[i][j] << " ";
        }
        cout << endl;
    }
    return 0;
}