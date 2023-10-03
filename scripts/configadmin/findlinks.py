import re
import json
from datetime import datetime, timezone
import argparse

SHORT_URL_PREFIX = "https://q6o.to/"


def extract_urls(file_path,
                 json_data,
                 json_file_path,
                 tag):
    # Create a dictionary with expanded URLs
    # as keys and their corresponding shortened forms as values
    url_dict = {entry[1]: entry[0] for entry in json_data["values"]}

    new_entries_count = 0  # Counter for new entries added to the JSON
    replacements_count = 0  # Counter for replacements made in the markdown

    with open(file_path, 'r') as file:
        content = file.read()

    for url, shortened in url_dict.items():
        if not shortened:  # If the shortened form is empty
            continue
        markdown_link_pattern = r'(\[.*?\]\()' + re.escape(url) + r'(\))'
        shortened_url = SHORT_URL_PREFIX + shortened
        if re.search(markdown_link_pattern, content):
            content = re.sub(markdown_link_pattern, r'\1'
                             + shortened_url + r'\2', content)
            replacements_count += 1

    pattern = r'\[.*?\]\((https[s]?://(?!q6o\.to).*?)\)'
    unmatched_urls = re.findall(pattern, content)
    for url in unmatched_urls:
        # Check if URL is already in json_data during this run
        if any(entry for entry in json_data["values"] if entry[1] == url):
            continue
        print(url, "not found in JSON")
        current_time = datetime.now(timezone.utc)\
            .astimezone().strftime('%Y-%m-%dT%H:%M:%S%z')
        current_time = current_time[:-2] + current_time[-2:]
        json_data["values"].append(["", url, tag, current_time])
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
    print('Summary:\nNew entries added to JSON:'
          '{}\nReplacements made in markdown:'
          '{}'.format(new_entries_count, replacements_count))


def main():
    parser = argparse.ArgumentParser(description="Extract URLs"
                                     "from a markdown file and check "
                                     "against a JSON file.")
    parser.add_argument("markdown_file", help="Path to the markdown file.")
    parser.add_argument("json_file", help="Path to the JSON file.")
    parser.add_argument("tag", help="Tag to use for new "
                        "entries in the JSON data.")
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
