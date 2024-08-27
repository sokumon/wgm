# How to use this tool

- Install wireguard cli for Ubuntu
```bash 
sudo wgm install
```

- Create initial config (wg.conf)
```bash 
sudo wgm init
```

- Create client.conf
```bash
sudo wgm create-client <name>
```

- List client.conf
```bash
sudo wgm list
```

- Delete client.conf
```
sudo wgm delete-client <name>
sudo wgm delete-client --all
```

