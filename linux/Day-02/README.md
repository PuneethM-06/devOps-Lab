## DAY-02 OF LINUX

1. #### Absolute path and Realative path - 
```
- cd /workspaces/devOps-Lab is an example of path
So here; 
Absolute path is: cd /workspaces/devOps-Lab
Relative path is: cd devOps-Lab
To move one level up: cd .
```
## LINUX FILE SYSTEM HIERARCHY (INTERVIEW QUESTION)

```
/
├── home
├── etc
├── var
├── usr
├── tmp
├── root
```
1. `/` -  It is the root of the file system.
2. `/home` - It is the Users personal space. 
- Example: `/home/puneeth/devOps_Lab`
3. `/etc` - It contains systems settings and configuration files
- Example: 
```
/etc
├── hosts
├── passwd
├── ssh
└── nginx
```
4. `/var` - This folder contains variable data, meaning data that can change frequently 
- Example: Logs, cache etc
```
/var
├── log
├── cache
└── spool
```
5. `/usr` - Here, it does not mean user instead it is the section where softwares are installed and stored.
- Example:
```
/usr/bin
/usr/lib
/usr/share
```
6. `/tmp` - This is the section where program uses to store temporary files
7. `/root` - This is not same as `/`
- This is the Home directory of the root user
- Example:
```
/home/puneeth -> Puneeth's room
/root -> Admin's private room
```
## PERMISSIONS
- Here:
    - r = READ
    - w = WRITE
    - x = EXECUTE
- Example:
```
-rw-rw-rw- = Owner - Group - Others
```
## SYMBOLIC chmod
1. #### ADD EXECUTE
```
Example:
chmod +x script.sh
```
- This means for file change the mode to add the execute permission

```
BEFORE:
-rw-r--r--
```
```
AFTER:
-rwxr-xr-x
```
Now linux allows you to, run the `script.sh` file
```
/script.sh
```
2. #### REMOVE EXECUTE
```
Example:
- chmod -x script.sh

BEFORE
-rwxr-xr-x
AFTER
-rw-r--r--
```
## ADD PERMISSIONS TO USERS AND OTHER

1. #### ADD PERMISSION TO OWNER
```
chmod u+x script.sh
```
2. #### ADD PERMISSION TO GROUP
```
chmod g+x script.sh
```
3. #### ADD PERMISSION TO ALL
```
chmod a+x script.sh
```
4. #### ADD PERMISSION TO OTHERS
```
chmod o+x script.sh
```

## NUMERIC chmod

- The numeric numbering for each permission goes like this 
```
r = 4
w = 2
x = 1
```
1. #### READ-ONLY
- Numeric value = 4
2. #### WRITE-ONLY
- Numeric value = 2
3. #### EXECUTE-ONLY
- Numeric value = 1
4. #### READ+WRITE
- Numeric value = 4 + 2 = 6
So and so forth

## EXAMPLE
```
chmod 755 script.sh

7 = rwx (owner)
5 = r-x (group and others)
```
## Execute a Script
```
- mkdir project
- cd project && touch script.sh
- echo "pip3 install python3 " > script.sh
- ./script.sh -> This is the command for executing 
```

## NOTES VERY IMPORTANT
- `#!/bin/bash` means make use of this to execute the script
```
600 = SSH private keys / secrets
644 = Config files / source code
755 = Scripts / executables / directories
```
## SHEBANG 
` #!/bin/bash` This at top of the script is called as Shebang.
- So this specifies the interpreter.
- When a user does `./deploy.sh` Linux asks which program should execute this script and then shebang answers that 
