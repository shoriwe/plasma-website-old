import os
import os.path


def remove_bom(path):
    with open(path, "rb") as file_object:
        contents = file_object.read()
    with open(path, "wb") as file_object:
        new_contents = contents.replace(b"\r", b"")
        file_object.write(new_contents)
    return len(contents) - len(new_contents)


def main():
    total = 0
    for root, folders, files in os.walk("."):
        for file in files:
            if any(ext in file for ext in (".go", ".py", ".txt", ".md", ".html", ".css", ".js")):
                total += remove_bom(os.path.join(root, file))
    print(f"Removed {total}")


if __name__ == "__main__":
    main()