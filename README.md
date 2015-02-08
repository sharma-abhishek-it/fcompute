# fcompute

[![Build Status](https://travis-ci.org/sharma-abhishek-it/fcompute.svg?branch=master)](https://travis-ci.org/sharma-abhishek-it/fcompute)
[![Coverage Status](https://coveralls.io/repos/sharma-abhishek-it/fcompute/badge.svg?branch=master)](https://coveralls.io/r/sharma-abhishek-it/fcompute?branch=master)

<br>
This repository contains the computation part of fCompute project written in Golang.

To setup first install [Docker](https://docs.docker.com/) on your environment

The project is assembled as

- lib/ - contains utilities like querying redis, precomputing data, models to be used etc
- main/ - http server and controllers. Not using any specific framework rather a very lightweight http middleware named [goji](https://goji.io/). Golang already providers a solid http server library and hence using that right now


-----------------
Setup steps:
1. Install and run Docker (boot2docker on osx)
2. Install [fig](http://www.fig.sh/)
3. In this dir run `fig build app` then `fig up`
4. To run test cases `fig run app go test`
