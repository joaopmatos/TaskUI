# Definition of the project and scope

The project will encompass two front-ends: a CLI and a TUI over the same engine. This utility will be able to replace basic functionality of the Windows Task Scheduler.

This project **will be**:
- Small by design
- Performant and simple
- Focused on giving a better UX for repeated task management on Windows

This project **will not be**:
- Encompassing all the possible functionalities of the Windows Task Scheduler GUI
- Cross-platform: its focus is on one sole platform

# Functionality and Technical Details

We'll take advantage of Windows COM API **ITaskService** to interact with system tooling. This will prove to be more complex than parsing text from console output, for example, but it will also be less prone to error and will allow running with privileges if desired.

## Technical design
The TUI and the CLI are two thin front-ends over the same engine. The TUI does not shell out to the CLI; both call the engine directly.

There is a single binary. If the user runs:
``` powershell
taskui.exe
```

It'll open the TUI. However, if the user runs something like:
```
taskui.exe list
```

The output of the list command will be printed directly to the command line.

## Interfacing

For interfacing with Windows COM (Component Object Model), we'll make use of the package [github.com/go-ole/go-ole](https://github.com/go-ole/go-ole).

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
- **List/Table**: the application is showing the loaded list of tasks.
- **Error**: in case of an error, the application will occupy the whole screen with the error message and either show a confirmation dialog allowing the user to bypass the error, or ask the user to restart the application with privileges (if that's what's causing difficulties).

## Functionality: Run tasks
The app will be able to run currently configured tasks.
Each run will either be **successful** or **unsuccessful**, with a corresponding exit code.

# Technological stack
We'll build this tool in Go, taking advantage of its cross-compilation (the development environment is on Fedora) and using the Bubble Tea TUI framework and Bubbles (visual components made for Bubble Tea).

## Technical Constraints
- **CGO:** Builds must target `CGO_ENABLED=0` to ensure seamless cross-compilation from Fedora to Windows using the pure-Go OLE library (`go-ole`).