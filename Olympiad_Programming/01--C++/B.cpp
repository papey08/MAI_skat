#include <iostream>
#include <iomanip>
using namespace std;
int
main ()
{
  double a, b, c;
  cin >> a >> b >> c;
  double d = (a + b) * c;
  cout << fixed << setprecision (6) << d << endl;
  return 0;

}
