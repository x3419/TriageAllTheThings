# Digital Forensics Triage Tool

For rapid, multi-platform incident response

NOTE: The Windows tools included in this repository serve as a starting example configuration. It is recommended that you add your own tools and set the arguments to your own specifications. I have not included the non-default GNU/Linux or Mac OSX tools in this repository because a compilation process is necessary for most open source tools, and different distributions and architectures will compile using different configurations.

Lastly, please pay special attention to the arguments and be sure to set --portable=false when using it without portability 

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
  - --portable = false
    - Default is true
    
##### Making a portable executable
  - Bundles tools within the executable or as an archive
  - Implemented using go.rice
    - https://github.com/GeertJohan/go.rice
    - go get the package and go build within the github.com/GeertJohan/rice path
    - add the github.com/GeertJohan/rice path to your environmental variables
    - cd back into Capstone
    - rice go-embed
      - for tools being within the executable
    - rice embed-syso
      - for generating a coff .syso file archive that must be in the same folder as the executable
    - go build
  - NOTE: When using go.rice to bundle the tools within the executable, you must have your configuration file within the "Configuration" folder
  

#### Cross compilation
 If you would like to cross compile (compile for an operating system other than the one currently running):
  - Mac OSX
    - GOOS=darwin GOARCH=386 go build main.go
  - Windows
    - GOOS=windows GOARCH=386 go build main.go
  - GNU/Linux
    - GOOS=linux GOARCH=386 go build main.go
