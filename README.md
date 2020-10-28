## CSV Processor

## High level Overview of the implementation

The idea behind the implementation is to build a simple data pipeline, which takes a CSV, mutates some of teh fields
and dumps it into an elastic index. With each day, as we pull fresh data from the google ads api, we process the csv in the pipeline and perform an update operation on the elastic index. The older documents remain and new ones are added. The project has a logstash configuration to dump data into elastic index which would be the ideal way to go about it.

To populate the index using logstash:
1. Install logstash
2. Fire up elastic search using `docker-compose up elasticsearch`
2. Run `sudo bin/logstash -f <your path>/go/src/github.com/gunjan01/data_pipeline/logstash/logstash.conf`

The code for logstash pipeline can be found in the logstash.conf in this project. For the purpose of the assignment,
I have written a simple function that reads a CSV line by line and dumps it into an index.

After loading the data into elasticsearch, I have exposed a simple endpoint to query it. Once the project is fired,
a simple http server runs on `9090` and the endpoint can be hit to see the squashed data for a specific datae range. A sample would like: `http://localhost:9090/data?start_date=2020-06-01&end_date=2020-06-02`

The heart of the implementation is in the search package inside source. The package has aggregation and filter queries that we perform on the index to squash data according to the requirement.

## Setup

Run following commands to try it out:

* Clone the project.
* Install the dependencies.
* Fire up elastic search using `docker-compose up elasticsearch`
* Run `make build` to build and run binaries. A sample built binary can be found under bin.
  You could even use the dockerfile to run the project.
* Hit the endpoint at running at `localhost:9090` and pass appropriate parameters to see the
  collated data returned in the response.

## Manual

```
  λ  make

  build     Build the binaries
  lint      Check for linting errors
  format    Check for formatting errors
  tools     Installs tools
  test      Run tests

```
