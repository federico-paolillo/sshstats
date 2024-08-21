# What's in here

This folder will contain all `<nodename>.json` files with login attempts statistics. Such files will be downloaded by the CI pipelines and exposed for consumption to Nuxt. This removes the need to call to provide the API Keys to the client. These files will be served using `serve` in debugging or when generating the website.

Run `npm run data` to load the static file server that will serve the entire `/data` folder.
