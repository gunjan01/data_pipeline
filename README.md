## CSV Processor

## High level Overview of the implementation

The idea behind the implementation is to build a simple data pipeline, which takes a CSV, mutates some of the fields
and dumps it into an elastic index. With each day, as we pull fresh data from the google ads api, we process the csv in the pipeline and perform an update operation on the elastic index. The older documents remain and new ones are added. The project has a logstash configuration to dump data into elastic index which would be the ideal way to go about it.

To populate the index using logstash:
1. Install logstash
2. Fire up elastic search using `docker-compose up elasticsearch`
2. Run `sudo bin/logstash -f <your path>/go/src/github.com/gunjan01/data_pipeline/logstash/logstash.conf`

The code for logstash pipeline can be found in the logstash.conf in this project. For the purpose of this assignment,
I have written a simple function that reads a CSV line by line and dumps it into an index once you start the project.

After loading the data into elasticsearch, I have exposed a simple endpoint to query it. Once the project is fired,
a simple http server runs on `9090` and the endpoint can be hit to see the squashed data for a specific date range. A sample would like: `http://localhost:9090/data?start_date=2020-06-01&end_date=2020-06-02`

The heart of the implementation is in the search package inside source. The package has aggregation and filter queries that are performed on the index to squash data according to the requirement. Data is squashed according to
the date range passed. The implementation aggregates the data on the dimensions and applies a sum to it.

If you rebuild the project, you can use `curl -XDELETE http://localhost:9200/dimensions` to delete the previous index.

Disclaimer: While loading the CSV into the index, I have used the first value of type float. This is because if elastic encounters an Integer value on the first insertion field, it dynamically maps the field to type long which will end up rounding off sum aggregation results. Therefore, I changed the first field to a floating number in the CSV and opted to not write a seperate mapping file for dumping data for the purpose of this assignment. Elastic will
dynamically map the right types on insertion.

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
