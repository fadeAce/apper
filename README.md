#### Apper is a high-performance scrapper

###### for starters to build project apper

###### you should first install go gb

    go get github.com/constabulary/gb/...

###### and make sure $GOPATH/bin is added to your $PATH

###### then cd $PROJECT_PATH (apper project location)

    gb vendor restore
    bash run.sh

it'll generate a apper execute file and run it in -i option automatically
if you want to run it in daemon mode as apper server please see ./apper -h for more information
###### Finally , add $PROJECT_PATH/vendor and $PROJECT_PATH to $GOPATH

----

###### the architect of pipes on processing task

```
        pip ---+ task +------- finish
         +		  |				+
         |		  |				|
         |		  +				|
         |		fragment ---+ done
         |						|
         |						|
         |						+
         +----------------- in progress
```

