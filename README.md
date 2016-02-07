# clustereasy

## Input

```yaml
cluster:
  token: rio
  state: new

  nodes:
  - name: campos
    interface: enp3s4
    ip: 192.168.0.1

  - name: caxias
    interface: enp3s4
    ip: 192.168.0.2

  - name: macae
    interface: enp3s4
    ip: 192.168.0.3

  - name: resende
    interface: enp3s4
    ip: 192.168.0.4

  - name: arraial
    interface: enps3s4
    ip: 192.168.0.5

  kubernetes:
    deploy: true
    master: 
    - campos

  units:
  - name: sshd.socket
    command: start
```
