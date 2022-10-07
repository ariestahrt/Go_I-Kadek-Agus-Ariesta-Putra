# SUMMARY

## 1. About Docker & Container

Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. A container is a process with file system isolation. A container is not a virtual machine.

## 2. Differences between Container and Virtual Machines

| Container | Virtual Maachines |
| - | - |
| Abstraction at the app layer | Abstraction of physical hardware |
| Containers take up less space than VMs (contaainer images are typically tens of MBs in size) | Each VM includes a full copy of an operating system |
| Handle more applications and require fewer VMs and Operating Systems | Also be slow to boot |

## 3. Docker Basic

### Syntax

| Syntax | Definition |
| - | - |
| FROM | Getting image from the docker registry |
| RUN | Execute bash command when building container |
| ENV | Set varible inside container |
| ADD | Copy the file with some other process |
| COPY | Copy the file |
| WORKDIR | Set working file directory |
| ENTRYPOINT | Execute command when finish building container |
| CMD | Execute command but can be overwrite |
