// Print a table of a given function f, computed by taylor series

// function to compute
let f x : float = float (System.Math.E ** (2. * x))

let a = 0.1
let b = 0.6
let n = 10

let rec fact' acc =
    function
    | n when n = 0 -> acc
    | n when n = 1 -> acc
    | n -> fact' (acc * n) (n - 1)

let fact x = fact' 1 x


let rec taylor_naive' (x: float) n (acc: float) =
    let eps = 0.000001

    if abs (f x - acc) < eps then
        acc, n
    else
        let newAcc = acc + ((2. * x) ** float (n)) / float (fact n)
        taylor_naive' x (n + 1) newAcc


// Define a function to compute f using naive taylor series method
let taylor_naive (x: float) = taylor_naive' x 0 0.


let rec taylor' (x: float) n (prev: float) (acc: float) =
    let eps = 0.000001

    if abs (f x - acc) < eps then
        acc, n
    else
        let newAcc = acc + prev * 2. * x / float n
        let newPrev = prev * 2. * x / float n
        taylor' x (n + 1) newPrev newAcc

// Define a function to do the same in a more efficient way
let taylor (x: float) = taylor' x 1 1. 1.

let main =
    for i = 0 to n do
        let x = a + (float i) / (float n) * (b - a)
        let x1, n1 = taylor_naive x
        let x2, n2 = taylor x
        printfn "%5.2f    %10.6f  %10.6f %d  %10.6f %d" x (f x) x1 n1 x2 n2
// make sure to improve this table to include the required number of iterations
// for each of the methods

main
