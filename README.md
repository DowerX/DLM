# DLM
Sorts downloaded files.

## Setup

### On Windows:
Clone into C:/ (or you will need to edit some dlm.reg and dlm-addon/background.js).\
Run dlm.reg.\
Load the dlm-addon folder into Chrome in developer mode.\
Edit config.yml as you please.\

### On Linux:
Clone and go build.\
Edit the dlm-addon/background.js file so it points to your own config.yml.\
Set up the dlm: URL protocol to use your cloned files.\
Load the dlm-addon folder into Chrome in developer mode.\
Edit config.yml as you please.

## config.yml:
```
downloads: <download folder here>
conditions:
  - ext: <file extension here>
    path: <where to store that filetype>
  - ext: <2nd file extension here>
    path: <where to store that 2nd filetype>
...

```
