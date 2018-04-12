# Digital Forensics Triage Tool

For rapid, multi-platform incident response

### How to build

  - Install go
    - https://golang.org/
  - Navigate to $GOPATH/src
    - [For help run:] go env
  - Clone the repository
    - go get
    - git clone https://github.com/x3419/Capstone.git
  - Set your preferred options in the configuration file
    - Located in Configuraiton/config.txt
    - Formatted in TOML
      - https://github.com/BurntSushi/toml
    - Enable the tools of interest by setting 'enabled' to true or false (boolean, not a string)
    - Set RelativePath (boolean) to set whether file paths are relative (Capstone/Tools directory) or not
    - Set the tools respective argument
  - Build the project
    - cd Capstone
    - go build
  - Run the executable

##### Arguments
  - --config myConfigFile.txt
  - Default is the location on github
    - Capstone/Configuration/config.txt

#### Cross compilation
 If you would like to cross compile (compile for an operating system other than the one currently running):
  - Mac OSX
    - GOOS=darwin GOARCH=386 go build main.go
  - Windows
    - GOOS=windows GOARCH=386 go build main.go
  - GNU/Linux
    - GOOS=linux GOARCH=386 go build main.go
