# wc-tool

## To run
You can run the program without flags

```
go run main.go <file_name>
```

The output would be

```
>> [number of lines]  [number of words]  [number of bytes] <file_name>
```

## The flags

Flags are optional they are used in the following way

```
go run main.go -[flag] <file_name>
```

The available flags are:
- c: Counts the number of bytes in the file
- l: Counts the number of lines in the file
- w: Counts the number of words in the file
- m: Counts the number of characters in the file

## Read from standard input
You can also read the data in from standard input

```
cat text.txt | go run main.go -l
```
