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

-----------------
Inetgration steps
Make sure that both repos fcompute and fcompute_frontend_server are inside the same directory
which we will call as root for time being. This will be our working dir for the entire process

- Stop and kill all containers `docker stop $(docker ps -a -q)` and `docker rm $(docker ps -a -q)`
- Copy *Data* directory to *fcompute_frontend_server/tmp*
- Copy *fig_integration.yml*, *fig_test.yml*, *fig_benchmark.yml* to the one level up(root) dir
- Copy *fig_sidekiq.yml* from frontend_web_server to root dir
- Run `fig -f fig_integration.yml run sidekiq bundle exec rake db:setup`
- To simply start server for checking api requests do `fig -f fig_integration.yml up`. Make sure to make your first request only after sidekiq has run once to get the Data into Redis. It runs every minute but that can be changed in code. On successful completion of sidekiq process Data dir is deleted from tmp.

For benchmarking or testing
- Copy *Data* directory to *fcompute_frontend_server/tmp*
- Run `fig -f fig_sidekiq.yml run sidekiq bundle exec rake db:setup`
- Run `fig -f fig_sidekiq.yml run sidekiq`
- Wait for sidekiq to run once i.e Data dir deleted from tmp
- Stop this sidekiq
- For testing do `fig -f fig_test.yml run fcompute`
- For benchmarking do `fig -f fig_benchmark.yml run fcompute`
