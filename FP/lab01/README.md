[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-c66648af7eb3fe8bc4f294546bfd86ef473780cde1dea487d3c4ff354943c9ae.svg)](https://classroom.github.com/online_ide?assignment_repo_id=10292812&assignment_repo_type=AssignmentRepo)
# FP: Lab 1. Computing Things Functionally

## Solving simple computational problems in functional style

**Objective:** Getting acquainted with IDE and functional programming language, solving simple iterative tasks in functional style.

## Task 0: Get Your F# Setup

Make sure you have configured your F# environment and are able to execute code. All options to set up F# on your operating system are described at [F# Official Web Site](https://fsharp.org/) [**USE** menu option]. We recommend the following options (in the recommended order):
* Install [Visual Studio Code](https://code.visualstudio.com/) with [Ionide Extension](http://ionide.io/)
* Install [Full Visual Studio](http://visualstudio.com) with F# support (then use **F# Interactive** window to execute files line-by-line interactively)
* Set up local [Jupyter Notebook environment with .NET Support](https://github.com/dotnet/interactive/blob/main/docs/NotebooksLocalExperience.md)
* [Run F# Online in the browser](https://fsharp.org/use/browser/) (for example, using [Repl.it](http://repl.it))

Once you are able to execute F# scripts, make the assignments below.

## Task 1: Taylor Series 

Write a program in a functional language that will print the table of values of some elementary function **f(x)** on the interval **[a,b]**, calculated in the following ways:
 
 * Using built-in functions
 * Using **dumb Taylor series**, where each term is calculated according to the formula
 * Using **smart Taylor series**, where each term is calculated from the previous one (this would yield better efficiency)

In both cases, you should continue adding terms to the series until next term becomes sufficiently small (eg. smaller than a given value **eps**). You should print the number of terms in the table as well.

Here is how the output table should look like:

| x | Builtin | Smart Taylor | # terms | Dumb Taylor | # terms |
|---|---------|--------------|---------|-------------|---------|
| a | ... | ... | ... | ... |
| ... |
| b | ... | ... | ... | ... |

Actual functions and corresponding Taylor series are given in the [PDF Handout](Lab1.pdf). Your personal task is defined by your number in the class list (please refer to the teacher for specific instructions where to find this list).

## Task 2: Solving Equations

Develop a function to solve transcendental algebraic equations numerically, using the following [root-finding methods](https://en.wikipedia.org/wiki/Root-finding_algorithms): 

 * [Bisection method](https://en.wikipedia.org/wiki/Bisection_method) (also called *dichotomy method*)
 * [Iterations method](http://www.simumath.com/library/book.html?code=Alg_Equations_Iterations)
 * [Newton's method](https://en.wikipedia.org/wiki/Newton%27s_method)

All equations should be specified as functional parameters, to make functions sufficiently generic. Also, take advantage of the fact that Newton's method is a special case for more generic method of iterations, and make sure to re-use the code. 

Please apply three procedures for three methods to three consecutive equations in the table 2 of [the handout](Lab1.pdf), starting from the row specified by your personal number from the class list. You should produce a table with a total of 9 solutions.
