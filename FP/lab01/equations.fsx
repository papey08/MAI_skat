// Define functions to solve algebraic equations

let rec dichotomy f (a: float) (b: float) =
    let eps = 0.000001
    let xn = (a + b) / 2.
    let fa = f a
    let fb = f b

    if abs (f xn) < eps then
        xn
    else if fa < fb then
        if (f xn) < 0. then dichotomy f xn b else dichotomy f a xn
    else 
        if (f xn) > 0. then dichotomy f xn b else dichotomy f a xn


let rec iterations phi x0 =
    let eps = 0.000001

    if abs (x0 - (phi x0)) < eps then
        x0
    else
        let next = phi x0
        iterations phi next


let newthon f f' x0 =
    let phi x : float = x - (f x) / (f' x)
    iterations phi x0
// make sure to use function 'iterations' here

// Solve 3 equations using three methods defined above
let f1 x : float = 0.1 * (x ** 2.) - x * System.Math.Log x // #20

let f2 x : float = // #21
    System.Math.Tan x - ((System.Math.Tan x) ** 3.) / 3.
    + ((System.Math.Tan x) ** 5.) / 5.
    - (1. / 3.)

let f3 x : float = // #22
    System.Math.Acos x - sqrt (1. - 0.3 * (x ** 3.))


let f1' x : float = 0.2 * x - System.Math.Log x - 1. // #20

let f2' x : float = // #21
    0.125 * (3. * System.Math.Cos 4. * x + 5.) * ((System.Math.Cos x) ** (-6.))

let f3' x : float = // #22
    (0.45 * (x ** 2.)) / (sqrt (1. - 0.3 * (x ** 3.))) - 1. / (sqrt (1. - x ** 2.))


let phi1 x : float = System.Math.E ** (0.1 * x) // #20

let phi2 x : float = // #21
    System.Math.Atan((System.Math.Tan(x) ** 3.) / 3. - (System.Math.Tan(x) ** 5.) / 5. + 1. / 3.)

let phi3 x : float = // #22
    System.Math.Cos(sqrt (1. - 0.3 * (x ** 3.)))

let main =
    printfn "%10.5f  %10.5f  %10.5f" (dichotomy f1 1. 2.) (iterations phi1 1.5) (newthon f1 f1' 1.5)
    printfn "%10.5f  %10.5f  %10.5f" (dichotomy f2 0. 0.8) (iterations phi2 0.4) (newthon f2 f2' 0.4)
    printfn "%10.5f  %10.5f  %10.5f" (dichotomy f3 0. 1.) (iterations phi3 0.5) (newthon f3 f3' 0.5)

main
