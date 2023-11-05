CREATE TABLE user (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       name VARCHAR(32),
       picture VARCHAR
);

CREATE TABLE recipe (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       author INTEGER,
       name VARCHAR,
       thumbnail VARCHAR,
       preperation_time INTEGER,
       preperation VARCHAR,
       FOREIGN KEY (author) REFERENCES user(id)
);

CREATE TABLE ingredient (
       id INTEGER PRIMARY KEY AUTOINCREMENT,
       name VARCHAR,
       unit VARCHAR(16)
);

CREATE TABLE recipe_ingredient (
       recipe_id INTEGER,
       ingredient_id INTEGER,
       amount INTEGER,
       FOREIGN KEY (recipe_id) REFERENCES recipe(id),
       FOREIGN KEY (ingredient_id) REFERENCES ingredient(id)
);

CREATE TABLE shopping_list (
       user_id INTEGER,
       ingredient_id INTEGER,
       total_amount INTEGER,
       FOREIGN KEY (user_id) REFERENCES user(id),
       FOREIGN KEY (ingredient_id) REFERENCES ingredient(id)
);
