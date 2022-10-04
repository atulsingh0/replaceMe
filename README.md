# replaceme
This utility replaces string in a file from the passed JSON or the Environmnt var and write into output file

# Inspiration
In k8s world, sometime we have to replace values in config map at run time. This utility can help to do that if used in init container to replace all the instances of particular String.
Work is still in progress, Feel free to drop any enhancement and fixes.

## Usage

```bash
Usage of ./replaceme:
  -i string   Input file path
  -o string   Output file path [optional]
  -m map      key-value map where key will be replace by value in inputfile [optional]
              if key will overwrite the env var if having the same name
```

```bash
# Input - /tmp/data_in
# Output - /tmp/data_out
# Key-Map - '{"NAME": "atul", "JOB" : "SE" }'
# It will replace all instances of NAME, JOB and any instance of Env var with its value and create an output file
$ ./replaceme -i "/tmp/data_in" -o "/tmp/data_out" -m '{"NAME": "atul", "JOB" : "SE" }'

# Input - /tmp/test
# Output - /tmp/test
# Key-Map - '{"NAME": "atul", "JOB" : "SE" }'
# It will replace all instances of NAME, JOB and any instance of Env var with its value in input file only
$ ./replaceme -i "/tmp/test" -m '{"NAME": "atul", "JOB" : "SE" }'

# Input - /tmp/test
# Output - /tmp/test
# It will replace any instance of Env var with its value in input file only
$ ./replaceme -i "/tmp/test"
```
