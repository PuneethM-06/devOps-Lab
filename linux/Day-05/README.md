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
