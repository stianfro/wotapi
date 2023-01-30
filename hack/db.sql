CREATE TABLE manga (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  author TEXT NOT NULL,
  magazine TEXT NOT NULL,
  publisher TEXT NOT NULL
) CREATE TABLE mangaVolumes (
  id TEXT PRIMARY KEY,
  mangaID TEXT NOT NULL,
  number INTEGER NOT NULL,
  title TEXT NOT NULL,
  releaseDate TEXT NOT NULL,
  isbn TEXT NOT NULL,
  FOREIGN KEY (mangaID) REFERENCES manga(id)
) CREATE TABLE mangaChapters (
  id TEXT PRIMARY KEY,
  volumeID TEXT NOT NULL,
  number INTEGER NOT NULL,
  title TEXT NOT NULL,
  FOREIGN KEY (volumeID) REFERENCES mangaVolumes(id)
)