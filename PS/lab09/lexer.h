//                 lexer.h 2021
#ifndef LEXER_H
#define LEXER_H
#include "baselexer.h"
class tLexer:public tBaseLexer{
public:
 tLexer():tBaseLexer(){

//________________________________________

    addrange(Adec,0,'0','9',1);
    addstr  (Adec,1,".",   2);
    addrange(Adec,2,'0','9',2);
  Adec.final(2);
//________________________________________

    addrange(Aid,0,'a','z',1);
    addrange(Aid,1,'a','z',1);
    addstr  (Aid,1,"-!",   1);
  Aid.final(1);
//________________________________________

    addstr  (Aidq,0,"?",   1);
  Aidq.final(1);
}
};
#endif

