# Day 02 - Linux Filesystem, Permissions & Scripts Interview Questions

## Paths and Navigation

1. What is an absolute path?
2. What is a relative path?
3. What is the difference between an absolute path and a relative path?
4. Give an example of an absolute path.
5. Give an example of a relative path.
6. What does `cd /` do?
7. What does `cd ~` do?
8. How would you navigate to a directory using a relative path?
9. What is the difference between `/` and `~`?

## Linux Filesystem Hierarchy

10. What is the Linux filesystem hierarchy?
11. What is the purpose of the root directory (`/`)?
12. What is stored in `/home`?
13. What is stored in `/etc`?
14. Why is `/etc` important for system administrators?
15. What kind of files are typically found in `/etc`?
16. What is stored in `/var`?
17. Why are application logs usually stored under `/var/log`?
18. What is the purpose of `/usr`?
19. What is stored in `/usr/bin`?
20. What is the purpose of `/tmp`?
21. Can files in `/tmp` be deleted automatically by the system?
22. What is `/root`?
23. What is the difference between `/root` and `/`?
24. What is the difference between `/root` and `/home/puneeth`?

## Linux Permissions

25. What do the permission characters `r`, `w`, and `x` represent?
26. What is the meaning of the permission string below?

```text
-rwxr-xr-x
```

27. Who are the three permission categories in Linux?
28. What permissions does the owner have in `755`?
29. What permissions does the group have in `755`?
30. What permissions do others have in `755`?
31. What is the difference between read and write permission?
32. What does execute permission allow a user to do?

## chmod (Symbolic Mode)

33. What does `chmod` do?
34. What does the following command do?

```bash
chmod +x script.sh
```

35. What does the following command do?

```bash
chmod -x script.sh
```

36. What does `chmod u+x script.sh` do?
37. What does `chmod g+x script.sh` do?
38. What does `chmod o+x script.sh` do?
39. What does `chmod a+x script.sh` do?
40. What is the difference between `u`, `g`, `o`, and `a` in chmod?

## chmod (Numeric Mode)

41. What numeric value represents read permission?
42. What numeric value represents write permission?
43. What numeric value represents execute permission?
44. Why does `7` represent `rwx`?
45. Why does `5` represent `r-x`?
46. What does the command below do?

```bash
chmod 755 script.sh
```

47. What does the command below do?

```bash
chmod 644 app.py
```

48. What permissions are represented by `600`?
49. What permissions are represented by `644`?
50. What permissions are represented by `755`?

## Scripts and Execution

51. What is a shell script?
52. How do you execute a script from the current directory?
53. Why do we use `./script.sh` instead of just `script.sh`?
54. What happens if a script does not have execute permissions?
55. What is the purpose of the command below?

```bash
chmod +x script.sh
```

56. What is the difference between creating a script and executing a script?

## Shebang

57. What is a shebang?
58. What does the following line mean?

```bash
#!/bin/bash
```

59. Why is a shebang important?
60. What happens when Linux encounters a shebang at the top of a script?
61. What interpreter is used when a script begins with `#!/bin/bash`?
62. Can a script run without a shebang? If yes, under what conditions?

## Scenario-Based Questions

63. A script cannot be executed using `./script.sh`. What are the first things you would check?

64. You need to make a script executable for everyone. Which command would you use?

65. You need a configuration file to be readable by everyone but writable only by the owner. Which permission would you choose?

66. Why are SSH private keys commonly set to permission `600`?

67. You accidentally gave a script permission `777`. Why could that be a security concern?

68. Your application logs are filling up disk space. Which directory would you inspect first?

69. You need to find the Nginx configuration file. Which directory would you start looking in?

## Frequently Asked DevOps Questions

70. What is the difference between `644`, `755`, and `600`?

71. Why should scripts usually have `755` permissions?

72. Why are private keys often restricted to `600` permissions?

73. What is the difference between symbolic chmod and numeric chmod?

74. Explain the permission string below:

```text
-rw-r--r--
```

75. Explain what happens internally when you run:

```bash
./deploy.sh
```
