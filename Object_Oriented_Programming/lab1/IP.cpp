#include "IP.h"
#include <iostream>
#include <stdio.h>

IPAddress::IPAddress() : a(0), b(0), c(0), d(0)
    {}

IPAddress::IPAddress(unsigned char _a, unsigned char _b, unsigned char _c, unsigned char _d) : a(_a), b(_b), c(_c), d(_d)
    {}

void IPAddress::Print()
{
    printf("%d %d %d %d\n", a, b, c, d);
}

bool IPAddress::Check(IPAddress Addr, IPAddress Mask)
{
    if (Mask.a == 0)
    {
        return ((Addr.a == 0)&&(Addr.b == 0)&&(Addr.c == 0)&&(Addr.d == 0));
    }
    if (Mask.a < 255)
    {
        return ((Addr.b == 0)&&(Addr.c == 0)&&(Addr.d == 0)&&(Mask.a + a - 255 == Addr.a));
    }
    if (Mask.b == 0)
    {
        return ((a == Addr.a)&&(Addr.b == 0)&&(Addr.c == 0)&&(Addr.d == 0));
    }
    if (Mask.b < 255)
    {
        return ((a == Addr.a)&&(Mask.b + b - 255 == Addr.b)&&(Addr.c == 0)&&(Addr.d == 0));
    }
    if (Mask.c == 0)
    {
        return ((a == Addr.a)&&(b == Addr.b)&&(Addr.c == 0)&&(Addr.d == 0));
    }
    if (Mask.c < 255)
    {
        return ((a == Addr.a)&&(b == Addr.b)&&(Mask.c + c - 255 == Addr.c)&&(Addr.d == 0));
    }
    if (Mask.d == 0)
    {
        return ((a == Addr.a)&&(b == Addr.b)&&(c == Addr.c)&&(Addr.d == 0));
    }
    if (Mask.d < 255)
    {
        return ((a == Addr.a)&&(b == Addr.b)&&(c == Addr.c)&&(Mask.d + d - 255 == Addr.d));
    }
    return true;
}

IPAddress operator+(IPAddress A, IPAddress B)
{
    unsigned _a = (A.a + B.a) % 256;
    unsigned _b = (A.b + B.b) % 256;
    unsigned _c = (A.c + B.c) % 256;
    unsigned _d = (A.d + B.d) % 256;
    return IPAddress(_a, _b, _c, _d);
}

IPAddress operator-(IPAddress A, IPAddress B)
{
    int _a = (A.a - B.a) % 256;
    int _b = (A.b - B.b) % 256;
    int _c = (A.c - B.c) % 256;
    int _d = (A.d - B.d) % 256;
    return IPAddress(_a, _b, _c, _d);
}

bool operator==(IPAddress A, IPAddress B)
{
    return ((A.a == B.a)&&(A.b == B.b)&&(A.c == B.c)&&(A.d == B.d));
}

bool operator!=(IPAddress A, IPAddress B)
{
    return !((A.a == B.a)&&(A.b == B.b)&&(A.c == B.c)&&(A.d == B.d));
}

bool operator>(IPAddress A, IPAddress B)
{
    if ((A.a == B.a)&&(A.b == B.b)&&(A.c == B.c))
    {
        return A.d > B.d;
    }
    if ((A.a == B.a)&&(A.b == B.b))
    {
        return A.c > B.c;
    }
    if (A.a == B.a)
    {
        return A.b > B.b;
    }
    return A.a > B.a;
}

bool operator<(IPAddress A, IPAddress B)
{
    if ((A.a == B.a)&&(A.b == B.b)&&(A.c == B.c))
    {
        return A.d < B.d;
    }
    if ((A.a == B.a)&&(A.b == B.b))
    {
        return A.c < B.c;
    }
    if (A.a == B.a)
    {
        return A.b < B.b;
    }
    return A.a < B.a;
}

bool operator>=(IPAddress A, IPAddress B)
{
    if ((A.a == B.a)&&(A.b == B.b)&&(A.c == B.c))
    {
        return A.d >= B.d;
    }
    if ((A.a == B.a)&&(A.b == B.b))
    {
        return A.c > B.c;
    }
    if (A.a == B.a)
    {
        return A.b > B.b;
    }
    return A.a > B.a;
}

bool operator<=(IPAddress A, IPAddress B)
{
    if ((A.a == B.a)&&(A.b == B.b)&&(A.c == B.c))
    {
        return A.d <= B.d;
    }
    if ((A.a == B.a)&&(A.b == B.b))
    {
        return A.c < B.c;
    }
    if (A.a == B.a)
    {
        return A.b < B.b;
    }
    return A.a < B.a;
}
