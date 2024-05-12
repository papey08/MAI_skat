//               cppid.cpp
#include <iostream>
#include <fstream>
#include <iomanip>
#include "fsm.h"
using namespace std;

int main()
{
  tFSM fsm;
///////////////////////
  addstr(fsm,0,"_",1);
  addrange(fsm,0,'A','Z',1);
  addrange(fsm,0,'a','z',1);
  addstr(fsm,1,"_",1);
  addrange(fsm,1,'A','Z',1);
  addrange(fsm,1,'a','z',1);
  addrange(fsm,1,'0','9',1);
  fsm.final(1);
///////////////////////

  cout << "Number of states = " << fsm.size()
       << "\n";
  char buf[1000];
  do{
   char* name="cppid.ss";
    ifstream file(name);
    if(!file){
 cout<<"Can't open file "<< name << " !\n";
 continue;
            }
    while(file){
     *buf=0;
     file.getline(buf,1000);
     cout<< ">"<<buf<<endl;
  int res = fsm.apply(buf);
  cout << setw(res?res+1:0) << "^"
       << endl;
     }//whil
 } while(false);
 return 0;
}

