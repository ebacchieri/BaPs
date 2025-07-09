apk_url = "https://drive.google.com/uc?export=download&id=1a9YEYLw3hRhcmlk2FZKZOKVYeCujG_cp"
def download_file() -> str:
    import glob
    import os
    TEMP_DIR = "Temp"
    os.makedirs(TEMP_DIR, exist_ok=True)
    apk_dir = glob.glob(f"./{TEMP_DIR}/*.zip")
    if len(apk_dir) > 0:
        return apk_dir[0].replace("\\", "/")
    from lib.downloader import FileDownloader
    from lib.console import ProgressBar, notice
    from os import path
    notice("Downloading zip...")
    if not (
        (
            apk_req := FileDownloader(
                apk_url,
                request_method="get",
                use_cloud_scraper=True,
                verbose=True,
            )
        )
        and (apk_data := apk_req.get_response(True))
    ):
        raise LookupError("Cannot fetch zip info.")

    apk_path = path.join(
        TEMP_DIR,
        "Dump.zip"
    )
    apk_size = int(apk_data.headers.get("Content-Length", 0))

    if path.exists(apk_path) and path.getsize(apk_path) == apk_size:
        return apk_path

    FileDownloader(
        apk_url,
        request_method="get",
        enable_progress=True,
        use_cloud_scraper=True,
    ).save_file(apk_path)

    return apk_path.replace("\\", "/")
