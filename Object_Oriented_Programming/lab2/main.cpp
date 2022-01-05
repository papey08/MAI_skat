#include <iostream>
#include <stdio.h>
#include <string>
#include "IP.h"

std::string operator "" _with_dots(const char* s, size_t size)
{
    std::string str; 
    for (int i = 0; i < 16; ++i)
    {
        if (s[i] == ' ')
        {
            str.push_back('.');
        }
        else
        {
            str.push_back(s[i]);
        }
    }
    return str;
}

int main()
{
    std::cout << "Enter A and B IPAddresses:\n";
    int a, b, c, d;
    std::cin >> a >> b >> c >> d;
    IPAddress A(a, b, c, d);
    std::cin >> a >> b >> c >> d;
    IPAddress B(a, b, c, d);
    if (A == B)
    {
        std::cout << "IPAdresses are equal\n";
    }
    if (A > B)
    {
        std::cout << "A is greater than B\n";
    }
    if (A < B)
    {
        std::cout << "B is greater than A\n";
    }
    std::cout << "Sum of A and B is ";
    (A + B).Print();
    std::cout << "Difference of A and B is ";
    (A - B).Print();
    std::cout << "Enter the mask: \n";
    std::cin >> a >> b >> c >> d;
    IPAddress M(a, b, c, d);
    if (A.Check(B, M))
    {
        std::cout << "A belongs to B\n";
    }
    else
    {
        std::cout << "A do not belongs to B" << std::endl;
    }
    std::cout << "Using of the literal: \n";
    //char* s = M.Get();
    std::cout << "192 168 10 12"_with_dots << std::endl;
    return 0;
}
