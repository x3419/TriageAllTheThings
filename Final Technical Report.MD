# Final Technical Report

Digital Forensics Triage Tool

Alex Bernier

04/30/2018

### Abstract
For my capstone project, I decided to create a digital forensics triage tool. The motivation behind this project was illuminated through my experience at my internship over the summer - I found that, while learning to use various digital forensics tools, it would require extra programming to bundle different tools together for deployment; this was a static process that would inevitably need to be modified when adding tools and features. While manually programming a way in which tools can be triaged (organized based on priority) is definitely a way of precisely providing a software solution for this problem, the static nature of how it was implemented had room for improvement. This project is aiming to solve that problem, by providing an interface that can automate the process of deploying various tools dynamically, without providing extra programming. Though this project has application beyond digital forensics, this was my initial motivation so I have catered the UI and features around this target audience. Furthermore, this tool can be used with or without the use of deployment. In my testing, I have found that the results of this project reflect my initial goal, and I hope that the users of this software find that as well. 

##### Keywords:
Digital 
forensic 
triage 
cyber 
information security 
tool 
incident response 
IR 
infosec 
deployment 
deploy 
bundle 
package 
go 
golang 
multi platform 
ui 
capstone 
appalachian state

### Table of Contents
Introduction and Project Overview

Design, Development and Test

Results 

Conclusions and Future Work

References

Document Quality

### Introduction and Project Overview


##### Background and Alternative Solutions
As a digital forensic or incident respont analyst, your job is to examine computers generally involved in some type of criminal behavior. This is done by analyzing various forensic artifacts on the computer; frequently (almost always in one way or another) when something is clicked, written or read on a computer, there is a log of it. It's not always a human friendly log like you might expect, and may require advanced methodologies in order to understand the logs, but they do exist. They're stored in various places by the operating system (such as Windows, Mac OSX, GNU/Linux), both in RAM and on the harddrive. An example of a popular location of forensic artifacts on Windows are hives on the Windows registry, where Windows stores information so it remembers things like what your most recently opened folders were on explorer. Since many of these forensic artifacts must be found in ways more complex than navigating to a folder and finding a text file, the open source community has provided various projects geared toward locating these logs and many times making them more readable for humans. As a result, forensic analysts must often search for relevant projects that suit their needs, and use them one at a time, wasting time and energy. Furthermore, a frequent problem with this process is that there are times when multiple (sometimes hundreds or thousands) of computers need to be investigated at once, which makes this an even more difficult problem to tackle. The scope of this capstone project is to address these problems and to provide an effective solution that would be valuable to digital forensic analysts performing an incident response investigation. There are alternative projects that exist online that have the goal of addressing the same problem, such as the following projects:
- https://github.com/mantarayforensics/mantaray
- https://github.com/AJMartel/IRTriage
- https://github.com/travisfoley/dfirtriage
- https://github.com/Yelp/osxcollector

While some of these projects do address all of the same problems as I have illuminated, many do not. Furthermore, in the following "features summary" section I will provide a summary that will help in demonstrating why this tool can be more effective than the aforementioned alternative solutions, as many of the features provided in this tool are not provided in the others.

