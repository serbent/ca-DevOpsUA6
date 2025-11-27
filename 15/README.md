# 15. Git

```bash
ssh-keygen -t ecdsa
```

```bash
cat id_ecdsa.pub
```

```bash
git clone git@github.com:serbent/ca-DevOpsUA6-pub.git # Repleace with your git repo
```
```bash
cd GIT_REPO_DIR
```
Create New Branch
```bash
git checkout -b bug/1111
```
Switch to _main_ branch again
```bash
git checkout main
```
After making changes
```bash
git add . 
git commit -m 'MESSAGE'
git push
```

To merge another branch to your branch: 
```bash
git merge main
```

To see git changes made: 
```bash
git status
```
Can give info like this: 
```
❯ git status
On branch bug/1111
Your branch is up to date with 'origin/bug/1111'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        instructions.md

no changes added to commit (use "git add" and/or "git commit -a")
```

Lazygit function: 
```bash
# Git to pull/commit/push
function lazygit {
    echo "Git pull:"
    git pull
    echo "Git add:"
    git add .
    echo "git commit -am :"
    git commit -a -m "$1"
    echo "git push"
    git push
}
```