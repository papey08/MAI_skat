#include <iostream>
#include <cmath>

using namespace std;

int main() {
int x1,y1,x2,y2,x3,y3,x4,y4,s,sf,sff,gg;
cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3 >> x4 >> y4;
int p1 = max(x1,x2);
int p2 = max(x3,x4);
int p3 = min(p1,p2);
int p4 = min(x1,x2);
int p5 = min(x3,x4);
int p6 = max(p4,p5);
int s1 = max(y1,y2);
int s2 = max(y3,y4);
int s3 = min(s1,s2);
int s4 = min(y1,y2);
int s5 = min(y3,y4);
int s6 = max(s4,s5);

if ((p3 - p6 < 0) || (s3 - s6 < 0)) {
s = 0;
}
else {
s = (p3-p6)*(s3-s6);
}

sf = (abs(x2-x1)) * (abs(y2-y1));
sff = (abs(x4-x3))*(abs(y4-y3));
gg = sf+sff-s;
cout << gg << endl;
}
