def arabic_to_roman(arabic):
    roman_numerals = {
        1: 'I',
        4: 'IV',
        5: 'V',
        9: 'IX',
        10: 'X',
        40: 'XL',
        50: 'L',
        90: 'XC',
        100: 'C',
        400: 'CD',
        500: 'D',
        900: 'CM',
        1000: 'M'
    }

    roman_numeral = ""
    for value, numeral in sorted(roman_numerals.items(), key=lambda x: x[0], reverse=True):
        while arabic >= value:
            roman_numeral += numeral
            arabic -= value

    return roman_numeral

def roman_to_arabic(roman):
    roman_numerals = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000
    }

    arabic_numeral = 0
    prev_value = 0

    for numeral in reversed(roman):
        value = roman_numerals[numeral]

        if value < prev_value:
            arabic_numeral -= value
        else:
            arabic_numeral += value

        prev_value = value

    return arabic_numeral

roman_nums = []

for i in range(1, 4000):
    roman_nums.append(arabic_to_roman(i))

roman_nums.sort()
t = int(input())

for _ in range(t):
    roman_number = input()
    idx = roman_to_arabic(roman_number)
    print(roman_nums[idx-1])
