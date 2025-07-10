f_url1 = "https://drive.google.com/uc?export=download&id=1wCpIqTbCusU323VXGM0Px4jsM85XG_MD"
f_url2 = "https://drive.google.com/uc?export=download&id=1nH11IQpMRIj1NpidiwFymvTbpSxpOr-q"
f_url3 = "https://drive.google.com/uc?export=download&id=1bYEtTSVJWhDlhKQrgExrMwXYnVJX1vDG"
f_url4 = "https://drive.google.com/uc?export=download&id=1_Aj1omcWG4a23cqNEAHSoDqeomp3BXxT"
f_url = [f_url1, f_url2, f_url3, f_url4]
def download_files(extract_path: str) -> list[str]:
    import glob
    import os
    import shutil
    TEMP_DIR = "Temp"
    os.makedirs(TEMP_DIR, exist_ok=True)
    # apk_dir = glob.glob(f"./{TEMP_DIR}/*.zip")
    # if len(apk_dir) > 0:
    #     return apk_dir[0].replace("\\", "/")
    from lib.downloader import FileDownloader
    from lib.console import ProgressBar, notice
    from os import path
    notice("Downloading dump files...")    

    if not (
        (
            apk_req := FileDownloader(
                f_url1,
                request_method="get",
                use_cloud_scraper=True,
                verbose=True,
            )
        )
        and (apk_data := apk_req.get_response(True))
    ):
        raise LookupError("Cannot fetch zip info.")

    path1 = path.join(
        extract_path,
        "dump.cs"
    )
    path2 = path.join(
        extract_path,
        "il2cpp.h"
    )
    path3 = path.join(
        extract_path,
        "script.json"
    )
    path4 = path.join(
        extract_path,
        "stringliteral.json"
    )
    paths = [path1, path2, path3, path4]
    # apk_size = int(apk_data.headers.get("Content-Length", 0))

    # if path.exists(apk_path) and path.getsize(apk_path) == apk_size:
    #     return apk_path
    for x in range(0, 4):
        notice("Downloading url: "+f_url[x])
        FileDownloader(
            f_url[x],
            request_method="get",
            enable_progress=True,
            use_cloud_scraper=True,
        ).save_file(paths[x])    
    return '|'.join(paths).replace("\\", "/")
