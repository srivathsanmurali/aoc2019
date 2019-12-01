# AOC 2019
My solutions to the problems.
*Input is given as stdin*
```
dayXX
|- go
  |- p1.go
  |- p2.go
|- c
  |- p1.c
  |- p2.c
|- elixir
  |- p1.exs
  |- p2.exs
|- input.txt
```

**To run the go code**
```
cd dayXX/go
go run p1.go < ../input.txt
go run p2.go < ../input.txt
```

**To run the c code**
```
cd dayXX/c
make all
./p1 < ../input.txt
./p2 < ../input.txt
make clean
```

**To run the elixir code**
```
cd dayXX/elixir
elixir p1.exs < ../input.txt
elixir p2.exs < ../input.txt
```
