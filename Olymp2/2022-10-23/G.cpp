#include <iostream>
#include <vector>

int gcd(int a, int b) {
    while (a != 0) {
        b %= a;
        std::swap(a, b);
    }
    return b;
}

int main(){
    int N, M;
    std::cin >> N >> M;
    int maxSize = gcd(N, M);
    if(maxSize == 1){
        std::cout << "NO";
        return 0;
    }
    std::cout << "YES\n";
    int minSize = maxSize;
    for(int i = 2; i < maxSize; ++i){
        if(N % i == 0 && M % i == 0){
            minSize = i;
            break;
        }
    }
    //std::cout << minSize << "\n";
    int amountInRow = M / minSize;
    int amountInColumn = N / minSize;
    // устанавливаем ураганы
    std::vector<std::vector<bool>> isX(N, std::vector<bool>(M));
    //std::cout << "aft vectors\n";
    bool beginWithB = true;
    for(int i = 0; i < N; i = i + minSize){
        for(int j = (beginWithB ? minSize : 0); j < M; j = j + 2 * minSize){
            //std::cout << i << ',' << j << std::endl;
            for(int k = i; k < i + minSize; ++k){
                for(int l = j; l < j + minSize; ++l){
                    isX[k][l] = true;
                }
            }
        }
        beginWithB = !beginWithB;
    }

    for(int i = 0; i < N; ++i){
        for(int j = 0; j < M; ++j){
            std::cout << (isX[i][j] ? 'X' : 'B');
        }
        std::cout << std::endl;
    }
    return 0;
}
