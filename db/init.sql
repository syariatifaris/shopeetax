CREATE TABLE IF NOT EXISTS shopee_tax(
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    expression VARCHAR(255) NOT NULL    
);

INSERT INTO shopee_tax (code, description, expression) VALUES 
    ('Tobacco', 'tax for tobacco', '0.1 * ${price}'),
    ('Food', 'tax for food items', '10 + (0.02 * ${price})');

CREATE TABLE IF NOT EXISTS  shopee_taxable_product(
    id INT PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255) NOT NULL,
    tax_category_id INT NOT NULL,
    price FLOAT
);

CREATE TABLE IF NOT EXISTS shoppe_tax_component(
    id INT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    var_type VARCHAR(255) NOT NULL
);

INSERT INTO shoppe_tax_component (code, description, var_type) VALUES
    ('price', 'product price component', 'float32');

SELECT * FROM shopee_tax;
SELECT * FROM shoppe_tax_component;
SELECT * FROM shopee_taxable_product;