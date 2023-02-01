CREATE TABLE manga (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  author TEXT NOT NULL,
  magazine TEXT NOT NULL,
  publisher TEXT NOT NULL
);
CREATE TABLE manga_volumes (
  id TEXT PRIMARY KEY,
  manga_id TEXT NOT NULL,
  number INTEGER NOT NULL,
  title TEXT NOT NULL,
  release_date TEXT NOT NULL,
  isbn TEXT NOT NULL,
  FOREIGN KEY (manga_id)
    REFERENCES manga(id)
);
CREATE TABLE manga_chapters (
  id TEXT PRIMARY KEY,
  volume_id TEXT NOT NULL,
  number INTEGER NOT NULL,
  title TEXT NOT NULL,
  FOREIGN KEY (volume_id)
    REFERENCES manga_volumes(id)
);