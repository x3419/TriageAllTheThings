# **Capstone Final Proposal**


### Metadata

**Name:** Digital Forensics and Incident Response Triage Tool

**Developer:** Alex Bernier

### Project Overview

In today’s world, information security is one of the top growing industries due to reasons that you can probably guess – intellectual property, research and development, or anything else that can be stored on computers are prized assets. With cyber intrusion headlines sweeping the news every day, it’s clear that the approach of most people toward tech security leaves a lot to be desired. When a breach inevitably happens, what does one do? If you’re smart and can afford it, you might consider hiring a cyber security firm to deal with the problem for you – what do they do, though?

When approached with this problem, cyber security firms generally perform an Incident Response (IR) engagement. They take hard drive images of all potentially infected machines (or investigate them physically in-person), use a variety of open source tools to dig into the operating system to find forensic artifacts, and they analyze those artifacts to determine valuable intelligence. This intelligence can provide insight into how the adversaries breached the machines, whether they established lateral movement to other machines, and what activity they performed – specifically, an important detail on their activity would be what type of data they were able to exfiltrate, which allows the company to make important decisions on future steps to take. Lastly, after identifying all this information, they can help clean up the infected systems and ensure that future attacks can be mitigated.

**What is the purpose?**
  - Aggregate many artifact collection tools into a single tool, thereby eliminating an IR analyst having to run each tool individually, saving time and manpower
    - Each tool has various command line arguments, which is a large roadblock when running these tools individually
- I believe that this project would give me a great way to learn various forensic artifacts, their importance, and the type of analysis that would be conducted using them. There are similar triage tools that exist, and they might work just as well as what I would make, but it would be a great learning opportunity.
- What makes this project different than existing ones, is that it's created using Go. This provides a lean, dependency-free, multi-platform environment which is ideal for incident response scenarios - you never know what type of computer might be infected.


**Features and specifics:**
  - Group A
    - Operating System compatibility: Windows
    - Configuration file specifies:
        - Which forensic artifacts/tools to use
        - Destination directory
        - Other miscellaneous tool options
    - Export artifacts to a server (SFTP)
        - If large in size, will export to disk (such as USB drive) instead
        - Uses identifier (such as machine name) to separate and sort forensic evidence
    - Tool integreation:
        - bulk_extractor
            - a computer forensics tool that scans a disk image, a file, or a directory of files and extracts useful information without parsing the file system or file system structures.
        - fiwalk
            - a batch forensics analysis program written in C that uses a forensics toolkit called SleuthKit
        - frag_find
            - a program for finding blocks of one or more master files in a disk image file
        - tcpflow
            - a program that captures data transmitted as part of TCP connections (flows), and stores the data in a way that is convenient for protocol analysis and debugging
        - afxml
            - converts metadata for disk images into a readable format
        - SleuthKit
            - a toolkit with a variety of applications
                - memory analysis
                - file system forensics
                - registry hive dumping
        - MFTDump
            - dumps the master file table
        - NirSoft WinPrefetch
            - dumps the windows prefetch, which contains information on active and historial processes running
        - RegRipper
            - an open source tool, written in Perl, for extracting/parsing information (keys, values, data) from the Registry and presenting it for analysis
  - Group B
    - Operating System compatibility: GNU/Linux
    - Tool integreation:
        - ps
            - dump a list of running processes
        - mrutools
            - linux MRU (most recently used) forensic tool
        - Linux-Forensics-Tool
            - gathers data using:
                - date
                - netstat
                - ps
                - lsof
                - route
                - arp
                - ifconfig
                - top
                - w
                - last
                - uname
                - lsmod
    - BulkExtractor
        - a computer forensics tool that scans a disk image, a file, or a directory of files and extracts useful information without parsing the file system or file system structures.
    - ENT
        - a tool for calculating entropy on the filesystem, indicating whether files are encrypted
    - Fdupes
        - a program for identifying duplicate files residing
within specified directorie
    - Foremost
        - a tool used to recover files based on their headers, footers, and internal data structures
  - Group C
    - Operating System compatibility: Mac OSX
    - Combine artifact formats together uniformally 
    - Tool integreation:
        - OS X Audiotr
            - a free Mac OS X computer forensics tool
        - Knock Knock
            - displays persistent items (scripts, commands, binaries, etc.), that are set to execute automatically on OS X
        - Pac4Mac
            - a forensics framework allowing extraction and analysis session informations in highlighting the real risks in term of information leak
        - Mac OS X Keychain Forensic Tool
            - can extract user credential in a Keychain file with Master Key or user password in forensically sound manner
        - OSXCollector
            -  another broad forensic evidence collection & analysis toolkit for OSX
        
**What this project will not do:**
  - Make your computer secure
  - Automatically analyze forensic artifacts
    - This step must be taken manually by an analyst


**Who's the audience?**
Anyone interesting in forensically analyzing a computer. This could be done for learning puproses, or in a real-world IR professional scenario. 

**Similar Existing Work:**
- https://github.com/mantarayforensics/mantaray
- https://github.com/travisfoley/dfirtriage
- https://github.com/AJMartel/IRTriage
- https://github.com/rshipp/ir-triage-toolkit

**Technology:**
- Language: Go
    - Multi-platform
    - Dependency-free
- Testing
    - Go unit tests
    - https://golang.org/pkg/testing/
- Project Depndencies
    - Should synchronize and resolve dependencies automatically when using "go install"
- Code Style
    - Gofmt is a tool that automatically formats Go source code
    - https://blog.golang.org/go-fmt-your-code
- Quality Control
    - Gofmt and unit tests will confirm that the code is functioning and formatted uniformally and as expected.
    
**Risk Areas**
- Combining artifact formats uniformally
    - This is a huge area of confusion and risk, because there are MANY types of forensic artifacts on various operating systems. I have not tried all the tools that I am going to utilize, so I don't know what format type each tool will output. This endeavour will be experimental and may not be completed. 
- Go unit testing
    - I don't have any experience doing unit tests in Go (or much in any language, for that manner), so testing will be an area of uncertainty. I have located the appropriate documentation, and have the resources to learn what will be necessary. 
- Tool integreation
    - Many of the tools that I plan on integrating I have never used, so there's a high likelihood that problems will arise. When this occurs, I'll continue to troubleshoot. If problems persist, I will naturally have to replace the tool with another candidate that accomplishes the same goal. 

**Un-answered Questions**
- Have you ever used Kali Linux? (random, but seems relevant)
    - Yes! I grew up messing around with Backtrack Linux (which eventually turned into Kali). I had no idea what I was doing but slowly became less of a script kiddie the more I messed around.
- What will using your tool look like? GUI? Unified command line interface?
    - This tool will be a unified command line interface. This is done because in an IR scenario, it's important to have a mechanism by which the user can quickly use the tool in an explicit manner. Furthermore, I want to allow the possibility for future developers to add to this project and implement it within other tools.    
