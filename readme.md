Following the tutorial here:
https://cloud.google.com/appengine/docs/go/gettingstarted/helloworld

Make sure that myapp is under a go [workspace](https://golang.org/doc/code.html#Workspaces)

e.g. Your folder structure should be:

```
echo $GOPATH
# /Users/lancefisher/go

echo $PWD
# /Users/lancefisher/go/src/myapp
```

Starting the app:

```
goapp serve app
```

Local Datastore:

http://localhost:8000/instances

Deploying
---------------

Manage the app here: https://console.developers.google.com/project/trusty-server-724
It's at: http://trusty-server-724.appspot.com

Deploy with:

```
appcfg.py --oauth2 update src/MyApp
```