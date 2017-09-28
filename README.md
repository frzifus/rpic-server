# Info
This project was created for learning purposes.

# Source

##### golang:
```sh
go get git@github.com:Frzifus/rpic-server.git
```

##### ssh:
```sh
git clone git@github.com:Frzifus/rpic-server.git
```

##### https:
```sh
git clone https://github.com:Frzifus/rpic-server.git
```

# Building

You will need golang (1.6 or newer) and a golang protobuf.
Currently supported/tested are:
 - Raspberry PI 2
##### Build => "./build/bin/":
```sh
make
```

# Test

Write output to ./build/log/test_[date].log
##### Run tests:
```sh
make test
```

##### Connection test with nc:
```sh
# listen
nc -l -p 4444
# send
cat ./test/vehicleY1000.bin | nc 127.0.0.1 4445
```

# Camera
