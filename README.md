# treego

An implementation of the Unix [tree](https://linux.die.net/man/1/tree) command in Go.

Just a way of practicing Go a bit more.

![screenshot](https://yld.me/raw/baBn.png)

### Options implemented

```
	--help
	-a 	    All files are listed
	-d 	    List directories only
	-f      Print full path prefixes
	-i      Do not print any indendation prefixes
	-s      Print size of each file
```

### TODO

1. There's a bunch more flags listed on the man page. I might add some more simple ones like checking uid or gid.

2. Show/expand symlinks (i.e., fakefile.txt -> /usr/share/fakefile.txt)

3. Add colors?

I hope to write this in Rust and maybe C eventually.
