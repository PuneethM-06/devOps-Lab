# Day 04 - Log Analysis & File Search Interview Questions

## Reading Files

1. What is the purpose of the `cat` command?
2. What is the purpose of the `less` command?
3. What is the difference between `cat` and `less`?
4. When would you prefer `less` over `cat`?
5. How do you search for a word inside `less`?
6. How do you move to the next search result in `less`?
7. How do you move to the previous search result in `less`?
8. How do you exit `less`?
9. How do you move forward one page in `less`?
10. How do you move backward one page in `less`?

## head Command

11. What is the purpose of the `head` command?
12. By default, how many lines does `head` display?
13. How would you display the first 5 lines of a file?
14. What does the following command do?

```bash
head -n 20 app.log
```

15. When would `head` be useful during troubleshooting?

## tail Command

16. What is the purpose of the `tail` command?
17. By default, how many lines does `tail` display?
18. How would you display the last 20 lines of a file?
19. What does the following command do?

```bash
tail -n 50 app.log
```

20. What is the purpose of `tail -f`?
21. What does the `-f` option stand for?
22. Why is `tail -f` commonly used for log monitoring?
23. What happens when new lines are written to a file being monitored by `tail -f`?

## grep Command

24. What is the purpose of the `grep` command?
25. How would you search for the word ERROR in a log file?

```bash
grep "ERROR" app.log
```

26. What does the `-i` option do in grep?
27. What does the `-n` option do in grep?
28. What does the `-c` option do in grep?
29. What does the `-r` option do in grep?
30. What does the `-E` option do in grep?
31. How would you search for both ERROR and INFO messages?

```bash
grep -E "ERROR|INFO" app.log
```

32. How would you count the number of ERROR messages in a log file?
33. How would you display line numbers while searching logs?
34. How would you perform a case-insensitive search?

## wc Command

35. What is the purpose of the `wc` command?
36. What information is displayed by the following output?

```text
25 152 1272 practice.log
```

37. What does `wc -l` do?
38. What does `wc -w` do?
39. What does `wc -c` do?
40. How would you count the number of lines in a file?
41. How would you count the number of ERROR entries in a file?

```bash
grep "ERROR" app.log | wc -l
```

42. What is the difference between line count and word count?

## sort Command

43. What is the purpose of the `sort` command?
44. What does the following command do?

```bash
grep "ERROR" app.log | sort
```

45. What does `sort -n` do?
46. What does `sort -r` do?
47. What is the difference between alphabetical sorting and numeric sorting?
48. Why might `sort -n` be necessary when sorting numbers?

## uniq Command

49. What is the purpose of the `uniq` command?
50. Does `uniq` remove all duplicates automatically?
51. Why is `sort` often used before `uniq`?
52. What does the following command do?

```bash
sort app.log | uniq
```

53. What does `uniq -c` do?
54. How would you count duplicate log entries?

## find Command

55. What is the purpose of the `find` command?
56. How would you find all `.log` files in the current directory and subdirectories?

```bash
find . -name "*.log"
```

57. How would you find all Markdown files?

```bash
find . -name "*.md"
```

58. What does the `.` mean in the find command?
59. What does the `-name` option do?
60. Why is `find` useful for DevOps engineers?

## Pipes and Command Chaining

61. What is a pipe (`|`) in Linux?
62. What happens in the following command?

```bash
grep "ERROR" app.log | wc -l
```

63. What happens in the following command?

```bash
grep "ERROR" app.log | sort | uniq -c
```

64. Why are pipes useful in Linux?

## Scenario-Based Questions

65. A log file is 5 GB in size. Would you use `cat` or `less`? Why?

66. A production application is currently running. How would you monitor new logs as they are generated?

67. You need to find all ERROR messages and count them. Which command would you use?

68. You need to identify duplicate log messages. Which commands would you use?

69. You need to find every `.log` file in a project. Which command would you use?

70. You want to search for "database" regardless of upper or lower case. Which grep option would you use?

71. You want to know which line contains a specific error. Which grep option would you use?

72. You need to count how many lines are present in a log file. Which command would you use?

## Frequently Asked DevOps Questions

73. What is the difference between `cat`, `less`, `head`, and `tail`?

74. Why is `tail -f` one of the most commonly used DevOps commands?

75. What is the difference between `grep -i` and `grep -n`?

76. What is the difference between `sort` and `uniq`?

77. Why is `sort | uniq` commonly seen together?

78. How would you find all occurrences of ERROR in a large log file?

79. How would you monitor logs in real time?

80. Explain the command below step-by-step:

```bash
grep "ERROR" app.log | sort | uniq -c
```
