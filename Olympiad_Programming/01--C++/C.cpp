#include <iostream> //Made by Matvey Popov MAI Ì8Î-108Á-20
#include <cmath>
using namespace std;

int
main ()
{
  double D, X, Y, d;
  cin >> D >> X >> Y;
  d=sqrt(X*X + Y*Y);
  if (d<=D) {
      cout << "YES";
  }
    else {
        cout << "NO";
    }


  return 0;
}
