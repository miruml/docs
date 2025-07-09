#!/bin/sh

git_repo_root_dir=$(git rev-parse --show-toplevel)
source_docs="$git_repo_root_dir/cli/gen"
dest_docs="$git_repo_root_dir/snippets/cli-reference/gen"


# Remove the existing generated docs
rm -rf "$dest_docs"
# Create the directory for the generated docs
mkdir -p "$dest_docs"

rename_synopsis_to_description() {
    file="$1"
    sed -i 's/# Synopsis/# Description/g' "$file"
}

remove_everything_before_description() {
    file="$1"
    sed -i '/# Description/,$!d' "$file"
}

remove_SEE_ALSO() {
    file="$1"
    sed -i '/SEE ALSO/,$d' "$file"
}

add_bash_to_opening_code_fences() {
    file="$1"
    awk '
        FNR==NR {if ($0 ~ /^```$/) total++; next}
        {
            if ($0 ~ /^```$/) {
                count++
                if (count % 2 == 1 && count < (total-4)) 
                    print "```bash"
                else 
                    print $0
            } else 
                print $0
        }
    ' "$file" "$file" > temp && mv temp "$file"
}

move_command_to_top() {
    file="$1"

    awk 'BEGIN { 
        printing = 0;
        cmd = "";
        desc = "";
        rest = "";
        found_first_bash = 0;
    }

    # Only capture the first bash block that appears before ### Description
    /^```bash/ {
        if (!found_first_bash) {
            found_first_bash = 1;
            cmd = $0 "\n";
            next;
        }
    }

    # After finding first ```bash, capture until closing ```
    {
        if (found_first_bash == 1) {
            cmd = cmd $0 "\n";
            if ($0 ~ /^```$/) {
                found_first_bash = 2;
            }
            next;
        }
    }

    # Capture description section
    /^### Description/,/^### Examples/ {
        if ($0 !~ /^### (Description|Examples)/) {
            desc = desc $0 "\n";
        }
    }

    # Capture everything from Examples onwards
    /^### Examples/,EOF {
        rest = rest $0 "\n";
    }

    END {
        printf "### Command\n\n";
        printf "%s\n", cmd;
        printf "### Description\n";
        printf "%s", desc;
        printf "%s", rest;
    }' "$file" > temp && mv temp "$file"
}

split_examples() {
    file="$1"
    awk '
    BEGIN { in_example=0; in_code_block=0 }

    /^### Examples$/ { in_example=1; print; next }
    /^### / && in_example { in_example=0 }  # Exit example section on next header

    in_example && /^# / { 
        if (in_code_block) print "```\n\n";  # Close previous code block before opening a new one
        print "```bash " substr($0, 3); 
        in_code_block=1;
        next 
    }

    in_example && /^```bash$/ { next } # delete existing ```bash markers

    # Skip empty lines within code blocks
    in_example && in_code_block && /^[[:space:]]*$/ { next }

    { print } # print everything else normally
    ' "$file" > temp && mv temp "$file"
}

# Copy and rename all generated MD files to MDX
for file in "$source_docs"/*.md; do
    if [ -f "$file" ]; then
        # Extract just the filename without path and extension
        basename=$(basename "$file")
        filename="${basename%.md}"
        dest_file="$dest_docs/${filename}.mdx"

        # copy the files as mdx
        cp "$file" "$dest_file"

        rename_synopsis_to_description "$dest_file"

        remove_everything_before_description "$dest_file"

        remove_SEE_ALSO "$dest_file"

        add_bash_to_opening_code_fences "$dest_file"

        move_command_to_top "$dest_file"

        split_examples "$dest_file"
    fi
done

