
//               adec.cpp
#include <iostream>
#include <fstream>
#include <iomanip>
#include "fsm.h"
using namespace std;

int main()
{
  tFSM Adec;
///////////////////////
//������� ���� ��� � �������� �����
  addrange(Adec,0,'0','9',1);
  addrange(Adec,1,'0','9',1);
  addstr(Adec,1,"e",2);
  addstr(Adec, 2, "+", 3);
  addstr(Adec, 2, "-", 3);
  addrange(Adec,3,'0','9',4);
  addrange(Adec,4,'0','9',4);
 Adec.final(4);
///////////////////////
  cout << "Number of states = " << Adec.size()
       << "\n";

  char buf[1000];
  do{
   char* name="dec.ss";
    ifstream file(name);
    if(!file){
 cout<<"Can't open file "<< name << " !\n";
 continue;
            }
    while(file){
     *buf=0;
     file.getline(buf,1000);
     cout<< ">"<<buf<<endl;
  int res = Adec.apply(buf);
  cout << setw(res?res+1:0) << "^"
       << endl;
     }//whil
 } while(false);
 return 0;
}

