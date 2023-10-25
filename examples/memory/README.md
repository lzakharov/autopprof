# Memory Usage Monitoring

This example starts memory monitoring with a threshold of 1MB and allocates
memory every 100 milliseconds for 10 seconds.

Run example:

```sh
go run main.go
```

Output top entries from the saved profile:

```sh
# note that the file name will be different
go tool pprof -top 20231025T121810_profile.pb.gz 
```

Result:

```
Type: inuse_space
Time: Oct 25, 2023 at 12:18pm (MSK)
Showing nodes accounting for 2.58MB, 100% of 2.58MB total
      flat  flat%   sum%        cum   cum%
    2.58MB   100%   100%     2.58MB   100%  main.allocate
         0     0%   100%     2.58MB   100%  main.main
         0     0%   100%     2.58MB   100%  runtime.main
```
