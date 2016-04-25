sample usage:
```
$> cat file.txt 
http://example.com/
http://example.com/
http://example.com/
```
```
$> ./slowpoke file.txt  | json_pp 
{
   "3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423" : {
      "Example Domain" : [
         "http://example.com/",
         "http://example.com/",
         "http://example.com/"
      ]
   }
}
$> 
```

if no file is given, slowpoke will read input from stdin.
