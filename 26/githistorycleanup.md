echo node_modules > .gitignore

git add .gitignore

git commit -m "Add node_modules to .gitignore"

sudo apt-get install git-filter-repo

git filter-repo --path node_modules --invert-paths

git commit -am "Remove node_modules from history"

git push -f 