//   BASE-COMPILER.h 2021
#ifndef BASE_COMPILER_H
#define BASE_COMPILER_H
#include <iostream>
#include <iomanip>
#include <string>
#include <vector>
#include "lexer.h"
#include "slr.h"

class tBC{
public:
 typedef tGramma::tSymb tSymb;
 typedef tLR::tState tState;
 tBC(const char* gramma_name);
 virtual ~tBC();
 int rewrite(const char* source_name);
 const std::string& getObject()const
                      {return fobject;}
 const std::string& getMessage()const
                    {return ferror_message;}
 const tGramma& getGramma()const{return gr;}
 operator bool()const{return gr;}
 bool PARSER_DEBUG;
 std::string Authentication;

 static std::string Uint_to_str(unsigned int);
 tLexer       lex;


protected:
 std::string  ferror_message;
 std::string  fobject;
 virtual void init(){}

// атрибуты
 struct tSA{
  std::string line;
  std::string name;
  int    count;
  int    types;
  std::string obj;
//constructor
  tSA(
     std::string aline=std::string(),
     std::string aname=std::string(),
     int         acount=0,
     int         atypes=0,
     std::string aobj=std::string()):line(aline),
     obj(aobj),name(aname),count(acount),types(atypes){}

  void print(){
   std::cout << "\t["<< std::setw(2) << line <<"| " 
       <<name<<"|"
       << count <<"| "
       << std::hex << types << std::dec <<"| "
       << obj <<" ]\n";
  }
 };//tSA

 tSA *S1,*S2,*S3,*S4,*S5,*S6;

private:
 tGramma            gr;
 tLR                lr;
 std::vector<tSA> ast;// стeк атрибутов

 tBC(const tBC& );
 tBC& operator=(const tBC& );
 tGramma::tSymb getTerm();

// таблица указателей на 
// семантические подпрограмммы
 typedef int(tBC::* tSemPointer)();
 std::vector<tSemPointer> links;

 int call_sem(size_t production){
  if(production < 1 || production >= links.size())
                           return 0;
  tSemPointer sptr=links[production];
  return sptr ? 
//   вызов функции класса через указатель
                (this->*sptr)()
              : 0;
 }
 void make_links(){
  tSemPointer tmp[]={0
  ,&tBC::p01,&tBC::p02,&tBC::p03,&tBC::p04,&tBC::p05
  ,&tBC::p06,&tBC::p07,&tBC::p08,&tBC::p09,&tBC::p10
  ,&tBC::p11,&tBC::p12,&tBC::p13,&tBC::p14,&tBC::p15
  ,&tBC::p16,&tBC::p17,&tBC::p18,&tBC::p19,&tBC::p20
  ,&tBC::p21,&tBC::p22,&tBC::p23,&tBC::p24,&tBC::p25
  ,&tBC::p26,&tBC::p27,&tBC::p28,&tBC::p29,&tBC::p30
  ,&tBC::p31,&tBC::p32,&tBC::p33,&tBC::p34,&tBC::p35
  ,&tBC::p36,&tBC::p37,&tBC::p38,&tBC::p39,&tBC::p40
  ,&tBC::p41,&tBC::p42,&tBC::p43,&tBC::p44,&tBC::p45
  ,&tBC::p46,&tBC::p47,&tBC::p48,&tBC::p49,&tBC::p50
  ,&tBC::p51,&tBC::p52,&tBC::p53,&tBC::p54,&tBC::p55
  ,&tBC::p56,&tBC::p57,&tBC::p58,&tBC::p59,&tBC::p60
  ,&tBC::p61,&tBC::p62,&tBC::p63,&tBC::p64,&tBC::p65
  ,&tBC::p66,&tBC::p67,&tBC::p68,&tBC::p69,&tBC::p70
  ,&tBC::p71,&tBC::p72,&tBC::p73,&tBC::p74,&tBC::p75
  ,&tBC::p76,&tBC::p77,&tBC::p78,&tBC::p79,&tBC::p80
  ,&tBC::p81,&tBC::p82,&tBC::p83,&tBC::p84,&tBC::p85
  ,&tBC::p86,&tBC::p87,&tBC::p88,&tBC::p89,&tBC::p90
  ,&tBC::p91,&tBC::p92,&tBC::p93,&tBC::p94,&tBC::p95
  ,&tBC::p96,&tBC::p97,&tBC::p98,&tBC::p99,&tBC::p100
  ,&tBC::p101,&tBC::p102,&tBC::p103,&tBC::p104,&tBC::p105
  ,&tBC::p106,&tBC::p107,&tBC::p108,&tBC::p109,&tBC::p110
  };
     links =
//  Вектор конструируется из массива
    std::vector<tSemPointer>(tmp, tmp +
           sizeof(tmp)/sizeof(tSemPointer));
 }
protected:

