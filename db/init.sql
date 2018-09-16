CREATE TABLE IF NOT EXISTS shopee_tax(
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    expression VARCHAR(255) NOT NULL,
    min_price FLOAT DEFAULT 0.0
);

INSERT INTO shopee_tax (code, description, expression, min_price) VALUES 
    ('Food', 'tax for tobacco', '0.1 * ${price}', -1),
    ('Tobaco', 'tax for food items', '10 + (0.02 * ${price})', -1),
    ('Entertainment', 'tax for entertainment', '0.1 * (${price} - 100)', 99);


CREATE TABLE IF NOT EXISTS  shopee_taxable_product(
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255) NOT NULL,
    tax_category_id INT NOT NULL,
    price FLOAT DEFAULT 0.0,
    tax_price FLOAT DEFAULT 0.0,
    total_price FLOAT DEFAULT 0.0
);

CREATE TABLE IF NOT EXISTS shopee_tax_component(
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    var_type VARCHAR(255) NOT NULL
);

INSERT INTO shopee_tax_component (code, description, var_type) VALUES
    ('price', 'product price component', 'float32');

SELECT * FROM shopee_tax_component;