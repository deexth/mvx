# MVX

* mvx -p path/dir/dir this will create the directory(dir) that does not exist 
* mvx -cd path/dir adding the -cd flag will create and cd into that directory
* mvx -f path/dir/file this will create the file at the end of the path and also any parent directory that doesn't exist. If the -f flag is not provided the file at the end will considered a directory
* mvx file.txt /path/dir will create the path location and any parent which doesn't exist and then move the file.txt to that location
* mvx -c file.txt /path/dir will create the path location and any parent which doesn't exist and then copy the file.txt to the location

For both mvx -c and mvx; both commands will also move/copy not only files but directories and their children

```go

mvx source target
mvx source1 source2 target_dir/

Flags:
  -c, --copy        copy instead of move
  -p, --parents     create parent directories (default true)
  -n, --dry-run     show what would happen
  -v, --verbose

```
