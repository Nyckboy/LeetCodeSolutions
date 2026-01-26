import os
import re

ROOT_DIR = '.'
# UPDATE: Changed source directories
ALGO_DIR = 'algorithms'
SYS_DIR = 'system-design'
README_PATH = os.path.join(ROOT_DIR, 'README.md')

# Regex remains the same
ALGO_PATTERN = re.compile(r'^(\d+)-(.+)$')
SYS_PATTERN = re.compile(r'^Q(\d+)-(.+)$')

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

def generate_table(directory, pattern, prefix_path):
    # Check if directory exists
    if not os.path.exists(directory):
        return "| No questions found |"
        
    table_data = [] 
    header = "| # | Title | Solution | Difficulty |\n|---|---|---|---|"

    for entry in os.listdir(directory):
        full_path = os.path.join(directory, entry)
        if not os.path.isdir(full_path):
            continue
        
        match = pattern.match(entry)
        if match:
            problem_id_int = int(match.group(1))
            slug = match.group(2)
            title = slug.replace('-', ' ').title()
            
            # Correct relative path for links
            rel_path = f"./{prefix_path}/{entry}"
            
            sol_link = "No Code"
            if os.path.exists(f"{full_path}/Solution.java"):
                sol_link = f"[Java]({rel_path}/Solution.java)"
            elif os.path.exists(f"{full_path}/Solution.js"):
                sol_link = f"[JS]({rel_path}/Solution.js)"
            
            difficulty = get_difficulty(full_path)
            row = f"| {problem_id_int} | [{title}]({rel_path}) | {sol_link} | {difficulty} |"
            table_data.append((problem_id_int, row))

    table_data.sort(key=lambda x: x[0])
    rows = [item[1] for item in table_data]
    
    if not rows:
        return "| No questions found |"
    return header + "\n" + "\n".join(rows)

def update_section(content, start_marker, end_marker, new_text):
    start_index = content.find(start_marker)
    end_index = content.find(end_marker)

    if start_index == -1 or end_index == -1:
        print(f"⚠️ Warning: Could not find markers {start_marker} and {end_marker}")
        return content

    pre_content = content[:start_index]
    post_content = content[end_index + len(end_marker):]
    return f"{pre_content}{start_marker}\n\n{new_text}\n\n{end_marker}{post_content}"

def update_readme():
    if not os.path.exists(README_PATH):
        return

    with open(README_PATH, 'r', encoding='utf-8') as f:
        content = f.read()

    # 1. Generate Algo Table (Scanning 'algorithms' folder)
    algo_table = generate_table(ALGO_DIR, ALGO_PATTERN, ALGO_DIR)
    content = update_section(content, "<!-- SOL_TABLE_START_ALGO -->", "<!-- SOL_TABLE_END_ALGO -->", algo_table)

    # 2. Generate System Design Table (Scanning 'system-design' folder)
    sys_table = generate_table(SYS_DIR, SYS_PATTERN, SYS_DIR)
    content = update_section(content, "<!-- SOL_TABLE_START_SYSDES -->", "<!-- SOL_TABLE_END_SYSDES -->", sys_table)
    
    with open(README_PATH, 'w', encoding='utf-8') as f:
        f.write(content)
    print("✅ README updated successfully.")

if __name__ == "__main__":
    update_readme()