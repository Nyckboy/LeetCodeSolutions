import os
import re

# Config
ROOT_DIR = '.'
README_PATH = os.path.join(ROOT_DIR, 'README.md')
# Regex: Matches any number of digits at the start (e.g., "1-two-sum", "10-regular-expression")
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
    table_data = [] # Store tuple: (id_int, row_string)

    # 1. Header
    header = "| # | Title | Solution | Difficulty |\n|---|---|---|---|"

    # 2. Scan directories
    for entry in os.listdir(ROOT_DIR):
        if not os.path.isdir(entry):
            continue
        
        match = FOLDER_PATTERN.match(entry)
        if match:
            problem_id_str = match.group(1)
            problem_id_int = int(problem_id_str) # Convert to int for sorting
            slug = match.group(2)
            title = slug.replace('-', ' ').title()
            
            sol_link = "No Code"
            if os.path.exists(f"{entry}/Solution.java"):
                sol_link = f"[Java](./{entry}/Solution.java)"
            elif os.path.exists(f"{entry}/Solution.js"):
                sol_link = f"[JS](./{entry}/Solution.js)"
            
            difficulty = get_difficulty(entry)
            
            row = f"| {problem_id_str} | [{title}](./{entry}) | {sol_link} | {difficulty} |"
            table_data.append((problem_id_int, row))

    # 3. Sort numerically by ID (so 2 comes before 10)
    table_data.sort(key=lambda x: x[0])

    # 4. Join rows
    rows = [item[1] for item in table_data]
    return header + "\n" + "\n".join(rows)

def update_readme():
    if not os.path.exists(README_PATH):
        return

    with open(README_PATH, 'r', encoding='utf-8') as f:
        content = f.read()

    table_content = generate_table()
    
    pattern = re.compile(
        r'().*()', 
        re.DOTALL
    )
    
    new_content = pattern.sub(f"\\1\n\n{table_content}\n\n\\2", content)
    
    with open(README_PATH, 'w', encoding='utf-8') as f:
        f.write(new_content)
    print("✅ README updated successfully.")

if __name__ == "__main__":
    update_readme()