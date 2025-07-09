import os
import shutil
from os import path
from download_file import download_filezip
from utils.util import ZipUtils
EXTRACT_DIR = "Extracted"
DUMP_PATH = "Dumps"
if not path.exists(path.join(EXTRACT_DIR, "FlatData")):
    import setup_apk
    from lib.dumper import IL2CppDumper
    from lib.console import notice
    from utils.util import FileUtils
    from extractor import compile_python
    IL2CPP_NAME = "libil2cpp.so"
    METADATA_NAME = "global-metadata.dat"

    TEMP_DIR = "Temp"

    save_path = TEMP_DIR
    dumper = IL2CppDumper()
    dumper.get_il2cpp_dumper(save_path)

    il2cpp_path = FileUtils.find_files(TEMP_DIR, [IL2CPP_NAME], True)
    metadata_path = FileUtils.find_files(
        TEMP_DIR, [METADATA_NAME], True
    )

    if(il2cpp_path and metadata_path):
        notice("il2cpp_path: " + il2cpp_path[0])
        notice("metadata_path: " + metadata_path[0])

    if not (il2cpp_path and metadata_path):
        raise FileNotFoundError(
            "Cannot find il2cpp binary file or global-metadata file. Make sure exist."
        )
    abs_il2cpp_path = path.abspath(il2cpp_path[0])
    abs_metadata_path = path.abspath(metadata_path[0])
    if(il2cpp_path and metadata_path):
        notice("il2cpp_path: " + abs_il2cpp_path)
        notice("metadata_path: " + abs_metadata_path)
    
    extract_path = path.abspath(path.join(EXTRACT_DIR, DUMP_PATH))

    print("Try to dump il2cpp...")
    dumper.dump_il2cpp(
        extract_path, abs_il2cpp_path, abs_metadata_path, 5
    )
    notice("Dump il2cpp binary file successfully.")
    DUMP_DIR = "/home/runner/work/BaPs/BaPs/Dump"
    notice("Looking into extracted folder...")
    DUMP_LINK="https://drive.google.com/uc?export=download&id=1ZC3FvMwFmTGNiUnUGO4P4xrsES6j3qeP"
    def extract_zip_file(zip_path: str) -> None:    
        notice("Zip Path: "+zip_path)
        apk_files = ZipUtils.extract_zip(
            zip_path, path.join(TEMP_DIR), keywords=["zip"]
        )

        ZipUtils.extract_zip(
            apk_files, EXTRACT_DIR, zips_dir=TEMP_DIR
        )


    zip_path = download_filezip()
    extract_zip_file(zip_path)
    
    compile_python(path.join(extract_path, "dump.cs"), EXTRACT_DIR)
    notice("Generated FlatData to dir: " + EXTRACT_DIR)