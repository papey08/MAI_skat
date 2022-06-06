; t6
(define(try x)(set! x(f x 2))x)
(define(f x y)(expt x y))
(try 3)
