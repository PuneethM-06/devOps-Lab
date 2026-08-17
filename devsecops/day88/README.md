# DAY 88 - gitLeaks 

1. ### WHAT IS GIT LEAKS?
- Gitleaks is a tool that is used to identify secrets that are accidentally in Git repositories and code 
- The core purpose is **Catch secrets before they become exposed in a Git repository, Git history, or deployed code.**

2. ### WHAT COUNTS AS A SECRET?
- A secret is a sensitive information that should not be exposed to unauthorized people 
- common example:
```
Secrets
│
├── API keys
├── Passwords
├── Database credentials
├── Access tokens
├── GitHub tokens
├── AWS access keys
├── Private keys
└── Connection strings containing credentials
```
> Config values such as ports etc, are not really secrets 

3. ### WHY .gitignore is not enough
- .gitignore can help secrets not going to github for untracked files (files inside .gitignore). But secrets can go when they are places inside a tracked file 

4. ### HOW GITLEAKS DETECTS SECRETS 
- It uses **detection rules to identiy a secret** 
1. **Known patterns**
- Some credentials follow known patterns 
- Example: 
```
AWS_ACCESS_KEY_ID=AKIA...
GITHUB_TOKEN=ghp_...
```

2.**Keywords and context**
- Variable names can provide context
```
password
secret
token
api_key
access_key
```
3. **ENTROPY**
- some secrets are long strings, which can be indicated as a secret 
- Example: `skdjf83jdKLSJDF9sdfKJ23klsdf...`

4 