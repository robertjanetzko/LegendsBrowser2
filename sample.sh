#!/bin/bash

curl http://localhost:8081/debug/pprof/heap > heap.0.pprof
go tool pprof heap.0.pprof