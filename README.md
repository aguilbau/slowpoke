sample usage:
```
$> cat file.txt 
http://example.com/
http://example.com/
http://example.com/
http://google.com/
```
```
$> ./slowpoke file.txt  | json_pp 
{
   "9be2aaa40d05ef5fbe66e44f82f76dc12740dfbd7759c7bd9881eb58b16b1108" : {
      "Google" : [
         "http://google.com/"
      ]
   },
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
