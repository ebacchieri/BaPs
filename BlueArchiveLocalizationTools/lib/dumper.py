"""Dump il2cpp file to csharp file."""

import json
import os
from os import path
from lib.console import notice
from lib.downloader import FileDownloader
from utils.util import CommandUtils, FileUtils, ZipUtils

IL2CPP_ZIP = "https://github.com/akio7624/Il2CppDumper/releases/download/v25.06.25/Il2CppDumper-net8-akiobuild-v25.06.25.zip"
IL2CPP_FOLDER = ""
ZIP_NAME = "Il2CppDumper-net8-akiobuild-v25.06.25.zip"


class IL2CppDumper:

    def __init__(self) -> None:
        self.project_dir = ""

    def get_il2cpp_dumper(self, save_path: str) -> None:
        FileDownloader(IL2CPP_ZIP).save_file(path.join(save_path, ZIP_NAME))
        ZipUtils.extract_zip(path.join(save_path, ZIP_NAME), save_path)
        notice("Config path to search: "+path.join(save_path, IL2CPP_FOLDER))
        notice("Looking into extracted folder...")
        for dir_path, _, files in os.walk(save_path):
            for file in files:
                    notice(os.path.join(dir_path, file))
        if not (
            config_path := FileUtils.find_files(
                path.join(save_path, IL2CPP_FOLDER), ["config.json"], True
            )
        ):
            raise FileNotFoundError(
                "Cannot find config file. Make sure il2cpp-dumper exsist."
            )

        with open(config_path[0], "r+", encoding="utf8") as config:
            il2cpp_config: dict = json.load(config)
            il2cpp_config["RequireAnyKey"] = False
            il2cpp_config["GenerateDummyDll"] = False
            config.seek(0)
            config.truncate()
            json.dump(il2cpp_config, config)

        self.project_dir = path.dirname(config_path[0])

    def dump_il2cpp(
        self,
        extract_path: str,
        il2cpp_path: str,
        global_metadata_path: str,
        max_retries: int = 1,
    ) -> None:
        """Dump il2cpp with using il2cpp-dumper.
        Args:
            extract_path (str): Absolute path to extract dump file.
            il2cpp_path (str): Absolute path to il2cpp lib.
            global_metadata_path (str): Absolute path to global metadata.
            max_retries (int): Max retry count for dump when dump failed.


        Raises:
            RuntimeError: Raise error when dump unsuccess.
        """
        notice("Creating dir...")
        os.makedirs(extract_path, exist_ok=True)
        notice("Dir created: "+extract_path)
        notice("Dumper dir here: "+self.project_dir)

        success, err = CommandUtils.run_command(
           "dotnet",
           "exec",
            "/home/runner/work/BaPs/BaPs/"+self.project_dir+"/Il2CppDumper.dll",
            il2cpp_path,
            global_metadata_path,
            extract_path,
            cwd=self.project_dir,
        )
        notice("Dumper work:  "+err)
        if not success:
            notice("Failed to extract, trying again... ")
            notice("Error: "+err)
            if max_retries == 0:
                raise RuntimeError(
                    f"Error occurred during dump the lib2cpp file. Retry might solve this issue. Info: {err}"
                )
            return self.dump_il2cpp(
                extract_path, il2cpp_path, global_metadata_path, max_retries - 1
            )

        return None
