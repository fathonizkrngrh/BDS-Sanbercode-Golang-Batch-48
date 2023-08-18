CREATE TABLE `data_mhs`.`nilai_mhs` (
  `id` INT NOT NULL,
  `nama` VARCHAR(255) NOT NULL,
  `mata_kuliah` VARCHAR(255) NOT NULL,
  `nilai` INT NOT NULL,
  `indeks_nilai` VARCHAR(2) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`));
