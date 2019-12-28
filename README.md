# Find blocking syscalls in `strace` output

A simple little tool to filter out brief syscalls from `strace` output and find what is blocking a process.

Just use `make` to build it (go 1.13 is recommended).

Run `strace` with `-T` flag and send it's output to stdin of the tool:

    $ strace -tT -o '!./bin/skip-brief-syscalls -d 0.5s' sleep 1
    22:35:36 nanosleep({1, 0}, NULL)        = 0 <1.000180>

You could also attach `strace` to existing process (nginx worker in this example) and find it's blocking spots:

    # strace -tT -o '!./skip-brief-syscalls -d 0.1s' -p 32222
    strace: Process 32779 attached
    23:11:46 epoll_wait(604, [{EPOLLOUT, {u32=1011222333, u64=140674777075222}}], 512, 128) = 1 <0.121433>
    23:11:46 epoll_wait(604, [], 512, 108)  = 0 <0.109150>
    23:11:46 epoll_wait(604, [{EPOLLOUT, {u32=1013336688, u64=140674777482777}}], 512, 308) = 1 <0.187595>
    ^Cstrace: Process 32779 detached
