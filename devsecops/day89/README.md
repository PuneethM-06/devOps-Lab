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
