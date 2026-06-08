## DAY 3 OF LINUX

### PROCESS IN LINUX
- Process can be defined as running instance of a program.
- Example: if we have `app.py` which is a program and when we do `python app.py` we are executing this program and hence it becomes a process.
- Linux needs a way to manage CPU, RAM etc. and hence it makes use of processes.
#### command for process:
```
- ps
- output:
PID TTY          TIME CMD
1234 ttys000  0:00.02 zsh
5678 ttys000  0:00.01 ps
```
1. **PID** - Process ID and is unique 
2. **TTY** - It is the terminal where the process is running 
3. **TIME** - CPU time consumed
4. **CMD** - command that started the process
- Example: zsh - Z shell started the process

#### NOTE: `ps` shows the processes that are running in the current terminal

#### command to see all the processes
`ps aux`

### BREAKING DOWN `PS AUX` 
1. **USER** - Who owns the process
2. **PID** - Process ID
3. **%CPU** - CPU usage
4. **%MEM** - RAM usage
5. **VSZ** - Total virtual size consumed in kb
6. **RSZ** - Actual physical RAM consumed by the process in kb

## Process States (`STAT`) in `ps aux`

| STAT | Meaning | Description |
|------|---------|-------------|
| R | Running | Process is running or ready to run on the CPU. |
| S | Sleeping | Waiting for an event to complete (most common state). |
| D | Uninterruptible Sleep | Waiting for I/O (disk/network); cannot be interrupted. |
| T | Stopped | Process has been stopped, usually by a signal or debugger. |
| Z | Zombie | Process has terminated but its parent has not collected its exit status. |
| I | Idle | Idle kernel thread. |

### Additional STAT Flags

| Flag | Meaning | Description |
|------|---------|-------------|
| s | Session Leader | Process is the leader of a session. |
| l | Multi-threaded | Process has multiple threads. |
| + | Foreground Process | Running in the foreground process group. |
| < | High Priority | Higher-than-normal scheduling priority. |
| N | Low Priority | Lower-than-normal scheduling priority (nice value). |

### Examples

| STAT | Interpretation |
|------|---------------|
| `Ss` | Sleeping, Session Leader |
| `Sl` | Sleeping, Multi-threaded |
| `R+` | Running in Foreground |
| `Ssl` | Sleeping, Session Leader, Multi-threaded |
| `Z` | Zombie Process |
| `D` | Waiting for I/O (Uninterruptible Sleep) |

### Common Interview Question

**What is a Zombie Process?**

A zombie process is a process that has finished execution but still has an entry in the process table because its parent process has not yet collected its exit status using `wait()` or `waitpid()`. It appears with the state `Z` in the `STAT` column.

## PROCESS LIFECYCLE
#### BIRTH
- Example:
```
python app.py
```
#### RUNNING
- Linux schedules process time 

#### WAITING
- Maybe waiting for network response, User input, Database query etc.

#### DEATH
```
- kill PID or exit
```
### EXAMPLE OF CREATING A PROCESS
```
- sleep 300 - creates a process and makes it sleep for 300 seconds
Open another terminal and 
- ps aux
```
#### EXAMPLE TO KILL A PROCESS
`kill <PID>`
Example:
```
kill 19088
```
## TOP IN LINUX
- Top can be considered as the LIVE stats of processes.
- Top command is used to find out which process is taking a lot of time so that we can understand from where is the issue being caused 
```
Example of top can be;
- run `top` in a terminal 
- run `yes > /dev/null` purposely to ensure it consumes a lot of memory and you can see it work 
```
### NOTE: make use of `htop` for better understanding

## SIGNALS IN LINUX
- signal is a message that is sent by the user to the process to do certain tasks like kill, stop and etc.

### KILL COMMAND
`kill PID`
- when we do the above it sends a `SIGTERM(15)` and NOT `SIGKILL(9)` 

### WHAT IS SIGTERM(15)
- It means that `Please stop your work gracefully`
- Process gets a chance to:
    - Save data
    - close files
    - write logs etc.
    and dies
### COMMAND TO SIGTERM(15)
` kill -15 PID`

### WHAT IS SIGKILL(9)
- It means `STOP THE JOB IMMEDIATELY`
- Process doesn't get a chance to shutdown gracefully
### COMMAND TO SIGKILL(9)
`kill -9 PID`

## COMMON SIGNALS TO KNOW 
| Signal  | Number | Purpose                               |
| ------- | ------ | ------------------------------------- |
| SIGINT  | 2      | Interrupt (Ctrl+C)                    |
| SIGKILL | 9      | Force kill                            |
| SIGTERM | 15     | Graceful termination                  |
| SIGHUP  | 1      | Reload/restart config (many services) |
| SIGSTOP | 19     | Pause process                         |
| SIGCONT | 18     | Continue paused process               |

## BACKGROUND JOBS
- `sleep 100` doesn't let you run any other commands in the terminal for 100 seconds
- Using `sleep 100 &` lets you run other jobs in the same terminal and you can keep track of the jobs using commands `jobs`
- `ctrl + z` can pause a job
- `bg%1` sends a job to background
` fg%1` gets a job to foreground

## DIFFERENCE BETWEEN GREP AND PGREP
- grep is a general search tool where as,
- pgrap is a process searc tool

Example:
```
Instead of 
ps aux | grep python
we can do,
pgrep python
```

### NOTE: echo $$ gives you the PID of current terminal
