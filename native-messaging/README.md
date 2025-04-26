# Native messaging

This demonstrates communication between browser extensions and native applications, so it requires installing three pieces:
- The browser extension
    - This should be installed the normal way for unpacked/debug extensions in your browser
- The native application
    - This is a Go application, you can build it by running `task build`.
- The native messaging manifest
    - Installation requirements vary by browser. If you're using Firefox on Windows or MacOS, `install_native_messaging_manifest.py` should do it, otherwise you'll need to execute the steps necessary for your browser and platform.
    
Running `task install` should build and install the native application and native messaging manifest automatically.