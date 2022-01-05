#include <unistd.h>
#include <iostream>
#include <vector>
#include <string>
#include <fstream>

using namespace std;

int main()
{
    int fd[2];
    pipe(fd);
    int id = fork();
    if (id == -1)
    {
        return -1;
    }
    if (id == 0)
    {
        string filename;
        int length;
        read(fd[0], &length, sizeof(int));
        for (int i = 0; i < length; i++)
        {
            char c;
            read(fd[0], &c, sizeof(char));
            filename.push_back(c); 
        }
        ofstream outfile(filename);
        int t;
        read(fd[0], &t, sizeof(int));
        int amount;
        read(fd[0], &amount, sizeof(int));
        int sum = 0;
        for (int i = 0; i < amount; i++)
        {
            int n;
            read(fd[0], &n, sizeof(int));
            sum += n;
        }
        for (int i = 0; i < t; i++)
        {
            int amount;
            read(fd[0], &amount, sizeof(int));
            int sum = 0;
            for (int i = 0; i < amount; i++)
            {
                int n;
                read(fd[0], &n, sizeof(int));
                sum += n;
            }
            outfile << sum << endl;
        }
        outfile.close();
        close(fd[0]);
        close(fd[1]);
    }
    else
    {   
        cout << "Parent's PID: " << getpid() << endl;
        cout << "Child's PID: " << id << endl;
        vector<int> numbers;
        string filename;
        cout << "Enter the file name:\n";
        cin >> filename;
        int length = filename.length();
        write(fd[1], &length, sizeof(int));
        for (int i = 0; i < length; i++)
        {
            write(fd[1], &filename[i], sizeof(char));
        }
        cout << "Enter amount of commands:\n";
        int t;
        cin >> t;
        t++;
        write(fd[1], &t, sizeof(int));
        for (int i = 0; i < t; i++)
        {
            string s;
            getline(cin, s);
            vector<int> numbers;
            string n;
            for (int i = 0; i < s.length(); i++)
            {
                if ((s[i] != ' ')||(s[i] != '\0'))
                {
                    n.push_back(s[i]);
                }
                if ((s[i] == ' ')||(s[i] == '\0')||(s[i] == '\n'))
                {
                    int num = stoi(n);
                    n = "";
                    numbers.push_back(num);
                }
            }
            int amount = numbers.size();
            write(fd[1], &amount, sizeof(int));
            for (int i = 0; i < amount; i++)
            {
                write(fd[1], &numbers[i], sizeof(int));
            }
        }
        close(fd[1]);
        close(fd[0]);
    }
    return 0;
}