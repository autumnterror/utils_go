# Breezy Innovations (Brz inv.) Utility Suite

Welcome to the official documentation for the **Breezy Innovations (Brz inv.) Utility Suite**! This collection of Go packages is designed to streamline your development process, offering a set of elegant and powerful tools to handle everyday tasks with ease. Let's dive in!

---

## 🎨 Package `log`

Logging is the heartbeat of any application. Our `log` package not only makes it simple but also adds a splash of color to your console, making debugging a more intuitive and visually pleasing experience.

You can create a `Logger` instance or use the package-level functions directly for quick and easy logging.

```go
// Create a new logger instance
logger := log.NewLogger()
logger.Info("Auth", "User logged in")

// Or use the package-level function
log.Success("Payment", "Payment processed successfully")
```
## Leveled Logging
These functions provide context and severity levels to your logs, complete with beautiful colors.


| Function	           | Description                                                                             | 	Color |
|---------------------|-----------------------------------------------------------------------------------------|--------|
| Info(op, msg)	      | Logs general information. Perfect for tracking application flow.                        | 	Blue  |
| Success(op, msg)	   | Indicates that an operation completed successfully.                                     | 	Green |
| Warn(op, msg, err)	 | Highlights a potential issue that doesn't break the application. The err is option al.	 | Yellow |
| Error(op, msg, err) | 	Logs a critical error that requires attention. The err is optional but recommended.	   | Red    |

## Simple Color Logging
Want to add a quick splash of color without the formal structure of leveled logging? Use these handy functions.

- `Blue(in ...any)`
- `Yellow(in ...any)`
- `Green(in ...any)`
- `Red(in ...any)`
## Standard Logging Wrappers
For convenience, the package also wraps the standard Go log functions.
- `Println(in ...any)`
- `Printf(format string, v ...any)`
- `Panic(in ...any)`
---
## 🧮 Package `alg`
A collection of handy, generic algorithms to keep your code DRY and readable.

| Function                         | Description                                                                                  |
|----------------------------------|----------------------------------------------------------------------------------------------|
| `MinInt(a, b int)              `  | Returns the smaller of two integers. No more inline if statements cluttering your code!      |
| `MaxInt(a, b int)              `  | Returns the larger of two integers.                                                          |
| `SliceCopy[T any](s, start, end`	 | Creates a safe, shallow copy of a slice from a start to an end index. Generic and type-safe! |
| `IsIn[T comparable](obj, sl)	 `   | A generic way to check if an item obj exists within a slice sl.                              |

## ✨ Package `format`
Handle data formatting and error wrapping with style and clarity.

- `Error(op string, err error) error`
Wraps a standard error with an operational context. This standardizes your error messages, making them much easier to trace. It returns a new error in the format: OP: <op>: ERROR: <original_error>.

- `Struct(in any) string`
Turns any struct into a beautifully indented JSON string. It's perfect for debugging and visualizing complex data structures. As a fun little touch, we use a cat emoji (🐱) for indentation!

Example Output:
```
JSON
{
🐱"name": "Breezy Logger",
🐱"version": "1.0",
🐱"features": [
🐱🐱"Colorful",
🐱🐱"Leveled"
🐱]
}
```
## 🆔 Package `uid`
A straightforward wrapper around Google's UUID library for generating and validating unique identifiers.

| Function	                 | Description                                                                       |
|---------------------------|-----------------------------------------------------------------------------------|
| New() string	             | Generates a new, random, V4 UUID and returns it as a string.                      |
| Validate(id string) bool	 | Checks if a given string is a valid UUID. Returns true if valid, false otherwise. |

## 🛡️ Package `validate`
Simple and essential validation logic for common use cases.

### `Password(password string) bool`
Validates a password against a set of strong security criteria. A password is considered valid if it meets all the following conditions:

- Is between 5 and 20 characters long.
- Contains at least one uppercase letter.
- Contains at least one digit.
- Contains at least one special character (e.g., !, @, #, ?).