##### Features Summary
Many of the alternative solutions provide a good platform to perform incident response. They have various tools statically bundled into the project, and do an excellent job of utilizing them in a uniform-fashion that makes it friendly for novice users and those that are in a hurry by not requiring manual configuration of each tool.
This project does this in a different way - while each tool must be independently configured (through command line arguments), it allows as many tools as the user would like to be added without manually modifying the codebase. The configuration file, formatted in Tom's Obvious Minimal Language (https://github.com/toml-lang/toml for more information) allows the user to specify which tools they would like to be ran, the arguments, whether they're enabled, and the path to the tool itself. While this may be a tedious process, it gives the user more control over how a forensic investigation is done by only running what is absolutely necessary. The next useful feature for forensic analysts is the ability to deploy a solution to multiple servers - as previously discussed, if an analyst needs to figure out how to transport the various tools along with the triage tool, this makes their life a lot more difficult. By integrating with go.rice (https://github.com/GeertJohan/go.rice for more information), the analyst can instead bundle as many tools as necessary inside the executable so that ther is only a single file for deployment. This saves time and energy by avoiding extra steps such as checksum confirmations and whatnot. Furthermore, if an analyst has many tools that result in the single executable having a file size larger than they would like, go.rice has the ability to, instead of bundling the files within the executable itself, add a single additional file that must be within the same folder as the executable. Next, the ability for this tool to run in various operating systems provides a dynamic platform that avoids having to find multiple compatible triage tools; golang was a stable multi platform language to use for this endeavour. For Windows, there is a GUI provided that aids in the analysis process - it provides a simple yet effective interface that analysts can use to understand the status and output of the various forensic tools being ran. On GNU/Linux and Mac OSX, there is a similarly effective command line interface that instead outputs the results of the various forensic tools to disk, while still providing constant progress information while the tools are being executed. This execution process is done through a feature of golang called "goroutines" that provide concurrency to functions being executed - they are cheap compared to threads, and only a few kilobytes in stack size. This was another reason that golang was chosen as the language for this project, and *Figure 1* demonstrates the basic funcitonality of how a goroutine works.

![Figure 1](http://2.bp.blogspot.com/-n7-qbNilhsQ/TfDLMX0tyFI/AAAAAAAABH8/wX04pyqEVaQ/s1600/parallel-execution.png)

**Figure 1**


### Design, Development and Test
##### Design
When performing incident response, it's very important that development occurs under the following constraints:
* Small in size
    * Deployment often occurs over the internet
    * Deployment often occurs on machines with potentially low disk space
* Dependency-free (as much as possible)
    * Installing additional components complicates the deployment process, especially when different types of machine architecture need to be configured differently

These were the important facets to consider when deciding which libraries to integrate into the development process of this project. With these constraints in mind, I tried to be conservative and use as few additional components as possible to ensure that there is stable compatibility across all operating systems and system architectures. 
To implement the configuration file, I used the TOML (Tom's Obvious Minimal Language) golang library to parse what the user specifies in regards to each individual digital forensic tool settings. All that this this library was used for was to serialize the text file into an object that can then be used like any other struct in golang. 

The go.rice library was used in the development process to, when enabled in the executable argument, allow the various digital forensic tools to be bundled within the executable or into an additional archive file. Go.rice allows developers to create "boxes", which convert a folder on the disk that contains files into a programming file with the ".go" file extension. Upon "unboxing", this file is then extracted and in our case written to disk so the tools can then be executed.  Due to the fact that go.rice parses the go files during its execution, it requires that string literals be used when specifying "box" locations. For this reason, the "Capstone/Tools" path is the only directory that can be used to temporarily store the tools before the boxing process occurs. There are potential workarounds for this - I could always create a different main file, and instead change that string literal dynamically before executing the real main file. This is a feature that I plan to implement in the future, but was outside of the scope of this capstone - for now, it's simple enough to just put the tools you plan to bundle into a specific folder.  

For the GUI, I experimented with a few different multi platform GUI frameworks and experienced a multitude of problems, from compilation to compatibility. In the end, I found that andlabs/ui (https://github.com/andlabs/ui for more information) seemed to work well on both Windows architectures, and it claimed to be cross platform compatible. During the development process, I found that there was a feature (multi line text box) that if needed (it was) would have to be implemented manually by myself. After searching around, I found that there was a fork of this library called ProtonMail/ui that already implemented that feature, so I decided to use that fork instead of the original. This framework was utililzed by first creating all necessary components based on the enabled tools within the configuration file, and the dynamic components of the GUI (tool progress, for example) was passed around to various functions that needed it for read/write capabilities. The passing of these components was done through the structure UIComp (located in the Structure folder) that contains two of the necessary GUI components. Lastly, the classes for running operating system specific code all implement the ToolRunner interface, and are broken into various classes with the filename beginning with "Osutil_" followed by the operating system name specified by golang specifications. For a detailed depiction of this entire design, please refer to *Figure 2*. Please note, however, that the UML diagram in *Figure 2* only truly represents the design specified for Windows, though all other operating system implementations are the same, just without the Pronton/ui GUI integration.

![Figure 2](https://i.imgur.com/l7Kwzti.png)

**Figure 2**

##### Development

This project began with a rough design of how it would be implemented, and more and more pieces were integrated as time went on. After refactoring multiple times throughout this process, the final design was successfully completed. Originally the project began with compatibility only for Windows. I had no GUI, and almost all of the code was in the main file. Configuration data was implemented through JSON; multiple tools had been statically integrated, and the communication about tool command line output had been passed through goroutine channels. After talking to Professor Waldon, he described how golang has the ability to differentiate between different operating systems for a single file by appending an "_[Operating System]" to the file name. Armed with this new information, I integrated the basic empty implementation for each operating system. It was at this point that I set about researching various GUI frameworks and eventually found Proton/ui. The flow of GUI components described in *Figure 2* was then implemented, and the GUI worked great - the goroutine channels had been removed as well, since the output would now be displayed directly through the GUI. After tackling the GUI, I then went back to multi platform implementation; I soon found that while the GUI looked great on Windows, it did not for Mac OSX. Furthermore, I found that the framework would not even compile on GNU/Linux despite spending a week trying to find an effective solution. I decided that since the GUI worked well on a certain version of Mac OSX and very strange on others, that I would disable the GUI for Mac OSX as well as for GNU/Linux. After talking with Professor Waldon again, he told me about TOML and I thought that it was a great language to replace JSON, that I had already implemented. I changed the configuration library, and refactored my code so that tools were no longer static - instead, alternative tools could be added without changing the codebase. It was at this point that I had the final design that you see in *Figure 2*.

##### Test

This project was tested as follows:

* Configuration 
    * Various example configuration file data was testing against the parsing function to test against anomolies. This was done to test "bounds" and outliers that could potentially break things, as well as to determine whether the results were as to be expected.
* Functions
    * There were a few functions that while they weren't a "feature", I wanted to ensure the results of certain calls were as expected. This was done by passing in various appropriate and innapropriate parameters and testing against the results.
* GUI
    * Unfortunately most of the testing that I intended to implement for the ProtonMail/ui framework did get integrated. I tried for a week straight to figure out how to test for things such as mouse clicks and active components, but this did not seem to be well implemented into the framework. However, I have provided a test to check whether the behavior of the function actually handling the building of the GUI results in expected behavior when bad parameters are passed in.
* Tool execution
    * I have provided tests to ensure that the functions handling execution of the digital forensic tools work as expected. This was also done by passing in anomolous and normal parameters and checking whether the results are what they should be.


### Results
Overall, I think that I met all of the main objectives that I outlined at the beginning of this project. I hit many road blocks along the way, and ended up with a much smaller project than I had hoped for, but I'm very satisfied with the results. I expected that I would crank out code very quickly, but soon realized after a few weeks that golang is a harder language than I thought; I had a little experience over the summer learning it, but that was my first time being exposed to it. Had I written this in Java, I think I could have made something far more complex and larger in size. However, that wouldn't have made sense for this project for a number of reasons - it's bloated and has dependencies, though it is multi platform. Regardless, I think that the multiple refactoring and "getting stuck" multiple times set me back quite a bit. Had I contacted Professor Waldon upon initially getting stuck, I would have had more upward mobility; I thought that if I kept tinkering around and thinking that I would have come up with solutions to my problems, but this never actually happen two times throughout the project. Thankfully he was a great resource when I finally contacted him, and helped me get around those problems that I was stuck on.

To speak more about specifics, I will now outline the results for each main feature of the project:

**UI**
Unfortunately the GUI was only successfully implemented on Windows. It was working on Mac OSX, but many of the components of the UI looked strange and nothing like on Windows, as I have discussed in other sections of this final technical report. However, I'm very satisfied with how the GUI turned out on Windows. It's user friendly, and what I would describe as "bare bones" - forensic analysts only need what is absolutely necessary, because many times incident response engagement are very time sensitive where every second counts. With this in mind, I believe that the GUI was optimal for the target audience. Furthermore, through the usability testing (simply messing around with the GUI trying to break things), the results have been very consistent and indicate a stable build. In terms of the command line interface for Mac OSX and GNU/Linux, I'm also very happy with the results. The only real necessity for incident response analysts is understanding the progress of specific tools; this is accomplished successfully through this project. It would be more helpful, however, to know the specific progress of a given tool in terms of percentage, but this is impossible to do if I would like to keep the dynamic nature of this project. However, it would be great in the future to perhaps contribute to open source digital forensic tools, create a way of tracking progress by percent, and integrating that specific build into this project in the future. 

**Cross platform compatibility**
This feature was implemented very successfully, thanks to the language that it was written in, golang. The codebase for GNU/Linux and Mac OSX is the same, since they're both derivatives of Unix, but it's an excellent complement to this project. By implementing cross platform compatibility, this project reaches a larger audience and provides more of an opportunity for developers to contribute to the open source project in the future. In my testing, I haven't come across a single issue in regards to compatibility (beyond attempting to implement the GUI) which means the build is stable and consistent. In the future, I plan to implement the GUI for the operating systems that it was not implemented for, through alternative GUI frameworks - perhaps each operating system will need it's own, but that's why I designed the code the way that I did. 

**Deployment**
I'm very happy with the results of this project in regards to the ability to deploy the software with various digital forensic tools bundled within the executable or as an archive. Go.rice was an excellent library to use in this endeavour, and it's provided a stable and consistent addition to the other features. As I've mentioned in the previous sections of this paper, in the future I would like to not have the directory of the configuration file hard-coded when initially bundling the software using go.rice. Also, deployment can become difficult when using many forensic tools of large size (because nobody wants to transfer around a 3 gigabyte file), but this is just the nature of computers. I don't see any reason why this feature would not be extremely helpful to digital forensic analysts, and I think that the results met my objective for the feature.

**Configuration**
The configuration of this software through a configuration file is helpful and straightforward. This is a simple feature, but it successfully met my objective; TOML was an excellent language as opposed to JSON as I had before, because it is much more user friendly. Through my testing, I found that the functions I use to parse the file are stable and consistent, and I would have a hard time understanding how this feature wouldn't be simple and effective for my target audience.

**Performance**
Though this isn't a feature, I believe that it's necessary to discuss the performance of this software due to it being an important factor for incident response analysts. Though I haven't done any programical performance testing, I have tested this software on three machines of mine that I believe represent a basic spectrum from "very slow" to "very fast". What I have found is that this software runs very well, consistently, on all operating systems, architectures, of both high and low performing machines. This is a very good result to find, because incident response investigations will be used on all different types of machines. Though certain digital forensic tools themselves may not be as stable and consistent as this project itself, it's still important that this tool never crashes to ensure hours of time are not wasted. For this reason, I'm very happy with the results in terms of performance and I think the audience of this software will feel the same way.


This is a summary of the feature results:
| Feature/Factor       | Implemented as expected | Plan to be improved  |
| ------------- |:-------------:| -----:|
| UI      | Yes | Yes |
| Cross platform compatibility      | Yes      |   No |
| Deployment | Yes      |    Yes |
| Configuration | Yes      |    No |
| Performance | Yes      |    No |
**Figure 3**

### Conclusions and Future Work
The problem that this project aimed to address was the process of incident response - with many tools with various complex parts, having a solution that would automate this process would provide a valuable platform for incident response analysts. In particular, having a multitude of different tools being executed separately with many different command line windows open is not the best solution. Often companies must invest time and energy toward creating their own custom software to do what this project does, but in a static and often times confusing manner. This frequently occurs especially for scenarios in which this type of triage must occur on multiple machines using deployment. Avoiding this problem is important, because I consider myself very passionate about cyber security and stopping criminals from stealing intellectual property - if this project can assist in the investigation of a single criminal, then that is something that I can truly be proud of. I believe that what this project has provided will solve this problem by creating a single way in which that execution can occur. The results are effective, efficient, and stable, which are very important to the target audience of this project. Furthermore, the design encourages open source developers to contribute in the future, and to provide feedback that is invaluable to future adaptations of the codebase. What I learned throughout this project is that if I or anyone else ever gets stuck on a problem, it is okay to ask for help - the open source as well as academic community is a great resource and is always willing to give their opinion. Furthermore, I found that it is very important to consider how your target audience will react to different facets of a piece of software, and that if your software is not helpful to that target audience, it's often not useful to anyone. In the future, I hope to work on the "plan to be improved" features specified in *Figure 3* - if I had more time, I would have continued to research various GUI frameworks for golang and find ones that work great for GNU/Linux and Mac OSX. Similarly, I would like to fix that small "string literal" issue in regards to deployment. Luckily the future is bright and programming is a skill that must be honed over time, so I will absolutely work on these deficits in the future and continue to become a better programmer.


