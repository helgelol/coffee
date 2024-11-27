CREATE TABLE IF NOT EXISTS beans (
  id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  country varchar(255) NOT NULL,
  region varchar(255) NOT NULL,
  producer varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  process varchar(255) NOT NULL,
  flavours varchar(255) NOT NULL
);
INSERT INTO beans (
    country,
    region,
    producer,
    name,
    process,
    flavours
  )
VALUES (
    'Ethiopia',
    'Yirgacheffe',
    'Munin',
    'Burtukaana Wote Konga',
    'Natural',
    'Round, Citric, With Stone fruit notes, and a good structure.'
  ),
  (
    'Ethiopia',
    'Sidamo',
    'Munin',
    'Samii Bensa',
    'Washed',
    'Herbal, tropical, fruity acidity.'
  ),
  (
    'Ethiopia',
    'Yirgacheffe',
    'Munin',
    'Burtukaana Danche',
    'Natural',
    'Round, Sweety, Citric, with stone fruit notes'
  ),
  (
    'El Salvador',
    'Apaneca-Ilamatepec',
    'Jacobsen og Svart',
    'El Limo',
    'Anaerobisk Ferment',
    'Vinous, floral, tropical fruit and caramel'
  ),
  (
    'Ethiopia',
    'Shakisso, Guji',
    'Jacobsen og Svart',
    'Koromii',
    'Berry dried',
    'Strawberry, fruit, mandarin'
  ),
  (
    'Kenya',
    'Kirinyaga',
    'Munin',
    'PB Kiri',
    'Washed',
    'Pleasant citrus, fruity, sweet, hints of tea'
  ),
  (
    'Peru',
    'San Ignazio',
    'Munin',
    'Geisha Maximiliano Gracia',
    'Washed',
    'Complex, sweet, floral, fruit'
  ),
  (
    'Ethiopia',
    'Gedeb',
    'Munin',
    'Diima Chelchele',
    'Natural',
    'Layered and complex. Peach, sweet orange and raspberries.'
  ),
  (
  'Honduras',
  'Marcala',
  'Jacu',
  'Caballeros Lomas',
  'Berry dried',
  'Strawberry, dried fruits and caramel.'
  ),
  (
  'El Salvador',
  'Apaneca-Ilamatepec',
  'The Coast Coffee',
  'Finca Las Brisas',
  'Honey',
  'Hazelnut, Toasted Almond, Dark Chocolate.'
  ),
  (
  'Colombia',
  'Valle de Cauca',
  'The Coast Coffee',
  'Pink Bourbon',
  'Washed',
  'Cherry, Red wine, Apple, Cacao Nibs'
  );
