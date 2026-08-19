# DAY 89 - DISTROLESS IMAGES 

1. ### WHAT IS A DISTROLESS IMAGE
- **A distroless image strips out the unnecessary general-purpose toolig and keeps only what is needed to run the application**
```
Your application
        +
Required runtime libraries
        ↓
Distroless image
```
- By default there will be no shell manager or no package manager 
- **By this way, the attack surface will be reduced but it also comes with the cons** 

2. ### TRADITIONAL BASE IMAGES vs DISTROLESS IMAGES
-  A traditional image such as ubunut, alpine and debian contain general purpose tooling by default.

| Traditional Image        | Distroless Image         |
| ------------------------ | ------------------------ |
| Usually has a shell      | Usually no shell         |
| May have package manager | No package manager       |
| More utilities           | Minimal utilities        |
| Easier to debug inside   | Harder to debug directly |
| Larger attack surface    | Smaller attack surface   |

### 3. WHY DOES A SMALLER ATTACK SURAFACE MATTER 
- An attack surface is basically the collection of possible entry points for the attackers to enter and breach the application

- More additional component can result in:
    1. A vulnerability
    2. Misconfigurations
    3. Unnecessary binary attack 
**PRINCIPLE**: If a component is not needed then do not include at all 

4. ### NO SHELL: WHAT DOES THAT ACTUALLY MEAN 
- In traditional images we can do something like ``docker exec -it my-container bash`` where as in a distroless image that cannot e done because distroless image does not include a shell like `sh` or `bash`  

- The logic is, shell is used for debugging, if the goal is to run the application then why do we need a shell 

5. ### HOW DO WE DEBUG A DISTROLESS CONTAINER
1. **OBSERVABILITY FIRST**
```
Application logs
        ↓
Metrics
        ↓
Tracing
        ↓
Kubernetes events/logs
```
- Here we leverage, application logs, k8s logs and metrics from observaility tools 

2. **DEBUG OR EPHEMERAL CONTAINERS**
- In k8s, we can attach epehemeral debug containers that has tools such as:
```
sh
curl
ping
nslookup
ps
```
```
Pod
│
├── Application Container
│      └── Distroless
│          └── No shell
│
└── Ephemeral Debug Container
       └── Has debugging tools
```
3. **SEPERATE DEBUG IMAGE**
- Another approach is to build a sepearate debug image with the same application runnning 

6. ### DISTROLESS + GO
- Go and distroless image go hand in hand because Go application can be often compiled into a **self-contained executable binary**
```
Go source code
      ↓
go build
      ↓
Compiled binary
      ↓
Minimal runtime image
      ↓
Distroless container
```
- **Unlike interpreted languages, the production container does not need the Go compiler or source code to run the compiled application.**

7. ### SCRATCH vs DISTROLESS
- Scratch is basically an **empty base image** and we build the image with components that our app needs 
- Where as a Distroless image is not a base image and it comes with a minimum runtime 

| `scratch`                                     | Distroless                           |
| --------------------------------------------- | ------------------------------------ |
| Essentially empty                             | Minimal runtime environment          |
| You provide everything needed                 | Provides selected runtime components |
| Can be extremely small                        | Usually slightly larger              |
| Can require more setup                        | Often more convenient                |
| Good for fully self-contained/static binaries | Good for minimal production runtimes |

8. ### SECURITY TRADE OFFS
- **Distroless can reduce the attack surface but that does not mean we can automatically make our application more secure**
- A distroless image can still contain:
    1. Vulnerable application dependencies
    2. vul runtime libraries
    3. Application code vul
    4. miconfig
