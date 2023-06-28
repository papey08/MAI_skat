// Solution

// #10 Define if list is a geometrical progression

// Method 1: Library Function
let isGeometricProgression1 (lst: float list) =
    match lst with
    | [] | [_] -> true
    | x::y::t -> 
        let q = y / x
        lst |> List.pairwise |> List.forall (fun (a, b) -> b / a = q)


// Method 2: Recursion
let rec isGeometricProgression2 (lst: float list) = 
    match lst with
    | [] | [_] | [_;_] -> true
    | x::y::z::t -> 
        if y / x = z / y then
            isGeometricProgression2 ([y]@[z]@t)
        else
            false


// Method 3: Tail Rec
let rec tailRec (acc: float) (lst: float list) =
    match lst with
    | [] | [_] -> true
    | x::y::t -> 
        let q = y / x
        if q = acc then
            tailRec q (y::t)
        else
            false

let isGeometricProgression3 (lst: float list) =
    match lst with
    | [] | [_] -> true
    | x::y::t -> tailRec (y / x) (y::t)
