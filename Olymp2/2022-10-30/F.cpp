#include <iostream>
#include <string>
 
int main() {
    std::string input;
    std::getline(std::cin, input);
    while((input[0] == ' ' || input[0] == '0') && input.size() != 1){
        input.erase(0, 1);
        //std::cout << input << "\n";
    }
    for(char ch : input){
        if(int(ch) < 48 || int(ch) > 57){
            std::cout << "invalid input";
            return 0;
        }
    }
 
    std::cout << input;
    return 0;
}