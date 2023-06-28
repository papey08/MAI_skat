fib = chich(n) {
	if n < 2 {
		return n
	}
	fib(n-1) + fib(n-2)
}

println(fib(4))
println(fib(5))
println(fib(6))

_tail_fib = chich(n, acc1, acc2) {
    if n < 2 {
        return acc1
    }
    _tail_fib(n-1, acc1 + acc2, acc1)
}

tail_fib = chich(n) {
    _tail_fib(n, 1, 0)
}

println(tail_fib(4))
println(tail_fib(5))
println(tail_fib(6))
