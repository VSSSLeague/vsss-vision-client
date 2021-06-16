[![Organization](https://img.shields.io/:Very%20Small%20Size%20Soccer-vsss--vision--client-82B2CD.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAADjElEQVQ4jSXS2U/bBQDA8R/rapEB6woUaCljSDnKIdfAMRgtlLa0/HofUDqgG8cKCGTcKMcQNgKyRdlGNiUmSsY0Bvc0X6YxmvhgzOKLjz4sMZq4qItxYVx+fdi/8MlHaDDqsVvN2EUbLpcDv99PRUU5ZWUl6PKzKSzKRZefTUF+Dt99+zV7u/9iNtcAf6LLTUPwOG3091xkeKCX0cFB+i70MBNYQCKRsLP8mLLSQvJ1uRQW5KDL05KlzSAjIx2t9jV2nz1BCAbcDPV1MTs1wgw3eEU4gmN8H1m0BOv4X1RUFHGwe51DNsnSZpCenoZSqSQh4QSVZ8oRmn1O+nsukqw8gW/sFyqrbCSnJBATE01mZgZbWxt88+gBwhGBQ+6h0ajZ/uJTFAo5J09qEDrON+P2iBi8GhJTZOzvPcfnd+Jy2mgw6lGrU3n01Zf8/uRnrK5i/P4q0tJUaDRqTmVqEOAQq83EIavss0ioxcL8zDCDb4bp6GjGbDKgUispyMtCFBv47el9MrWJaLMzKC0pQGgNeZFIoujoaKM95GN6vI+F6WEW5kYZu9xDOOTHYKhHpVby2eZtSopyODj4ldJSHXX6MgSv101MTDRJSQm8PTnExx/dZn1tieX5WQ4PduCuCoeop6AwF7U6ldd1WdTWVmIwnMFkrEQINHtQKOQkJSXw/TMnA70RShtbubEyx8zUAEG/SJPZgEFfzd7eC+qN5zCZTFistdTX1SCsvHuNO3fWiYoSiJUdJXL/KaHzzVxdmGCgt51wWwCvy4pTNJGmVnLs2KuIosjVx/9xurwYYX5+DpUqBZlMilQqoSXoA/aZmujHYq2mMxwk1OKmyWZCoZAjlUoYCRUTHx+LtdGIMDE5gsfjIj4+Frk8nlp9Fd2dzdy6OcdPtPLiYIu2Vjdnq07zz99/0Bvp5PnOQ4IBOwGviBC51ElPTxc5OVoc9kbaO4Iv7xfmUnlWhbFJiyxGQld3mKUfDthePUdvd5BIZ5juCwGElqCXtvYW2tpb8HjtWG31KBRycvNOMTs3gaXRiCAIxMUd54Mf4WicwOKVMSZHI4wMdSE4nFY8Xjs+vxOnw0p1zRskJychVxwnJTWRlJSXcDKZlPffW2X9+jU+XFtmdXGKd2aHEZpEM6LdgtlSh9vVhKlBj1204PfYuTx4iem3BlldmmdzY42Hn3/Cg3sbbG/e5ebyHLdWrvA/5UEMH5awhQgAAAAASUVORK5CYII=)](https://github.com/VSSSLeague)
[![CircleCI](https://circleci.com/gh/VSSSLeague/vsss-vision-client.svg?style=shield)](https://circleci.com/gh/VSSSLeague/vsss-vision-client)
![Release](https://badgen.net/github/release/VSSSLeague/vsss-vision-client)

# vsss-vision-client

## Introduction
The vsss-vision-client is an tool created to **IEEE Very Small Size Soccer League**, aiming the possibility to receive vision multicast packages and shows them in a web ui.
This tool was built on the [work](https://github.com/RoboCup-SSL/ssl-vision-client) done for the [Small Size League](https://github.com/RoboCup-SSL) community.

## Requirements
 * Go >= 1.14
 * Node >= 10
 * Yarn

## Usage
If you just want to use this app, you can simply download the latest [release binary](https://github.com/VSSSLeague/vsss-vision-client/releases) that is available for your OS. Note that the binary **is self-contained**, so, no dependencies are required.

**Note:** After download it, you will need to enable the binary to be executed by terminal using `chmod +x ./vsss-vision-client_...` command.

By default, the UI will be available at `http://localhost:8082`, but you can configure this using the args (run the binary with `-h` arg to see them).
 
## Development

### Download
Download and install to GOPATH:

```
go get -u github.com/VSSSLeague/vsss-vision-client/...
```

Switch to project root directory

```
cd $GOPATH/src/github.com/VSSSLeague/vsss-vision-client/...
```

Download dependencies for frontend

```
yarn install
```

### Run
Run the backend:

```
go run cmd/main.go
```

Run the UI:

```
yarn serve
```

### Build self-contained binary
First, build the UI resources

```
yarn build
```

Then build the backend with `packr` module
```
# get packr
go get github.com/gobuffalo/packr/packr

# install the binary
cd cmd/
packr install
```
