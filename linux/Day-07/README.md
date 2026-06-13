## DAY-07 OF LINUX

1. ### WHAT IS SHELL?
-  Shell is an interface between user and the operating system.
- Example:
    - zsh
    - bash
    - sh
 #### NOTE: `echo $SHELL` is the command you can use to check the terminal/interface that you are using 

 2. ### WHAT IS SHELL SCRIPTING?
 -  Shell scripting is simply a file of containing linux commands 

 ## SHEBANG
- `#!/shell/bash` is called as SHEBANG 
- This command mentions the **Interpreter** in which shell scripting should run. 

- `#!` mentions the interpreter it should use after this command 
- '/shell/bash` is the interpreter in which we should run. 
- Example of different interpreters
```
1. #!/bin/bash
2. #!/bin/sh
3. #!/bin/zsh
```

#### BEST WAY TO USE SHEBANG
- `#!/usr/bin/env bash`
- env finds it automatically
----

### CREATING YOUR FIRST SCRIPT
create:
```
vim linux.sh
```
- Now, to insert input, use `i`
- To exit, use `Esc` and `:wq`

### To run the script
- `./linux.sh`
----

### HOW SHELL SCRIPTING ACTUALLY WORKS
1. Checks for permission
2. Reads `shebang` - `#!/bin/bash`
3. starts bash 
4. Executes each command line by line 
5. Displays output

------
### METHODS OD RUNNING A SCRIPT
1. `./deploy.sh`
2. `bash deploy.sh`
3. `source deploy.h`

### COMMENTS 
- `#` we make use of this for a comment 
- Example:
```
# This is a comment in script.sh
```

## VARIABLE
1. ### CREATE VARIABLE 
- Example
```
name="Puneeth"
```
- There will be no spaces when we create variables

2. ### ACCESS VARIABLES
- Example
```
echo $name
```
3. ### CURRENT USER
```
echo $USER
```
4. ### CURRENT SHELL
```
echo $SHELL
```
5. ### CURRENT DIRECTORY
```
echo $PWD
```

## COMMAND SUBSTITUTION
- Example
```
today=$(date)
 or 
user=$(whoami)
```
---
## QUOTING 

1. ### DOUBLE QUOTES
- Example:
```
name="Puneeth"
echo "$name"
```
- output:
```
Puneeth
```

2. ### SINGLE QUOTES
- Example:
```
name="Puneeth"
echo '$name'
```
- output:
```
'$name'
```
## READ INPUT FROM USER

```
read name 
```
- Example:
```
#!/bin/bash 

echo "Whats your name"

read name

echo "Hello, $name!"
```
## IF ELSE CONDITION

- syntax 
```
if [ condition ]; then
    commands
else
    commands
fi
```
## If Elif CONDITION
- syntax
```
if [ condition ]; then
    commands
elif [ condition ]; then 
    commands
else 
    commands
fi
```

## CONDITIONS 
1. `-f` - checks if a regular file exists
2. `-d` - checks if a directory exists
3. `-e` - checks if anything exists
4. `-z` - checks if length of a string is empty
5. `-n` - checks if length of a string is not 0

## CASE STATEMENTS
- syntax
```
case "$VAR" in 
    pattern1)
        commands
    pattern2)
        commands
    *)
        default statement
        ;;
esac
```
## FOR LOOPS
```
for i in 1 2 3 4 
do
    echo 
done 
```

## WHILE LOOP
- syntax
``` 
while read linme
do
    echo $line
done 
```
