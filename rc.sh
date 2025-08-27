NC="\033[0m"       # Text Reset
RED="\033[0;31m"          # Red
GREEN="\033[0;32m"        # Green

# Function to display usage information
usage() {
    echo "Usage: $0 <regex_pattern> [-s <input_string>]"
    echo "If -s is not provided, input string will be read from stdin."
    echo "  -e  Regular expression pattern"
    echo "  -s  Input string to match against (optional)"
    exit 1
}

# Parse command-line arguments
expression=""
input=""
while getopts "e:s:" opt; do
    case $opt in
        e) expression="$OPTARG" ;;
        s) input="$OPTARG" ;;
        \?) usage ;;
    esac
done

# Check if expression is provided
if [ -z "$expression" ]; then
    usage
fi

# If input string is not provided via -s, read from stdin
if [ -z "$input" ]; then
    echo "Enter input string: "
    read -r input
fi

# Check if input is empty (in case of stdin read failure)
if [ -z "$input" ]; then
    echo "Error: No input string provided"
    exit 1
fi

# Perform regex match
if [[ $input =~ $expression ]]; then
    echo "${GREEN}Pattern match${NC}"
else
    echo "${RED}No match${NC}"
fi
