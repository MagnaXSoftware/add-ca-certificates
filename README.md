# add-ca-certificates

[![.github/workflows/test.yml](https://github.com/MagnaXSoftware/add-ca-certificates/actions/workflows/test.yml/badge.svg)](https://github.com/MagnaXSoftware/add-ca-certificates/actions/workflows/test.yml)

`add-ca-certificates` updates the _ca-certificates.crt_ bundle, often located at _/etc/ssl/certs/ca-certificates.crt_ on linux distributions.

Contrary to `update-ca-certificates`, which can remove CA certificates from the bundle, `add-ca-certificates` only ever **adds** new certificates to the bundle.

This project came to be due to an issue in k3OS ([#518](https://github.com/rancher/k3os/issues/518)), where running `update-ca-certificates` would cause all existing trusted ca-certificates to be removed from the bundle.
This is an issue in many corporate environments, as the entreprise PKI should be trusted, but the public PKI should not be distrusted.

## Usage

    add-ca-certificates [--bundle path-to-the-bundle] [--local-path path/to/the/locally/trusted/certificates]

`add-ca-certificate` will now maintain the existing order of the cert bundle and insert the new certificates at the end, this will reduce the instability of the resulting cert bundle.
