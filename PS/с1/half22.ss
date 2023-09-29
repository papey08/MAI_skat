;half22 for 208
(define (root a b)
 (define temp 0)
 
 (set! temp (half-interval a b (f a)(f b)))
  (display"Total number of iteranions=")
  (display total-iterations)(newline)
 (display"[")
 (display a)
 (display" , ")
 (display b)
 (display"]")
      temp 
)
(define (half-interval a b fa fb)
 (define root 0)
 (set! total-iterations 0)
   (set! root(cond
               ((and(not(>= fa 0))(and (>= fb 0) (not(= fb 0))))      (try a b))
               (else (cond
                         ((and(and (>= fa 0) (not(= fa 0)))(not(>= fb 0)))      (try b a))        
                         (else(+ b 1)))
                         )
                    )
               )
               (newline)
     root
    )
  

(define(try neg-point pos-point)
 (define midpoint 0)
 (define test-value 0)
     (set! midpoint (average neg-point pos-point))
     (cond((close-enough? neg-point pos-point) midpoint)
          (else 
               
               (let() (set! test-value (f midpoint)) 
               (display "+")
               (set! total-iterations (+ total-iterations 1))
               (cond((and (>= test-value 0) (not(= test-value 0)))(try neg-point midpoint))
                    (else (cond((not(>= test-value 0))(try midpoint pos-point)) (else midpoint)))
                    
               ))
          )
     )
)
(define (close-enough? x y)
  (not(>=(abs (- x y))tolerance)))
(define (average x y)(*(+ x y)(/  2e+0)))

(define tolerance 1e-3)
(define total-iterations 0)
(define(f z)
  (- (expt (cos z) 2) (expt (sin z) 2))
  )
 (display"Variant 208-20\n")
;      a b
 (root 157e-2 3e+0)
 (display"(c) Popov Matvey 2022\n")
 