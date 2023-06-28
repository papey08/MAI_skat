#include <iostream>
#include <set>

struct ttime {
    int hh;
    int mm;
};

ttime int_to_time(int m) {
    ttime res;
    res.hh = m / 60;
    res.mm = m % 60;
    return res;
}

int time_to_int(ttime t) {
    int m = 0;
    m += t.hh * 60;
    m += t.mm;
    return m;
}

bool check(ttime t1, ttime t2) {
    std::set<int> nums;
    nums.insert(t1.hh / 10);
    nums.insert(t1.hh % 10);
    nums.insert(t1.mm / 10);
    nums.insert(t1.mm % 10);
    nums.insert(t2.hh / 10);
    nums.insert(t2.hh % 10);
    nums.insert(t2.mm / 10);
    nums.insert(t2.mm % 10);
    return nums.size() == 8;
}

int main() {
    int k;
    ttime ans;
    bool no_pass = false;
    std::cin >> k;
    for (int i = 0; i <= 23*60 + 59; ++i) {
        ttime start = int_to_time(i);
        ttime finish = int_to_time((i + k) % (24 * 60));
        if (check(start, finish)) {
            ans = start;
            no_pass = true;
            break;
        }
    }
    if (no_pass) {
        if (ans.hh < 10) {
            std::cout << "0";
        }
        std::cout << ans.hh << ":" << ans.mm << std::endl;
    } else {
        std::cout << "PASS\n";
    }
    return 0;
}
