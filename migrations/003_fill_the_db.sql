INSERT INTO users (email, password, username) VALUES
    ('zhenya_krylov', 'superclass228', 'Evgeniy Krylov');
INSERT INTO users (email, password, username) VALUES
    ('solovei_razboynik', 'gachanaski2005', 'Solovey Razboynik');
INSERT INTO users (email, password, username) VALUES
    ('zhenya_krylov', 'superclass228', 'Evgeniy Krylov');
INSERT INTO items(name, description, price, stock, country) VALUES ('AKM Assault Rifle', 'Developed in the late 1950s, the Kalashnikov modernized automatic rifle is an upgraded version of the classic AK rifle. It is the most ubiquitous variant of the entire AK ' ||
                                                                                            'series of firearms with a rich history of military applications.', 499, 14, 'Russia');
ALTER TABLE items ADD COLUMN country VARCHAR(100);