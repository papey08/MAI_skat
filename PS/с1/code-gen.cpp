/* $b20 */
#include "code-gen.h"
using namespace std;
void tCG::init(){declarations.clear();
 Authentication = "PMR";
//                  ^ 
// replace with your initials!!! 
}

int tCG::p01(){ // S -> PROG
  string header ="/*  " + Authentication +"   */\n";
  header += "#include \"mlisp.h\"\n";
  header += declarations;
  header += "//________________ \n";
  S1->obj = header + S1->obj;
    return 0;}

int tCG::p02(){ //     PROG -> CALCS
  S1->obj = "int main(){\n " + S1->obj + "std::cin.get();\n\t return 0;\n\t }\n";
  return 0;}

int tCG::p03(){ //     PROG -> DEFS
  S1->obj +=
      "int main(){\n "
      "display(\"No calculations!\");\n\t newline();\n\t "
      "std::cin.get();\n\t return 0;\n\t }\n";
  return 0;}

int tCG::p04(){ //     PROG -> DEFS CALCS
    S1->obj += "int main(){\n ";
    S1->obj += S2->obj;
    S1->obj += " std::cin.get();\n\t return 0;\n\t }\n";
    return 0;}

int tCG::p05(){ //        E -> $id
    S1->obj = decor(S1->name);
    return 0;}

int tCG::p06(){ //        E -> $int
    S1->obj = S1->name + ".";
    return 0;}

int tCG::p07(){ //        E -> $dec
    S1->obj = S1->name;
    return 0;}

int tCG::p08(){ //        E -> AREX
    return 0;}

int tCG::p09(){ //        E -> COND
    return 0;}

int tCG::p10(){ //        E -> EASYLET
    return 0;}

int tCG::p11(){ //        E -> CPROC
    
    return 0;}

int tCG::p12(){ //  EASYLET -> HEASYL E )
    S1->obj += S2->obj + ")";
    return 0;}

int tCG::p13(){ //   HEASYL -> ( let ( )
    S1->obj = "(";
    return 0;}

int tCG::p14(){ //   HEASYL -> HEASYL INTER
    S1->obj += S2->obj;
    S1->obj += (S1->count % 2) ? ", " : "\n\t , ";
    return 0;}

int tCG::p15(){ //     AREX -> HAREX E )
    if ((S1->name == "/") && (S1->count == 0)) {
        S1->obj = "(1. " + S1->obj + " " + S2->obj + ")";
    }
    else {
        S1->obj = "(" + S1->obj + " " + S2->obj + ")";
    }
    return 0;}

int tCG::p16(){ //    HAREX -> ( AROP
    S1->obj = S2->obj;
    S1->name = S2->name;
    return 0;}

int tCG::p17(){ //    HAREX -> HAREX E
    if (S1->count == 0) {
        S1->obj = S2->obj + " " + S1->name;
    }
    else {
        S1->obj = S1->obj + " " + S2->obj + " " + S1->name;
    }
    ++S1->count;
    return 0;}

int tCG::p18(){ //     AROP -> +
    S1->obj = S1->name;
    return 0;}

int tCG::p19(){ //     AROP -> -
    S1->obj = S1->name;
    return 0;}

int tCG::p20(){ //     AROP -> *
    S1->obj = S1->name;
    return 0;}

int tCG::p21(){ //     AROP -> /
    S1->obj = S1->name;
    return 0;}

int tCG::p22(){ //     COND -> ( cond BRANCHES )
    S1->obj = "(" + S3->obj + ")";
    return 0;}

int tCG::p23(){ // BRANCHES -> CLAUS ELSE
    S1->obj = S1->obj + "\n\t: (" + S2->obj + ")";
    return 0;}

int tCG::p24(){ //    CLAUS -> ( BOOL E )
    S1->obj += S2->obj + "\n\t? " + S3->obj + "\n\t ";
    return 0;}

int tCG::p25(){ //     ELSE -> ( else E )
    S1->obj += S3->obj;
    return 0;}

int tCG::p26(){ //      STR -> $str
    S1->obj = S1->name;
    return 0;}

int tCG::p27(){ //      STR -> SCOND
    S1->obj = "(" + S1->obj + ")";
    return 0;}

int tCG::p28(){ //    SCOND -> ( cond SBRANCHES )
    S1->obj = "(" + S3->obj + ")";
    return 0;}

int tCG::p29(){ //SBRANCHES -> SELSE
    return 0;}

