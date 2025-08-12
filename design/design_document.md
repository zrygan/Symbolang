# Design of Symbolang

These are all implemented (and some planned) features of Symbolang. Each example has a C code equivalent (if none exist, a Python may be provided).

## Printing Emojis
There are two printing functions in Symbolang. One for basic print (`print`) and another for printing with a new line (`println`).

### Print
``` c
✏️ "Hello, World!" 🫷

// Equivalent C code:
printf("Hello, World!");
```
- ✏️: denotes the literal after it is to be printed.
- 🫷: denoted the end of the line.

The implementation above 1 will print "Hello, World!" without a new line.

### Println
```c
🖊️ "Hello, World!" 🫷

// Equivalent C code:
printf("Hello, World!\n");
```
- 🖊️: denotes the literal after it is to be printed with a new line.

### Multiple Print Statements
Let ⬜ denote the ✏️ or 🖊️ emoji, multiple print statements are of the form:
```c
⬜ <literal_1> ⬜ <literal_2> ⬜ ... 🫷

// Equivalent C code:
printf(<literal_1>);
printf(<literal_2>);
```

Example:
```c
🖊️ "Hello" ✏️ "World" ✏️ "!" 🫷

// Equivalent C code:
printf("Hello\n");
printf("World");
printf("!");
```

## Comments
```c
🖊️ "Hello, World!" 🫷 💭 This is a comment
💭 🖊️ "This will not work" 🫷
🖊️ "Okay, bye!" 🫷

// Equivalent C code:
printf("Hello, World!\n"); // This is a comment
// printf("This will not work\n");
printf("Okay, bye!\n");
```
- 💭: denotes a comment, makes all columns in the same line after it a comment.

Notice that the comment emoji is implicitly terminated "area-of-effect" at the end of the row.

## Variables and Constants
### Variables
```c
💭 Example: 📃 <identifier> <literal> 🫷
📃 name "Matthew" 🫷

// Equivalent C code:
// Example: <data_type> <identifier> = <literal>
char name[7] = "Matthew";
```
- 📃: denotes variable declaration and assignment.

We use the same syntax to redefine variables
```c
📃 name "Matthew" 🫷
📃 name "Zhean" 🫷

// Equivalent C code:
char name[7] = "Matthew";
strcpy(name, "Zhean"); // assume string.h is included
```

⚠️ **Data types are unimplemented**. Data types are implicit and dynamic, hence we can redefine a variable to a value of a different data type.

```c
📃 x "Matthew" 🫷
📃 x 100 🫷

// Equivalent Python code:
x = "Matthew"
x = 100
```

### Constants
A constant is a variable whose value cannot be changed.

```c
💭 Example: 🪨 <identifier> <literal> 🫷
🪨 name "Matthew" 🫷

// Equivalent C code:
// Example: const <data_type> <identifier> = <literal>
const char name[7] = "Matthew";
```
- 🪨: constant declaration.

Redefining a constant leads to an error:
```
🪨 name "Matthew" 🫷
🪨 name "Zhean" 🫷       💭 Error occurs here!
```

### Deletion
A variable or constant can be deleted.

```c
💭 Example: ✂️ <identifier> 🫷
🪨 name "Matthew" 🫷
✂️ name 🫷                 💭 Assuming name is defined
🖊️ name 🫷                 💭 Error occurs here `name` is does not exist!
🪨 name "Zhean" 🫷   💭 `name` can be defined again
🖊️ name 🫷
```
- ✂️: deletes a variable or constant.

The example above prints out `Zhean`.