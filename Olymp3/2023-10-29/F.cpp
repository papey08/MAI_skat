#include <iostream>
#include <vector>

struct member {
    int favourite;
    int ind;
};

int main() {
    int n;
    std::cin >> n;
    std::vector<member> members(n);
    
    for (int i = 0; i < n; i++) {
        std::cin >> members[i].favourite;
        members[i].ind = i;
    }
    int t = 0;
    int cur = 0;
    while (members.size() != 1) {
        std::vector<member> temp_vec(n-cur-1);
        int temp = temp_vec.size() + 1;
        t = (t+members[t].favourite+temp-1) % temp;
        if (t > 0) {
            std::move(members.begin(), members.begin()+t, temp_vec.begin());
        }
        std::move(members.begin()+t+1, members.end(), temp_vec.begin()+t);
        t = (t >= temp-1) ? 0 : t;
        members = temp_vec;
        ++cur;
    }
    std::cout << members[0].ind + 1 << std::endl;
    return 0;
}