int tCG::p30(){ //SBRANCHES -> SCLAUS SBRANCHES
    S1->obj += S2->obj;
    return 0;}

int tCG::p31(){ //   SCLAUS -> ( BOOL STR )
    S1->obj = S2->obj + "\n\t? " + S3->obj + "\n\t: ";
    return 0;}

int tCG::p32(){ //    SELSE -> ( else STR )
    S1->obj = "(" + S3->obj + ")";
    return 0;}

int tCG::p33(){ //    CPROC -> HCPROC )
    S1->obj += ")";
    return 0;}

int tCG::p34(){ //   HCPROC -> ( $id
    S1->obj = decor(S2->name) + "(";
    return 0;}

int tCG::p35(){ //   HCPROC -> HCPROC E
    if (S1->count) {
        S1->obj += (S1->count % 2) ? ", " : "\n\t , ";
    }
    S1->obj += S2->obj;
    ++S1->count;
    return 0;}

int tCG::p36(){ //     BOOL -> $bool
    S1->obj = (S1->name == "#t" ? "true" : "false");
    return 0;}

int tCG::p37(){ //     BOOL -> $idq
    S1->obj = decor(S1->name);
    return 0;}

int tCG::p38(){ //     BOOL -> REL
    return 0;}

int tCG::p39(){ //     BOOL -> ( not BOOL )
    S1->obj = "!(" + S3->obj + ")";
    return 0;}

int tCG::p40(){ //     BOOL -> CPRED
    return 0;}

int tCG::p41(){ //     BOOL -> AND
    return 0;}

int tCG::p42(){ //      AND -> ( and ANDARGS )
    S1->obj = "(" + S3->obj + ")";
    return 0;}

int tCG::p43(){ //  ANDARGS -> BOOL ANDARGS
    S1->obj += " && " + S2->obj;
    return 0;}

int tCG::p44(){ //  ANDARGS -> BOOL
    return 0;}

int tCG::p45(){ //    CPRED -> ( $idq )
    S1->obj += decor(S2->name) + "()";
    return 0;}

int tCG::p46(){ //    CPRED -> ( $idq PDARGS )
    S1->obj = decor(S2->name) + "(" + S3->obj + ")";
    return 0;}

int tCG::p47(){ //   PDARGS -> ARG
    return 0;}

int tCG::p48(){ //   PDARGS -> ARG PDARGS
    S1->obj += (S1->count % 2) ? ", " : "\n\t , ";
    S1->obj += S2->obj;
    return 0;}

int tCG::p49(){ //      ARG -> E
    return 0;}

int tCG::p50(){ //      ARG -> BOOL
    return 0;}

int tCG::p51(){ //      REL -> ( = E E )
    S1->obj = "(" + S3->obj + " == " + S4->obj + ")";
    return 0;}

int tCG::p52(){ //      REL -> ( >= E E )
    S1->obj = "(" + S3->obj + " >= " + S4->obj + ")";
    return 0;}

int tCG::p53(){ //      SET -> HSET E )
    S1->obj += S2->obj;
    return 0;}

int tCG::p54(){ //     HSET -> ( set! $id
    S1->obj = decor(S3->name) + " = ";
    return 0;}

int tCG::p55(){ //  DISPSET -> ( display E )
    S1->obj = "display(" + S3->obj + ")";
    return 0;}

int tCG::p56(){ //  DISPSET -> ( display BOOL )
    S1->obj = "display(" + S3->obj + ")";
    return 0;}

int tCG::p57(){ //  DISPSET -> ( display STR )
    S1->obj = "display(" + S3->obj + ")";
    return 0;}

int tCG::p58(){ //  DISPSET -> ( newline )
    S1->obj = "newline()";
    return 0;}

int tCG::p59(){ //  DISPSET -> SET
    return 0;}

int tCG::p60(){ //    INTER -> DISPSET
    return 0;}

int tCG::p61(){ //    INTER -> E
    return 0;}

int tCG::p62(){ //    CALCS -> CALC
    return 0;}

int tCG::p63(){ //    CALCS -> CALCS CALC
    S1->obj += S2->obj;
    return 0;}

int tCG::p64(){ //     CALC -> E
    S1->obj = "display(" + S1->obj + ");\n\t newline();\n\t ";
    return 0;}

int tCG::p65(){ //     CALC -> BOOL
    S1->obj = "display(" + S1->obj + ");\n\t newline();\n\t ";
    return 0;}

