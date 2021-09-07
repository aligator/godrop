# 01 Goal

GoDrop is a cloud file server similar to e.g. Nextcloud but much simpler.

## Milestone v0.1.0

The first step is to create a fully functional file upload server which can just
be started and then accessed through a specific port. It then provides a simple 
web UI where you can upload files, create folders and download files. 
Nothing more and nothing less.

__Functionality:__  
[x] a Go server which can serve the frontend and provide a GraphQL API  
[x] the server is implemented with a real filesystem as data storage, 
but it should be as easy as re-implementing a filesystem to replace the data source
[ ] for now metadata is stored next to the file/folder in a `metadata.json` containing description and mime type
[x] a simple frontend  
[x] UI can show a folder and navigate  
[x] UI can create a sub folder  
[x] UI can upload a file  
[ ] UI can download a file.  

