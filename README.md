# sre.monitor.metrics
Export Marvin related information in the prometheus format.

## Go 1.11 Modules

Please upgrade your go version to v1.11+ so that you can use go module. You have to set `GO111MODULE=on`. For more information, please see [golang/go](https://github.com/golang/go/wiki/Modules)
```sh
    $ export GO111MODULE=on
    $ env
```


## Build binary

Executes the following commands right under the root directory of this repository:
```sh
    $ go build -o marvin-exporter cmd/marvin/main.go
```
or
```sh
    $ go build -o marvin-exporter github.com/pnetwork/sre.monitor.metrics/cmd/marvin
```

This both generate the executable binary


## Build image

To build an image, use Dockerfile at the directory:

```sh
    $ docker build -f build/package/marvin/Dockerfile -t {your_image_path_with_tags} .
```

## Run via Docker

You can run like this:
```sh
    $ docker run -idt --name marvin-exporter -p 9987:9987 dockerhub.pentium.network/pnrick/marvin-exporter
```


## Skaffold

This builds and deploy the image to [cr-dev.pentium.network/marvin-exporter] and [cr-dev.pentium.network/faas-exporter]:
```sh
    $ skaffold build --quiet > build.out
```

## Metrics

* http://host:9987/metrics

## Contact

Any questions and feedbacks are so welcome.
* siangye@pentium.network
* rick.lin@pentium.network
# nats.exporter.metrics
