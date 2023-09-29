//                 parser.cpp 2021
#include <sstream>
#include "parser.h"
using namespace std;

tParser::tParser(const char* gramma_name):lr(gr)
{
  PARSER_DEBUG=false;
  gr.loadFromFile(gramma_name);
  if(!gr){
   ferror_message = gr.getError();
      return;
  }
  SLRbuild(lr);
}

int tParser::parse(const char* source_name){
  if(lr.size()==0){
// испорчены управляющие таблицы
    return 1;
   }
   ferror_message.clear();
   if(!lex.Begin(source_name)){
    ferror_message = string("Can't open file ")+
                                     source_name;
    return 1;
    }

   std::ostringstream buf;
   vector<tSymb> stack;
   vector<tState> states;
   tState state = 0;
   tState next = 0;
   const tSymb start = gr.getStart();
   tSymb term = 1;// маркер
   stack.push_back(term);
   states.push_back(state);
   term = getTerm();
   if (!term){
     buf << "Lexis: unknown token!";
     }
//+++++++++++++++++++++++++++++++
   else while(true){
     next = lr.go(state,term);
     if(next==0){
       buf << "Syntax: unmatched token "
           << gr.decode(term) << "\nexpected: ";
      out_chain(buf, gr, lr.expected_tokens(state));
      break;
     }

     if(next>0){//перенос
       state = next;
       stack.push_back(term);
       states.push_back(state);
         if(PARSER_DEBUG){
            cout<<"  <-  "
                <<gr.decode(term)<<endl;
         }
         term = getTerm();
         if (!term){
                    buf << "Lexis: unknown token!";
                    break;
                    }
         continue;
     }//перенос
//свертка
      tGramma::tRule descr = tLR::unpack(next);
      const tGramma::tAlt& alt = gr.getAlt(descr);
      size_t n = alt.rp.size();
      for(size_t i=0; i<n; ++i){
        stack.pop_back();
        states.pop_back();
      }
      tSymb left = descr.left;
      state = lr.go(states.back(), left);
  if(PARSER_DEBUG)
        out_prod(cout,gr,descr);

// заменить основу символом левой части
      stack.push_back(left);
      states.push_back(state);
// проверить условие допустимости цепочки
      if(stack.size() == 2 &&
         left == start &&
         term == 1){// маркер коца
                     lex.End();
                     return 0;
        }
   }
//+++++++++++++++++++++++++++++++
// добавить к сообщению об ошибке номер
// строки и смещение
     buf<< endl;
     buf<< setw(4) << 
         lex.GetLineCount()<<"|"<<
         lex.GetLineText()<< endl;
     buf<< "     " << 
//          setw(1+lex.GetStartPos()) << "^"
          string(lex.GetStartPos(),' ') << "^"
        << endl;
     ferror_message += buf.str();

  if(PARSER_DEBUG){
    cout <<   "Stack:";
    for( size_t i=0; i<stack.size(); ++i)
                  cout<<" "<<gr.decode(stack[i]);
    cout<<"  <-  "<<gr.decode(term)<<endl;
    }
     
     lex.End();
     return 1;
}
