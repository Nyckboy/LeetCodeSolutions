# 🚀 LeetCode Solutions

A collection of my LeetCode solutions in **Java** and **JavaScript**, created to track my progress in Data Structures and Algorithms. 

The repository focuses on:
* **Clean Code:** Solutions are written to be readable and maintainable.
* **Efficiency:** Time and Space complexity analysis included for each problem.
* **Automation:** Contains custom shell/PowerShell scripts to streamline the folder creation process.

## 📂 Structure
Each problem is located in its own directory (e.g., `1-two-sum/`) containing:
1.  `README.md`: Problem description and constraints.
2.  `Solution.js`: The working code with complexity notes.

## 🛠️ Usage
This repo uses a custom script to generate boilerplate code.
### 1. Create a New Problem
Generates the folder structure, `README.md`, and solution file for a new LeetCode problem.
```bash
# Usage: ./new.sh <Problem ID> "<Problem Title>"
./new.sh 1 "Two Sum"
```
### 2. Sync & Push
Since the GitHub Action updates the README automatically, the remote repo is often ahead of local. Use this script to safely pull (rebase), stage, commit, and push in one command.
```bash
# Usage: ./push.sh "<Commit Message>"
./push.sh "feat: add solution for 1-two-sum"
```

## Solutions

### Algorithms

<!-- SOL_TABLE_START_ALGO -->

| # | Title | Solution | Difficulty |
|---|---|---|---|
| 1 | [Two Sum](./algorithms/1-two-sum) | [JS](./algorithms/1-two-sum/Solution.js) | Easy |
| 2 | [Add Two Numbers](./algorithms/2-add-two-numbers) | [JS](./algorithms/2-add-two-numbers/Solution.js) | Medium |
| 3 | [Longest Substring Without Repeating Characters](./algorithms/3-longest-substring-without-repeating-characters) | [JS](./algorithms/3-longest-substring-without-repeating-characters/Solution.js) | Medium |
| 7 | [Reverse Integer](./algorithms/7-reverse-integer) | [JS](./algorithms/7-reverse-integer/Solution.js) | Medium |
| 8 | [String To Integer (Atoi)](./algorithms/8-string-to-integer-(atoi)) | [JS](./algorithms/8-string-to-integer-(atoi)/Solution.js) | Medium |
| 11 | [Container With Most Water](./algorithms/11-container-with-most-water) | [JS](./algorithms/11-container-with-most-water/Solution.js) | Medium |
| 14 | [Longest Common Prefix](./algorithms/14-longest-common-prefix) | [JS](./algorithms/14-longest-common-prefix/Solution.js) | Easy |
| 17 | [Letter Combinations Of A Phone Number](./algorithms/17-letter-combinations-of-a-phone-number) | [JS](./algorithms/17-letter-combinations-of-a-phone-number/Solution.js) | Medium |
| 1200 | [Minimum Absolute Difference](./algorithms/1200-minimum-absolute-difference) | [JS](./algorithms/1200-minimum-absolute-difference/Solution.js) | Easy |

<!-- SOL_TABLE_END_ALGO -->

### Sys-design

<!-- SOL_TABLE_START_SYSDES -->

| # | Title | Solution | Difficulty |
|---|---|---|---|
| 1 | [Lru Cache](./system-design/Q1-lru-cache) | No Code | Unknown |

<!-- SOL_TABLE_END_SYSDES -->