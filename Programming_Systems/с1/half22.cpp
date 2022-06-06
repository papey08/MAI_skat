/*  PMR   */
#include "mlisp.h"
double root/*2*/ (double a, double b);
	 double half__interval/*15*/ (double a, double b
	 , double fa, double fb);
	 double __PMR__try/*31*/ (double neg__point, double pos__point);
	 bool close__enough_Q/*48*/ (double x, double y);
	 double average/*50*/ (double x, double y);
	 extern double tolerance/*52*/;
	extern double total__iterations/*53*/;
	double f/*54*/ (double z);
	 //________________ 
double root/*2*/ (double a, double b){
 double temp/*3*/(0.);
	temp = half__interval(a, b
	 , f(a), f(b));
	 display("Total number of iteranions=");
	 display(total__iterations);
	 newline();
	 display("[");
	 display(a);
	 display(" , ");
	 display(b);
	 display("]");
	 return
 temp;
	}

double half__interval/*15*/ (double a, double b
	 , double fa, double fb){
 double root/*16*/(0.);
	total__iterations = 0.;
	 root = ((!((fa >= 0.)) && ((fb >= 0.) && !((fb == 0.))))
	? __PMR__try(a, b)
	 
	: (((((fa >= 0.) && !((fa == 0.))) && !((fb >= 0.)))
	? __PMR__try(b, a)
	 
	: ((b + 1.)))));
	 newline();
	 return
 root;
	}

double __PMR__try/*31*/ (double neg__point, double pos__point){
 double midpoint/*32*/(0.);
	double test__value/*33*/(0.);
	midpoint = average(neg__point, pos__point);
	 return
 (close__enough_Q(neg__point
	 , pos__point)
	? midpoint
	 
	: ((test__value = f(midpoint)
	 , display("+")
	 , total__iterations = (total__iterations + 1.)
	 , (((test__value >= 0.) && !((test__value == 0.)))
	? __PMR__try(neg__point, midpoint)
	 
	: ((!((test__value >= 0.))
	? __PMR__try(midpoint, pos__point)
	 
	: (midpoint)))))));
	}

bool close__enough_Q/*48*/ (double x, double y){
 return !((abs((x - y)) >= tolerance));
	 }

double average/*50*/ (double x, double y){
 return
 ((x + y) * (1. / 2e+0));
	}

double tolerance/*52*/ = 1e-3;
	
double total__iterations/*53*/ = 0.;
	
double f/*54*/ (double z){
 return
 (expt(cos(z), 2.) - expt(sin(z), 2.));
	}
int main(){
 display("Variant 208-20\n");
	 display(root(157e-2, 3e+0));
	 newline();
	 display("(c) Popov Matvey 2022\n");
	  std::cin.get();
	 return 0;
	 }

