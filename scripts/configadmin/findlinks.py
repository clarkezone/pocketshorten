import re
import json
from datetime import datetime, timezone
import argparse

SHORT_URL_PREFIX = "http://q6o.to/"

def extract_urls(file_path, json_data, json_file_path, tag):
    # Create a dictionary with expanded URLs as keys and their corresponding shortened forms as values
    url_dict = {entry[1]: entry[0] for entry in json_data["values"]}

    new_entries_count = 0  # Counter for new entries added to the JSON
    replacements_count = 0  # Counter for replacements made in the markdown

    with open(file_path, 'r') as file:
        content = file.read()

    pattern = r'\[.*?\]\((http[s]?://(?!q6o\.com).*?)\)'
    urls = re.findall(pattern, content)
    for url in urls:
        # Skip URLs starting with the short form prefix
        if url.startswith(SHORT_URL_PREFIX):
            continue
        elif url in url_dict and url_dict[url]:
            # Only replace if the shortened form is not an empty string
            shortened_url = SHORT_URL_PREFIX + url_dict[url]
            content = content.replace(url, shortened_url)
            replacements_count += 1
        else:
            # Check if URL is already in json_data during this run
            if any(entry for entry in json_data["values"] if entry[1] == url):
                continue

            print(url, "not found in JSON")

            # Adjusting the datetime format to remove the ':' in the timezone portion
            current_time = datetime.now(timezone.utc).astimezone().strftime('%Y-%m-%dT%H:%M:%S%z')
            current_time = current_time[:-2] + current_time[-2:]  # Removing the ':' from the offset
            json_data["values"].append(["", url, tag, current_time])  # Use the tag here
            new_entries_count += 1

    # Only write the updated markdown content back if replacements were made
    if replacements_count > 0:
        with open(file_path, 'w') as file:
            file.write(content)

    # Only write the updated JSON data back if new entries were added
    if new_entries_count > 0:
        with open(json_file_path, 'w') as jfile:
            json.dump(json_data, jfile, indent=4)

    # Print the summary
    print(f"Summary:\nNew entries added to JSON: {new_entries_count}\nReplacements made in markdown: {replacements_count}")


def main():
    parser = argparse.ArgumentParser(description="Extract URLs from a markdown file and check against a JSON file.")
    parser.add_argument("markdown_file", help="Path to the markdown file.")
    parser.add_argument("json_file", help="Path to the JSON file.")
    parser.add_argument("tag", help="Tag to use for new entries in the JSON data.")  # New argument for the tag
    args = parser.parse_args()

    with open(args.json_file, 'r') as jfile:
        data = json.load(jfile)

    try:
        extract_urls(args.markdown_file, data, args.json_file, args.tag)
    except FileNotFoundError:
        print("Error: One of the provided files was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == "__main__":
    main()
