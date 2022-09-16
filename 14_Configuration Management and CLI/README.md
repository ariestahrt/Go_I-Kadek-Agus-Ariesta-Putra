# SUMMARY

## 1. Command Line

Command line is a quick powerful, text-based interface developers use to more effectively and efficiently communicate with computers to accomplish a wider set of tasks.

Command line is useful for
- Granular control of an OS or application
- Faster managemenet of a large number of operating system.
- Ability to store scripts to automate regular tasks.
- It interface knowledge to help with troubleshooting, such as network connection issues.

There are common CLI (Command Line Interface):

- Shell = CLI for OS' Services
    - UNIX Shell
    - Command Prompt (MSDOS)

- Other app CLI
    - Python REPL
    - MySQL CLI client
    - Mongo shell
    - redis-cli

## 2. Intro to UNIX Shell
**UNIX Shell Legend**:

```
me@localheart:~$ _



```

- "me" = current user
- "localheart" = hostname
- "~" = current path (~=home)
- "$" Shell for normal user, "#" for root user

**Most popular command**:

- Directory things: `pwd,ls,mkdir,cd,rm,cp,mv,ln`
- Files things: `touch, head,cat,tail,less,vim,nano,chown,chmod,diff`
- Network things: `ping,ssh,netstat,nmap,wget,curl,telnet,netcat`
- Utility: `man,env,echo,date,which,watch,sudo,history,grep,locate`

## 3. Shell Script
**Shell** is program that functions as a bridge between the user and the kernal.

**Shell Script**, a programming language compiled based on shell commands.

Shell script contains shell commands that compiled into a files. For example:

```sh
echo "HALO" >> test.txt

# Print directory
ls -la

isnumber(){
    re='^[0-9]+$'
    if [[ ${1} =~ $re ]] ; then
        echo "Y"
    else
        echo "N"
    fi
}

printf "[+] Input number : "
read num
isnumber=$(isnumber "$num")
while [[ $isnumber = "N" ]]; do
    printf "\t[-] ERROR | Please input valid number : "
    read num
    isnumber=$(isnumber "$num")
done
```