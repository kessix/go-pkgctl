# PKGCTL

## About

Binary GO to validate package format in zip format.

### Packages Format

- Zip package

{ApplicationName}_{Release}_{Version}.zip

Inside zip has a folder with the application files. The folder must be the same with zip {ApplicationName}.
If files are zipped without a folder inside the zip, the deploy app will break.

The name of the zip package must be {ApplicationName}_{Version}.zip, the {Release} field will be removed by pkgctl deuring deployment.

Therefore, pkgctl ensure the valid package format for deploy.

### Examples 

$ pkgctl check
$ pkgctl format





