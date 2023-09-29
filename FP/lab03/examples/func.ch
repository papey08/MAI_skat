min = chich(a, b) {
    if a < b {
        return a
    }
    b
}

max = chich(a, b) {
    if a > b {
        return a
    }
    b
}

comp = chich(a, b, f) {
    f(a, b)
}

println(comp(1, 2, min))
println(comp(1, 2, max))
