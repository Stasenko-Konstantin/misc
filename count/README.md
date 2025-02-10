# count (blazingly fast 🚀)

command line applications for recursive line counting in files

### todo:
- [x] recursive counting of lines in files
- [x] specifying the file/directory to be counted
- [x] exclude files/directories from counting 
- [ ] highlighting/ignoring certain file formats
- [ ] specifying how rows are counted:
  - [ ] counting all lines in files
  - [ ] counting lines without line breaks
  - [ ] specifying that different formats should be counted as one
  - [ ] regular expression specification
- [x] loading message
- [x] help message

### usage

requires [cargo](https://www.rust-lang.org/tools/install) for building executable file

```bash
git clone git@github.com:Stasenko-Konstantin/misc.git && cd misc/count
./build.sh     # requires sudo for cp executable file to /bin
               # reopen terminal
count -h                  
```

```bash
Usage: count [OPTIONS]

Options:
  -p, --paths <PATHS>  
  -e, --excludes <EXCLUDES>  Excludes specified file names and/or extensions
  -h, --help           Print help
  -V, --version        Print version
```