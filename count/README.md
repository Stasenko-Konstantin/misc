# count (blazingly fast 🚀)

command line applications for recursive line counting in files

### todo:
- [x] recursive counting of lines in files
- [x] specifying the file/directory to be counted
- [ ] exclude files/directories from counting 
- [ ] highlighting/ignoring certain file formats
- [ ] specifying how rows are counted:
  - [ ] counting all lines in files
  - [ ] counting lines without line breaks
  - [ ] specifying that different formats should be counted as one
  - [ ] regular expression specification
- [x] loading message

### usage

requires cargo for building executable file

```bash
git clone git@github.com:Stasenko-Konstantin/misc.git && cd misc/count
./build.sh     # requires sudo for cp executable file to /bin
               # reopen terminal
count          # prints files from current and sub- directories
count <file>   # prints <file>
count <dir>    # prints files from <dir>
```