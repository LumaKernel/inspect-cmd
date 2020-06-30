# inspect-cmd

Command line that tool tells us the real appearance of inputs.

```
$ go get -u github.com/LumaKernel/inspect-cmd
$ go install github.com/LumaKernel/inspect-cmd
```

```
$ inspect-cmd hello world!
Args length: 3
Args[0]:   11 bytes   11 runes
69 6e 73 70 65 63 74 2d  63 6d 64                 |inspect-cmd|
Args[1]:    5 bytes    5 runes
68 65 6c 6c 6f                                    |hello|
Args[2]:    6 bytes    6 runes
77 6f 72 6c 64 21                                 |world!|

stdin: not available
```

```
$ # from bash
$ echo "piped" | inspect-cmd -c <<<2*3~/"  "
Args length: 2
Args[0]:   11 bytes   11 runes
69 6e 73 70 65 63 74 2d  63 6d 64                 |inspect-cmd|
Args[1]:    2 bytes    2 runes
2d 63                                             |-c|

stdin: exists
32 2a 33 7e 2f 20 20 0a                           |2*3~/  .|
```
