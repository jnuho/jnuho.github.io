
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
```

### Linux File System

- check default shell
```sh
cat /etc/passwd | grep foo
echo $0
```

- wildcards
  - ```*``` : any string of none or more characters
  - ```?``` : single character
  - ```[]``` : any of the characters  inside the brackets
- metacharacters
  - ```>``` : ouput redirection
  - ```>>``` : ouput redirection - append
  - ```<``` : input redirection
  - ```<<``` :  input redirection

