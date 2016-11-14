# expect
light + simple test tool for my own small modules


### notes
- **fastest golang install**
 - ```mkdir temp; cd temp; sudo wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz; sudo tar -C /usr/local -xzf go1.7.3.linux-amd64.tar.gz; export PATH=$PATH:/usr/local/go/bin; cd ..; rm -rf temp; mkdir -p ./src/github.com/wambosa/```


### sensible external tests
- ```git clone git@github.com:username/REPO_NAME.git``` should be run in the username folder (ex: ```./src/github.com/wambosa/```)
- because I sometimes like to test APIs externally, golang import paths must be tricked
- develop in ```~/workspace/src/github.com/wambosa/REPO_NAME```
- package files must be stored at the fully qualified import path **even during dev time**
- git repo should be **inside each individual packageName folder**
- when I later use ```go get ...```, the repository will not know about the development environments extra folders
- ```go test -v``` will now work as expected both in dev environment and when obtained via ```go get```
- this is all because the default GOPATH will look at ```~/workspace/src/...``` for every single ```import``` statement