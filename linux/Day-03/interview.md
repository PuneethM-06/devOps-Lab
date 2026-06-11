# Day 03 - Linux Processes, Signals & Job Control Interview Questions

## Process Fundamentals

1. What is a process in Linux?
2. What is the difference between a program and a process?
3. What happens internally when you execute `python app.py`?
4. Why does Linux use processes?
5. Can multiple processes be created from the same program?

## Process Identification

6. What is a PID?
7. Why is a PID important?
8. Can two running processes have the same PID?
9. What command shows the PID of the current shell?
10. What does `echo $$` display?

## ps Command

11. What does the `ps` command do?
12. What is the difference between `ps` and `ps aux`?
13. Why does `ps` show fewer processes than `ps aux`?
14. What information is displayed in the following output?

```text
PID TTY          TIME CMD
1234 ttys000  0:00.02 zsh
```

15. What does the `TTY` column represent?
16. What does the `TIME` column represent?
17. What does the `CMD` column represent?

## Understanding ps aux

18. What does the `USER` column represent?
19. What does the `%CPU` column represent?
20. What does the `%MEM` column represent?
21. What is VSZ?
22. What is RSS?
23. Why is RSS often more useful than VSZ when troubleshooting memory issues?

## Process States

24. What does the process state `R` mean?
25. What does the process state `S` mean?
26. What does the process state `D` mean?
27. What does the process state `T` mean?
28. What does the process state `Z` mean?
29. What is a Zombie Process?
30. Why are zombie processes created?
31. How can you identify a zombie process?
32. What does `Ss` mean in the STAT column?
33. What does `Sl` mean?
34. What does `R+` mean?
35. What does `Ssl` mean?

## Process Lifecycle

36. Explain the lifecycle of a process.
37. What happens when a process is created?
38. What happens when a process enters a waiting state?
39. Give examples of events that can cause a process to wait.
40. How does a process terminate?

## Creating and Managing Processes

41. What does the following command do?

```bash
sleep 300
```

42. How can you verify that the sleep process is running?
43. How do you find the PID of a running process?
44. How do you terminate a process?

## top and Monitoring

45. What is the purpose of the `top` command?
46. How is `top` different from `ps aux`?
47. Why is `top` useful during production incidents?
48. What information can you observe using `top`?
49. What is `htop`?
50. Why do many Linux administrators prefer `htop` over `top`?

## Linux Signals

51. What is a signal in Linux?
52. Why are signals used?
53. What happens when you execute:

```bash
kill PID
```

54. Which signal is sent by default when running `kill PID`?
55. What is the difference between SIGTERM and SIGKILL?

## SIGTERM

56. What is SIGTERM?
57. What signal number is associated with SIGTERM?
58. Why is SIGTERM considered a graceful shutdown?
59. What actions can a process perform before exiting after receiving SIGTERM?
60. How do you explicitly send SIGTERM?

```bash
kill -15 PID
```

## SIGKILL

61. What is SIGKILL?
62. What signal number is associated with SIGKILL?
63. Why is SIGKILL considered a forceful termination?
64. Can a process ignore SIGKILL?
65. How do you explicitly send SIGKILL?

```bash
kill -9 PID
```

66. When would you use SIGKILL instead of SIGTERM?

## Common Signals

67. What does SIGINT do?
68. What key combination sends SIGINT?
69. What is SIGHUP commonly used for?
70. What does SIGSTOP do?
71. What does SIGCONT do?

## Background Jobs

72. What happens when you run:

```bash
sleep 100
```

73. What happens when you run:

```bash
sleep 100 &
```

74. What does the `&` operator do?
75. What command lists background jobs?

```bash
jobs
```

76. What does Ctrl+Z do?
77. What is the difference between Ctrl+C and Ctrl+Z?
78. How do you move a stopped job to the background?

```bash
bg %1
```

79. How do you bring a background job to the foreground?

```bash
fg %1
```

## grep vs pgrep

80. What is the difference between `grep` and `pgrep`?
81. Why might `pgrep` be preferred when searching for processes?
82. What is the advantage of:

```bash
pgrep python
```

over:

```bash
ps aux | grep python
```

83. What does the following command return?

```bash
pgrep nginx
```

## Scenario-Based Questions

84. An application is consuming excessive CPU. Which commands would you use to identify it?

85. A process is not responding. Would you use SIGTERM or SIGKILL first? Why?

86. You accidentally started a long-running command in the foreground. How would you move it to the background?

87. A process appears as `Z` in the STAT column. What does that indicate?

88. How would you find all running Python processes?

89. A service keeps restarting unexpectedly. Which process-related commands would you use to investigate?

90. A process is stuck waiting on disk I/O. Which process state would you expect to see?

## Frequently Asked DevOps Questions

91. What is the difference between a process and a thread?

92. What is a zombie process and how would you troubleshoot it?

93. Why should SIGTERM usually be attempted before SIGKILL?

94. What is the difference between foreground and background processes?

95. What command would you use to monitor processes in real time?

96. What command would you use to find a process by name?

97. What does `echo $$` return?

98. What is the purpose of process IDs in Linux?

99. Why is `top` useful during production outages?

100. Explain what happens internally when you run:

```bash
python app.py
```
