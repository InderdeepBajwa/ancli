# Ancli

Ancli is a CLI to create and maintain Ansible projects.

Inspired by the [office Ansible directory layout recommendations](https://docs.ansible.com/ansible/2.8/user_guide/playbooks_best_practices.html#directory-layout)

## Installation

1. Download [the latest binary release by clicking here](https://github.com/InderdeepBajwa/ancli/tags)
2. Drop the `ancli` binary into a preferred path
3. Binary can be run using `ancli` if added to the $PATH or `/path/to/ancli`. See commands below.

## Usage

### Create a new Ansible project and role(s)

`ancli create --name <name-of-project>`

Example: `ancli create --name myAnsibleProject`

Need to add roles? Simply reply with `y` to the prompts and create as many roles as needed.

### Create a new role in any Ansible project
`ancli new-role --name nameOfRole --project nameOfProject`

Example: `ancli new-role --name webservers --project primaryinfrastructure`

#### Full new project create example:

```console
$ ancli create --name testProject
Do you want to create Ansible role in this project? Y/n y
Enter the name of role: configureWebServers
Creating role configureWebServers...
Do you want to create Ansible role in this project? Y/n n
Project created successfully.
```
The above example creates the following project structure:

![image](https://user-images.githubusercontent.com/20612193/213058705-8844edfe-d7eb-467f-bf2b-7094e8daffb6.png)
