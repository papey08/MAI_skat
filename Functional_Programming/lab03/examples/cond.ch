year = 2023
if year == 2023 {
	println("2023")
}

year = 2020
if year >= 2023 {
    println("more or equal")
} else {
    println("less")
}

year = 2025
if year <= 2023 {
    println("less of equal")
} else {
    println("more")
}

year = 2020
if year > 2023 {
    println("more")
} else {
    println("less or equal")
}

year = 2025
if year < 2023 {
    println("less")
} else {
    println("more or equal")
}

year = 2020
if year > 2000 && year < 2100 {
    println("21st century")
}

year = 2020
if (year == 2016 || year == 2020 || year == 2024) && (year != 2100) {
    println("leap year")
}
