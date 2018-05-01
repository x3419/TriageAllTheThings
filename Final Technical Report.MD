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

As a digital forensic or incident respont analyst, your job is to examine computers generally involved in some type of criminal behavior. This is done by analyzing various forensic artifacts on the computer; frequently (almost always in one way or another) when something is clicked, written or read on a computer, there is a log of it. It's not always a human friendly log like you might expect, and may require advanced methodologies in order to understand the logs, but they do exist. They're stored in various places by the operating system (such as Windows, Mac OSX, GNU/Linux), both in RAM and on the harddrive. An example of a popular location of forensic artifacts on Windows are hives on the Windows registry, where Windows stores information so it remembers things like what your most recently opened folders were on explorer. Since many of these forensic artifacts must be found in ways more complex than navigating to a folder and finding a text file, the open source community has provided various projects geared toward locating these logs and many times making them more readable for humans. As a result, forensic analysts must often search for relevant projects that suit their needs, and use them one at a time, wasting time and energy. Furthermore, a frequent problem with this process is that there are times when multiple (sometimes hundreds or thousands) of computers need to be investigated at once, which makes this an even more difficult problem to tackle. The scope of this capstone project is to address these problems and to provide an effective solution that would be valuable to digital forensic analysts performing an incident response investigation.  




