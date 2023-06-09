import os
from define_emotion.test_neyro import gen_rand_description


def load_to_ssd(data: bytes, filename: str) -> str:
    file_path = f"saved_images/{filename + gen_rand_description(5)}.png"
    try:
        with open(file_path, "wb") as f:
            f.write(data)
    except (PermissionError, FileNotFoundError, OSError, TypeError) as ex:
        print("error occured: ", ex)
        return ""
    return file_path


def delete_from_ssd(filePaths: [str]):
    for fp in filePaths:
        try:
            os.remove(fp)
        except (OSError, TypeError) as ex:
            print("error while deleting file:", ex)
