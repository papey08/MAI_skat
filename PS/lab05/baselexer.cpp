//                 baselexer.cpp 2021
#include "baselexer.h"
using namespace std;

tBaseLexer::tBaseLexer():start(buf),end(buf),
                       line(0){
 *buf = 0;
    addstr  (Aint,0,"+-", 2);
    addrange(Aint,2,'0','9',1);
    addrange(Aint,0,'0','9',1);
    addrange(Aint,1,'0','9',1);
  Aint.final(1);
    const char SP=' ';
    addstr  (Astr,0,"\"",   1);
    addstr  (Astr,1,"\"",   2);
    addrange(Astr,1,SP,'!',1);
    addrange(Astr,1,'#','[',1);
    addrange(Astr,1,']','~',1);
    addrange(Astr,1,'\x80','\xff',1);
    addstr  (Astr,1,"\\",   3);
// \'  \"  \?  \\  \a  \b  \f  \n  \r  \t  \v
    addstr  (Astr,3,"\"\\?'abfnrtv",   1);
  Astr.final(2);
//________________________________________

    addstr  (Aoper,0,"+-*/=",   1);
    addstr  (Aoper,0,"<>",   2);
    addstr  (Aoper,2,"=",   1);
  Aoper.final(1);
  Aoper.final(2);
//________________________________________

    addstr  (Abool,0,"#",   1);
    addstr  (Abool,1,"tf",  2);
  Abool.final(2);
//________________________________________
}

bool tBaseLexer::Begin(const char* filename){
   line = 0;
   start = buf;
   end = buf;
   *buf = 0;
   lexeme = "";
   fsource.clear();
   fsource.open(filename);
   return bool(fsource);
  }

string tBaseLexer::GetToken(){
  lexeme = "";
  for (;;){
   if(*end==0 || *end==';'){
     fsource.getline(buf,bufsize);
     end = buf;
     ++line;
     if(!fsource) {
                  *end=0;
                   start=end;
                   return "#";
                  }
// skip Racket metadata
  if(*buf=='#'&&!(*(buf+1)=='t'||*(buf+1)=='f'))*buf=0;
     continue;
   }
   if(!std::isspace(*end)) break;
   ++end;
  }//for...

   start = end;
   if(*start == '('|| *start == ')'){
     lexeme = string(1,*start);
     ++end;
     return lexeme;
   }
  if(*start == '"'){
   int lstr=Astr.apply(start);
   if(lstr != 0){
    end += lstr;
   lexeme = string(start,end);
   return "$str";
   }
   lexeme = string(start);
   end += lexeme.length();
   return "?";
  }
   while(*(++end)!=0){
    if(std::isspace(*end)||
       *end == ';'||
       *end == '"'||
       *end == '('||
       *end == ')' ) break;
    }//while...
   lexeme = string(start,end);
   int total        =lexeme.length();
   const char* input=lexeme.c_str();
   if(Aid.apply(input)==total) return "$id";
   if(Adec.apply(input)==total) return "$dec";
   if(Aint.apply(input)==total) return "$int";
   if(Abool.apply(input)==total) return "$bool";
   if(Aidq.apply(input)==total) return "$idq";
   if(Aoper.apply(input)==total) return lexeme;
   return "?";
 }

