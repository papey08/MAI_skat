//                 code-gen.h 2014
#ifndef CODE_GEN_H
#define CODE_GEN_H
#include <string>
#include "base-compiler.h"

class tCG:public tBC{
public:
//конструктор
 tCG(const char* gramma_name) :tBC(gramma_name){}
private:
 std::string declarations;
 std::string decor(const std::string& id);

protected:
 void init();

  int p01();  int p02();  int p03();  int p04();
  int p05();  int p06();  int p07();  int p08();
  int p09();  int p10();  int p11();  int p12();
  int p13();  int p14();  int p15();  int p16();
  int p17();  int p18();  int p19();  int p20();
  int p21();  int p22();  int p23();  int p24();
  int p25();  int p26();  int p27();  int p28();
  int p29();  int p30();  int p31();  int p32();
  int p33();  int p34();  int p35();  int p36();
  int p37();  int p38();  int p39();  int p40();
  int p41();  int p42();  int p43();  int p44();
  int p45();  int p46();  int p47();  int p48();
  int p49();  int p50();  int p51();  int p52();
  int p53();  int p54();  int p55();  int p56();
  int p57();  int p58();  int p59();  int p60();
  int p61();  int p62();  int p63();  int p64();
  int p65();  int p66();  int p67();  int p68();
  int p69();  int p70();  int p71();  int p72();
  int p73();  int p74();  int p75();  int p76();
  int p77();  int p78();  int p79();  int p80();
  int p81();  int p82();  int p83();  int p84();
  int p85();  int p86();  int p87();  int p88();
  int p89();  int p90();  int p91();  int p92();
  int p93();  int p94();  int p95();  int p96();
  int p97();  int p98();  int p99();  int p100();
  int p101();  int p102();  int p103();  int p104();
  int p105();  int p106();  int p107();  int p108();
  int p109();  int p110();

};
 std::string tCG::decor(const std::string& id){
  static const char* cpp_reserved[]={
   "try",
//????????????????????????
   "main"
  };
  static const size_t n =
              sizeof(cpp_reserved)/sizeof(char*);
  for(size_t i=0; i<n; ++i)
  if( id == cpp_reserved[i])
     return "__"+
            Authentication +
            "__"+id;
// заменить '-','!','?'
 std::string tmp;
 for(size_t i=0; i<id.size(); ++i){
  switch(id[i]){
    case '?': tmp+="_Q"; break;
    case '!': tmp+="_E"; break;
    case '-': tmp+="__"; break;
    default : tmp+=id[i];
  }
 }
 return tmp;
 }
#endif

