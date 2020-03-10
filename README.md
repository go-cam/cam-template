# Install

#### 1. Clone cam-template from github

    git clone --depth=1 https://github.com/go-cam/cam-template.git -b v0.3.0

#### 2. Rename folder to your project name

    mv cam-template my-project
    
#### 3. Update dependency Library

    cd my-project
    go mod tidy

#### 4. Build and run server module

    cd server
    go build main.go
    ./main

#### 5. Check whether it runs successfully

Open the browser and visit: http://127.0.0.1:8800/hello


# More
[Document](https://github.com/go-cam/cam-doc)
