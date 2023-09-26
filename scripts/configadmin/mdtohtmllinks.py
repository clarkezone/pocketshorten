import re
import sys

def markdown_to_html_link(filepath):
    with open(filepath, 'r', encoding='utf-8') as file:
        content = file.read()

        # Replace markdown style links with HTML style links that open in a new tab
        content_new, num_replacements = re.subn(r'\[(.*?)\]\((.*?)\)', r'<a href="\2" target="_blank">\1</a>', content)

    # Write the modified content back to the file only if there were replacements
    if num_replacements > 0:
        with open(filepath, 'w', encoding='utf-8') as file:
            file.write(content_new)
        print(f"{num_replacements} replacements were made.")
    else:
        print("No replacements were made.")

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("Usage: python script_name.py <path_to_markdown_file>")
        sys.exit(1)
    filepath = sys.argv[1]
    markdown_to_html_link(filepath)
