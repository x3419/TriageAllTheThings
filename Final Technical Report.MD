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
This project does this in a different way - while each tool must be independently configured (through command line arguments), it allows as many tools as the user would like to be added without manually modifying the codebase. The configuration file, formatted in Tom's Obvious Minimal Language (https://github.com/toml-lang/toml for more information) allows the user to specify which tools they would like to be ran, the arguments, whether they're enabled, and the path to the tool itself. While this may be a tedious process, it gives the user more control over how a forensic investigation is done by only running what is absolutely necessary. The next useful feature for forensic analysts is the ability to deploy a solution to multiple servers - as previously discussed, if an analyst needs to figure out how to transport the various tools along with the triage tool, this makes their life a lot more difficult. By integrating with go.rice (https://github.com/GeertJohan/go.rice for more information), the analyst can instead bundle as many tools as necessary inside the executable so that ther is only a single file for deployment. This saves time and energy by avoiding extra steps such as checksum confirmations and whatnot. Furthermore, if an analyst has many tools that result in the single executable having a file size larger than they would like, go.rice has the ability to, instead of bundling the files within the executable itself, add a single additional file that must be within the same folder as the executable. Next, the ability for this tool to run in various operating systems provides a dynamic platform that avoids having to find multiple compatible triage tools; golang was a stable multi platform language to use for this endeavour. For Windows, there is a GUI provided that aids in the analysis process - it provides a simple yet effective interface that analysts can use to understand the status and output of the various forensic tools being ran. On GNU/Linux and Mac OSX, there is a similarly effective command line interface that instead outputs the results of the various forensic tools to disk, while still providing constant progress information while the tools are being executed. This execution process is done through a feature of golang called "goroutines" that provide concurrency to functions being executed - they are cheap compared to threads, and only a few kilobytes in stack size. This was another reason that golang was chosen as the language for this project, and Figure 1 demonstrates the basic funcitonality of how a goroutine works.

![Figure 1](http://2.bp.blogspot.com/-n7-qbNilhsQ/TfDLMX0tyFI/AAAAAAAABH8/wX04pyqEVaQ/s1600/parallel-execution.png)

**Figure 1**

### Design, Development and Test
