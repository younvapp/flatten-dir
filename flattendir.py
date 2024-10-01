import os
import shutil
import sys


def flatten_dir(src_dir, dst_dir):
    if not os.path.exists(src_dir):
        print("Source directory does not exist.")
        return
    if not os.path.exists(dst_dir):
        os.makedirs(dst_dir)
    for root, dirs, files in os.walk(src_dir):
        print("Processing directory: %s" % root)
        for file in files:
            src_file = os.path.join(root, file)
            dst_file = os.path.join(dst_dir, file)
            shutil.copy2(src_file, dst_file)


if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: python flattendir.py src_dir dst_dir")
        sys.exit(1)
    src_dir = sys.argv[1]
    dst_dir = sys.argv[2]
    flatten_dir(src_dir, dst_dir)
    print("Flatten directory successfully.")
