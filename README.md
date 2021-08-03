# add-ca-certificates

`add-ca-certificates` updates the _ca-certificates.crt_ bundle, often located at _/etc/ssl/certs/ca-certificates.crt_ on linux distributions.

Contrary to `update-ca-certificates`, which can remove CA certificates from the bundle, `add-ca-certificates` only ever **adds** new certificates to the bundle.

This project came to be due to an issue in k3OS ([#518](https://github.com/rancher/k3os/issues/518)), where running `update-ca-certificates` would cause all existing trusted ca-certificates to be removed from the bundle.
This is an issue in many corporate environments, as the entreprise PKI should be trusted, but the public PKI should not be distrusted.

## Usage

    add-ca-certificates [--bundle path-to-the-bundle] [--local-path path/to/the/locally/trusted/certificates]

The first time that `add-ca-certificates` runs, it may re-order the certificates already present in the bundle, as the ordering used internally might not match with the previously used ordering.
