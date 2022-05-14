# Legends Browser 2 #

Legends Browser 2 is an multi-platform, open source, legends viewer for dwarf fortress 0.47 written in go.
It is a complete rewrite of the original Legends Browser, available at https://github.com/robertjanetzko/LegendsBrowser

### Features ###

* Works in the browser of your choice (just launch an open http://localhost:58881)
* Recreates Legends Mode from dwarf fortress, with objects being accessible as links
* Add several statistics and overviews

### Using Legends Browser ###

* Download the latest release from the downloads page https://github.com/robertjanetzko/LegendsBrowser2/releases
* Run the application
* An browser window should open, if not navigate to http://localhost:58881
* Open a legends export by navigating your file system
* ready to load exports should show up in green
* after loading finished you should see an overview over all civilizations

### Command Line Options ###

```
-p,--port <arg>     use specific port
-s,--serverMode     run in server mode (disables file chooser)
-u,--subUri <arg>   run on /<subUri>
-w,--world <arg>    path to legends.xml or archive
```

### Important Note ###

* some features require the legends_plus.xml from dfhack (run 'exportlegends info')

### Troubleshooting ###

* If you find any bugs, feel free to open an issue here on github
* If you have questions there is a forum thread http://www.bay12forums.com/smf/index.php?topic=155307.0
