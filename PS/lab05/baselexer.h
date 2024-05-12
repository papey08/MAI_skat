//                 baselexer.h 2020
#ifndef BASELEXER_H
#define BASELEXER_H
#include <iostream>
#include <fstream>
#include <string>
#include "fsm.h"

class tBaseLexer{
 std::ifstream fsource;
 int         line;
 char        buf[1000];
 static const size_t bufsize=1000;
 char*       start;
 char*       end;
 std::string lexeme;

public:
 tBaseLexer();
 ~tBaseLexer() {}
 bool Begin(const char* filename);
 void End(){fsource.close();}

 std::string GetToken();
 std::string GetLexeme()const{return lexeme;}
 int GetLineCount()const{return line;}
 int GetStartPos()const{return start-buf;}
 int GetEndPos()const{return end-buf;}
 std::string GetLineText()const{
                     return std::string(buf);}

protected:
  tFSM Aint;
  tFSM Adec;
  tFSM Aid;
  tFSM Aidq;
private:
  tFSM Aoper;
  tFSM Abool;
  tFSM Astr;
};
#endif

