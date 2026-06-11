# Linux Day 5 — Text Processing Practice

## Practice Log

Create:

```bash
cat > practice.log << 'EOF'
INFO Server started
INFO Request received
ERROR Database connection failed
WARNING Memory high
ERROR Timeout on API call
ERROR Database connection failed
INFO Request completed
WARNING Disk usage high
ERROR Authentication failed
ERROR Database connection failed
EOF
```

---

# Practice Exercises

## grep

### Exercise 1

Find all ERROR lines.

### Exercise 2

Find all WARNING lines.

### Exercise 3

Find all lines containing the word Database.

### Exercise 4

Count the number of ERROR lines.

---

## awk

### Exercise 5

Print the first word of every line.

### Exercise 6

Print the second word of every line.

### Exercise 7

Print line number and entire line.

---

## sed

### Exercise 8

Convert every occurrence of:

ERROR

to:

CRITICAL

without modifying the original file.

---

## sort + uniq

### Exercise 9

Find the most common message in the log.

Hint:

* sort
* uniq -c
* sort -rn

---

# Additional Practice

## tee

### Exercise 10

Display and save output at the same time.

Print:

Application Started

and save it to:

output.log

---

### Exercise 11

Append a new line to output.log using tee.

---

## xargs

### Exercise 12

Delete all .log files found in a directory tree.

---

### Exercise 13

List all .txt files using find and xargs.

---

### Exercise 14

Create three files using echo and xargs touch.

Example output:

file1.txt
file2.txt
file3.txt

---

# Mini Project

Create:

```
day-05/
├── practice.log
├── report.txt
└── README.md
```

Generate report.txt using Linux commands only.

Required output:

```text
Total INFO:
Total WARNING:
Total ERROR:

Most Common Error:
Most Common Warning:
```

Rules:

* Use grep
* Use awk if needed
* Use sort
* Use uniq
* Use pipes
* No manual counting

---

# Interview Questions

## grep

### 1. What is the difference between grep and awk?

### 2. What does grep -v do?

### 3. What does grep -n do?

### 4. What does grep -c do?

---

## cut

### 5. What does:

```bash
cut -d',' -f2 employee.csv
```

do?

### 6. What is the purpose of the -d option?

### 7. What is the purpose of the -f option?

---

## awk

### 8. What does:

```bash
awk '{print $1}'
```

do?

### 9. What does:

```bash
awk '{print $2}'
```

do?

### 10. What does:

```bash
awk '{print $0}'
```

do?

### 11. What does NR mean in awk?

### 12. Difference between $0 and $1?

### 13. Difference between grep and awk?

---

## sed

### 14. Difference between sed and awk?

### 15. What does:

```bash
sed 's/error/ERROR/g'
```

do?

### 16. What does the g mean?

### 17. What does:

```bash
sed '/INFO/d'
```

do?

### 18. What is the purpose of the d command in sed?

---

## sort and uniq

### 19. Why is sort often used before uniq?

### 20. What does:

```bash
uniq -c
```

do?

### 21. What does:

```bash
sort -r
```

do?

### 22. What does:

```bash
sort -n
```

do?

### 23. What does:

```bash
sort -rn
```

do?

---

## tee

### 24. What does tee do?

### 25. Difference between:

```bash
>
```

and

```bash
tee
```

### 26. What does:

```bash
tee -a
```

do?

---

## xargs

### 27. What does xargs do?

### 28. Explain:

```bash
find . -name "*.log" | xargs rm
```

### 29. Why might:

```bash
find . -name "*.txt" | xargs rm
```

be dangerous?

### 30. What is the safer version using -print0 and -0?

---

## Pipes

### 31. What does a pipe (|) do?

### 32. Explain:

```bash
grep ERROR app.log | sort | uniq -c | sort -rn
```

line by line.

### 33. Explain:

```bash
ps aux | awk '{print $1,$11}'
```

### 34. Explain:

```bash
grep ERROR app.log | wc -l
```

---

# Goal

By the end of Day 5, you should comfortably use:

* grep
* cut
* awk
* sed
* sort
* uniq
* wc
* tee
* xargs
* pipes

and combine them together to analyze logs and process text efficiently.
