# go-ec2-instance-metadata
An example in GO to retrieve meta-data from a AWS EC2 instance like the official tool: [https://aws.amazon.com/pt/code/ec2-instance-metadata-query-tool/](https://aws.amazon.com/pt/code/ec2-instance-metadata-query-tool/)

## Building and running

To build the project simple run the __env__ command followed by the __go build__ command.

For Linux instances:

```sh
$ env GOOS=linux GOARCH=amd64 go build ec2-instance-metadata
```

Copy the binary to your instance and set the execute flag to it:

```sh
$ chmod +x ec2-instance-metadata
```

To execute it:

```sh
$ ./ec2-instance-metadata
```

It will display some of the available options, for a complete options list click [here](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)


