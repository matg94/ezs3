# Easy S3

Simple Golang CLI / Library to upload, download, or delete files from S3.

## Installation

You can install this CLI using the following command:
`go install github.com/matg94/ezs3`

## Usage

Once installed, make sure it is set up correctly in your Gopath.
You can run `ezs3 -h` to see more information about usage.

### Setting up Credentials

This CLI will use credentials in `~/.aws/credentials`, to authenticate with S3.

### Flags

These are the flags available to use with the CLI.
You can Download and Delete at the same time (downloads, and then deletes from S3), but you can't
upload and download or upload and delete.

The `-f <filepath>` flag, will be the `origin` path, so if you wanted to upload:
`~/testfile.txt`, you would use `-f ~/testfile.txt`, but would need to set the target path in S3 using `-t` e.g `-t test/testfile.txt`.

If you do not set a `-t` flag, the default will be whatever you've put as `-f`.

deletion just uses to `-f` path, and will delete the file in S3 at that location.

`ezs3` can use the bucket and endpoint from your environment. These would need to be set under the names
`AWS_BUCKET`, and `AWS_ENDPOINT`, respectively.

You will need to set the `-e endpoint -b bucketname` flags as well.


```
Usage of ezs3:
  -b string
        Specify the S3 target Bucket. (default "-")
  -delete
        Deletes the origin file from S3.
  -download
        Downloads the target file from S3.
  -e string
        Specify the S3 target Endpoint. (default "-")
  -f string
        Specify the origin path for the file you'd like to upload / download. (default "-")
  -t string
        Specify the destination path where you would like the file saved. Default is origin. (default "-")
  -upload
        Uploads target file to S3.
```

## Examples

Uploading `config/prod.yaml` to S3 folder `config/production/application.yaml`:

`ezs3 -upload -e <endpoint> -b <bucketname> -f config/prod.yaml -t config/production/application.yaml`


Download `config/production/application.yaml` to `config/application.yaml`:

`ezs3 -download -e <endpoint> -b <bucketname> -f config/production/application.yaml -t config/application.yaml`


Deleting `config/production/application.yaml` in S3:

`ezs3 -delete -e <endpoint> -b <bucketname> -f config/production/application.yaml`