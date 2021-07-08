# simplesite

## TODO

- get Vugu static HTML output working (template and body separate, probably header and footer too)
  - start with a simple single page that is outside of the content thing in it's own handler, hacked into the simplesite project
- spend a few hours fiddling with the design so it looks half-way decent, maybe ask DC to help
- hammer out a working initial version of the read-only aspect of the CMS outputting pages
  - loads files from disk
  - caches in memory (with interfaces possible for other stores later)
  - do not worry about editing yet
  - should be able to make a basic site show up just by editing data files on disk, config
  - will need to decide on data format (must preserve comments!)
- see if we can get Yoast-like functionality in here with SEO stuff happening
- regroup, decide if we want to start on Admin UI or work on the CMS API endpoints

- Maybe we need a separate server as a sort of watchdog application that handles filesystem changes and reversion, etc. This way if a source code change renders the project un-buildable, you can unbrick your app.
- document how to roll this into a container
  - needs to handle edits in dev scenario
  - could be two different approaches, one that is static and simple and does compilation during the container build and another that supports edits and does the generate/compile/etc inside the running container
