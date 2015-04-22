# go-compositor performance testing

## Requirements

* go 1.4.2
* libmagickcore

Installing Go is left as an exercise for the reader

To get the latter on OSX, `brew install imagemagick`.

On Ubuntu, `sudo apt-get -y imagemagick libmagickcore-dev pkg-config`

## Usage

The default test is to build a 35-page book preview with 4 layers per
page. On my machine (late 2013 MBP 15"):

```
$ go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkGeneratePage                  20       69766361 ns/op
BenchmarkGeneratePagesSync             1        2469893604 ns/op
BenchmarkGeneratePagesAsync            1        1809076247 ns/op
ok      github.com/Lostmyname/go-compositor     5.793s
```

```
$ GOMAXPROCS=8 go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkGeneratePage-8                20       69413212 ns/op
BenchmarkGeneratePagesSync-8           1        2414899855 ns/op
BenchmarkGeneratePagesAsync-8          1        1840543020 ns/op
ok      github.com/Lostmyname/go-compositor     5.750s
```

Notably, no speed increase whatsoever is achieved by increasing
GOMAXPROCS, suggesting that the composition is CPU bound.

## CPU Profiling

To investigage this, go's built-in cpu profiling output was used.
Unfortunately this doesn't work on OSX without a kernel hack, so a Linux
VM was used (this is not ideal but can give broad brush-stroke
measurements):

```
$ go test -c

$ ./go-compositor.test -test.cpuprofile=cpu.out
-test.bench=GeneratePagesAsync
testing: warning: no tests to run
PASS
BenchmarkGeneratePagesAsync-8          1        2266496545 ns/op

$ go tool pprof go-compositor.test cpu.out
Entering interactive mode (type "help" for commands)
(pprof) top20
1380ms of 1380ms total (  100%)
      flat  flat%   sum%        cum   cum%
    1360ms 98.55% 98.55%     1360ms 98.55%  runtime.cgocall_errno
      20ms  1.45%   100%       20ms  1.45%  runtime.rtsigprocmask
         0     0%   100%       20ms  1.45%  System
         0     0%   100%     1360ms 98.55%  github.com/Lostmyname/go-compositor.GeneratePage
         0     0%   100%     1360ms 98.55%  github.com/Lostmyname/go-compositor.func·001
         0     0%   100%      280ms 20.29%  github.com/Lostmyname/magick.(*MagickImage).Compose
         0     0%   100%       20ms  1.45%  github.com/Lostmyname/magick.(*MagickImage).ReplaceImage
         0     0%   100%      810ms 58.70%  github.com/Lostmyname/magick.(*MagickImage).ToFile
         0     0%   100%      270ms 19.57%  github.com/Lostmyname/magick.NewFromFile
         0     0%   100%      260ms 18.84%  github.com/Lostmyname/magick._Cfunc_ComposeSourceWithImage
         0     0%   100%       20ms  1.45%  github.com/Lostmyname/magick._Cfunc_DestroyImage
         0     0%   100%      270ms 19.57%  github.com/Lostmyname/magick._Cfunc_ReadImage
         0     0%   100%      810ms 58.70%  github.com/Lostmyname/magick._Cfunc_WriteImages
         0     0%   100%     1360ms 98.55%  runtime.goexit
(pprof)
```

This confirms the suspicion above that we're IO bound - 80% of the
runtime is spent in ToFile() and NewFromFile().