int tCG::p66(){ //     CALC -> STR
    S1->obj = "display(" + S1->obj + ");\n\t newline();\n\t ";
    return 0;}

int tCG::p67(){ //     CALC -> DISPSET
    S1->obj = S1->obj + ";\n\t ";
    return 0;}

int tCG::p68(){ //     DEFS -> DEF
    return 0;}

int tCG::p69(){ //     DEFS -> DEFS DEF
    S1->obj = S1->obj + "\n" + S2->obj;
    return 0;}

int tCG::p70(){ //      DEF -> PRED
    return 0;}

int tCG::p71(){ //      DEF -> VAR
    declarations += "extern " + S1->name + ";\n\t";
    S1->obj = S1->name + " = " + S1->obj + ";\n\t";
    return 0;}

int tCG::p72(){ //      DEF -> PROC
    return 0;}

int tCG::p73(){ //     PRED -> HPRED BOOL )
    S1->obj += S2->obj + ";\n\t }\n";
    return 0;}

int tCG::p74(){ //    HPRED -> PDPAR )
    S1->obj += ")";
    declarations += S1->obj + ";\n\t ";
    S1->obj += "{\n return ";
    S1->count = 0;
    return 0;}

int tCG::p75(){ //    PDPAR -> ( define ( $idq
    S1->obj = "bool " + decor(S4->name) + "/*" + S4->line + "*/ (";
    S1->count = 0;
    return 0;}

int tCG::p76(){ //    PDPAR -> PDPAR $idq
    if (S1->count) {
        S1->obj += (S1->count % 2) ? ", " : "\n\t , ";
    }
    S1->obj += "double " + decor(S2->name);
    ++S1->count;
    return 0;}

int tCG::p77(){ //    PDPAR -> PDPAR $id
    if (S1->count) {
        S1->obj += (S1->count % 2) ? ", " : "\n\t , ";
    }
    S1->obj += "double " + decor(S2->name);
    ++S1->count;
    return 0;}

int tCG::p78(){ //      VAR -> ( define $id VARINI )
    S1->name = "double " + decor(S3->name) + "/*" + S3->line + "*/";
    S1->obj = S4->obj;
    return 0;}

int tCG::p79(){ //   VARINI -> $int
    S1->obj = S1->name + ".";
    return 0;}

int tCG::p80(){ //   VARINI -> $dec
    S1->obj = S1->name;
    return 0;}

int tCG::p81(){ //     PROC -> HPROC E )
    S1->obj += "return\n " + S2->obj + ";\n\t" + "}\n";
    return 0;}

int tCG::p82(){ //    HPROC -> PCPAR )
    S1->obj += ")";
    declarations += S1->obj + ";\n\t ";
    S1->obj += "{\n ";
    return 0;}

int tCG::p83(){ //    HPROC -> HPROC INTER
    S1->obj += S2->obj + ";\n\t ";
    return 0;}

int tCG::p84(){ //    HPROC -> HPROC VAR
    S1->obj += S2->name + "(" + S2->obj + ");\n\t";
    return 0;}

int tCG::p85(){ //    PCPAR -> ( define ( $id
    S1->obj = "double " + decor(S4->name) + "/*" + S4->line + "*/ (";
    S1->count = 0;
    S1->name = S4->name;
    return 0;}

int tCG::p86(){ //    PCPAR -> PCPAR $id
    if (S1->count) {
        S1->obj += (S1->count % 2) ? ", " : "\n\t , ";
    }
    S1->obj += "double " + decor(S2->name);
    ++S1->count;
    return 0;}

//_____________________
int tCG::p87(){return 0;} int tCG::p88(){return 0;} 
int tCG::p89(){return 0;} int tCG::p90(){return 0;} 
int tCG::p91(){return 0;} int tCG::p92(){return 0;} 
int tCG::p93(){return 0;} int tCG::p94(){return 0;} 
int tCG::p95(){return 0;} int tCG::p96(){return 0;} 
int tCG::p97(){return 0;} int tCG::p98(){return 0;} 
int tCG::p99(){return 0;} int tCG::p100(){return 0;} 
int tCG::p101(){return 0;} int tCG::p102(){return 0;} 
int tCG::p103(){return 0;} int tCG::p104(){return 0;} 
int tCG::p105(){return 0;} int tCG::p106(){return 0;} 
int tCG::p107(){return 0;} int tCG::p108(){return 0;} 
int tCG::p109(){return 0;} int tCG::p110(){return 0;} 


