/*  PMR   */
#include "mlisp.h"
double day__of__week/*2*/ ();
	 double zeller/*11*/ (double d, double m
	 , double y, double c);
	 double neg__to__pos/*21*/ (double d);
	 double birthday/*26*/ (double dw);
	 extern double dd/*44*/;
	extern double mm/*45*/;
	extern double yyyy/*46*/;
	//________________ 
double day__of__week/*2*/ (){
 return
 zeller(dd, (8. + (!(((mm >= 2.) && !((mm == 2.))))
	? (mm + 2.)
	 
	: ((mm - 010.))))
	 , remainder((1. + (!((mm >= 3.))
	? (yyyy - 2.)
	 
	: ((yyyy - 1.)))), 100.), quotient(((2. >= mm)
	? (yyyy - 1.)
	 
	: (yyyy)), 100.));
	}

double zeller/*11*/ (double d, double m
	 , double y, double c){
 return
 neg__to__pos(remainder((d + y + quotient(((26. * m) - 2.), 10.) + quotient(y, 4.) + quotient(c, 4.) + (2. * (- c))), 7.));
	}

double neg__to__pos/*21*/ (double d){
 return
 (!((d >= 0.))
	? (d + 7.)
	 
	: (d));
	}

double birthday/*26*/ (double dw){
 display("Matvey Popov was born on ");
	 display((((dw == 0.)
	? "Sunday "
	: (dw == 1.)
	? "Monday "
	: (dw == 2.)
	? "Tuesday "
	: (dw == 3.)
	? "Wednesday "
	: (dw == 4.)
	? "Thursday "
	: (dw == 5.)
	? "Friday "
	: ("Saturday "))));
	 display(dd);
	 display(".");
	 display(mm);
	 display(".");
	 return
 yyyy;
	}

double dd/*44*/ = 29.;
	
double mm/*45*/ = 11.;
	
double yyyy/*46*/ = 2002.;
	int main(){
 display(birthday(day__of__week()));
	 newline();
	  std::cin.get();
	 return 0;
	 }

