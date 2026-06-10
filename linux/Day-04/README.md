## DAY 04 OF LINUX

#### LESS IN LINUX
- It is similar to `cat` where we can read files but the difference between `less` and `cat` is that 
- `cat` is used for reading small files and it dumps everything which can be difficult to read at time 
- `less` it is like a open book where we can read page by page 

In less; 

- space - moves to next page
- b - moves backward
- / - search for a word
- n - next match in searching 
- N - goes to previous match
- q - quit / exit 

#### HEAD AND TAIL IN LINUX

`HEAD` - head in linux is used for printing or showing the initial few lines of a file 
- Example:
```
head practice.log - shows the first 10 lines
head -n 6 practice.log - shows specified number of lines
```

`Tail` - tail in linux is used for printing or showing the last few lines of a file
- Example:
```
tail practice.log - shows the last 10 lines
tail -n 6 practice.log - shows specified number of lines
```
`Tail -f` 
- Here -f stands for follow
- This is very useful for reading logs as we follow. 

### SEARCH LOGS
- we can make use of grep for doing this 
- Example
```
grep "ERROR" practice.log
```
- Example for case insensitive
```
grep -i "ERROR"  practice.log
```
- Example of recursive search
```
grep -r "ERROR" practice.log
```
-r stands for recursive search where it searches for all the files 
- Example for showing line numbers
```
grep -n "ERROR" practice.log
```
- Example for count matches
```
grep -c "ERROR" practice.log
```
- Example of multiple matches
```
grep -E "ERROR|INFO" practice.log
```

### COUNT LINES IN LINUX
`wc` is the command that can be used to count the number of words in a file. however, it can be combined with other commands
```
Example
wc practice.log
output:
25  152 1272 practice.log
Lines words character
```
#### count lines
- Example
```
wc -l practice.log
or
grep "ERROR" practice.log | wc -l
```

#### count words
- Example
```
wc -w practice.log
or
grep "ERROR" practice.log | wc -w
```
#### count characters
- Example
```
wc -c practice.log
or grep "ERROR" practice.log | wc -c
```
### sort them 
```
Example
grep "ERROR" practice.log | sort
```

#### UNIQUE ERROR
`-uniq` is used to remove duplicate error lines or log lines 
- Example
```
grep "ERROR" practice.log | uniq
or grep "ERROR" practice.log | uniq -c
```
### FINDING FILES
`find` is the word that is used to find files 
```
find . -name "*.log"
or
find . -name "*.md"
```

### sort -n
this command is to sort 
```
Example
grep "ERROR" | sort -n
```
-n means sort numerically

### sort -r
this command is to sort 
```
Example
grep "ERROR" | sort -r
```
-n means sort reverse



