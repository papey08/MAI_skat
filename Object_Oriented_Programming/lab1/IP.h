#ifndef IP_H
#define IP_H

class IPAddress
{
public:
    IPAddress();
    IPAddress(unsigned char _a, unsigned char _b, unsigned char _c, unsigned char _d);

    friend IPAddress operator+(IPAddress A, IPAddress B);
    friend IPAddress operator-(IPAddress A, IPAddress B);
    friend bool operator==(IPAddress A, IPAddress B);
    friend bool operator!=(IPAddress A, IPAddress B);
    friend bool operator<(IPAddress A, IPAddress B);
    friend bool operator>(IPAddress A, IPAddress B);
    friend bool operator<=(IPAddress A, IPAddress B);
    friend bool operator>=(IPAddress A, IPAddress B);

    void Print();

    bool Check(IPAddress Addr, IPAddress Mask);

private:
    unsigned char a, b, c, d;
};

#endif
