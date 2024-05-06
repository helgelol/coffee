CREATE DATABASE IF NOT EXISTS coffee;
USE coffee;

CREATE TABLE IF NOT EXISTS beans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    country VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    producer VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    process VARCHAR(255) NOT NULL,
    flavours VARCHAR(255) NOT NULL
);


INSERT INTO beans (country, region, producer, name, process, flavours) VALUES
  ("Ethiopia","Yirgacheffe","Munin","Burtukaana Wote Konga","Natural","Round, Citric, With Stone fruit notes, and a good structure."),
  ("Ethiopia","Sidamo","Munin","Samii Bensa","Washed","Herbal, tropical, fruity acidity."),
  ("Ethiopia", "Yirgacheffe", "Munin", "Burtukaana Danche","Natural", "Round, Sweety, Citric, with stone fruit notes"),
  ("El Salvador", "Apaneca-Ilamatepec", "Jacobsen og Svart", "El Limo", "Anaerobisk Ferment", "Vinous, floral, tropical fruit and caramel"),
  ("Ethiopia", "Shakisso, Guji", "Jacobsen og Svart", "Koromii", "Berry dried", "Strawberry, fruit, mandarin"),
  ("Kenya", "Kirinyaga", "Munin","PB Kiri", "Washed", "Pleasant citrus, fruity, sweet, hints of tea");
