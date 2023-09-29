/* $m21 */
#include "code-gen.h"
using namespace std;
void tCG::init(){declarations.clear();
 Authentication = "PMR";
//                  ^ 
// should be replaced with your initials!!! 
}
int tCG::p01(){ // S -> PROG
  string header ="/*  " + Authentication +"   */\n";
  header += "#include \"mlisp.h\"\n";
  header += declarations;
  header += "//________________ \n";
  S1->obj = header + S1->obj;
	return 0;}
int tCG::p02(){ //  PROG -> CALCS
 S1->obj = "int main(){\n " + S1->obj
    + "std::cin.get();\n\t return 0;\n\t }\n";
	return 0;}
int tCG::p03(){ //  PROG -> DEFS
 S1->obj += "int main(){\n "
    "display(\"No calculations!\");\n\t newline();\n\t "
    " std::cin.get();\n\t return 0;\n\t }\n";
	return 0;}
int tCG::p04(){ //  PROG -> DEFS CALCS
//?????????
	return 0;}
int tCG::p05(){ // CALCS -> CALC
	return 0;}
int tCG::p06(){ // CALCS -> CALCS CALC
// S1->obj += S2->obj;
	return 0;}
int tCG::p07(){ //  CALC -> E
 S1->obj = "display("+S1->obj+");\n\t newline();\n\t ";
	return 0;}
int tCG::p08(){ //     E -> $id
 S1->obj = decor(S1->name);
	return 0;}
int tCG::p09(){ //     E -> $int
//?????????
	return 0;}
int tCG::p10(){ //     E -> CPROC
	return 0;}
int tCG::p11(){ // CPROC -> HCPROC )
//????
	return 0;}
int tCG::p12(){ //HCPROC -> ( $id
//????
	return 0;}
int tCG::p13(){ //HCPROC -> HCPROC E
//????
	return 0;}
int tCG::p14(){ //   SET -> HSET E )
//????
	return 0;}
int tCG::p15(){ //   SET -> ( set! $id
//????
	return 0;}
int tCG::p16(){ //   DEF -> PROC
	return 0;}
int tCG::p17(){ //  DEFS -> DEF
	return 0;}
int tCG::p18(){ //  DEFS -> DEFS DEF
//????
	return 0;}

int tCG::p19(){ //  PROC -> HPROC E )
 S1->obj += "return\n " + S2->obj+";\n\t }\n";
	return 0;}
int tCG::p20(){ // HPROC -> PCPAR )
 S1->obj += ")";
 declarations += S1->obj + ";\n"; //!!!
 S1->obj += "{\n ";
	return 0;}
int tCG::p21(){ // HPROC -> HPROC SET
//????
	return 0;}
int tCG::p22(){ // PCPAR -> ( define ( $id
 S1->obj =  "double " + decor(S4->name) +
 "/*" + S4->line + "*/ (";
 S1->count = 0;
 S1->name = S4->name;
	return 0;}
int tCG::p23(){ // PCPAR -> PCPAR $id
 if(S1->count)S1->obj += S1->count%2 ? ", " : "\n\t , ";
 S1->obj += "double " + decor(S2->name);
 ++(S1->count);
	return 0;}
//_____________________
int tCG::p24(){return 0;} int tCG::p25(){return 0;} 
int tCG::p26(){return 0;} int tCG::p27(){return 0;} 
int tCG::p28(){return 0;} int tCG::p29(){return 0;} 
int tCG::p30(){return 0;} int tCG::p31(){return 0;} 
int tCG::p32(){return 0;} int tCG::p33(){return 0;} 
int tCG::p34(){return 0;} int tCG::p35(){return 0;} 
int tCG::p36(){return 0;} int tCG::p37(){return 0;} 
int tCG::p38(){return 0;} int tCG::p39(){return 0;} 
int tCG::p40(){return 0;} int tCG::p41(){return 0;} 
int tCG::p42(){return 0;} int tCG::p43(){return 0;} 
int tCG::p44(){return 0;} int tCG::p45(){return 0;} 
int tCG::p46(){return 0;} int tCG::p47(){return 0;} 
int tCG::p48(){return 0;} int tCG::p49(){return 0;} 
int tCG::p50(){return 0;} int tCG::p51(){return 0;} 
int tCG::p52(){return 0;} int tCG::p53(){return 0;} 
int tCG::p54(){return 0;} int tCG::p55(){return 0;} 
int tCG::p56(){return 0;} int tCG::p57(){return 0;} 
int tCG::p58(){return 0;} int tCG::p59(){return 0;} 
int tCG::p60(){return 0;} int tCG::p61(){return 0;} 
int tCG::p62(){return 0;} int tCG::p63(){return 0;} 
int tCG::p64(){return 0;} int tCG::p65(){return 0;} 
int tCG::p66(){return 0;} int tCG::p67(){return 0;} 
int tCG::p68(){return 0;} int tCG::p69(){return 0;} 
int tCG::p70(){return 0;} int tCG::p71(){return 0;} 
int tCG::p72(){return 0;} int tCG::p73(){return 0;} 
int tCG::p74(){return 0;} int tCG::p75(){return 0;} 
int tCG::p76(){return 0;} int tCG::p77(){return 0;} 
int tCG::p78(){return 0;} int tCG::p79(){return 0;} 
int tCG::p80(){return 0;} int tCG::p81(){return 0;} 
int tCG::p82(){return 0;} int tCG::p83(){return 0;} 
int tCG::p84(){return 0;} int tCG::p85(){return 0;} 
int tCG::p86(){return 0;} int tCG::p87(){return 0;} 
int tCG::p88(){return 0;} int tCG::p89(){return 0;} 
int tCG::p90(){return 0;} int tCG::p91(){return 0;} 
int tCG::p92(){return 0;} int tCG::p93(){return 0;} 
int tCG::p94(){return 0;} int tCG::p95(){return 0;} 
int tCG::p96(){return 0;} int tCG::p97(){return 0;} 
int tCG::p98(){return 0;} int tCG::p99(){return 0;} 
int tCG::p100(){return 0;} int tCG::p101(){return 0;} 
int tCG::p102(){return 0;} int tCG::p103(){return 0;} 
int tCG::p104(){return 0;} int tCG::p105(){return 0;} 
int tCG::p106(){return 0;} int tCG::p107(){return 0;} 
int tCG::p108(){return 0;} int tCG::p109(){return 0;} 
int tCG::p110(){return 0;} 