 virtual int p01()=0; virtual int p02()=0;
 virtual int p03()=0; virtual int p04()=0;
 virtual int p05()=0; virtual int p06()=0;
 virtual int p07()=0; virtual int p08()=0;
 virtual int p09()=0; virtual int p10()=0;
 virtual int p11()=0; virtual int p12()=0;
 virtual int p13()=0; virtual int p14()=0;
 virtual int p15()=0; virtual int p16()=0;
 virtual int p17()=0; virtual int p18()=0;
 virtual int p19()=0; virtual int p20()=0;
 virtual int p21()=0; virtual int p22()=0;
 virtual int p23()=0; virtual int p24()=0;
 virtual int p25()=0; virtual int p26()=0;
 virtual int p27()=0; virtual int p28()=0;
 virtual int p29()=0; virtual int p30()=0;
 virtual int p31()=0; virtual int p32()=0;
 virtual int p33()=0; virtual int p34()=0;
 virtual int p35()=0; virtual int p36()=0;
 virtual int p37()=0; virtual int p38()=0;
 virtual int p39()=0; virtual int p40()=0;
 virtual int p41()=0; virtual int p42()=0;
 virtual int p43()=0; virtual int p44()=0;
 virtual int p45()=0; virtual int p46()=0;
 virtual int p47()=0; virtual int p48()=0;
 virtual int p49()=0; virtual int p50()=0;
 virtual int p51()=0; virtual int p52()=0;
 virtual int p53()=0; virtual int p54()=0;
 virtual int p55()=0; virtual int p56()=0;
 virtual int p57()=0; virtual int p58()=0;
 virtual int p59()=0; virtual int p60()=0;
 virtual int p61()=0; virtual int p62()=0;
 virtual int p63()=0; virtual int p64()=0;
 virtual int p65()=0; virtual int p66()=0;
 virtual int p67()=0; virtual int p68()=0;
 virtual int p69()=0; virtual int p70()=0;
 virtual int p71()=0; virtual int p72()=0;
 virtual int p73()=0; virtual int p74()=0;
 virtual int p75()=0; virtual int p76()=0;
 virtual int p77()=0; virtual int p78()=0;
 virtual int p79()=0; virtual int p80()=0;
 virtual int p81()=0; virtual int p82()=0;
 virtual int p83()=0; virtual int p84()=0;
 virtual int p85()=0; virtual int p86()=0;
 virtual int p87()=0; virtual int p88()=0;
 virtual int p89()=0; virtual int p90()=0;
 virtual int p91()=0; virtual int p92()=0;
 virtual int p93()=0; virtual int p94()=0;
 virtual int p95()=0; virtual int p96()=0;
 virtual int p97()=0; virtual int p98()=0;
 virtual int p99()=0; virtual int p100()=0;
 virtual int p101()=0; virtual int p102()=0;
 virtual int p103()=0; virtual int p104()=0;
 virtual int p105()=0; virtual int p106()=0;
 virtual int p107()=0; virtual int p108()=0;
 virtual int p109()=0; virtual int p110()=0;
};//tBC


#endif
