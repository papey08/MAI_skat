#include <iostream>
#include "triangle.h"
#include "square.h"
#include "rectangle.h"

using namespace std;

int main()
{
    bool s = 1;
    while (s == 1)
    {
        cout << "Select the figure:\n";
        cout << "1) Triangle\n";
        cout << "2) Square\n";
        cout << "3) Rectangle\n";
        int f;
        cin >> f;
        if (f == 1)
        {
            
            cout << "Enter 3 points:\n";
            Triangle t(cin);
            t.Print(cout);
            cout << "Triangle contains " << t.VertexesNumber() << " vertices.\n";
            cout << "Area: " << t.Area() << endl; 
        }
        if (f == 2)
        {
            cout << "Enter 4 points:\n";
            Square s(cin);
            s.Print(cout);
            cout << "Square contains " << s.VertexesNumber() << " vertices.\n";
            cout << "Area: " << s.Area() << endl;
        }
        if (f == 3)
        {
            cout << "Enter 4 points:\n";
            Rectangle r(cin);
            r.Print(cout);
            cout << "Rectangle contains " << r.VertexesNumber() << " vertices.\n";
            cout << "Area: " << r.Area() << endl;
        }
        cout << "Want to continue? (1 or 0)\n";
        cin >> s;
    }
    cout << "Finished.\n";
    return 0;
}