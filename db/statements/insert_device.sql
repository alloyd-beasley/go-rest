INSERT INTO device (manufacturer_d_address_1, manufacturer_d_state, manufacturer_d_postal_code, manufacturer_d_city, manufacturer_d_country, manufacturer_d_name, lot_number, model_number, generic_name, brand_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id

