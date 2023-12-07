import Foundation

if let input = readLine() {
    let n = input.components(separatedBy: " ")
    if let a = Int(n[0]), let b = Int(n[1]) {
        print(a + b)
    }
}
