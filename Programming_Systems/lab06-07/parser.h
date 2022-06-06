//                         parser.h 2019
#ifndef PARSER_H
#define PARSER_H
#include "lexer.h"
#include "slr.h"

class tParser{
public:
 typedef tGramma::tSymb tSymb;
 typedef tLR::tState tState;
 tParser(const char* gramma_name);
 ~tParser(){}

 int parse(const char* source_name);
 const std::string& getMessage()const
                    {return ferror_message;}
 operator bool()const{return gr;}
 bool PARSER_DEBUG;
 const tGramma& getGramma()const{return gr;}

private:
// блокировать копирующие функции
 tParser(const tParser& );
 tParser& operator=(const tParser& );

 tLexer             lex;
 tGramma            gr;
 tLR                lr;
 std::string        ferror_message;

// функция 
 tGramma::tSymb getTerm();

};

inline tGramma::tSymb tParser::getTerm(){
 std::string token=lex.GetToken();
 tGramma::tSymb term = gr.encode(token);

//  из токена $id извлекаются ключевые слова
 if(token == "$id"){
   std::string ident = lex.GetLexeme();
   tGramma::tSymb keyword=gr.encode(ident);//поиск в алфавите
// если идентификатор найден и является терминалом,
   if(keyword && gr.terminal(keyword))
//                        это ключевое слово
                          term = keyword;
  }
 return term;
}
#endif
