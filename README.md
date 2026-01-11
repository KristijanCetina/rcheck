# rcheck

A simple Go CLI tool to check if a regular expression pattern matches an input string.

## Installation

### From Source

```bash
git clone https://github.com/KristijanCetina/rcheck.git
cd rcheck
go build -o rcheck main.go
sudo mv rcheck /usr/local/bin/
```

## Usage

```bash
rcheck -e <regex_pattern> [-s <input_string>]
```

If `-s` is omitted, input is read from stdin.

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `-e` | required | Regular expression pattern to match |
| `-s` | optional | Input string to match against. If omitted, reads from stdin |
| `-v` | optional | Print the time spent checking |
| `-h` | optional | Show help message |

## Examples

### Basic matching with `-s` flag

```bash
rcheck -e "^hello.*" -s "hello world"
# Output: The pattern matches the input string!
```

### Reading from stdin

```bash
echo "hello world" | rcheck -e "^hello.*"
# Output: The pattern matches the input string!

cat file.txt | rcheck -e "^[A-Z].*"
```

### No match example

```bash
rcheck -e "^goodbye" -s "hello world"
# Output: The pattern does not match the input string.
```

### With verbose timing

```bash
rcheck -e "^hello.*" -s "hello world" -v
# Output:
# The pattern matches the input string!
# Time spent checking: 3.625µs
```

### Complex regex patterns

```bash
# Match email addresses
rcheck -e "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$" -s "user@example.com"

# Match dates (YYYY-MM-DD format)
rcheck -e "^\d{4}-\d{2}-\d{2}$" -s "2024-01-15"

# Match phone numbers
rcheck -e "^\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$" -s "555-123-4567"
```

### Piping from file

```bash
grep -r "error" /var/log/ | rcheck -e "^\[.*ERROR.*\]"
```

## Exit Codes

- `0` - Pattern matches the input
- `1` - Pattern does not match, or an error occurred

## License

MIT
