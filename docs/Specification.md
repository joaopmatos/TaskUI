# Definition of the project and scope

The project will encompass two things: a cli tool and a TUI on top of it. This utility will be able to replace basic functionality of the Windows Task Scheduler.

This project **will be**:
- Small by design
- Focused on giving a better UX for doing repeated task management on windows

This project **will not be**:
- Encompassing all the possible functionalities of the Windows Task Scheduler GUI
- Cross platform; it's focus is on one sole platform

# Functionalities And Technical details

We'll take advantage of Windows's COM API **ITaskService** to interact with system tooling. This will prove to be more complex than parsing text from a console output, for example, but it will also be les prone to error and will allow to run with privileges if desired.

## Technical design
The TUI will be, essentially, a wrapper for the cli that'll be sitting underneath.

This means that if the user runs:
``` powershell
taskui.exe
```

It'll open the TUI. However, if the user run something like:
```
taskui.exe list
```

It'll get the output of the list command directly in the command line.

## Interfacing

For interfacing with windows's COM (Component Object Model), we'll make use of the package go-ole/go-ole.

## Functionality: List tasks
The app will be able to list the currently configured tasks.
In said list, it'll show:
- The task's Name
- Last run
- Next run
- Current status
- Last run status
### Possible statuses of the UI
- **Loading**: the application is loading the list of tasks. 
- **List/Table**: the application is showing the loaded list of tasks
- **Error**: in case of an error, the application will occupy the whole screen with the error message and either show a confirmation dialog allowing to bypass the error, or force the user to restart the application with privileges (if that's what's causing difficullties).
## Functionality: Run tasks
The app will be able to run currently configured tasks. 
Said run will either be **successful** or **unsuccessful** with a corresponding exit code.

# Technological stack
We'll build this tool in Go, taking advantage of it's cross compilation (the development environment is on Fedora) and using the BubbleTea TUI Framework and Bubbles (visual components made for BubbleTea).

## Technical Constraints
- **CGO:** Environment must target `CGO_ENABLED=0` to ensure seamless cross-compilation from Fedora to Windows using pure-Go OLE libraries (`go-ole`). 