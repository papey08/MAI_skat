//                 BASE-COMPILER.cpp 2020
#include <sstream>
#include "base-compiler.h"
using namespace std;

std::string tBC::Uint_to_str(unsigned int n)
 {
  std::string s;
  while(n){
  s = std::string(1,n%10+'0')+ s;
  n /= 10;
  }
  return s.empty() ? "0" : s;
 }
//...
tGramma::tSymb tBC::getTerm(){
 std::string token=lex.GetToken();
 tGramma::tSymb term = gr.encode(token);

//  из токена $id извлекаютс€ ключевые слова
 if(token == "$id"){
   std::string ident = lex.GetLexeme();
   tGramma::tSymb keyword=gr.encode(ident);//поиск в алфавите
// если идентификатор найден и €вл€етс€ терминалом,
   if(keyword && gr.terminal(keyword))
//                        это ключевое слово
                          term = keyword;
  }
 return term;
}

 tBC::tBC(const char* gramma_name):lr(gr)
 {
  S1=S2=S3=S4=S5=S6=0;
  PARSER_DEBUG=false;
 Authentication = "???";
  gr.loadFromFile(gramma_name);
  if(!gr){
   ferror_message = gr.getError();
      return;
  }
  SLRbuild(lr);
  make_links();
}
 tBC::~tBC(){return;}

int tBC::rewrite(const char* source_name){
//+++++++++++++++++++
  fobject.clear();//+
   init();           //+
   ast.clear();      //+
   int stepcount=0;
   tSA atr;
//+++++++++++++++++++
  if(lr.size()==0){
// испорчены управл€ющие таблицы
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
   ast.push_back(atr);

   term = getTerm();
   if (!term){
     buf << "Lexis: unknown token!\n";
     }
//+++++++++++++++++++++++++++++++
   else while(true){
     next = lr.go(state,term);
     if(next==0){
       buf << "Syntax: unmatched token "
           << gr.decode(term) << "\nexpected: ";
      out_chain(buf, gr, lr.expected_tokens(state));
      buf<<endl;
      break;
     }

     if(next>0){//перенос
       state = next;
       stack.push_back(term);
       states.push_back(state);
//+++++++++++++++++++++++++++++++++++++
         atr = tSA(Uint_to_str(lex.GetLineCount()),lex.GetLexeme());
         ast.push_back(atr);
         if(PARSER_DEBUG){
            cout<<"  <-  "
                <<gr.decode(term)<<endl;
           cout<< setw(3)<< ++stepcount;
           atr.print();
         }
//+++++++++++++++++++++++++++++
         term = getTerm();
         if (!term){
                    buf << "Lexis: unknown token!\n";
                    break;
                    }
         continue;
     }//перенос
//свертка
      tGramma::tRule descr = tLR::unpack(next);
      const tGramma::tAlt& alt = gr.getAlt(descr);
      size_t k = alt.rp.size();
      for(size_t i=0; i<k; ++i){
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
//+++++++++++++++++++++++++++++++++++++++++++++++++
// вызов семантической подпрограммы       	//+
  S1=S2=S3=S4=S5=S6=0;
  int base = ast.size()-k;
  switch(k){
  case 6: S6 = &ast[base+5];
  case 5: S5 = &ast[base+4];
  case 4: S4 = &ast[base+3];
  case 3: S3 = &ast[base+2];
  case 2: S2 = &ast[base+1];
  case 1: S1 = &ast[base];
  }
   if(call_sem(alt.hndl)) break;        //+
  --k;						//+
   for(int i=0; i< k; ++i) ast.pop_back();	//+
//+++++++++++++++++++++++++++++++++++++++++++++++++
// проверить условие допустимости цепочки
      if(stack.size() == 2 &&
         left == start &&
         term == 1){// маркер коца
//++++++++++++++++++++++++++++++++++++
          fobject = ast.back().obj;//+
//++++++++++++++++++++++++++++++++++++
                     lex.End();
                     return 0;
        }
     if(PARSER_DEBUG){
      cout<< setw(3)<< ++stepcount;
      ast.back().print();
      }
   }
//+++++++++++++++++++++++++++++++
     buf<< setw(4) << 
         lex.GetLineCount()<<"|"<<
         lex.GetLineText()<< endl;
     buf<< "     " << 
          setw(1+lex.GetStartPos()) << "^"
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
