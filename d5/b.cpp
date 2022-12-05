#include <fstream>
#include <regex>
#include <iostream>
#include <string>
#include <stack>
#include <vector>

int main(int argc, char * argv[]) {
  std::string data = "./input.txt";
  if (argc > 1) {
    data = argv[1];
  }


  std::string line;
  std::fstream file(data);
  // std::getline(file, line);
  std::vector<std::string> ivec;
  while (std::getline(file, line) && line[1] != '1') {
    ivec.push_back(line);
  }
  std::vector<std::stack<char>> stacks(line.size() / 4 + 1);
  // Populate the stacks
  for (int i = ivec.size() - 1; i >= 0; i--) {
    for (int j = 0; j < stacks.size(); j ++) {
      // Handle offset for <space> <[> <]> <space>
      const int k = j * 4 + 1;
      if (ivec[i][k] == ' ') continue;
      stacks[j].push(ivec[i][k]);
    }
  }
  // get empty line and discard
  std::getline(file, line);
  const std::regex pattern(R"(move ([0-9]+) from ([0-9]+) to ([0-9]+))");
  while (std::getline(file, line)) {
    std::smatch match;
    std::regex_match(line, match, pattern);
    const auto crates = std::stoi(match[1]);
    const auto from = std::stoi(match[2]) - 1;
    const auto to = std::stoi(match[3]) - 1;
    std::vector<char> to_move;
    for (int i = 0; i < crates; i++) {
            to_move.push_back(stacks[from].top());
            std::cout << stacks[from].top() << '\n';
            stacks[from].pop();
    }
    for (int i = to_move.size() - 1; i >= 0; i--) {
            stacks[to].push(to_move[i]);
    }
  }

  for (const auto stack : stacks) {
    if (!stack.empty()) {
      std::cout << stack.top();
      continue;
    }
    std::cout << ' ';
  }
  std::cout << '\n';
  return 0;
}
