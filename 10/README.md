# 10 Bash Scripting

## Usefull stuff

`set -euo pipefail` - set at the start of your script  
`-u` — Treat unset variables as errors  
`-e` — Exit immediately on error  
`-o pipefail` — Fail if any command in a pipeline fails  
`if [ -w/-r/-x/-f/-d $SOMETHING ]` only one check per if statement.  
Check if  
* -w file/directory is writable
* -r file/directory is readable
* -x file/directory is executable
* -f file/directory is a file
* -d file/directory is a directory


## Exercises

### Script 02

Go to [Mockaroo](https://mockaroo.com/) and generate some random data.

**Fields:**
- Full Name  
- Age  
- Email  
- Gender  
- Country  
- Job Title  
- Salary  

---

### Tasks

1. Create a procedure to display all **countries** in the report.  
2. Create a procedure to display all **job titles** in the report.  
3. Create a procedure to calculate **average salary** in a country.  
4. Create a procedure to calculate **average salary by job title** and display the **highest-paying job**.  
5. Create a procedure to calculate **average age by gender and country**.  

---

### Menu

Create a multi-choice menu:

1. Get countries  
2. Get job titles  
3. Get average salary by country  
4. Get average salary by job title  
5. Get average age by gender and country  
q. Exit  

The script should exit **only** when entering the choice `"q"`.  

After executing a function, it should display the results and then return to the starting menu.
