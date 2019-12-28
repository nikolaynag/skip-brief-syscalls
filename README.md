# Skip lines with brief syscalls in `strace` output

A simple little tool to filter out brief syscalls from `strace` output.
Just execute `make` to build it (go 1.13 or newer is recommended).

Use it to find blocking syscalls in `strace` output.

Run `strace` with `-T` flag and send it's output to stdin of the tool:

    $ strace -tT -o '!./bin/skip-brief-syscalls -d 0.5s' sleep 1
    22:35:36 nanosleep({1, 0}, NULL)        = 0 <1.000180>
