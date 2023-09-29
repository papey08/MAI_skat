#include <iostream>
#include <string>
#include <vector>

const int STR_SIZE = 64;
const int MAX_KEY = 65536;

class TPair {
    public:
        TPair() : key(0), value("")
        {}

        TPair(unsigned short n, std::string s) : key(n), value(s)
        {}

        TPair operator=(TPair b) {
            key = b.key;
            value = b.value;
            return *this;
        }

        int GetKey() {
            return key;
        }

        void Print() {
            std::cout << key << '\t' << value;
        }

    private:
        unsigned short key;
        std::string value;
};

void Count(std::vector<TPair>& raw, std::vector<int>& count) {
    for (int i = 0; i < raw.size(); ++i) {
        ++count[raw[i].GetKey()];
    }
    for (int i = 1; i < MAX_KEY; ++i) {
        count[i] += count[i - 1];
    }
}

void Sort(std::vector<TPair>& raw, std::vector<int>& count, std::vector<int>& ans) {
    for (int i = raw.size() - 1; i >= 0; --i) {
        --count[raw[i].GetKey()];
        ans[count[raw[i].GetKey()]] = i;
    }
}

int main() {
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    std::cout.tie(nullptr);
    std::vector<TPair> raw;
    unsigned short inputKey;
    std::string inputString;
    while(std::cin >> inputKey) {
        /* char inputChar;
        scanf("%c", &inputChar);
        do {
            scanf("%c", &inputChar);
            inputString.push_back(inputChar);
        } while (inputChar != '\n'); */
        std::cin >> inputString;
        raw.push_back({inputKey, inputString});
        // inputString = "";
    }

    std::vector<int> count(MAX_KEY, 0);
    std::vector<int> ans(raw.size());
    Count(raw, count);
    Sort(raw, count, ans);
    for (int i = 0; i < ans.size(); ++i) {
        raw[ans[i]].Print();
        std::cout << "\n";
    }
    return 0;
}
