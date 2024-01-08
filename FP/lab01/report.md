# Lab 1 report

## Functional Programming

### Student: 
|Name        |Group      | # in group list |Mark|
|------------|-----------|-----------------|----|
|Matvey Popov|М8О-308Б-20| 20              |    |

## Solving simple computational problems in functional style

### Task 1

According to my number in group list, I have the following function:
$$f(x) = e^{2x}$$
and it's Taylor expansion
$$f(x) = 1 + {2x \over 1!} + … + {(2x)^n \over n!}$$

Calculation of every member of Taylor series requires factorial operation which 
time complexity is $O(n)$, therefore time complexity of "dumb" scheme is 
$O(n^2)$. The main idea of effective scheme is we don't need to count 
factorial for every member because it's possible to count nth member of Taylor 
series using (n-1)th member:
$$t_n = t_{n-1}{2x \over n}$$
where $t_n$ is nth member of Taylor series. Time complexity of calculating 
every single member is constant now so calculating of the whole series is 
$O(n)$ which is the way better than $O(n^2)$ of the naive method. 


The result of my F# program [taylor.fsx](taylor.fsx)
is the following table:

| x    | Builtin  | Smart Taylor | # terms | Dumb Taylor | # terms |
|------|----------|--------------|---------|-------------|---------|
| 0.10 | 1.221403 | 1.221400     | 5       | 1.221400    | 5       |
| 0.15 | 1.349859 | 1.349858     | 6       | 1.349858    | 6       |
| 0.20 | 1.491825 | 1.491819     | 6       | 1.491819    | 6       |
| 0.25 | 1.648721 | 1.648720     | 7       | 1.648720    | 7       |
| 0.30 | 1.822119 | 1.822113     | 7       | 1.822113    | 7       |
| 0.35 | 2.013753 | 2.013751     | 8       | 2.013751    | 8       |
| 0.40 | 2.225541 | 2.225536     | 8       | 2.225536    | 8       |
| 0.45 | 2.459603 | 2.459602     | 9       | 2.459602    | 9       |
| 0.50 | 2.718282 | 2.718279     | 9       | 2.718279    | 9       |
| 0.55 | 3.004166 | 3.004159     | 9       | 3.004159    | 9       |
| 0.60 | 3.320117 | 3.320115     | 10      | 3.320115    | 10      |

### Task 2

Three equations I've got to solve:
$$0.1x^2 - x\ln x = 0$$
$$\tan x - {1\over 3}\tan^3 x + {1\over 5}\tan^5 x - {1\over 3} = 0$$
$$\arccos x - \sqrt{1 - 0.3x^3} = 0$$

### Bisection method

This method requires boundaries of segment where root of the equation is. Also 
function should be monotonous on this segment. Then we calculate value of the 
function in the middle of this segment and replace right or left boundary of 
segment with it's middle (demands on funcion increases or decreases on the 
segment and value in the middle is positive or negative). With each iteration 
length of the segment would decrease twice and soon we get root of the equation 
with required approximation.

### Iterations method

First of all we need to express $x = \varphi (x)$ from the original equation 
$f(x) = 0$:
$$x = e^{0.1x}$$
$$x = {1 \over 3}\arctan(\tan(x)^3) - {1 \over 5}\arctan(\tan(x)^5) + {1 \over 3}$$
$$x = \cos\sqrt{1 - 0.3x^3}$$
Then we need to calculate $x_n = \varphi(x_{n - 1})$ choosing $x_0$ close 
to the original root. The more is $n$ — the better precision of solution.

### Newton method

Newton method is a special case of iterations method, where we choose 
$\varphi(x) = x - {f(x)\over f'(x)}$. We calculate differential of 
every function:
$$f'(x) = 0.2x - \ln x - 1$$
$$f'(x) = 0.125(3\cos4x + 5)(\cos x)^{-6}$$
$$f'(x) = {0.45x^2\over \sqrt{1 - 0.3x^3}} - {1\over \sqrt{1 - x^2}}$$

The result of my F# program [equations.fsx](equations.fsx)
is the following table:

|# |Bisection|Iterations|Newton |
|--|---------|----------|-------|
|20| 1.11833 | 1.11833  |1.11833|
|21| 0.33326 | 0.33326  |0.33326|
|22| 0.56293 | 0.56293  |0.56293|

### Conclution

After doing the laboratory work, I got acquainted with functional programming 
using the F# language as an example, and also remembered how to find the value 
of a function using its expansion in a Taylor series and three methods for 
solving equations: bisection method, iterations method and Newton method.
