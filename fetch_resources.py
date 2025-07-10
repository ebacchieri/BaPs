import sys
from pathlib import Path
import shutil
sys.path.insert(0, str(Path("./BlueArchiveLocalizationTools")))
import setup_flatdata
from extractor import TablesExtractor
TablesExtractor('Extracted', './resources/TableBundles').extract_tables()

shutil.move(Path("Extracted/Table/Excel"), Path("resources/"))
shutil.move(Path("Extracted/Table/ExcelDB"), Path("resources/"))