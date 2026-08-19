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
- In traditional images we can do something like ``docker exec -it my-container bash`` where as in a distroless image that cannot e done because distroless image does not include a shell like `sh` or `]`  