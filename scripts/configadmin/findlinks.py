import re
import argparse
import json

def extract_urls(file_path, json_data):
    # Create a dictionary with expanded URLs as keys and their corresponding shortened forms as values
    url_dict = {entry[1]: entry[0] for entry in json_data["values"]}

    with open(file_path, 'r') as file:
        content = file.read()
        pattern = r'\[.*?\]\((http[s]?://(?!q6o\.to).*?)\)'
        urls = re.findall(pattern, content)
        dict2 = {"foo": "bar"}
        dict2['https://twitter.com/clarkezone'] = "bing"
        print("foo" in dict2)
        print("https://twitter.com/clarkezone" in dict2)
        print(dict2)
        print(url_dict)
        print(len(url_dict))
        print()
        print()
        for url in urls:
            if url in url_dict:
                print(url, "yes, shortened form:", url_dict[url])
            else:
                print(url, "no")

def main():
    parser = argparse.ArgumentParser(description='Check URLs from a markdown file against a JSON list.')
    parser.add_argument('md_file', help='Path to the markdown file.')
    parser.add_argument('json_file', help='Path to the JSON file containing URLs.')
    args = parser.parse_args()

    # Load the JSON data
    with open(args.json_file, 'r') as jfile:
        json_data = json.load(jfile)

    try:
        extract_urls(args.md_file, json_data)
    except FileNotFoundError:
        print("Error: One of the provided files was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()
