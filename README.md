# goscanner
 A port scanner in Go for situations where transferring nmap to the target machine is too much trouble.

goscanner uses multithreading to speed the scan, whitch has pros (speed) and cons (network noise). I'll adress the cons by adding stealthier modes and options on the future.

Usage:
goscanner [target IP]
 
