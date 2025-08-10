# Design of Symbolang

## Printing Emojis

Simple `Hello, World` implementation:

```
✏️ "Hello, World!" 🫷
```
- ✏️: denotes the literal after it is to be printed.
- 🫷: denoted the end of the line.

The implementation above 1 will print "Hello, World!" without a new line.

In C, the equivalent implementation is

```c
printf("Hello, World!");
```

```
📢 "Hello, World!" 🫷
```
- 📢: denotes the literal after it is to be printed with a new line.

Its translation in C is
```
printf("Hello, World!\n");
``` 

Printing multiple literals is of the form:
```
✏️ <literal_1> ✏️ <literal_2> ✏️ ... 🫷
```

Which is equivalent to having multiple `printf()` statements. Notice that since we use the ✏️ emoji, these will not have new lines between each print statement.

```
✏️ 'A' ✏️ 'B' ✏️ 'C' 🫷 
```

Will print `ABC` not `A B C` or the three letters spaced by a new line. Also notice that we used runes here (or single character strings).

We can use a combination of ✏️ and 📢 for printing characters. The only rule is that there is a literal after each printing emoji and it ends with the stop 🫷 emoji.

