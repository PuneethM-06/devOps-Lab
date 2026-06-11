## DAY 05 LINUX

### CUT IN LINUX
- `cut` helps you extract certain or specific columns from structured data
```
Excel Sheet
-------------------------
Name      Role      Country
Puneeth   Engineer  India
```
We can extract Puneeth if we want that alone 

Syntax / Example:
```
cut -d',' -f1
```
here;
- `-d` is the delimeter we are referring to 
- `f1` refers to the field number 

#### MULTIPLE FIELDS
Suppose we want multiple fields
```
echo "puneeth, engineer, india"
cut -d',' f1,3
So we get,
puneeth,india
```
### Command:

cut -d',' -f1,3,5 employee.csv
> Question 9: Extract Unique Countries

Expected Output:

Canada
Germany
India
USA

Command:

cut -d',' -f4 employee.csv | sort | uniq
> Question 10: Count Number of Employees

Expected Output:

5

Command:

wc -l employee.csv
Question 11: Extract Names Using Pipe

Expected Output:

Puneeth

Command:

echo "Puneeth,Frontend Engineer,India,85000" | cut -d',' -f1
Question 12: Extract Country and Salary Using Pipe

Expected Output:

India,85000

Command:

echo "Puneeth,Frontend Engineer,India,85000" | cut -d',' -f3,4
----

# Practice Dataset

```text
101,Puneeth,Frontend Engineer,India,85000
102,Rahul,DevOps Engineer,India,95000
103,John,Backend Engineer,USA,120000
104,Alice,Data Engineer,Germany,110000
105,Sara,SRE,Canada,130000
```

## Question 1: Extract Employee Names

### Expected Output

```text
Puneeth
Rahul
John
Alice
Sara
```

### Command

```bash
cut -d',' -f2 employee.csv
```

---

## Question 2: Extract Employee Roles

### Expected Output

```text
Frontend Engineer
DevOps Engineer
Backend Engineer
Data Engineer
SRE
```

### Command

```bash
cut -d',' -f3 employee.csv
```

---

## Question 3: Extract Countries

### Expected Output

```text
India
India
USA
Germany
Canada
```

### Command

```bash
cut -d',' -f4 employee.csv
```

---

## Question 4: Extract Salaries

### Expected Output

```text
85000
95000
120000
110000
130000
```

### Command

```bash
cut -d',' -f5 employee.csv
```

---

## Question 5: Extract Employee ID and Name

### Expected Output

```text
101,Puneeth
102,Rahul
103,John
104,Alice
105,Sara
```

### Command

```bash
cut -d',' -f1,2 employee.csv
```

---

## Question 6: Extract Name and Country

### Expected Output

```text
Puneeth,India
Rahul,India
John,USA
Alice,Germany
Sara,Canada
```

### Command

```bash
cut -d',' -f2,4 employee.csv
```

---

## Question 7: Extract Role and Salary

### Expected Output

```text
Frontend Engineer,85000
DevOps Engineer,95000
Backend Engineer,120000
Data Engineer,110000
SRE,130000
```

### Command

```bash
cut -d',' -f3,5 employee.csv
```

---

## Question 8: Extract Employee ID, Role and Salary

### Expected Output

```text
101,Frontend Engineer,85000
102,DevOps Engineer,95000
103,Backend Engineer,120000
104,Data Engineer,110000
105,SRE,130000
```

### Command

```bash
cut -d',' -f1,3,5 employee.csv
```

---

## Question 9: Extract Unique Countries

### Expected Output

```text
Canada
Germany
India
USA
```

### Command

```bash
cut -d',' -f4 employee.csv | sort | uniq
```

---

## Question 10: Count Number of Employees

### Expected Output

```text
5
```

### Command

```bash
wc -l employee.csv
```

---

## Question 11: Extract Name Using Pipe

### Input

```text
Puneeth,Frontend Engineer,India,85000
```

### Expected Output

```text
Puneeth
```

### Command

```bash
echo "Puneeth,Frontend Engineer,India,85000" | cut -d',' -f1
```

---

## Question 12: Extract Country and Salary Using Pipe

### Input

```text
Puneeth,Frontend Engineer,India,85000
```

### Expected Output

```text
India,85000
```

### Command

```bash
echo "Puneeth,Frontend Engineer,India,85000" | cut -d',' -f3,4
```

---

## Key Learnings

### `-d`

Specifies the delimiter.

Example:

```bash
cut -d',' -f2 employee.csv
```

Here the delimiter is:

```text
,
```

### `-f`

Specifies the field number(s) to extract.

Examples:

```bash
-f1
-f2
-f3
-f1,3,5
```

### Common Pattern

```bash
cut -d'<delimiter>' -f<field_numbers> <file>
```

Example:

```bash
cut -d',' -f2,4 employee.csv
```
------
#### NOTE: TO REMOVE DUPLICATE WE NEED TO DO
```
cut -d',' -f1 employee.csv | sort | uniq
```
#### UINIQ ALWAYS REMOVES ADJACENT DUPLICATES
SOLUTION: always use sort before using uniq
-----


## AWK in linux
- If `grep` can find lines and `cut` can extract columns 
- `awk` can filter, calculate, format and do other things column by column 

- By default, awk splits each line by whitespace 
Example:
```
awk '{print $1} file.txt
```
- $1 is the field we are referring to and 
- $0 is a special character the prints every line 
```
awk '{print $1} file.txt
or 
awk -F',' {print $1} file.txt
```

#### NR -> record numbers (Line numbers)
Example:
```
awk '{print NR, $1} file.txt
```
## SED IN LINUX

- grep is used for finding lines
- find is used for finding files
- cut is used for finding columns and fields
- awk is used to process data columns
- sed is used for editing the text streams 

example:
```
sed 's/error/ERROR/g' log.text
```
Here;
s = substitue
error in the example is old
ERROR is the new replacement 

- /g = does for all the occurances in the file 
- `sed 's/error/ERROR` does only for the first occurance in the word

### SAVE OUTPUT FOR ANOTHER FILE
```
sed 's/error/ERROR/g' file.txt > modified.txt
```

### DELETE MATCHING LINES
Example:
```
sed '/INFO/d' file.txt
```

### DELETING SPECIFIC LINE NUMBERS

```
sed '3d' file.txt
```
### DELETING MULTIPLE LINES AT A TIME 
```
sed '2,4d' file.txt
```
Deletes all the lines from 2 to 4

### PRINT SPECIFC LINES

```
sed -n '3p' file.txt
or
sed -b '3,5p' file.txt
```

## tee IN LINUX
- tee is a command that is used to send the input to the file and also display on the screen 

Example:
```
echo "Hello! Welcome to DevOps" | tee modified.log
or
echo "Application STOPPED" | tee -a modified log
```
#### NOTE - `-a` doesnt rewrite instead it appends to the exisiting one 

## xargs IN LINUX
- `xargs` takes an input and turns it into command arguement
```
echo "file 1 file 2 file 3" | xargs rm 
```

Real world Example:
```
echo "practice.log" | xargs find . -name 
or 
find . -name "practice.log" | xargs touch modified.log | cp practice.log modified.log
```
