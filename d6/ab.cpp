#include <iostream>
#include <map>
#include <vector>
#include <string>
#include <fstream>

bool checkSignals(std::vector<char> window) {
        std::map<char, int> occs;
        for (const auto item : window) {
                if (occs.find(item) == occs.end()) {
                        occs[item] = 1;
                        continue;
                }
                occs[item] += 1;
        }
        for (const auto kv : occs) {
                //std::cout << kv.first << ": " << kv.second;
                if (kv.second > 1) {
                       return false;
                }
        }
        return true;
}

const int START_OF_PACKET = 4;
const int START_OF_MESSAGE = 14;

int main(int argc, char * argv[]) {
        std::string data = "./input.txt";
        if (argc > 1) {
                data = argv[1];
        }
        std::string line;
        std::fstream file(data);

        while (std::getline(file, line)) {
                std::vector<char> window;
                int idx = -1;
                for (const auto signal : line) {
                        idx++;
                        if (window.size() < START_OF_MESSAGE) {
                                window.insert(window.begin(), signal);
                                continue;
                        }
                        if (checkSignals(window)) {
                                std::cout << idx << '\n';
                                break;
                        }
                        window.pop_back();
                        window.insert(window.begin(), signal);
                        // std::cout << signal << '\n';
                }
                std::cout << '\n';
        }
}

