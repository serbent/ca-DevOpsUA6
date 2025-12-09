# 20. Ansible. 

## Roles

Get help for Andible roles command: 
```bash 
ansible-galaxy role -h
```

List all available role in ansible Galaxy repository:
```bash
ansible-galaxy role list
```

Search for a specific role:
```bash
ansible-galaxy role search nginx
```

Search for a specific role by author:
```bash
ansible-galaxy role search nginx --author geerlingguy
```

Install a role globaly:
```bash
ansible-galaxy role install geerlingguy.nginx
```
