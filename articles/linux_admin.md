- Mastering Linux Administration

### Create User

```sh
su root
useradd foo
pass foo

# Add user to the sudoers group
usermod -aG wheel foo
su - foo
whoami
```

### Package Management
```sh
# apt is a frontend to dpkg
dpkg --list {name}
dpkg --remove {name}
apt remove {name}
````

### Linux File System

- check default shell

```sh
cat /etc/passwd | grep foo
echo $0
```

- wildcards
  - `*` : any string of none or more characters
  - `?` : single character
  - `[]` : any of the characters inside the brackets
- metacharacters
  - `>`: Output redirection.
  - `>>`: Output redirection – append.
  - `<`: Input redirection.
  - `<<`: Input redirection.
  - `*`: File substitution wildcard (explained previously).
  - `?`: File substitution wildcard (explained previously).
  - `[ ]`: File substitution wildcard (explained previously).
  - `|`: Pipe for using multiple commands.
  - `;`: Command execution sequence.
  - `( )`: Group of commands in the execution sequence.
  - `||`: Conditional execution (OR).
  - `&&`: Conditional execution (AND).
  - `&`: Run a command in the background.
  - `#`: Use a command directly in the shell.
  - `$`: Variable value expansion.
  - `\`: The escape character.
  - `cmd`: Command substitution.
  - `$(cmd)`: Command substitution.

```sh
ls -l /etc/ | less

touch myfile-1
touch myfile-2
# brace expansion
rm ~/{myfile-1, myfile-2}
```

- Bash Shell variables
  - HOME
  - LOGNAME : login name
  - PWD
  - OLDPWD
  - PATH
    - This variable is essential in Linux. It helps the shell know where all the programs are located
    - When you enter a command into your Bash shell, it first has to search for that command through the Linux filesystem
    - To make any changes permanent, you must modify the PATH variable inside a file called `~/.bash_profile` or `~/.bashrc`
  - SHELL
  - USER
  - TERM

```sh
printenv | grep PATH
echo $PATH
```

- Links
  - Symbolic Links
  - Hard Links
  - [Difference between soft vs hard](https://stackoverflow.com/a/185903)

```sh
# Soft Link
# ln -s [original_filename] [link_filename]
touch a.txt
ln -s a.txt b.txt
echo "aaa" >> a.txt
cat a.txt
cat b.txt
# return the name of the file the symbolic link points to
readlink b.txt

# Hard Link
# different virtual file that points to the original file
# they are physically the same
touch a.txt
ln a.txt b.txt
echo "aaa" >> a.txt
cat a.txt
cat b.txt
readlink b.txt

# more detailed info-name,size,type
stat a.txt
# report type of a file
file a.txt
```


- `less` print one page and allows navigation
- `head` print first 10 lines
- `tail` print last 10 lines
  - `-f` output appends as the file grows


