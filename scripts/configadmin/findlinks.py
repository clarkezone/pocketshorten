import re
import argparse

def extract_urls(file_path):
    with open(file_path, 'r') as file:
        content = file.read()
        pattern = r'\[.*?\]\((http[s]?://(?!q6o\.to).*?)\)'
        urls = re.findall(pattern, content)
        for url in urls:
            print(url)
        print("done")

def main():
    parser = argparse.ArgumentParser(description='Extract q6o.com URLs from a markdown file.')
    parser.add_argument('md_file', help='Path to the markdown file.')
    args = parser.parse_args()

    try:
        extract_urls(args.md_file)
    except FileNotFoundError:
        print("Error: The provided markdown file was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()
