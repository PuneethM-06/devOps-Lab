# Day 01 - Linux Basics Interview Questions

## Basic Linux Questions

1. What is Linux?
2. What is the difference between Linux and Windows?
3. What is a Linux distribution?
4. What is a shell?
5. What is the purpose of the terminal in Linux?

## User and Directory Questions

6. What does the `whoami` command do?
7. What does the `pwd` command do?
8. What is the difference between an absolute path and a relative path?
9. What does the `~` symbol represent in Linux?
10. What happens when you execute `cd` without any arguments?
11. How do you navigate to the root directory?
12. How do you navigate to your home directory?

## File and Directory Commands

13. What is the purpose of the `ls` command?
14. What is the difference between `ls` and `ls -la`?
15. Why is `ls -la` commonly used by Linux administrators?
16. What are hidden files in Linux?
17. How can you view hidden files and directories?
18. What does the `mkdir` command do?
19. What does the `touch` command do?
20. What is the difference between creating a file and creating a directory?

## Understanding ls -la Output

21. What does the `d` at the beginning of a permission string indicate?
22. What does the `-` at the beginning of a permission string indicate?
23. What do the permission characters `r`, `w`, and `x` mean?
24. What are owner permissions?
25. What are group permissions?
26. What are other-user permissions?
27. Explain the output below:

```text
drwxr-xr-x
```

28. In the output of `ls -la`, what does the file size column represent?
29. In the output of `ls -la`, what does the owner column represent?
30. In the output of `ls -la`, what does the group column represent?

## File Content Commands

31. What does the `cat` command do?
32. How can you display the contents of a file?
33. What does the `echo` command do?
34. What does the `>` symbol do in Linux?
35. What is the difference between `echo "Hello" > file.txt` and `echo "Hello" >> file.txt`?

## Copy Operations

36. What does the `cp` command do?
37. How do you copy a file from one name to another?
38. Can `cp` be used to copy directories? If yes, how?

## Scenario-Based Questions

39. You are currently in `/home/user/projects` and want to move to the root directory. Which command would you use?

40. You need to create a directory called `logs` and a file called `app.log` inside it. Which commands would you use?

41. You accidentally created a file in the wrong directory. How would you create a copy of that file in another directory?

42. How would you verify that a file was successfully created?

43. How would you check which user is currently logged into the system?

44. You are given a path and need to verify your current location before running commands. Which command would you use?

## Common Interview Tricky Questions

45. What is the difference between `/` and `~`?

46. What is the difference between a file and a directory?

47. Does `touch` always create a new file?

48. What happens if you run `mkdir` on a directory that already exists?

49. Why might a DevOps engineer frequently use `ls -la` instead of just `ls`?

50. If you are in your home directory and execute `cd /`, where will you end up?
