## DAY 1 - LINUX BASICS COMMANDS

1. #### whoami - This shows the current logged in user.
2. #### pwd - shows the current working directory
3. #### ls - Gives lists of files and folders
4. #### ls -la - This is one of the most useful commands 
   ```
   Example output:
   @PuneethM-06 ➜ /workspaces/devOps-Lab (main) $ ls -la
    total 16
    drwxrwxrwx+ 4 codespace root      4096 Jun  6 09:59 .
    drwxr-xrwx+ 5 codespace root      4096 Jun  6 09:20 ..
    drwxrwxrwx+ 8 codespace root      4096 Jun  6 09:22 .git
    drwxrwxrwx+ 2 codespace codespace 4096 Jun  6 09:59 linux
    @PuneethM-06 ➜ /workspaces/devOps-Lab (main) $ 

    NOTE,
    d - directory
    - File
    rwx - owner permission
    r-x - Group permissions
    r-x - Other permissions

    we can think it as:
    [PERMISSIONS] [OWNER] [GROUP] [SIZE] [DATE] [NAME]

    Also, -l (detailed information)
          -a (hidden files)
    ```
5. #### cd - Change of directory
      - To go to **root directory** we can do `cd /`
6. #### cd ~ - Go home
7. #### mkdir <folder_name> - Create a new folder
8. #### touch <file_name> -  Create a new file 
9. #### echo - write into a file 
```
Example of echo
echo "print("hello world!)" > app.py
```
10. #### cat - Read into the file 
11. #### cp - copy file 
```
Example of cp:  cp app.py app_backend.py
```
11. 