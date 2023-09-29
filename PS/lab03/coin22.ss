; coin22.ss

(define VARIANT 20)
(define COINS 3)

(define (largest coins-set)
    (cond
        ((= coins-set 1) 1)
        ((= coins-set 2) 2)
        ((= coins-set 3) 3)
        (else 0)
    )
)

(define (count-change amount)
    (display "______\n amount: ")
    (display amount)
    (newline)
    (display "COINS: ")
    (display COINS)
    (newline)
    (cond((or (<= amount 0) (< COINS 1) (= (largest COINS) 0)) (display "Improper parameter value!\ncount-change= ") -1)
        (#t (display "List of coin denominations: ") (denomination-list COINS) (display "count-change= ") (cc amount COINS))
    )
)

(define (Shaeffer? x? y?)
    (or (not x?) (not y?))
)

(define (cc amount coins-set)
    (cond
        ((= amount 0) 1)
        ((Shaeffer? (>= amount 1) (> coins-set 0)) 0)
        (else (+ (cc amount (- coins-set 1)) (cc (- amount (largest coins-set)) coins-set)))
    )
)

(define (denomination-list coins-set)
    (cond
        ((= coins-set 0) 0)
        (else (display (largest coins-set))
            (display " ")
            (denomination-list (- coins-set 1))
        -1)
    )
)

(display "Variant ")
(display VARIANT)
(newline)
(display (count-change 100)) (newline)
(set! COINS 13)
(display (count-change 100)) (newline)
(display "(C) Popov Matvey 2022\n")