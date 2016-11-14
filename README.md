# expect
light + simple test tool for my own small modules


### notes
- **fastest Cloud9+golang install**
 - GOPATH && GOROOT are preconfigured
 - ```mkdir temp; cd temp; sudo wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz; sudo tar -C /usr/local -xzf go1.7.3.linux-amd64.tar.gz; export PATH=$PATH:/usr/local/go/bin; cd ..; rm -rf temp; mkdir -p ./src/github.com/wambosa/; cd ./src/github.com/wambosa```


### sensible external tests
- because I sometimes like to test APIs externally, golang import paths must be tricked
- package files must be stored at the fully qualified import path **even during dev time**
- ```git clone git@github.com:username/REPO_NAME.git``` should be run in the username folder (ex: ```cd ./src/github.com/wambosa/```)
- develop in ```cd ~/workspace/src/github.com/wambosa/REPO_NAME```
- git repo should be **inside each individual packageName folder**
- when I later use ```go get ...```, the repository will not know about the development environments extra folders
- ```go test -v``` will now work as expected both in dev environment and when obtained via ```go get```
- this is all because the default GOPATH will look at ```~/workspace/src/...``` for every single ```import``` statement
- for internal tests, this is less of an issue, since ew do not import the package being tested. also has full access to internal funcs
- 

### example dev environment
```
~/workspace
    |
    - src
        |
        - github.com
            |
            - wambosa
                |
                - expect
                    |
                    - .git
                    - .gitignore
                    - expect.go
                    - expect_test.go
                    - README.md
```