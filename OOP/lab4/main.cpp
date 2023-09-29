#include <iostream>
#include <string>
#include "TBinaryTree.h"

using namespace std;

int main () 
{
    cout << "Enter TEST to check program quickly\n";
    cout << "Else enter MASTER\n";
    string command;
    cin >> command;
    if (command == "TEST")
    {
        TBinaryTree TREE;
        Point o(0, 0);
        Point ax(1, 0);
        Point ay(0, 1);
        Point bx(2, 0);
        Point by(0, 2);
        Point cx(3, 0);
        Point cy(0, 3);
        Triangle A(o, ax, ay);
        Triangle B(o, bx, by);
        Triangle C(o, cx, cy);
        cout << "Triangle A: " << A << endl;
        cout << "Triangle B: " << B << endl;
        cout << "Triangle C: " << C << endl;
        TREE.Push(B);
        TREE.Push(A);
        TREE.Push(C);
        cout << "Push triangle B\nPush triangle A\nPush triangle C\n";
        cout << "Print tree:\n" << TREE << endl;
        cout << "GetItemNotLess 1:\n";
        Triangle R = TREE.GetItemNotLess(1);
        cout << R << endl;
        cout << "Count triangles with the same area with (0, 0) (2, 0) (0, 1):\n";
        Triangle D(o, bx, ay);
        cout << TREE.Count(D) << endl;
        cout << "Pop triangle C\n";
        TREE.Pop(C);
        cout << "Print tree:\n" << TREE << endl;
        cout << "Is tree empty?\n";
        if (TREE.Empty() == 1)
        {
            cout << "Yes\n";
        }
        else
        {
            cout << "No\n";
        }
        cout << "Done\n";
        return 0;
    }
    if (command == "MASTER")
    {
        cout << "Commands:\n";
        cout << "PUSH -- adds triangle into the tree\n";
        cout << "GINL -- returns triangle with area >= than yours\n";
        cout << "COUNT -- calculates amount of triangles with the same area in the tree\n";
        cout << "POP -- removes triangle from the tree\n";
        cout << "EMPTY -- returns is tree is empty\n";
        cout << "PRINT -- prints the tree\n";
        cout << "END -- clears the tree and ends program\n";
        cout << "TEST -- run test script to check the program\n";
        cout << "Enter your first command:" << endl;
        cin >> command;
        TBinaryTree TREE;
        while (command != "END")
        {
            if (command == "PUSH")
            {
                cout << "Enter chords of 3 points of triangle to PUSH: \n";
                Triangle T(cin);
                TREE.Push(T);
                cout << "Enter next command:\n";
                cin >> command;
            }
            if (command == "GINL")
            {
                cout << "Enter area: \n";
                double a;
                cin >> a;
                a -= 0.0000001;
                Triangle R = TREE.GetItemNotLess(a);
                cout << "Result:\n";
                cout << R << endl;
                cout << "Enter next command:\n";
                cin >> command;
            }
            if (command == "COUNT")
            {
                cout << "Enter chords of 3 points of triangle to COUNT: \n";
                Triangle T(cin);
                unsigned r = TREE.Count(T);
                cout << "Result is " << r << endl;
                cout << "Enter next command:\n";
                cin >> command;
            }
            if (command == "POP")
            {
                cout << "Enter chords of 3 points of triangle to POP: \n";
                Triangle T(cin);
                TREE.Pop(T);
                cout << "Enter next command:\n";
                cin >> command;
            }
            if (command == "EMPTY")
            {
                if (TREE.Empty() == 1)
                {
                    cout << "Tree is empty\n";
                }
                else
                {
                    cout << "Tree is not empty\n";
                }
                cout << "Enter next command:\n";
                cin >> command;
            }
            if (command == "PRINT")
            {
                cout << TREE << endl;
                cout << "Enter next command:\n";
                cin >> command;
            }
        }
        TREE.Clear();
        cout << "Done\n";
        return 0;
    }
}
