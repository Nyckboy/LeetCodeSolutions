import os
import re

ROOT_DIR = '.'
README_PATH = os.path.join(ROOT_DIR, 'README.md')
# Regex for finding folders (kept this because it works well)
FOLDER_PATTERN = re.compile(r'^(\d+)-(.+)$')

def get_difficulty(folder_path):
    readme = os.path.join(folder_path, 'README.md')
    if not os.path.exists(readme):
        return "Unknown"
    with open(readme, 'r', encoding='utf-8') as f:
        content = f.read()
        match = re.search(r'\*\*Difficulty:\*\*\s*(\w+)', content)
        if match:
            return match.group(1)
    return "Unknown"

def generate_table():
    table_data = [] 
    header = "| # | Title | Solution | Difficulty |\n|---|---|---|---|"

    for entry in os.listdir(ROOT_DIR):
        if not os.path.isdir(entry):
            continue
        match = FOLDER_PATTERN.match(entry)
        if match:
            problem_id_int = int(match.group(1))
            slug = match.group(2)
            title = slug.replace('-', ' ').title()
            
            sol_link = "No Code"
            if os.path.exists(f"{entry}/Solution.java"):
                sol_link = f"[Java](./{entry}/Solution.java)"
            elif os.path.exists(f"{entry}/Solution.js"):
                sol_link = f"[JS](./{entry}/Solution.js)"
            
            difficulty = get_difficulty(entry)
            row = f"| {problem_id_int} | [{title}](./{entry}) | {sol_link} | {difficulty} |"
            table_data.append((problem_id_int, row))

    table_data.sort(key=lambda x: x[0])
    rows = [item[1] for item in table_data]
    return header + "\n" + "\n".join(rows)

def update_readme():
    if not os.path.exists(README_PATH):
        print("❌ Error: README.md not found.")
        return

    with open(README_PATH, 'r', encoding='utf-8') as f:
        content = f.read()

    # Define Markers
    START_MARKER = "<!-- SOL_TABLE_START -->"
    END_MARKER = "<!-- SOL_TABLE_END -->"

    # Find the positions of the markers
    start_index = content.find(START_MARKER)
    end_index = content.find(END_MARKER)

    if start_index == -1 or end_index == -1:
        print(f"❌ Error: Could not find both {START_MARKER} and {END_MARKER} in README.md")
        print("Make sure they are exactly as written above.")
        return

    # Generate the new table
    table_content = generate_table()

    # Rebuild the file safely:
    # 1. Everything BEFORE the start tag (Your Bio/Header)
    # 2. The Start Tag itself
    # 3. The New Table
    # 4. The End Tag itself
    # 5. Everything AFTER the end tag (if any)
    
    pre_content = content[:start_index]
    post_content = content[end_index + len(END_MARKER):]

    new_content = f"{pre_content}{START_MARKER}\n\n{table_content}\n\n{END_MARKER}{post_content}"
    
    with open(README_PATH, 'w', encoding='utf-8') as f:
        f.write(new_content)
    print("✅ README updated successfully.")

if __name__ == "__main__":
    update_readme()