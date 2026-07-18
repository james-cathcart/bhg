# Metasploit Minimal

## Metasploit Documentation

https://docs.rapid7.com/metasploit/rpc-api/

**Start Metasploit***
```
msfconsole
```

**Start Local RPC Server**

```
msf > load msgrpc Pass=s3cr3t ServerHost=127.0.0.1
```

**Output**

```
[*] MSGRPC Service:  127.0.0.1:55552
[*] MSGRPC Username: msf
[*] MSGRPC Password: s3cr3t
[*] Successfully loaded plugin: msgrpc
msf >
```

## Go Lib
```
go get gopkg.in/vmihailenco/msgpack.v2​​
```

## Framework Handlers

Handlers include 
- core
- auth
- console
- module
- session
- plugin
- job
- db

## Metasploit Framwork RPC Server

The Metasploit Framework RPC server requires a username and password to be specified. This username and password combination can be used with the auth.login API to obtain a temporary token that will grant access to the rest of the API.