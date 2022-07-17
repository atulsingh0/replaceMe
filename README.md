# replaceme
This utility is used to replace a string in a file and write into output file

## Usage

```bash
Usage of ./replaceme:
  -i string   Input file path
  -o string   Output file path
  -m map      key-value map where key will be replace by value in inputfile
```

```bash
$ ./replaceme -i "/tmp/data_in" -o "/tmp/data_out" -m '{"NAME": "atul", "JOB" : "SE" }'

$ ./replaceme -i "/tmp/test" -m '{"NAME": "atul", "JOB" : "SE" }'
```
