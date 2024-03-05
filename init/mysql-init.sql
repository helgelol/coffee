CREATE DATABASE IF NOT EXISTS coffee;
USE coffee;

CREATE TABLE IF NOT EXISTS beans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    country VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    producer VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    process VARCHAR(255) NOT NULL,
    flavours VARCHAR(255) NOT NULL,
);


INSERT INTO beans (country, region, producer, name, process, flavours) VALUES
  ("Ethiopia","Yirgacheffe","Munin","Burtukaana Wote Konga","Natural","Round, Citric, With Stone fruit notes, and a good structure."),
  ("Ethiopia","Sidamo","Munin","Samii Bensa","Washed","Herbal, tropical, fruity acidity.");
