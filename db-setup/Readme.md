#### Docker commands:

**!** Sql script is only for setup in development not actual usage.

Build image:

From folder db-setup run

`docker build .`

Create local development container:

`docker run -d -p 5432:5432 -v {path to local storage}:/var/lib/postgresql/data --name {name of container} {name of image}`
